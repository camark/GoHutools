package pool

import (
	"sync"
	"sync/atomic"
)

// Pool is generic object pool
type Pool struct {
	pool    sync.Pool
	factory func() interface{}
	count   int64
}

// New creates new pool with factory function
func New(factory func() interface{}) *Pool {
	p := &Pool{
		factory: factory,
	}
	p.pool.New = func() interface{} {
		atomic.AddInt64(&p.count, 1)
		return factory()
	}
	return p
}

// Get gets object from pool
func (p *Pool) Get() interface{} {
	return p.pool.Get()
}

// Put returns object to pool
func (p *Pool) Put(v interface{}) {
	p.pool.Put(v)
}

// Size returns pool size (approximate)
func (p *Pool) Size() int {
	return int(atomic.LoadInt64(&p.count))
}

// PoolWithReset is pool with reset function
type PoolWithReset struct {
	pool    sync.Pool
	factory func() interface{}
	reset   func(interface{})
	count   int64
}

// NewWithReset creates new pool with reset function
func NewWithReset(factory func() interface{}, reset func(interface{})) *PoolWithReset {
	p := &PoolWithReset{
		factory: factory,
		reset:   reset,
	}
	p.pool.New = func() interface{} {
		atomic.AddInt64(&p.count, 1)
		return factory()
	}
	return p
}

// Get gets object from pool
func (p *PoolWithReset) Get() interface{} {
	return p.pool.Get()
}

// Put returns object to pool after resetting it
func (p *PoolWithReset) Put(v interface{}) {
	if v != nil {
		p.reset(v)
	}
	p.pool.Put(v)
}

// Size returns pool size (approximate)
func (p *PoolWithReset) Size() int {
	return int(atomic.LoadInt64(&p.count))
}

// BufferPool is bytes buffer pool
type BufferPool struct {
	pool    sync.Pool
	bufSize int
	count   int64
}

// NewBufferPool creates new buffer pool
func NewBufferPool(bufSize int) *BufferPool {
	if bufSize <= 0 {
		bufSize = 4096 // Default 4KB
	}

	p := &BufferPool{
		bufSize: bufSize,
	}
	p.pool.New = func() interface{} {
		atomic.AddInt64(&p.count, 1)
		buf := make([]byte, bufSize)
		return &buf
	}
	return p
}

// Get gets buffer from pool
func (p *BufferPool) Get() []byte {
	buf := p.pool.Get().(*[]byte)
	return *buf
}

// Put returns buffer to pool
func (p *BufferPool) Put(buf []byte) {
	if buf == nil {
		return
	}

	// Reset buffer length but keep capacity
	buf = buf[:0]
	p.pool.Put(&buf)
}

// Size returns pool size (approximate)
func (p *BufferPool) Size() int {
	return int(atomic.LoadInt64(&p.count))
}

// BufSize returns the buffer size
func (p *BufferPool) BufSize() int {
	return p.bufSize
}

// ResetPool is a pool that resets objects on Put
type ResetPool struct {
	pool    sync.Pool
	factory func() interface{}
	reset   func(interface{})
	count   int64
}

// NewResetPool creates a new reset pool (alias for NewWithReset)
func NewResetPool(factory func() interface{}, reset func(interface{})) *ResetPool {
	p := &ResetPool{
		factory: factory,
		reset:   reset,
	}
	p.pool.New = func() interface{} {
		atomic.AddInt64(&p.count, 1)
		return factory()
	}
	return p
}

// Get gets object from pool
func (p *ResetPool) Get() interface{} {
	return p.pool.Get()
}

// Put returns object to pool after resetting it
func (p *ResetPool) Put(v interface{}) {
	if v != nil {
		p.reset(v)
	}
	p.pool.Put(v)
}

// Size returns pool size (approximate)
func (p *ResetPool) Size() int {
	return int(atomic.LoadInt64(&p.count))
}

// TypedPool is a type-safe generic pool wrapper
type TypedPool struct {
	pool    sync.Pool
	factory func() interface{}
	count   int64
}

// NewTyped creates a new typed pool
func NewTyped(factory func() interface{}) *TypedPool {
	p := &TypedPool{
		factory: factory,
	}
	p.pool.New = func() interface{} {
		atomic.AddInt64(&p.count, 1)
		return factory()
	}
	return p
}

// Get gets object from pool
func (p *TypedPool) Get() interface{} {
	return p.pool.Get()
}

// Put returns object to pool
func (p *TypedPool) Put(v interface{}) {
	p.pool.Put(v)
}

// Size returns pool size (approximate)
func (p *TypedPool) Size() int {
	return int(atomic.LoadInt64(&p.count))
}
