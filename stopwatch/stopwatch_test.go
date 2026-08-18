package stopwatch

import (
	"strings"
	"testing"
	"time"
)

func TestStartStop(t *testing.T) {
	sw := New()
	sw.Start("task1")
	time.Sleep(20 * time.Millisecond)
	sw.Stop()

	if sw.IsRunning() {
		t.Error("should not be running after Stop")
	}
	elapsed := sw.ElapsedMilliseconds()
	if elapsed < 15 || elapsed > 500 {
		t.Errorf("ElapsedMilliseconds = %d, want ~20", elapsed)
	}
	if sw.LastTaskName() != "task1" {
		t.Errorf("LastTaskName = %q", sw.LastTaskName())
	}
	if sw.LastTaskTime() < 15*time.Millisecond {
		t.Errorf("LastTaskTime = %v", sw.LastTaskTime())
	}
	if sw.TaskCount() != 1 {
		t.Errorf("TaskCount = %d", sw.TaskCount())
	}
}

func TestSplit(t *testing.T) {
	sw := New()
	sw.Start("s1")
	time.Sleep(10 * time.Millisecond)
	sw.Split("checkpoint") // keep running
	time.Sleep(10 * time.Millisecond)
	sw.Stop()

	if sw.TaskCount() != 1 {
		t.Errorf("Split should not add a task, TaskCount = %d", sw.TaskCount())
	}
	if len(sw.Splits()) != 1 || sw.Splits()[0].Name != "checkpoint" {
		t.Errorf("Splits = %v", sw.Splits())
	}
	// total includes both segments
	if sw.TotalTime() < 15*time.Millisecond {
		t.Errorf("TotalTime = %v", sw.TotalTime())
	}
}

func TestReset(t *testing.T) {
	sw := New()
	sw.Start("x")
	time.Sleep(5 * time.Millisecond)
	sw.Stop()
	sw.Reset()

	if sw.TaskCount() != 0 {
		t.Errorf("TaskCount after Reset = %d", sw.TaskCount())
	}
	if sw.TotalTime() != 0 {
		t.Errorf("TotalTime after Reset = %v", sw.TotalTime())
	}
	if sw.IsRunning() {
		t.Error("Reset should stop running state")
	}
}

func TestMultipleTasks(t *testing.T) {
	sw := New()
	sw.Start("a")
	time.Sleep(5 * time.Millisecond)
	sw.Stop()
	sw.Start("b")
	time.Sleep(8 * time.Millisecond)
	sw.Stop()

	if sw.TaskCount() != 2 {
		t.Errorf("TaskCount = %d, want 2", sw.TaskCount())
	}
	tasks := sw.Tasks()
	if len(tasks) != 2 || tasks[0].Name != "a" || tasks[1].Name != "b" {
		t.Errorf("Tasks = %v", tasks)
	}
	if tasks[0].Time <= 0 || tasks[1].Time <= tasks[0].Time {
		t.Errorf("task times wrong: a=%v b=%v", tasks[0].Time, tasks[1].Time)
	}
}

func TestPrettyPrint(t *testing.T) {
	sw := NewNamed("bench")
	sw.Start("init")
	time.Sleep(5 * time.Millisecond)
	sw.Stop()
	sw.Start("run")
	time.Sleep(5 * time.Millisecond)
	sw.Stop()

	pp := sw.PrettyPrint()
	if !strings.Contains(pp, "bench") {
		t.Errorf("PrettyPrint missing name: %q", pp)
	}
	if !strings.Contains(pp, "init") || !strings.Contains(pp, "run") {
		t.Errorf("PrettyPrint missing tasks: %q", pp)
	}
	if !strings.Contains(pp, "%") {
		t.Errorf("PrettyPrint missing share percent: %q", pp)
	}
}

func TestStartWhileRunning(t *testing.T) {
	sw := NewNamed("t")
	sw.Start("a")
	defer func() {
		if r := recover(); r == nil {
			t.Error("Start while running should panic")
		}
	}()
	sw.Start("b")
}

func TestStopWhenNotRunning(t *testing.T) {
	sw := NewNamed("t")
	sw.Stop() // no-op, shouldn't panic
	sw.Stop()
	if sw.TaskCount() != 0 {
		t.Errorf("TaskCount = %d", sw.TaskCount())
	}
}

func TestStopWatchString(t *testing.T) {
	sw := NewNamed("named")
	if !strings.Contains(sw.String(), "named") {
		t.Errorf("String = %q", sw.String())
	}
}