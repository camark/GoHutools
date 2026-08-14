package cron

import (
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	s := New()
	if s == nil {
		t.Fatal("New() returned nil")
	}
	if s.running {
		t.Error("New scheduler should not be running")
	}
	if len(s.tasks) != 0 {
		t.Errorf("Expected 0 tasks, got %d", len(s.tasks))
	}
}

func TestAddFunc(t *testing.T) {
	s := New()

	id, err := s.AddFunc(EveryMinute, func() {})
	if err != nil {
		t.Fatalf("AddFunc error: %v", err)
	}
	if id != 1 {
		t.Errorf("Expected id 1, got %d", id)
	}

	id, err = s.AddFunc(EveryHour, func() {})
	if err != nil {
		t.Fatalf("AddFunc error: %v", err)
	}
	if id != 2 {
		t.Errorf("Expected id 2, got %d", id)
	}

	if len(s.tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(s.tasks))
	}
}

func TestAddFuncInvalidCron(t *testing.T) {
	s := New()

	_, err := s.AddFunc("invalid", func() {})
	if err == nil {
		t.Error("Expected error for invalid cron expression")
	}

	_, err = s.AddFunc("* * *", func() {})
	if err == nil {
		t.Error("Expected error for 3-field cron expression")
	}
}

func TestAddTask(t *testing.T) {
	s := New()

	task := &Task{
		cron: EveryMinute,
		fn:   func() {},
	}

	id, err := s.AddTask(task)
	if err != nil {
		t.Fatalf("AddTask error: %v", err)
	}
	if id != 1 {
		t.Errorf("Expected id 1, got %d", id)
	}
	if task.id != 1 {
		t.Errorf("Expected task.id 1, got %d", task.id)
	}
}

func TestRemove(t *testing.T) {
	s := New()

	id1, _ := s.AddFunc(EveryMinute, func() {})
	id2, _ := s.AddFunc(EveryHour, func() {})

	if len(s.tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(s.tasks))
	}

	s.Remove(id1)

	if len(s.tasks) != 1 {
		t.Errorf("Expected 1 task after remove, got %d", len(s.tasks))
	}

	if s.tasks[0].id != id2 {
		t.Errorf("Expected remaining task id %d, got %d", id2, s.tasks[0].id)
	}
}

func TestRemoveNonExistent(t *testing.T) {
	s := New()
	s.AddFunc(EveryMinute, func() {})

	s.Remove(999) // Should not panic

	if len(s.tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(s.tasks))
	}
}

func TestStartStop(t *testing.T) {
	s := New()

	if s.IsRunning() {
		t.Error("New scheduler should not be running")
	}

	s.Start()

	if !s.IsRunning() {
		t.Error("Scheduler should be running after Start()")
	}

	s.Stop()

	if s.IsRunning() {
		t.Error("Scheduler should not be running after Stop()")
	}
}

func TestDoubleStart(t *testing.T) {
	s := New()

	s.Start()
	s.Start() // Should not panic

	if !s.IsRunning() {
		t.Error("Scheduler should be running")
	}

	s.Stop()
}

func TestDoubleStop(t *testing.T) {
	s := New()

	s.Start()
	s.Stop()
	s.Stop() // Should not panic

	if s.IsRunning() {
		t.Error("Scheduler should not be running")
	}
}

func TestTasks(t *testing.T) {
	s := New()

	s.AddFunc(EveryMinute, func() {})
	s.AddFunc(EveryHour, func() {})

	tasks := s.Tasks()
	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
}

func TestEvery(t *testing.T) {
	s := New()

	var count int32
	it := s.Every(100 * time.Millisecond)
	it.Do(func() {
		atomic.AddInt32(&count, 1)
	})

	if len(s.tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(s.tasks))
	}
}

