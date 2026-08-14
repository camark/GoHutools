package cron

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// Common cron expressions (6-field: second minute hour dayOfMonth month dayOfWeek)
const (
	EverySecond = "* * * * * *"
	EveryMinute = "0 * * * * *"
	EveryHour   = "0 0 * * * *"
	EveryDay    = "0 0 0 * * *"
	EveryWeek   = "0 0 0 * * 0"
	EveryMonth  = "0 0 0 1 * *"
	EveryYear   = "0 0 0 1 1 *"
)

// Scheduler is cron scheduler
type Scheduler struct {
	tasks   []*Task
	running bool
	stop    chan struct{}
	mu      sync.RWMutex
	nextID  int
}

// Task is scheduled task
type Task struct {
	id      int
	cron    string
	fn      func()
	lastRun time.Time
	running atomic.Bool // accessed by both the scheduler goroutine and the task goroutine
	stop    chan struct{}
}

// IntervalTask is interval-based task
type IntervalTask struct {
	scheduler *Scheduler
	interval  time.Duration
	fn        func()
}

// cronSchedule represents parsed cron fields
type cronSchedule struct {
	second     []int
	minute     []int
	hour       []int
	dayOfMonth []int
	month      []int
	dayOfWeek  []int
}

// New creates new scheduler
func New() *Scheduler {
	return &Scheduler{
		tasks:   make([]*Task, 0),
		running: false,
		stop:    make(chan struct{}),
		nextID:  1,
	}
}

// AddFunc adds function with cron expression
func (s *Scheduler) AddFunc(cron string, fn func()) (int, error) {
	_, err := parseCron(cron)
	if err != nil {
		return 0, err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.nextID
	s.nextID++

	task := &Task{
		id:   id,
		cron: cron,
		fn:   fn,
		stop: make(chan struct{}),
	}

	s.tasks = append(s.tasks, task)
	return id, nil
}

// AddTask adds task
func (s *Scheduler) AddTask(task *Task) (int, error) {
	_, err := parseCron(task.cron)
	if err != nil {
		return 0, err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.nextID
	s.nextID++

	task.id = id
	if task.stop == nil {
		task.stop = make(chan struct{})
	}

	s.tasks = append(s.tasks, task)
	return id, nil
}

// Remove removes task by id
func (s *Scheduler) Remove(id int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, task := range s.tasks {
		if task.id == id {
			close(task.stop)
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return
		}
	}
}

// Start starts scheduler
func (s *Scheduler) Start() {
	s.mu.Lock()
	if s.running {
		s.mu.Unlock()
		return
	}
	s.running = true
	s.stop = make(chan struct{})
	s.mu.Unlock()

	go s.run()
}

// Stop stops scheduler
func (s *Scheduler) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.running {
		return
	}

	s.running = false
	close(s.stop)

	for _, task := range s.tasks {
		select {
		case <-task.stop:
		default:
			close(task.stop)
		}
	}
}

// IsRunning checks if scheduler is running
func (s *Scheduler) IsRunning() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.running
}

// Tasks returns all tasks
func (s *Scheduler) Tasks() []*Task {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]*Task, len(s.tasks))
	copy(result, s.tasks)
	return result
}

// Every creates interval task
func (s *Scheduler) Every(interval time.Duration) *IntervalTask {
	return &IntervalTask{
		scheduler: s,
		interval:  interval,
	}
}

// Do sets function for interval task and registers it
func (t *IntervalTask) Do(fn func()) *IntervalTask {
	t.fn = fn

	// Create a cron expression for the interval
	cronExpr := intervalToCron(t.interval)

	t.scheduler.mu.Lock()
	defer t.scheduler.mu.Unlock()

	id := t.scheduler.nextID
	t.scheduler.nextID++

	task := &Task{
		id:   id,
		cron: cronExpr,
		fn:   fn,
		stop: make(chan struct{}),
	}

	t.scheduler.tasks = append(t.scheduler.tasks, task)
	return t
}

// intervalToCron converts duration to cron expression
func intervalToCron(d time.Duration) string {
	seconds := int(d.Seconds())

	if seconds < 1 {
		seconds = 1
	}

	if seconds < 60 {
		if seconds == 0 {
			seconds = 1
		}
		return fmt.Sprintf("*/%d * * * * *", seconds)
	}

	minutes := seconds / 60
	if minutes < 60 {
		return fmt.Sprintf("0 */%d * * * *", minutes)
	}

	hours := minutes / 60
	if hours < 24 {
		return fmt.Sprintf("0 0 */%d * * *", hours)
	}

	return "0 0 0 * * *"
}

// run is the main scheduler loop
func (s *Scheduler) run() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-s.stop:
			return
		case now := <-ticker.C:
			s.executeTasks(now)
		}
	}
}

