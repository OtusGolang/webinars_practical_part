// EXAMPLES OF HOW TO MODIFY AND EXTEND THIS PROJECT

package main

// ==================== EXAMPLE 1: Run Single Cache Strategy ====================
// In main.go, replace the loop with:
/*
func exampleSingleCache() {
	cache := NewLRUCache(20)
	gen := NewRequestGenerator(ProfileHotSpot, 100)
	api := NewAPIService(cache)

	for i := 0; i < 1000; i++ {
		req := gen.Next()
		api.GetUser(req.ID)
	}

	stats := cache.Stats()
	fmt.Printf("LRU with Hot Spot Pattern:\n")
	fmt.Printf("  Hit Rate: %.1f%%\n", float64(stats.Hits)/(float64(stats.Hits+stats.Misses))*100)
	fmt.Printf("  Evictions: %d\n", stats.Evictions)
}
*/

// ==================== EXAMPLE 2: Add a New Cache Strategy ====================
// Copy and modify the LRU implementation:
/*
type ClockCache struct {
	mu        sync.RWMutex
	maxSize   int
	data      map[string]interface{}
	hand      int
	items     []string
	metrics   CacheMetrics
}

func NewClockCache(maxSize int) *ClockCache {
	return &ClockCache{
		maxSize: maxSize,
		data:    make(map[string]interface{}),
		items:   make([]string, 0),
		metrics: CacheMetrics{},
	}
}

func (c *ClockCache) Name() string {
	return "Clock Cache (Page Replacement)"
}

func (c *ClockCache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if value, exists := c.data[key]; exists {
		c.metrics.Hits++
		// Mark as referenced (in real impl, set reference bit)
		return value, true
	}
	c.metrics.Misses++
	return nil, false
}

func (c *ClockCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.data) >= c.maxSize {
		// Evict using clock hand
		evicted := c.items[c.hand]
		delete(c.data, evicted)
		c.items = append(c.items[:c.hand], c.items[c.hand+1:]...)
		c.metrics.Evictions++
	}

	c.data[key] = value
	c.items = append(c.items, key)
	c.metrics.Size = len(c.data)
}

func (c *ClockCache) Stats() CacheMetrics {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.metrics
}
*/

// ==================== EXAMPLE 3: Add a New Request Pattern ====================
// In generator.go, add to the switch statement in Next():
/*
case "gaussian":
	// Bell curve around center
	center := g.maxID / 2
	deviation := g.maxID / 6
	id := fmt.Sprintf("user_%d", center + int(rand.NormFloat64()*float64(deviation)))
	return RequestProfile{ID: id, Count: 1}

case "temporal":
	// Access pattern changes over time
	phase := (g.sequential / 100) % 3
	if phase == 0 {
		// Phase 1: Hot spot 1 (keys 0-20)
		id := fmt.Sprintf("user_%d", rand.Intn(20))
	} else if phase == 1 {
		// Phase 2: Hot spot 2 (keys 50-70)
		id := fmt.Sprintf("user_%d", 50+rand.Intn(20))
	} else {
		// Phase 3: Random
		id := fmt.Sprintf("user_%d", rand.Intn(g.maxID))
	}
	g.sequential++
	return RequestProfile{ID: id, Count: 1}

case "burst":
	// Burst pattern: many requests to same key
	if rand.Float64() < 0.1 { // 10% chance of burst
		id := fmt.Sprintf("user_%d", rand.Intn(g.maxID))
		return RequestProfile{ID: id, Count: 10} // Simulate 10 requests
	}
	id := fmt.Sprintf("user_%d", rand.Intn(g.maxID))
	return RequestProfile{ID: id, Count: 1}
*/

// ==================== EXAMPLE 4: Add Concurrency Testing ====================
// Add to main.go:
/*
func testConcurrency() {
	cache := NewLRUCache(20)
	api := NewAPIService(cache)

	const numGoroutines = 10
	const requestsPerGoroutine = 100

	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			gen := NewRequestGenerator(ProfileRandom, 100)
			for j := 0; j < requestsPerGoroutine; j++ {
				req := gen.Next()
				api.GetUser(req.ID)
			}
		}(i)
	}

	wg.Wait()
	elapsed := time.Since(start)

	stats := cache.Stats()
	fmt.Printf("Concurrent Test (10 goroutines × 100 requests):\n")
	fmt.Printf("  Total Time: %v\n", elapsed)
	fmt.Printf("  Hit Rate: %.1f%%\n", float64(stats.Hits)/(float64(stats.Hits+stats.Misses))*100)
	fmt.Printf("  Throughput: %.0f ops/sec\n", float64(numGoroutines*requestsPerGoroutine)/elapsed.Seconds())
}
*/

