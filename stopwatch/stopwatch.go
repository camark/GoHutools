package stopwatch

import (
	"fmt"
	"strings"
	"time"
)

// Task is a recorded stopwatch segment.
type Task struct {
	Name string
	Time time.Duration
}

// Split is a checkpoint recorded while a task is still running.
type Split struct {
	Name string
	Time time.Duration // cumulative elapsed time at the moment of the split
}

// StopWatch measures the duration of one or more sequential tasks —
// a Go port of Hutool's StopWatch.
type StopWatch struct {
	name     string
	tasks    []Task
	splits   []Split
	running  bool
	curName  string
	curStart time.Time
	elapsed  time.Duration
}

// New creates a StopWatch with a default name.
func New() *StopWatch {
	return NewNamed("StopWatch")
}

// NewNamed creates a StopWatch with the given name (used in PrettyPrint).
func NewNamed(name string) *StopWatch {
	return &StopWatch{name: name}
}

// Start begins timing a task with the given name.
// It panics if called while a task is already running.
func (sw *StopWatch) Start(name string) {
	if sw.running {
		panic("stopwatch: Start called while a task is already running")
	}
	sw.running = true
	sw.curName = name
	sw.curStart = time.Now()
}

// Stop stops the current task and records it. It is a no-op when idle.
func (sw *StopWatch) Stop() {
	if !sw.running {
		return
	}
	d := time.Since(sw.curStart)
	sw.tasks = append(sw.tasks, Task{Name: sw.curName, Time: d})
	sw.elapsed += d
	sw.running = false
}

// Split records a checkpoint name with the current cumulative elapsed time
// without stopping the task. Panics when not running.
func (sw *StopWatch) Split(name string) {
	if !sw.running {
		panic("stopwatch: Split called while not running")
	}
	sw.splits = append(sw.splits, Split{Name: name, Time: sw.Elapsed()})
}

// Reset clears all recorded tasks, splits and the running state.
func (sw *StopWatch) Reset() {
	sw.tasks = nil
	sw.splits = nil
	sw.running = false
	sw.curName = ""
	sw.elapsed = 0
}

// IsRunning reports whether a task is currently being timed.
func (sw *StopWatch) IsRunning() bool { return sw.running }

// Elapsed returns the total measured time, including the active task if running.
func (sw *StopWatch) Elapsed() time.Duration {
	e := sw.elapsed
	if sw.running {
		e += time.Since(sw.curStart)
	}
	return e
}

// ElapsedMilliseconds returns the total measured time in milliseconds.
func (sw *StopWatch) ElapsedMilliseconds() int64 {
	return sw.Elapsed().Milliseconds()
}

// TotalTime is an alias of Elapsed.
func (sw *StopWatch) TotalTime() time.Duration { return sw.Elapsed() }

// LastTaskName returns the name of the most recently stopped task,
// or the running task's name when running.
func (sw *StopWatch) LastTaskName() string {
	if sw.running {
		return sw.curName
	}
	if n := len(sw.tasks); n > 0 {
		return sw.tasks[n-1].Name
	}
	return ""
}

// LastTaskTime returns the duration of the most recently stopped task.
func (sw *StopWatch) LastTaskTime() time.Duration {
	if n := len(sw.tasks); n > 0 {
		return sw.tasks[n-1].Time
	}
	return 0
}

// TaskCount returns the number of stopped tasks.
func (sw *StopWatch) TaskCount() int { return len(sw.tasks) }

// Tasks returns the recorded tasks (a snapshot).
func (sw *StopWatch) Tasks() []Task {
	out := make([]Task, len(sw.tasks))
	copy(out, sw.tasks)
	return out
}

// Splits returns the recorded checkpoints (a snapshot).
func (sw *StopWatch) Splits() []Split {
	out := make([]Split, len(sw.splits))
	copy(out, sw.splits)
	return out
}

// PrettyPrint renders a summary table of tasks with their share of the total.
func (sw *StopWatch) PrettyPrint() string {
	var sb strings.Builder
	total := sw.TotalTime().Milliseconds()
	if total <= 0 {
		total = 1
	}
	fmt.Fprintf(&sb, "StopWatch '%s': running time = %d ms\n", sw.name, sw.TotalTime().Milliseconds())
	fmt.Fprintf(&sb, "----------------------------------------\n")
	for i, t := range sw.tasks {
		pct := float64(t.Time.Milliseconds()) * 100 / float64(total)
		fmt.Fprintf(&sb, "%s : %5d ms, %6.2f%%\n", t.Name, t.Time.Milliseconds(), pct)
		if i < len(sw.tasks)-1 {
			sb.WriteString("----------------------------------------\n")
		}
	}
	return sb.String()
}

// String renders a compact one-line summary.
func (sw *StopWatch) String() string {
	return fmt.Sprintf("StopWatch '%s': %d tasks, total %d ms",
		sw.name, len(sw.tasks), sw.TotalTime().Milliseconds())
}