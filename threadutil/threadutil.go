package threadutil

import (
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// Sleep blocks the current goroutine for the given duration.
func Sleep(d time.Duration) {
	time.Sleep(d)
}

// SleepSec sleeps for n seconds.
func SleepSec(sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
}

// Execute runs fn in a new goroutine (fire-and-forget).
func Execute(fn func()) {
	go fn()
}

// ExecuteAsync runs fn in a new goroutine and returns a channel
// receiving its return value.
func ExecuteAsync[T any](fn func() T) <-chan T {
	ch := make(chan T, 1)
	go func() {
		ch <- fn()
	}()
	return ch
}

// Executor is a simple fixed-size goroutine pool mirroring Hutool's ThreadUtil.ExecutorBuilder.
type Executor struct {
	sem    chan struct{}
	wg     sync.WaitGroup
	closed atomic.Bool
}

// NewExecutor creates a fixed goroutine pool with size workers.
// size <= 0 defaults to runtime.NumCPU().
func NewExecutor(size int) *Executor {
	if size <= 0 {
		size = runtime.NumCPU()
	}
	return &Executor{sem: make(chan struct{}, size)}
}

// Submit queues fn to run on the pool. Blocking until a worker is free.
func (e *Executor) Submit(fn func()) {
	if e.closed.Load() {
		panic("threadutil: Submit called after Shutdown")
	}
	e.wg.Add(1)
	e.sem <- struct{}{}
	go func() {
		defer func() {
			<-e.sem
			e.wg.Done()
		}()
		fn()
	}()
}

// SubmitWithResult queues fn on e and returns a channel carrying its result.
func SubmitWithResult[T any](e *Executor, fn func() T) <-chan T {
	ch := make(chan T, 1)
	e.Submit(func() { ch <- fn() })
	return ch
}

// AwaitAll blocks until all submitted tasks have finished, or until timeout.
func (e *Executor) AwaitAll(timeout time.Duration) {
	done := make(chan struct{})
	go func() {
		e.wg.Wait()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(timeout):
	}
}

// Shutdown waits for queued tasks to finish and prevents new submissions.
func (e *Executor) Shutdown() {
	if e.closed.CompareAndSwap(false, true) {
		e.wg.Wait()
	}
}

// NumActive returns the number of currently executing tasks.
func (e *Executor) NumActive() int {
	return len(e.sem)
}

// PoolSize returns the configured worker count.
func (e *Executor) PoolSize() int {
	return cap(e.sem)
}

// DefaultExecutor is the shared pool used by Execute/ExecuteAsync.
var DefaultExecutor = NewExecutor(0)

// NewScheduledExecutor returns a scheduler that runs fns at fixed intervals.
// interval <= 0 is treated as 1 second.
func NewScheduledExecutor(interval time.Duration, fn func()) *ScheduledExecutor {
	if interval <= 0 {
		interval = time.Second
	}
	return &ScheduledExecutor{interval: interval, fn: fn}
}

// ScheduledExecutor mirrors Hutool's ThreadUtil.createScheduledExecutor.
type ScheduledExecutor struct {
	interval time.Duration
	fn       func()
	stopCh   chan struct{}
	done     chan struct{}
	mu       sync.Mutex
	running  bool
}

// Start begins periodic execution in a background goroutine.
func (s *ScheduledExecutor) Start() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.running {
		return
	}
	s.running = true
	s.stopCh = make(chan struct{})
	s.done = make(chan struct{})
	go func() {
		defer close(s.done)
		t := time.NewTicker(s.interval)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				s.fn()
			case <-s.stopCh:
				return
			}
		}
	}()
}

// Stop halts periodic execution and waits for the current tick to finish.
func (s *ScheduledExecutor) Stop() {
	s.mu.Lock()
	if !s.running {
		s.mu.Unlock()
		return
	}
	s.running = false
	close(s.stopCh)
	s.mu.Unlock()
	<-s.done
}

// Repeat runs fn count times in sequence on a new goroutine,
// sleeping sleepBetween between invocations.
func Repeat(count int, sleepBetween time.Duration, fn func()) {
	if count <= 0 {
		return
	}
	go func() {
		for i := 0; i < count; i++ {
			fn()
			if sleepBetween > 0 && i < count-1 {
				time.Sleep(sleepBetween)
			}
		}
	}()
}