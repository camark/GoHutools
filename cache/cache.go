package cache

import (
	"container/list"
	"sync"
	"time"
)

// Cache is cache interface
type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
	SetWithExpire(key string, value interface{}, expire time.Duration)
	Delete(key string)
	Clear()
	Size() int
	Keys() []string
	Contains(key string) bool
}

// lruItem is cache item for LRU
type lruItem struct {
	key       string
	value     interface{}
	expireAt  time.Time
	hasExpire bool
}

// LRUCache is LRU cache implementation
type LRUCache struct {
	capacity int
	items    map[string]*list.Element
	order    *list.List
	mu       sync.RWMutex
}

// NewLRU creates a new LRU cache with given capacity
func NewLRU(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		order:    list.New(),
	}
}

// Get retrieves an item from the LRU cache
func (c *LRUCache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[key]; ok {
		item := elem.Value.(*lruItem)
		if item.hasExpire && time.Now().After(item.expireAt) {
			c.order.Remove(elem)
			delete(c.items, key)
			return nil, false
		}
		c.order.MoveToFront(elem)
		return item.value, true
	}
	return nil, false
}

// Set adds an item to the LRU cache
func (c *LRUCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[key]; ok {
		c.order.MoveToFront(elem)
		item := elem.Value.(*lruItem)
		item.value = value
		return
	}

	if c.order.Len() >= c.capacity {
		c.removeOldest()
	}

	item := &lruItem{key: key, value: value}
	elem := c.order.PushFront(item)
	c.items[key] = elem
}

// SetWithExpire adds an item with expiration to the LRU cache
func (c *LRUCache) SetWithExpire(key string, value interface{}, expire time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[key]; ok {
		c.order.MoveToFront(elem)
		item := elem.Value.(*lruItem)
		item.value = value
		item.expireAt = time.Now().Add(expire)
		item.hasExpire = true
		return
	}

	if c.order.Len() >= c.capacity {
		c.removeOldest()
	}

	item := &lruItem{
		key:       key,
		value:     value,
		expireAt:  time.Now().Add(expire),
		hasExpire: true,
	}
	elem := c.order.PushFront(item)
	c.items[key] = elem
}

// Delete removes an item from the LRU cache
func (c *LRUCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[key]; ok {
		c.order.Remove(elem)
		delete(c.items, key)
	}
}

// Clear removes all items from the LRU cache
func (c *LRUCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[string]*list.Element)
	c.order.Init()
}

// Size returns the number of items in the LRU cache
func (c *LRUCache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.order.Len()
}

// Keys returns all keys in the LRU cache
func (c *LRUCache) Keys() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	keys := make([]string, 0, c.order.Len())
	for elem := c.order.Front(); elem != nil; elem = elem.Next() {
		item := elem.Value.(*lruItem)
		if !item.hasExpire || !time.Now().After(item.expireAt) {
			keys = append(keys, item.key)
		}
	}
	return keys
}

// Contains checks if a key exists in the LRU cache
func (c *LRUCache) Contains(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	elem, ok := c.items[key]
	if !ok {
		return false
	}
	item := elem.Value.(*lruItem)
	if item.hasExpire && time.Now().After(item.expireAt) {
		c.order.Remove(elem)
		delete(c.items, key)
		return false
	}
	return true
}

func (c *LRUCache) removeOldest() {
	elem := c.order.Back()
	if elem != nil {
		c.order.Remove(elem)
		item := elem.Value.(*lruItem)
		delete(c.items, item.key)
	}
}

// fifoItem is cache item for FIFO
type fifoItem struct {
	key       string
	value     interface{}
	expireAt  time.Time
	hasExpire bool
}

// FIFOCache is FIFO cache implementation
type FIFOCache struct {
	capacity int
	items    map[string]*list.Element
	order    *list.List
	mu       sync.RWMutex
}

// NewFIFO creates a new FIFO cache with given capacity
func NewFIFO(capacity int) *FIFOCache {
	return &FIFOCache{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		order:    list.New(),
	}
}

// Get retrieves an item from the FIFO cache
func (c *FIFOCache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[key]; ok {
		item := elem.Value.(*fifoItem)
		if item.hasExpire && time.Now().After(item.expireAt) {
			c.order.Remove(elem)
			delete(c.items, key)
			return nil, false
		}
		return item.value, true
	}
	return nil, false
}

// Set adds an item to the FIFO cache
func (c *FIFOCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[key]; ok {
		item := elem.Value.(*fifoItem)
		item.value = value
		return
	}

	if c.order.Len() >= c.capacity {
		c.removeOldest()
	}

	item := &fifoItem{key: key, value: value}
	elem := c.order.PushBack(item)
	c.items[key] = elem
}

