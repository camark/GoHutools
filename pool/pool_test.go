package pool

import (
	"sync"
	"testing"
)

func TestPoolNew(t *testing.T) {
	factory := func() interface{} {
		return make([]byte, 1024)
	}

	p := New(factory)
	if p == nil {
		t.Fatal("New() returned nil")
	}
}

func TestPoolGetPut(t *testing.T) {
	factory := func() interface{} {
		return make([]byte, 1024)
	}

	p := New(factory)

	// Get an object
	obj := p.Get()
	if obj == nil {
		t.Fatal("Get() returned nil")
	}

	buf, ok := obj.([]byte)
	if !ok {
		t.Fatal("Expected []byte")
	}
	if len(buf) != 1024 {
		t.Errorf("Expected buffer length 1024, got %d", len(buf))
	}

	// Put it back
	p.Put(obj)

	// Get again - should reuse
	obj2 := p.Get()
	if obj2 == nil {
		t.Fatal("Get() returned nil after Put")
	}
}

func TestPoolSize(t *testing.T) {
	factory := func() interface{} {
		return new(int)
	}

	p := New(factory)

	if p.Size() != 0 {
		t.Errorf("Expected initial size 0, got %d", p.Size())
	}

	// Get some objects to trigger creation
	obj1 := p.Get()
	obj2 := p.Get()
	obj3 := p.Get()

	if p.Size() != 3 {
		t.Errorf("Expected size 3, got %d", p.Size())
	}

	p.Put(obj1)
	p.Put(obj2)
	p.Put(obj3)

	// Size should still be 3 (objects are created, not destroyed)
	if p.Size() != 3 {
		t.Errorf("Expected size 3 after put, got %d", p.Size())
	}
}

func TestPoolConcurrent(t *testing.T) {
	factory := func() interface{} {
		return new(int)
	}

	p := New(factory)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			obj := p.Get()
			if obj == nil {
				t.Error("Get() returned nil")
				return
			}
			p.Put(obj)
		}()
	}

	wg.Wait()
}

func TestPoolWithResetNew(t *testing.T) {
	factory := func() interface{} {
		return &struct{ value int }{value: 42}
	}
	reset := func(obj interface{}) {
		obj.(*struct{ value int }).value = 0
	}

	p := NewWithReset(factory, reset)
	if p == nil {
		t.Fatal("NewWithReset() returned nil")
	}
}

func TestPoolWithResetGetPut(t *testing.T) {
	type testObj struct {
		value int
		data  []byte
	}

	factory := func() interface{} {
		return &testObj{value: 42, data: make([]byte, 100)}
	}
	reset := func(obj interface{}) {
		o := obj.(*testObj)
		o.value = 0
		o.data = o.data[:0]
	}

	p := NewWithReset(factory, reset)

	// Get an object
	obj := p.Get().(*testObj)
	if obj.value != 42 {
		t.Errorf("Expected initial value 42, got %d", obj.value)
	}

	// Modify it
	obj.value = 100
	obj.data = append(obj.data, 1, 2, 3)

	// Put resets the object in place before pooling it
	p.Put(obj)
	if obj.value != 0 {
		t.Errorf("Expected value reset to 0 after Put, got %d", obj.value)
	}
	if len(obj.data) != 0 {
		t.Errorf("Expected data reset to length 0 after Put, got %d", len(obj.data))
	}

	// Get returns a ready-to-use object. sync.Pool does not guarantee that
	// the just-Put item is recycled (items may be dropped at any GC, and
	// the race detector drops Puts randomly), so accept either the
	// recycled reset object or a fresh factory one.
	obj2 := p.Get().(*testObj)
	if obj2.value != 0 && obj2.value != 42 {
		t.Errorf("Expected reset (0) or fresh (42) value, got %d", obj2.value)
	}
	if len(obj2.data) != 0 && len(obj2.data) != 100 {
		t.Errorf("Expected reset (0) or fresh (100) data length, got %d", len(obj2.data))
	}
}

func TestPoolWithResetSize(t *testing.T) {
	factory := func() interface{} {
		return new(int)
	}
	reset := func(obj interface{}) {
		*(obj.(*int)) = 0
	}

	p := NewWithReset(factory, reset)

	if p.Size() != 0 {
		t.Errorf("Expected initial size 0, got %d", p.Size())
	}

	obj1 := p.Get()
	obj2 := p.Get()

	if p.Size() != 2 {
		t.Errorf("Expected size 2, got %d", p.Size())
	}

	p.Put(obj1)
	p.Put(obj2)

	if p.Size() != 2 {
		t.Errorf("Expected size 2 after put, got %d", p.Size())
	}
}