func TestParseCron(t *testing.T) {
	tests := []struct {
		expr    string
		wantErr bool
	}{
		{"* * * * * *", false},
		{"0 * * * * *", false},
		{"0 0 * * * *", false},
		{"0 0 0 * * *", false},
		{"0 0 0 * * 0", false},
		{"0 0 0 1 * *", false},
		{"0 0 0 1 1 *", false},
		{"*/5 * * * * *", false},
		{"0 0 0 1-15 * *", false},
		{"0 0 0 1,15 * *", false},
		{"0 0 0 * * 1-5", false},
		{"invalid", true},
		{"* * *", true},
		{"60 * * * * *", true},
		{"* 60 * * * *", true},
		{"* * 24 * * *", true},
		{"* * * 32 * *", true},
		{"* * * * 13 *", true},
		{"* * * * * 7", true},
		{"*/0 * * * * *", true},
	}

	for _, tt := range tests {
		_, err := parseCron(tt.expr)
		if (err != nil) != tt.wantErr {
			t.Errorf("parseCron(%q) error = %v, wantErr %v", tt.expr, err, tt.wantErr)
		}
	}
}

func TestParseCronField(t *testing.T) {
	tests := []struct {
		field   string
		min     int
		max     int
		want    []int
		wantErr bool
	}{
		{"*", 0, 5, []int{0, 1, 2, 3, 4, 5}, false},
		{"*/2", 0, 5, []int{0, 2, 4}, false},
		{"1-3", 0, 5, []int{1, 2, 3}, false},
		{"1,3,5", 0, 5, []int{1, 3, 5}, false},
		{"3", 0, 5, []int{3}, false},
		{"6", 0, 5, nil, true},
		{"invalid", 0, 5, nil, true},
		{"*/0", 0, 5, nil, true},
		{"1-10", 0, 5, nil, true},
	}

	for _, tt := range tests {
		got, err := parseCronField(tt.field, tt.min, tt.max)
		if (err != nil) != tt.wantErr {
			t.Errorf("parseCronField(%q, %d, %d) error = %v, wantErr %v", tt.field, tt.min, tt.max, err, tt.wantErr)
			continue
		}
		if !tt.wantErr {
			if len(got) != len(tt.want) {
				t.Errorf("parseCronField(%q, %d, %d) = %v, want %v", tt.field, tt.min, tt.max, got, tt.want)
			} else {
				for i := range got {
					if got[i] != tt.want[i] {
						t.Errorf("parseCronField(%q, %d, %d)[%d] = %d, want %d", tt.field, tt.min, tt.max, i, got[i], tt.want[i])
					}
				}
			}
		}
	}
}

func TestCronScheduleNext(t *testing.T) {
	schedule := cronSchedule{
		second:     []int{0},
		minute:     []int{0},
		hour:       []int{0},
		dayOfMonth: []int{1},
		month:      []int{1},
		dayOfWeek:  []int{0, 1, 2, 3, 4, 5, 6},
	}

	// Test from Dec 31, 2024 23:59:59
	from := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)
	next := schedule.Next(from)

	expected := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	if !next.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, next)
	}
}

func TestCronScheduleNextEveryMinute(t *testing.T) {
	schedule := cronSchedule{
		second:     []int{0},
		minute:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59},
		hour:       []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
		dayOfMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
		month:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		dayOfWeek:  []int{0, 1, 2, 3, 4, 5, 6},
	}

	from := time.Date(2024, 1, 1, 0, 0, 30, 0, time.UTC)
	next := schedule.Next(from)

	expected := time.Date(2024, 1, 1, 0, 1, 0, 0, time.UTC)
	if !next.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, next)
	}
}

func TestCronScheduleMatches(t *testing.T) {
	schedule := cronSchedule{
		second:     []int{0},
		minute:     []int{30},
		hour:       []int{12},
		dayOfMonth: []int{15},
		month:      []int{6},
		dayOfWeek:  []int{0, 1, 2, 3, 4, 5, 6},
	}

	// Should match
	matchTime := time.Date(2024, 6, 15, 12, 30, 0, 0, time.UTC)
	if !schedule.matches(matchTime) {
		t.Error("Expected match for June 15, 12:30:00")
	}

	// Should not match - wrong second
	noMatchTime := time.Date(2024, 6, 15, 12, 30, 1, 0, time.UTC)
	if schedule.matches(noMatchTime) {
		t.Error("Expected no match for second=1")
	}

	// Should not match - wrong minute
	noMatchTime = time.Date(2024, 6, 15, 12, 31, 0, 0, time.UTC)
	if schedule.matches(noMatchTime) {
		t.Error("Expected no match for minute=31")
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		slice []int
		val   int
		want  bool
	}{
		{[]int{1, 2, 3}, 2, true},
		{[]int{1, 2, 3}, 4, false},
		{[]int{}, 1, false},
		{[]int{1}, 1, true},
	}

	for _, tt := range tests {
		got := contains(tt.slice, tt.val)
		if got != tt.want {
			t.Errorf("contains(%v, %d) = %v, want %v", tt.slice, tt.val, got, tt.want)
		}
	}
}