// SetWithExpire adds an item with expiration to the FIFO cache
func (c *FIFOCache) SetWithExpire(key string, value interface{}, expire time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[key]; ok {
		item := elem.Value.(*fifoItem)
		item.value = value
		item.expireAt = time.Now().Add(expire)
		item.hasExpire = true
		return
	}

	if c.order.Len() >= c.capacity {
		c.removeOldest()
	}

	item := &fifoItem{
		key:       key,
		value:     value,
		expireAt:  time.Now().Add(expire),
		hasExpire: true,
	}
	elem := c.order.PushBack(item)
	c.items[key] = elem
}

// Delete removes an item from the FIFO cache
func (c *FIFOCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[key]; ok {
		c.order.Remove(elem)
		delete(c.items, key)
	}
}

// Clear removes all items from the FIFO cache
func (c *FIFOCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[string]*list.Element)
	c.order.Init()
}

// Size returns the number of items in the FIFO cache
func (c *FIFOCache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.order.Len()
}

// Keys returns all keys in the FIFO cache
func (c *FIFOCache) Keys() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	keys := make([]string, 0, c.order.Len())
	for elem := c.order.Front(); elem != nil; elem = elem.Next() {
		item := elem.Value.(*fifoItem)
		if !item.hasExpire || !time.Now().After(item.expireAt) {
			keys = append(keys, item.key)
		}
	}
	return keys
}

// Contains checks if a key exists in the FIFO cache
func (c *FIFOCache) Contains(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	elem, ok := c.items[key]
	if !ok {
		return false
	}
	item := elem.Value.(*fifoItem)
	if item.hasExpire && time.Now().After(item.expireAt) {
		c.order.Remove(elem)
		delete(c.items, key)
		return false
	}
	return true
}

func (c *FIFOCache) removeOldest() {
	elem := c.order.Front()
	if elem != nil {
		c.order.Remove(elem)
		item := elem.Value.(*fifoItem)
		delete(c.items, item.key)
	}
}

// lfuItem is cache item for LFU
type lfuItem struct {
	key       string
	value     interface{}
	freq      int
	expireAt  time.Time
	hasExpire bool
	element   *list.Element
}

// LFUCache is LFU cache implementation
type LFUCache struct {
	capacity int
	items    map[string]*lfuItem
	freq     map[int]*list.List
	minFreq  int
	mu       sync.RWMutex
}

// NewLFU creates a new LFU cache with given capacity
func NewLFU(capacity int) *LFUCache {
	return &LFUCache{
		capacity: capacity,
		items:    make(map[string]*lfuItem),
		freq:     make(map[int]*list.List),
		minFreq:  0,
	}
}

// Get retrieves an item from the LFU cache
func (c *LFUCache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, ok := c.items[key]
	if !ok {
		return nil, false
	}

	if item.hasExpire && time.Now().After(item.expireAt) {
		c.removeItem(item)
		return nil, false
	}

	c.incrementFreq(item)
	return item.value, true
}

// Set adds an item to the LFU cache
func (c *LFUCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if item, ok := c.items[key]; ok {
		item.value = value
		c.incrementFreq(item)
		return
	}

	if len(c.items) >= c.capacity {
		c.evictLFU()
	}

	item := &lfuItem{
		key:     key,
		value:   value,
		freq:    1,
		element: nil,
	}
	c.addToFreq(item)
	c.minFreq = 1
	c.items[key] = item
}

// SetWithExpire adds an item with expiration to the LFU cache
func (c *LFUCache) SetWithExpire(key string, value interface{}, expire time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if item, ok := c.items[key]; ok {
		item.value = value
		item.expireAt = time.Now().Add(expire)
		item.hasExpire = true
		c.incrementFreq(item)
		return
	}

	if len(c.items) >= c.capacity {
		c.evictLFU()
	}

	item := &lfuItem{
		key:       key,
		value:     value,
		freq:      1,
		expireAt:  time.Now().Add(expire),
		hasExpire: true,
		element:   nil,
	}
	c.addToFreq(item)
	c.minFreq = 1
	c.items[key] = item
}

// Delete removes an item from the LFU cache
func (c *LFUCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if item, ok := c.items[key]; ok {
		c.removeItem(item)
	}
}

// Clear removes all items from the LFU cache
func (c *LFUCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[string]*lfuItem)
	c.freq = make(map[int]*list.List)
	c.minFreq = 0
}

// Size returns the number of items in the LFU cache
func (c *LFUCache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.items)
}

