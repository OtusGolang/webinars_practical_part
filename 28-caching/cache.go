package main

import (
	"fmt"
	"sync"
	"time"
)

// CacheMetrics tracks cache statistics
type CacheMetrics struct {
	Hits      int
	Misses    int
	Evictions int
	Size      int
}

// CacheStrategy defines the interface for a caching strategy
type CacheStrategy interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
	Stats() CacheMetrics
	Name() string
}

// ==================== LRU Cache ====================

type LRUCache struct {
	mu        sync.RWMutex
	maxSize   int
	data      map[string]interface{}
	accessLog []string
	metrics   CacheMetrics
}

func NewLRUCache(maxSize int) *LRUCache {
	return &LRUCache{
		maxSize:   maxSize,
		data:      make(map[string]interface{}),
		accessLog: make([]string, 0),
		metrics:   CacheMetrics{},
	}
}

func (c *LRUCache) Name() string {
	return "LRU (Least Recently Used)"
}

func (c *LRUCache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if value, exists := c.data[key]; exists {
		// Move to end (most recent)
		c.accessLog = removeString(c.accessLog, key)
		c.accessLog = append(c.accessLog, key)
		c.metrics.Hits++
		return value, true
	}
	c.metrics.Misses++
	return nil, false
}

func (c *LRUCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// If already exists, remove from log
	if _, exists := c.data[key]; exists {
		c.accessLog = removeString(c.accessLog, key)
	}

	c.data[key] = value
	c.accessLog = append(c.accessLog, key)

	// Evict if necessary
	for len(c.data) > c.maxSize {
		lru := c.accessLog[0]
		c.accessLog = c.accessLog[1:]
		delete(c.data, lru)
		c.metrics.Evictions++
	}

	c.metrics.Size = len(c.data)
}

func (c *LRUCache) Stats() CacheMetrics {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.metrics
}

// ==================== LFU Cache ====================

type LFUCache struct {
	mu        sync.RWMutex
	maxSize   int
	data      map[string]interface{}
	frequency map[string]int
	metrics   CacheMetrics
}

func NewLFUCache(maxSize int) *LFUCache {
	return &LFUCache{
		maxSize:   maxSize,
		data:      make(map[string]interface{}),
		frequency: make(map[string]int),
		metrics:   CacheMetrics{},
	}
}

func (c *LFUCache) Name() string {
	return "LFU (Least Frequently Used)"
}

func (c *LFUCache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if value, exists := c.data[key]; exists {
		c.frequency[key]++
		c.metrics.Hits++
		return value, true
	}
	c.metrics.Misses++
	return nil, false
}

func (c *LFUCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
	if _, exists := c.frequency[key]; !exists {
		c.frequency[key] = 0
	}

	// Evict if necessary
	for len(c.data) > c.maxSize {
		lfu := c.findLFU()
		delete(c.data, lfu)
		delete(c.frequency, lfu)
		c.metrics.Evictions++
	}

	c.metrics.Size = len(c.data)
}

func (c *LFUCache) findLFU() string {
	var minKey string
	var minFreq int = 999999

	for key, freq := range c.frequency {
		if freq < minFreq {
			minFreq = freq
			minKey = key
		}
	}
	return minKey
}

func (c *LFUCache) Stats() CacheMetrics {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.metrics
}

// ==================== FIFO Cache ====================

type FIFOCache struct {
	mu      sync.RWMutex
	maxSize int
	data    map[string]interface{}
	queue   []string
	metrics CacheMetrics
}

func NewFIFOCache(maxSize int) *FIFOCache {
	return &FIFOCache{
		maxSize: maxSize,
		data:    make(map[string]interface{}),
		queue:   make([]string, 0),
		metrics: CacheMetrics{},
	}
}

func (c *FIFOCache) Name() string {
	return "FIFO (First In First Out)"
}

func (c *FIFOCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if value, exists := c.data[key]; exists {
		c.metrics.Hits++
		return value, true
	}
	c.metrics.Misses++
	return nil, false
}

func (c *FIFOCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Skip if already exists (FIFO doesn't update)
	if _, exists := c.data[key]; !exists {
		c.data[key] = value
		c.queue = append(c.queue, key)
	}

	// Evict if necessary
	for len(c.data) > c.maxSize {
		oldest := c.queue[0]
		c.queue = c.queue[1:]
		delete(c.data, oldest)
		c.metrics.Evictions++
	}

	c.metrics.Size = len(c.data)
}

func (c *FIFOCache) Stats() CacheMetrics {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.metrics
}

// ==================== Random Eviction Cache ====================

type RandomCache struct {
	mu      sync.RWMutex
	maxSize int
	data    map[string]interface{}
	metrics CacheMetrics
}

func NewRandomCache(maxSize int) *RandomCache {
	return &RandomCache{
		maxSize: maxSize,
		data:    make(map[string]interface{}),
		metrics: CacheMetrics{},
	}
}

func (c *RandomCache) Name() string {
	return "Random Eviction"
}

func (c *RandomCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if value, exists := c.data[key]; exists {
		c.metrics.Hits++
		return value, true
	}
	c.metrics.Misses++
	return nil, false
}

func (c *RandomCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value

	// Evict if necessary
	for len(c.data) > c.maxSize {
		// Find a random key to evict
		for k := range c.data {
			delete(c.data, k)
			c.metrics.Evictions++
			break
		}
	}

	c.metrics.Size = len(c.data)
}

func (c *RandomCache) Stats() CacheMetrics {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.metrics
}

// ==================== TTL Cache ====================

type TTLCache struct {
	mu        sync.RWMutex
	maxSize   int
	data      map[string]interface{}
	expiry    map[string]time.Time
	ttl       time.Duration
	metrics   CacheMetrics
	cleanupCh chan struct{}
}

func NewTTLCache(maxSize int, ttl time.Duration) *TTLCache {
	c := &TTLCache{
		maxSize:   maxSize,
		data:      make(map[string]interface{}),
		expiry:    make(map[string]time.Time),
		ttl:       ttl,
		metrics:   CacheMetrics{},
		cleanupCh: make(chan struct{}),
	}

	// Start cleanup goroutine
	go c.cleanupExpired()

	return c
}

func (c *TTLCache) Name() string {
	return fmt.Sprintf("TTL (Time To Live, %v)", c.ttl)
}

func (c *TTLCache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if exp, exists := c.expiry[key]; exists && time.Now().Before(exp) {
		c.metrics.Hits++
		return c.data[key], true
	}
	c.metrics.Misses++
	return nil, false
}

func (c *TTLCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
	c.expiry[key] = time.Now().Add(c.ttl)

	// Evict if necessary
	for len(c.data) > c.maxSize {
		var oldestKey string
		var oldestTime time.Time

		for k, exp := range c.expiry {
			if oldestKey == "" || exp.Before(oldestTime) {
				oldestKey = k
				oldestTime = exp
			}
		}

		delete(c.data, oldestKey)
		delete(c.expiry, oldestKey)
		c.metrics.Evictions++
	}

	c.metrics.Size = len(c.data)
}

func (c *TTLCache) Stats() CacheMetrics {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.metrics
}

func (c *TTLCache) cleanupExpired() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for key, exp := range c.expiry {
			if now.After(exp) {
				delete(c.data, key)
				delete(c.expiry, key)
				c.metrics.Evictions++
			}
		}
		c.metrics.Size = len(c.data)
		c.mu.Unlock()
	}
}

// Helper function
func removeString(slice []string, item string) []string {
	for i, v := range slice {
		if v == item {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