// executeTasks executes all due tasks
func (s *Scheduler) executeTasks(now time.Time) {
	s.mu.RLock()
	tasks := make([]*Task, len(s.tasks))
	copy(tasks, s.tasks)
	s.mu.RUnlock()

	for _, task := range tasks {
		schedule, err := parseCron(task.cron)
		if err != nil {
			continue
		}

		nextRun := schedule.Next(task.lastRun)
		if now.After(nextRun) || now.Equal(nextRun) {
			select {
			case <-task.stop:
				continue
			default:
			}

			if task.running.CompareAndSwap(false, true) {
				task.lastRun = now
				go func(t *Task) {
					defer t.running.Store(false)
					t.fn()
				}(task)
			}
		}
	}
}

// parseCron parses a 6-field cron expression
// Format: second minute hour dayOfMonth month dayOfWeek
func parseCron(expr string) (cronSchedule, error) {
	fields := strings.Fields(expr)
	if len(fields) != 6 {
		return cronSchedule{}, fmt.Errorf("invalid cron expression: expected 6 fields, got %d", len(fields))
	}

	schedule := cronSchedule{}
	var err error

	schedule.second, err = parseCronField(fields[0], 0, 59)
	if err != nil {
		return cronSchedule{}, fmt.Errorf("invalid second field: %v", err)
	}

	schedule.minute, err = parseCronField(fields[1], 0, 59)
	if err != nil {
		return cronSchedule{}, fmt.Errorf("invalid minute field: %v", err)
	}

	schedule.hour, err = parseCronField(fields[2], 0, 23)
	if err != nil {
		return cronSchedule{}, fmt.Errorf("invalid hour field: %v", err)
	}

	schedule.dayOfMonth, err = parseCronField(fields[3], 1, 31)
	if err != nil {
		return cronSchedule{}, fmt.Errorf("invalid dayOfMonth field: %v", err)
	}

	schedule.month, err = parseCronField(fields[4], 1, 12)
	if err != nil {
		return cronSchedule{}, fmt.Errorf("invalid month field: %v", err)
	}

	schedule.dayOfWeek, err = parseCronField(fields[5], 0, 6)
	if err != nil {
		return cronSchedule{}, fmt.Errorf("invalid dayOfWeek field: %v", err)
	}

	return schedule, nil
}

// parseCronField parses a single cron field
func parseCronField(field string, min, max int) ([]int, error) {
	// Handle wildcard
	if field == "*" {
		result := make([]int, max-min+1)
		for i := min; i <= max; i++ {
			result[i-min] = i
		}
		return result, nil
	}

	// Handle step (*/n)
	if strings.HasPrefix(field, "*/") {
		stepStr := field[2:]
		step, err := strconv.Atoi(stepStr)
		if err != nil || step <= 0 {
			return nil, fmt.Errorf("invalid step: %s", stepStr)
		}

		var result []int
		for i := min; i <= max; i += step {
			result = append(result, i)
		}
		return result, nil
	}

	// Handle range (n-m)
	if strings.Contains(field, "-") {
		parts := strings.Split(field, "-")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid range: %s", field)
		}

		rangeMin, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid range start: %s", parts[0])
		}

		rangeMax, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid range end: %s", parts[1])
		}

		if rangeMin < min || rangeMax > max || rangeMin > rangeMax {
			return nil, fmt.Errorf("range out of bounds: %s", field)
		}

		result := make([]int, rangeMax-rangeMin+1)
		for i := rangeMin; i <= rangeMax; i++ {
			result[i-rangeMin] = i
		}
		return result, nil
	}

	// Handle list (n,m,o)
	if strings.Contains(field, ",") {
		parts := strings.Split(field, ",")
		result := make([]int, 0, len(parts))

		for _, part := range parts {
			val, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				return nil, fmt.Errorf("invalid value: %s", part)
			}
			if val < min || val > max {
				return nil, fmt.Errorf("value out of bounds: %d", val)
			}
			result = append(result, val)
		}

		sort.Ints(result)
		return result, nil
	}

	// Handle single value
	val, err := strconv.Atoi(field)
	if err != nil {
		return nil, fmt.Errorf("invalid value: %s", field)
	}
	if val < min || val > max {
		return nil, fmt.Errorf("value out of bounds: %d", val)
	}

	return []int{val}, nil
}

// Next calculates the next run time after the given time
func (s cronSchedule) Next(t time.Time) time.Time {
	// Start from the next second
	next := t.Add(time.Second)
	next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), next.Second(), 0, next.Location())

	// Try up to 2 years of seconds to find next match
	maxIterations := 366 * 24 * 60 * 60 * 2
	for i := 0; i < maxIterations; i++ {
		if s.matches(next) {
			return next
		}
		next = next.Add(time.Second)
	}

	// Fallback: return next minute if no match found (shouldn't happen with valid schedule)
	return t.Add(time.Minute)
}

// matches checks if a time matches the cron schedule
func (s cronSchedule) matches(t time.Time) bool {
	sec := t.Second()
	min := t.Minute()
	hour := t.Hour()
	day := t.Day()
	month := int(t.Month())
	weekday := int(t.Weekday())

	return contains(s.second, sec) &&
		contains(s.minute, min) &&
		contains(s.hour, hour) &&
		contains(s.dayOfMonth, day) &&
		contains(s.month, month) &&
		contains(s.dayOfWeek, weekday)
}

// contains checks if a slice contains a value
func contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