func TestPoolWithResetConcurrent(t *testing.T) {
	type testObj struct {
		value int
	}

	factory := func() interface{} {
		return &testObj{value: 1}
	}
	reset := func(obj interface{}) {
		obj.(*testObj).value = 0
	}

	p := NewWithReset(factory, reset)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			obj := p.Get()
			if obj == nil {
				t.Error("Get() returned nil")
				return
			}
			o := obj.(*testObj)
			o.value = 42
			p.Put(obj)
		}()
	}

	wg.Wait()
}

func TestBufferPoolNew(t *testing.T) {
	p := NewBufferPool(4096)
	if p == nil {
		t.Fatal("NewBufferPool() returned nil")
	}
	if p.BufSize() != 4096 {
		t.Errorf("Expected bufSize 4096, got %d", p.BufSize())
	}
}

func TestBufferPoolDefaultSize(t *testing.T) {
	p := NewBufferPool(0)
	if p == nil {
		t.Fatal("NewBufferPool() returned nil")
	}
	if p.BufSize() != 4096 {
		t.Errorf("Expected default bufSize 4096, got %d", p.BufSize())
	}
}

func TestBufferPoolNegativeSize(t *testing.T) {
	p := NewBufferPool(-1)
	if p == nil {
		t.Fatal("NewBufferPool() returned nil")
	}
	if p.BufSize() != 4096 {
		t.Errorf("Expected default bufSize 4096 for negative input, got %d", p.BufSize())
	}
}

func TestBufferPoolGetPut(t *testing.T) {
	p := NewBufferPool(1024)

	// Get a buffer
	buf := p.Get()
	if buf == nil {
		t.Fatal("Get() returned nil")
	}
	if len(buf) != 1024 {
		t.Errorf("Expected buffer length 1024, got %d", len(buf))
	}
	if cap(buf) < 1024 {
		t.Errorf("Expected buffer capacity >= 1024, got %d", cap(buf))
	}

	// Use the buffer
	for i := range buf {
		buf[i] = byte(i % 256)
	}

	// Put it back (stored truncated to length 0)
	p.Put(buf)

	// Get again. sync.Pool does not guarantee that the just-Put buffer is
	// recycled (items may be dropped at any GC, and the race detector drops
	// Puts randomly), so accept either the recycled truncated buffer or a
	// fresh full-size one.
	buf2 := p.Get()
	if len(buf2) != 0 && len(buf2) != 1024 {
		t.Errorf("Expected recycled (length 0) or fresh (length 1024) buffer, got %d", len(buf2))
	}
}

func TestBufferPoolNilPut(t *testing.T) {
	p := NewBufferPool(1024)

	// Should not panic
	p.Put(nil)
}

func TestBufferPoolSize(t *testing.T) {
	p := NewBufferPool(1024)

	if p.Size() != 0 {
		t.Errorf("Expected initial size 0, got %d", p.Size())
	}

	buf1 := p.Get()
	buf2 := p.Get()
	buf3 := p.Get()

	if p.Size() != 3 {
		t.Errorf("Expected size 3, got %d", p.Size())
	}

	p.Put(buf1)
	p.Put(buf2)
	p.Put(buf3)

	if p.Size() != 3 {
		t.Errorf("Expected size 3 after put, got %d", p.Size())
	}
}

func TestBufferPoolConcurrent(t *testing.T) {
	p := NewBufferPool(1024)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			buf := p.Get()
			if buf == nil {
				t.Error("Get() returned nil")
				return
			}

			// Use the buffer
			for j := 0; j < len(buf); j++ {
				buf[j] = byte(j % 256)
			}

			p.Put(buf)
		}()
	}

	wg.Wait()
}

func TestResetPoolNew(t *testing.T) {
	factory := func() interface{} {
		return new(int)
	}
	reset := func(obj interface{}) {
		*(obj.(*int)) = 0
	}

	p := NewResetPool(factory, reset)
	if p == nil {
		t.Fatal("NewResetPool() returned nil")
	}
}

func TestResetPoolGetPut(t *testing.T) {
	type testObj struct {
		value int
		name  string
	}

	factory := func() interface{} {
		return &testObj{value: 1, name: "default"}
	}
	reset := func(obj interface{}) {
		o := obj.(*testObj)
		o.value = 0
		o.name = ""
	}

	p := NewResetPool(factory, reset)

	// Get an object
	obj := p.Get().(*testObj)
	if obj.value != 1 {
		t.Errorf("Expected initial value 1, got %d", obj.value)
	}
	if obj.name != "default" {
		t.Errorf("Expected initial name 'default', got %s", obj.name)
	}

	// Modify it
	obj.value = 42
	obj.name = "modified"

	// Put resets the object in place before pooling it
	p.Put(obj)
	if obj.value != 0 {
		t.Errorf("Expected value reset to 0 after Put, got %d", obj.value)
	}
	if obj.name != "" {
		t.Errorf("Expected name reset to '' after Put, got %s", obj.name)
	}

	// Get returns a ready-to-use object: either the recycled reset one or
	// a fresh factory one (sync.Pool gives no recycling guarantee).
	obj2 := p.Get().(*testObj)
	if obj2.value != 0 && obj2.value != 1 {
		t.Errorf("Expected reset (0) or fresh (1) value, got %d", obj2.value)
	}
	if obj2.name != "" && obj2.name != "default" {
		t.Errorf("Expected reset ('') or fresh ('default') name, got %s", obj2.name)
	}
}

