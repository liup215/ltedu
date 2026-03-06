// Package cache provides a generic, thread-safe LRU cache with optional TTL.
//
// Usage:
//
//	c := cache.New[string, *MyType](128, 5*time.Minute)
//	c.Set("key", value)
//	if v, ok := c.Get("key"); ok { ... }
//	c.Delete("key")
package cache

import (
	"container/list"
	"sync"
	"time"
)

// entry holds a cached key/value pair and an optional expiry time.
type entry[K comparable, V any] struct {
	key     K
	value   V
	expiry  time.Time // zero value means no expiry
	element *list.Element
}

// LRU is a thread-safe, fixed-capacity Least-Recently-Used cache with optional TTL.
// When the cache is full the least-recently-used item is evicted to make room.
// A TTL of 0 disables expiry.
type LRU[K comparable, V any] struct {
	mu       sync.Mutex
	capacity int
	ttl      time.Duration
	ll       *list.List
	items    map[K]*entry[K, V]
}

// New creates a new LRU cache with the given capacity and TTL.
// capacity must be > 0.
// ttl == 0 means entries never expire.
func New[K comparable, V any](capacity int, ttl time.Duration) *LRU[K, V] {
	if capacity <= 0 {
		capacity = 1
	}
	return &LRU[K, V]{
		capacity: capacity,
		ttl:      ttl,
		ll:       list.New(),
		items:    make(map[K]*entry[K, V], capacity),
	}
}

// Set inserts or updates the value associated with key.
func (c *LRU[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var expiry time.Time
	if c.ttl > 0 {
		expiry = time.Now().Add(c.ttl)
	}

	if e, ok := c.items[key]; ok {
		e.value = value
		e.expiry = expiry
		c.ll.MoveToFront(e.element)
		return
	}

	// Evict least-recently-used entry if at capacity.
	if c.ll.Len() >= c.capacity {
		c.evictOldest()
	}

	e := &entry[K, V]{key: key, value: value, expiry: expiry}
	e.element = c.ll.PushFront(e)
	c.items[key] = e
}

// Get retrieves the value for key. Returns (zero, false) if the key is absent
// or the cached entry has expired.
func (c *LRU[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	e, ok := c.items[key]
	if !ok {
		var zero V
		return zero, false
	}

	if !e.expiry.IsZero() && time.Now().After(e.expiry) {
		c.removeEntry(e)
		var zero V
		return zero, false
	}

	c.ll.MoveToFront(e.element)
	return e.value, true
}

// Delete removes the entry for key, if present.
func (c *LRU[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if e, ok := c.items[key]; ok {
		c.removeEntry(e)
	}
}

// Purge removes all entries from the cache.
func (c *LRU[K, V]) Purge() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.ll.Init()
	c.items = make(map[K]*entry[K, V], c.capacity)
}

// Len returns the number of entries currently in the cache (including expired
// entries that have not yet been evicted).
func (c *LRU[K, V]) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ll.Len()
}

// evictOldest removes the least-recently-used entry. Must be called with c.mu held.
func (c *LRU[K, V]) evictOldest() {
	if back := c.ll.Back(); back != nil {
		c.removeEntry(back.Value.(*entry[K, V]))
	}
}

// removeEntry removes a specific entry. Must be called with c.mu held.
func (c *LRU[K, V]) removeEntry(e *entry[K, V]) {
	c.ll.Remove(e.element)
	delete(c.items, e.key)
}