// ==================== EXAMPLE 5: Add Visualization ====================
// Create a visualization function:
/*
func visualizeCache(cache CacheStrategy, generator *RequestGenerator, numRequests int) {
	bars := make(map[string]int)

	for i := 0; i < numRequests; i++ {
		req := generator.Next()
		api.GetUser(req.ID)

		// Simulate counting for histogram
		bars[req.ID]++
	}

	// Print histogram
	for key, count := range bars {
		bar := strings.Repeat("█", count/5) // Scale factor
		fmt.Printf("%s: %s (%d)\n", key, bar, count)
	}
}
*/

// ==================== EXAMPLE 6: Benchmark Different Cache Sizes ====================
// Add to main.go:
/*
func benchmarkCacheSizes() {
	sizes := []int{5, 10, 20, 50, 100}

	fmt.Println("Cache Size vs Hit Rate:")
	fmt.Println()

	for _, size := range sizes {
		cache := NewLRUCache(size)
		gen := NewRequestGenerator(ProfileHotSpot, 100)
		api := NewAPIService(cache)

		for i := 0; i < 1000; i++ {
			req := gen.Next()
			api.GetUser(req.ID)
		}

		stats := cache.Stats()
		hitRate := float64(stats.Hits) / float64(stats.Hits+stats.Misses) * 100
		fmt.Printf("Size %3d: Hit Rate %5.1f%% │ Evictions %4d\n",
			size, hitRate, stats.Evictions)
	}
}
*/

// ==================== EXAMPLE 7: Add TTL Simulation ====================
// Modify the demo to simulate time-based expiration:
/*
func simulateTimeBasedWorkload() {
	cache := NewTTLCache(20, 100*time.Millisecond)
	gen := NewRequestGenerator(ProfileRandom, 100)
	api := NewAPIService(cache)

	for i := 0; i < 500; i++ {
		if i == 250 {
			// Wait for TTL to expire
			time.Sleep(150 * time.Millisecond)
			fmt.Println("TTL expired, cache cleared...")
		}

		req := gen.Next()
		api.GetUser(req.ID)
	}

	stats := cache.Stats()
	fmt.Printf("Time-based expiration test completed\n")
	fmt.Printf("  Hit Rate: %.1f%%\n", float64(stats.Hits)/(float64(stats.Hits+stats.Misses))*100)
}
*/

// ==================== EXAMPLE 8: Hybrid LRU + TTL ====================
// Combine multiple strategies:
/*
type HybridCache struct {
	lru *LRUCache
	ttl *TTLCache
	mu  sync.RWMutex
}

func NewHybridCache(maxSize int, ttl time.Duration) *HybridCache {
	return &HybridCache{
		lru: NewLRUCache(maxSize),
		ttl: NewTTLCache(maxSize/2, ttl),
	}
}

func (c *HybridCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// Check LRU first
	if value, found := c.lru.Get(key); found {
		return value, true
	}

	// Fallback to TTL cache
	return c.ttl.Get(key)
}

func (c *HybridCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.lru.Set(key, value)
	c.ttl.Set(key, value)
}

func (c *HybridCache) Stats() CacheMetrics {
	c.mu.RLock()
	defer c.mu.RUnlock()

	lruStats := c.lru.Stats()
	ttlStats := c.ttl.Stats()

	return CacheMetrics{
		Hits:      lruStats.Hits + ttlStats.Hits,
		Misses:    lruStats.Misses + ttlStats.Misses,
		Evictions: lruStats.Evictions + ttlStats.Evictions,
		Size:      lruStats.Size + ttlStats.Size,
	}
}

func (c *HybridCache) Name() string {
	return "Hybrid LRU + TTL"
}
*/

// ==================== EXAMPLE 9: Cache Warming ====================
/*
func warmCache(cache CacheStrategy, api *APIService, numWarmupRequests int) {
	gen := NewRequestGenerator(ProfileHotSpot, 100)

	fmt.Println("Warming up cache...")
	for i := 0; i < numWarmupRequests; i++ {
		req := gen.Next()
		api.GetUser(req.ID)
	}

	stats := cache.Stats()
	fmt.Printf("Cache warmed: %d items, %d hits\n", stats.Size, stats.Hits)
}
*/

// ==================== EXAMPLE 10: Real-time Metrics Dashboard ====================
/*
func liveMetrics(cache CacheStrategy, duration time.Duration) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	timeout := time.After(duration)
	gen := NewRequestGenerator(ProfileRandom, 100)
	api := NewAPIService(cache)

	for {
		select {
		case <-ticker.C:
			// Generate a batch of requests
			for i := 0; i < 50; i++ {
				req := gen.Next()
				api.GetUser(req.ID)
			}

			// Display metrics
			stats := cache.Stats()
			fmt.Printf("\rHits: %4d │ Misses: %4d │ Hit Rate: %5.1f%% │ Size: %2d/20",
				stats.Hits, stats.Misses,
				float64(stats.Hits)/float64(stats.Hits+stats.Misses)*100,
				stats.Size)
		case <-timeout:
			fmt.Println("\nLive metrics test completed")
			return
		}
	}
}
*/