func TestResetPoolSize(t *testing.T) {
	factory := func() interface{} {
		return new(int)
	}
	reset := func(obj interface{}) {
		*(obj.(*int)) = 0
	}

	p := NewResetPool(factory, reset)

	if p.Size() != 0 {
		t.Errorf("Expected initial size 0, got %d", p.Size())
	}

	obj1 := p.Get()
	obj2 := p.Get()

	if p.Size() != 2 {
		t.Errorf("Expected size 2, got %d", p.Size())
	}

	p.Put(obj1)
	p.Put(obj2)

	if p.Size() != 2 {
		t.Errorf("Expected size 2 after put, got %d", p.Size())
	}
}

func TestResetPoolConcurrent(t *testing.T) {
	type testObj struct {
		value int
	}

	factory := func() interface{} {
		return &testObj{value: 1}
	}
	reset := func(obj interface{}) {
		obj.(*testObj).value = 0
	}

	p := NewResetPool(factory, reset)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			obj := p.Get()
			if obj == nil {
				t.Error("Get() returned nil")
				return
			}
			o := obj.(*testObj)
			o.value = 42
			p.Put(obj)
		}()
	}

	wg.Wait()
}

func TestTypedPoolNew(t *testing.T) {
	factory := func() interface{} {
		return new(int)
	}

	p := NewTyped(factory)
	if p == nil {
		t.Fatal("NewTyped() returned nil")
	}
}

func TestTypedPoolGetPut(t *testing.T) {
	factory := func() interface{} {
		return new(int)
	}

	p := NewTyped(factory)

	// Get an object
	obj := p.Get()
	if obj == nil {
		t.Fatal("Get() returned nil")
	}

	ptr, ok := obj.(*int)
	if !ok {
		t.Fatal("Expected *int")
	}
	*ptr = 42

	// Put it back
	p.Put(obj)

	// Get again - should reuse
	obj2 := p.Get()
	if obj2 == nil {
		t.Fatal("Get() returned nil after Put")
	}
}

func TestTypedPoolSize(t *testing.T) {
	factory := func() interface{} {
		return new(int)
	}

	p := NewTyped(factory)

	if p.Size() != 0 {
		t.Errorf("Expected initial size 0, got %d", p.Size())
	}

	obj1 := p.Get()
	obj2 := p.Get()
	obj3 := p.Get()

	if p.Size() != 3 {
		t.Errorf("Expected size 3, got %d", p.Size())
	}

	p.Put(obj1)
	p.Put(obj2)
	p.Put(obj3)

	if p.Size() != 3 {
		t.Errorf("Expected size 3 after put, got %d", p.Size())
	}
}

func TestTypedPoolConcurrent(t *testing.T) {
	factory := func() interface{} {
		return new(int)
	}

	p := NewTyped(factory)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			obj := p.Get()
			if obj == nil {
				t.Error("Get() returned nil")
				return
			}
			p.Put(obj)
		}()
	}

	wg.Wait()
}

// Benchmark tests
func BenchmarkPoolGetPut(b *testing.B) {
	factory := func() interface{} {
		return make([]byte, 1024)
	}

	p := New(factory)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		obj := p.Get()
		p.Put(obj)
	}
}

func BenchmarkPoolGetPutParallel(b *testing.B) {
	factory := func() interface{} {
		return make([]byte, 1024)
	}

	p := New(factory)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			obj := p.Get()
			p.Put(obj)
		}
	})
}

func BenchmarkPoolWithResetGetPut(b *testing.B) {
	type testObj struct {
		data [1024]byte
	}

	factory := func() interface{} {
		return &testObj{}
	}
	reset := func(obj interface{}) {
		o := obj.(*testObj)
		for i := range o.data {
			o.data[i] = 0
		}
	}

	p := NewWithReset(factory, reset)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		obj := p.Get()
		p.Put(obj)
	}
}

func BenchmarkBufferPoolGetPut(b *testing.B) {
	p := NewBufferPool(4096)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := p.Get()
		p.Put(buf)
	}
}

func BenchmarkBufferPoolGetPutParallel(b *testing.B) {
	p := NewBufferPool(4096)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buf := p.Get()
			p.Put(buf)
		}
	})
}

func BenchmarkDirectAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := make([]byte, 4096)
		_ = buf
	}
}