// Keys returns all keys in the LFU cache
func (c *LFUCache) Keys() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	keys := make([]string, 0, len(c.items))
	now := time.Now()
	for key, item := range c.items {
		if !item.hasExpire || !now.After(item.expireAt) {
			keys = append(keys, key)
		}
	}
	return keys
}

// Contains checks if a key exists in the LFU cache
func (c *LFUCache) Contains(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, ok := c.items[key]
	if !ok {
		return false
	}
	if item.hasExpire && time.Now().After(item.expireAt) {
		c.removeItem(item)
		return false
	}
	return true
}

func (c *LFUCache) incrementFreq(item *lfuItem) {
	oldFreq := item.freq
	newFreq := oldFreq + 1

	// Remove from old frequency list
	if l, ok := c.freq[oldFreq]; ok {
		l.Remove(item.element)
		if l.Len() == 0 {
			delete(c.freq, oldFreq)
			if c.minFreq == oldFreq {
				c.minFreq = newFreq
			}
		}
	}

	item.freq = newFreq
	c.addToFreq(item)
}

func (c *LFUCache) addToFreq(item *lfuItem) {
	freq := item.freq
	l, ok := c.freq[freq]
	if !ok {
		l = list.New()
		c.freq[freq] = l
	}
	item.element = l.PushBack(item)
}

func (c *LFUCache) removeItem(item *lfuItem) {
	if l, ok := c.freq[item.freq]; ok {
		l.Remove(item.element)
		if l.Len() == 0 {
			delete(c.freq, item.freq)
			if c.minFreq == item.freq {
				c.minFreq = 0
			}
		}
	}
	delete(c.items, item.key)
}

func (c *LFUCache) evictLFU() {
	l, ok := c.freq[c.minFreq]
	if !ok || l.Len() == 0 {
		return
	}

	elem := l.Front()
	if elem != nil {
		item := elem.Value.(*lfuItem)
		l.Remove(elem)
		if l.Len() == 0 {
			delete(c.freq, c.minFreq)
		}
		delete(c.items, item.key)
	}
}

// timedItem is cache item for TimedCache
type timedItem struct {
	value    interface{}
	expireAt time.Time
}

// TimedCache is time-based cache
type TimedCache struct {
	defaultExpire time.Duration
	items         map[string]*timedItem
	mu            sync.RWMutex
	stopCleanup   chan struct{}
}

// NewTimed creates a new TimedCache with given default expiration
func NewTimed(defaultExpire time.Duration) *TimedCache {
	c := &TimedCache{
		defaultExpire: defaultExpire,
		items:         make(map[string]*timedItem),
		stopCleanup:   make(chan struct{}),
	}
	go c.cleanup()
	return c
}

// Close stops the background cleanup goroutine and clears all items.
// The cache must not be used after Close is called.
func (c *TimedCache) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()

	select {
	case <-c.stopCleanup:
		// already closed
	default:
		close(c.stopCleanup)
	}
	c.items = make(map[string]*timedItem)
}

// Get retrieves an item from the TimedCache
func (c *TimedCache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, ok := c.items[key]
	if !ok {
		return nil, false
	}

	if time.Now().After(item.expireAt) {
		delete(c.items, key)
		return nil, false
	}

	return item.value, true
}

// Set adds an item with default expiration to the TimedCache
func (c *TimedCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = &timedItem{
		value:    value,
		expireAt: time.Now().Add(c.defaultExpire),
	}
}

// SetWithExpire adds an item with custom expiration to the TimedCache
func (c *TimedCache) SetWithExpire(key string, value interface{}, expire time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = &timedItem{
		value:    value,
		expireAt: time.Now().Add(expire),
	}
}

// Delete removes an item from the TimedCache
func (c *TimedCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}

// Clear removes all items from the TimedCache
func (c *TimedCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[string]*timedItem)
}

// Size returns the number of items in the TimedCache
func (c *TimedCache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.items)
}

// Keys returns all keys in the TimedCache
func (c *TimedCache) Keys() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	keys := make([]string, 0, len(c.items))
	now := time.Now()
	for key, item := range c.items {
		if !now.After(item.expireAt) {
			keys = append(keys, key)
		}
	}
	return keys
}

// Contains checks if a key exists in the TimedCache
func (c *TimedCache) Contains(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, ok := c.items[key]
	if !ok {
		return false
	}

	if time.Now().After(item.expireAt) {
		delete(c.items, key)
		return false
	}

	return true
}

// cleanup periodically removes expired items
func (c *TimedCache) cleanup() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			now := time.Now()
			for key, item := range c.items {
				if now.After(item.expireAt) {
					delete(c.items, key)
				}
			}
			c.mu.Unlock()
		case <-c.stopCleanup:
			return
		}
	}
}
