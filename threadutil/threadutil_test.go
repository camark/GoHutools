package threadutil

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	start := time.Now()
	Sleep(50 * time.Millisecond)
	elapsed := time.Since(start)
	if elapsed < 40*time.Millisecond {
		t.Errorf("Sleep(50ms) woke up after %v", elapsed)
	}
}

func TestExecute(t *testing.T) {
	var done atomic.Bool
	Execute(func() {
		done.Store(true)
	})
	deadline := time.Now().Add(2 * time.Second)
	for !done.Load() && time.Now().Before(deadline) {
		time.Sleep(5 * time.Millisecond)
	}
	if !done.Load() {
		t.Error("Execute task never ran")
	}
}

func TestExecuteAsync(t *testing.T) {
	ch := ExecuteAsync(func() int {
		time.Sleep(20 * time.Millisecond)
		return 42
	})
	select {
	case v := <-ch:
		if v != 42 {
			t.Errorf("ExecuteAsync returned %d, want 42", v)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("ExecuteAsync timed out")
	}
}

func TestNewExecutor(t *testing.T) {
	exec := NewExecutor(2)
	defer exec.Shutdown()
	if got := exec.NumActive(); got > 2 {
		t.Errorf("NumActive = %d, want <= 2", got)
	}
}

func TestExecutorSubmit(t *testing.T) {
	exec := NewExecutor(4)
	defer exec.Shutdown()

	var wg sync.WaitGroup
	results := make([]int, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		idx := i
		exec.Submit(func() {
			results[idx] = idx * idx
			wg.Done()
		})
	}
	wg.Wait()
	for i := 0; i < 10; i++ {
		if results[i] != i*i {
			t.Errorf("Submit result[%d] = %d, want %d", i, results[i], i*i)
		}
	}
}

func TestExecutorSubmitWithResult(t *testing.T) {
	exec := NewExecutor(2)
	defer exec.Shutdown()

	f := SubmitWithResult(exec, func() int { return 7 })
	select {
	case v := <-f:
		if v != 7 {
			t.Errorf("SubmitWithResult = %d, want 7", v)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("SubmitWithResult timed out")
	}
}

func TestExecutorShutdown(t *testing.T) {
	exec := NewExecutor(2)
	var ran atomic.Bool
	exec.Submit(func() {
		time.Sleep(10 * time.Millisecond)
		ran.Store(true)
	})
	exec.Shutdown()
	if !ran.Load() {
		t.Error("Shutdown returned before queued task completed")
	}
}

func TestAwaitAll(t *testing.T) {
	exec := NewExecutor(3)
	var count atomic.Int32
	for i := 0; i < 9; i++ {
		exec.Submit(func() {
			time.Sleep(10 * time.Millisecond)
			count.Add(1)
		})
	}
	exec.AwaitAll(5 * time.Second)
	if c := count.Load(); c != 9 {
		t.Errorf("AwaitAll: count = %d, want 9", c)
	}
}

func TestConcurrencyLimit(t *testing.T) {
	exec := NewExecutor(2)
	defer exec.Shutdown()

	var active atomic.Int32
	var maxActive atomic.Int32
	for i := 0; i < 20; i++ {
		exec.Submit(func() {
			cur := active.Add(1)
			for {
				m := maxActive.Load()
				if cur <= m || maxActive.CompareAndSwap(m, cur) {
					break
				}
			}
			time.Sleep(5 * time.Millisecond)
			active.Add(-1)
		})
	}
	exec.AwaitAll(5 * time.Second)
	if m := maxActive.Load(); m > 2 {
		t.Errorf("max concurrent = %d, want <= 2", m)
	}
	if m := maxActive.Load(); m < 1 {
		t.Errorf("max concurrent = %d, want >= 1", m)
	}
}

func TestPoolReuse(t *testing.T) {
	exec := NewExecutor(1)
	defer exec.Shutdown()
	runs := 3
	done := make(chan struct{}, runs)
	for i := 0; i < runs; i++ {
		exec.Submit(func() { done <- struct{}{} })
	}
	for i := 0; i < runs; i++ {
		select {
		case <-done:
		case <-time.After(5 * time.Second):
			t.Fatal("pool failed to execute all tasks")
		}
	}
}

func TestScheduledExecutor(t *testing.T) {
	var ticks atomic.Int32
	sched := NewScheduledExecutor(20*time.Millisecond, func() {
		ticks.Add(1)
	})
	sched.Start()

	deadline := time.Now().Add(150 * time.Millisecond)
	for ticks.Load() < 3 && time.Now().Before(deadline) {
		time.Sleep(10 * time.Millisecond)
	}
	sched.Stop()
	if n := ticks.Load(); n < 3 {
		t.Errorf("expected >= 3 ticks, got %d", n)
	}

	// Start then Stop early: no ticks after Stop
	before := ticks.Load()
	_ = before
	sched.Start()
	time.Sleep(30 * time.Millisecond)
	sched.Stop()
	n := ticks.Load()
	time.Sleep(60 * time.Millisecond)
	if after := ticks.Load(); after != n {
		t.Errorf("ticks continued after Stop: %d -> %d", n, after)
	}
}

func TestScheduledExecutorNoDoubleStart(t *testing.T) {
	var ticks atomic.Int32
	sched := NewScheduledExecutor(10*time.Millisecond, func() { ticks.Add(1) })
	sched.Start()
	sched.Start() // no-op
	time.Sleep(35 * time.Millisecond)
	sched.Stop()
	if n := ticks.Load(); n < 1 {
		t.Errorf("expected ticks, got %d", n)
	}
}

func TestSleepSec(t *testing.T) {
	start := time.Now()
	SleepSec(0)
	if time.Since(start) > time.Second {
		t.Error("SleepSec(0) should not block for a second")
	}
}