func TestSchedulerExecution(t *testing.T) {
	s := New()

	var count int32
	var wg sync.WaitGroup
	wg.Add(3)

	_, err := s.AddFunc("*/1 * * * * *", func() {
		defer wg.Done()
		atomic.AddInt32(&count, 1)
	})
	if err != nil {
		t.Fatalf("AddFunc error: %v", err)
	}

	s.Start()

	// Wait for tasks to execute
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// Tasks executed
	case <-time.After(5 * time.Second):
		t.Fatal("Timeout waiting for tasks to execute")
	}

	s.Stop()

	finalCount := atomic.LoadInt32(&count)
	if finalCount < 3 {
		t.Errorf("Expected at least 3 executions, got %d", finalCount)
	}
}

func TestIntervalTaskExecution(t *testing.T) {
	s := New()

	var count int32
	var wg sync.WaitGroup
	wg.Add(2)

	s.Every(1 * time.Second).Do(func() {
		defer wg.Done()
		atomic.AddInt32(&count, 1)
	})

	s.Start()

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// Tasks executed
	case <-time.After(10 * time.Second):
		t.Fatal("Timeout waiting for interval tasks to execute")
	}

	s.Stop()

	finalCount := atomic.LoadInt32(&count)
	if finalCount < 2 {
		t.Errorf("Expected at least 2 executions, got %d", finalCount)
	}
}

func TestRemoveWhileRunning(t *testing.T) {
	s := New()

	var count int32
	id, _ := s.AddFunc("*/1 * * * * *", func() {
		atomic.AddInt32(&count, 1)
	})

	s.Start()
	time.Sleep(1 * time.Second)

	s.Remove(id)
	time.Sleep(1 * time.Second)

	s.Stop()

	// Count should have some executions but not too many after removal
	finalCount := atomic.LoadInt32(&count)
	if finalCount == 0 {
		t.Error("Expected some executions before removal")
	}
}

func TestIntervalToCron(t *testing.T) {
	tests := []struct {
		interval time.Duration
		want     string
	}{
		{30 * time.Second, "*/30 * * * * *"},
		{1 * time.Minute, "0 */1 * * * *"},
		{5 * time.Minute, "0 */5 * * * *"},
		{1 * time.Hour, "0 0 */1 * * *"},
		{24 * time.Hour, "0 0 0 * * *"},
	}

	for _, tt := range tests {
		got := intervalToCron(tt.interval)
		if got != tt.want {
			t.Errorf("intervalToCron(%v) = %q, want %q", tt.interval, got, tt.want)
		}
	}
}

func TestCronConstants(t *testing.T) {
	constants := []struct {
		name  string
		expr  string
		parts int
	}{
		{"EverySecond", EverySecond, 6},
		{"EveryMinute", EveryMinute, 6},
		{"EveryHour", EveryHour, 6},
		{"EveryDay", EveryDay, 6},
		{"EveryWeek", EveryWeek, 6},
		{"EveryMonth", EveryMonth, 6},
		{"EveryYear", EveryYear, 6},
	}

	for _, c := range constants {
		parts := len(strings.Fields(c.expr))
		if parts != c.parts {
			t.Errorf("%s has %d parts, want %d", c.name, parts, c.parts)
		}

		_, err := parseCron(c.expr)
		if err != nil {
			t.Errorf("%s parse error: %v", c.name, err)
		}
	}
}
