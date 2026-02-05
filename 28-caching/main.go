package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("╔════════════════════════════════════════════════════════════════╗")
	fmt.Println("║          CACHING EVICTION MECHANISMS DEMONSTRATION              ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	caches := []CacheStrategy{
		NewLRUCache(20),
		NewLFUCache(20),
		NewFIFOCache(20),
		NewRandomCache(20),
		NewTTLCache(20, 500*time.Millisecond),
	}

	generators := []string{
		ProfileSequential,
		ProfileRandom,
		ProfileHotSpot,
		ProfileSkewed,
	}

	// Run demonstrations
	for _, generator := range generators {
		runDemonstration(generator, caches)
		fmt.Println()
	}

	// Summary comparison
	printSummary()
}

func runDemonstration(generatorProfile string, caches []CacheStrategy) {
	fmt.Printf("┌─────────────────────────────────────────────────────────────────┐\n")
	fmt.Printf("│ REQUEST PATTERN: %s\n", padString(GetProfileName(generatorProfile), 58))
	fmt.Printf("└─────────────────────────────────────────────────────────────────┘\n")
	fmt.Println()

	numRequests := 500
	requestGen := NewRequestGenerator(generatorProfile, 100)

	for _, cache := range caches {
		// Reset metrics by creating new cache
		var testCache CacheStrategy
		switch cache.(type) {
		case *LRUCache:
			testCache = NewLRUCache(20)
		case *LFUCache:
			testCache = NewLFUCache(20)
		case *FIFOCache:
			testCache = NewFIFOCache(20)
		case *RandomCache:
			testCache = NewRandomCache(20)
		case *TTLCache:
			testCache = NewTTLCache(20, 500*time.Millisecond)
		}

		api := NewAPIService(testCache)
		api.ResetCallCount()

		// Generate requests
		startTime := time.Now()
		for i := 0; i < numRequests; i++ {
			req := requestGen.Next()
			_ = api.GetUser(req.ID)
		}
		elapsed := time.Since(startTime)

		stats := testCache.Stats()

		// Calculate hit rate
		totalRequests := stats.Hits + stats.Misses
		hitRate := 0.0
		if totalRequests > 0 {
			hitRate = float64(stats.Hits) / float64(totalRequests) * 100
		}

		// Print results
		fmt.Printf("  %-30s │ ", testCache.Name())
		fmt.Printf("Size: %2d/20 │ ", stats.Size)
		fmt.Printf("Hits: %3d │ ", stats.Hits)
		fmt.Printf("Misses: %3d │ ", stats.Misses)
		fmt.Printf("Evictions: %3d │ ", stats.Evictions)
		fmt.Printf("Hit Rate: %5.1f%% │ ", hitRate)
		fmt.Printf("DB Calls: %3d │ ", api.GetCallCount())
		fmt.Printf("%v\n", elapsed.Round(time.Millisecond))
	}

	fmt.Println()
}

func printSummary() {
	fmt.Println("┌─────────────────────────────────────────────────────────────────┐")
	fmt.Println("│                         SUMMARY GUIDE                            │")
	fmt.Println("└─────────────────────────────────────────────────────────────────┘")
	fmt.Println()

	fmt.Println("CACHE STRATEGIES:")
	fmt.Println()

	strategies := []struct {
		name    string
		pros    []string
		cons    []string
		bestFor string
	}{
		{
			name: "LRU (Least Recently Used)",
			pros: []string{
				"• Evicts least recently accessed items",
				"• Generalizes well for most workloads",
				"• Simple to implement and understand",
			},
			cons: []string{
				"• Cache pollution on sequential scans",
				"• Can be inefficient with large datasets",
			},
			bestFor: "General-purpose caching, working sets",
		},
		{
			name: "LFU (Least Frequently Used)",
			pros: []string{
				"• Removes least frequently accessed items",
				"• Good for long-term access patterns",
				"• Preserves popular items in cache",
			},
			cons: []string{
				"• Requires frequency tracking overhead",
				"• Can miss temporal locality",
				"• Expensive to maintain state",
			},
			bestFor: "Hot-spot patterns, temporal databases",
		},
		{
			name: "FIFO (First In First Out)",
			pros: []string{
				"• Simple and very fast",
				"• Minimal memory overhead",
				"• Predictable behavior",
			},
			cons: []string{
				"• Doesn't account for access patterns",
				"• May evict frequently accessed items",
				"• Poor hit rate on typical workloads",
			},
			bestFor: "Bounded queues, stream processing",
		},
		{
			name: "Random Eviction",
			pros: []string{
				"• Extremely simple, minimal overhead",
				"• No state to maintain",
				"• Prevents pathological cases",
			},
			cons: []string{
				"• Unpredictable behavior",
				"• Often poor hit rates",
				"• May evict important data",
			},
			bestFor: "Fast approximations, debugging",
		},
		{
			name: "TTL (Time To Live)",
			pros: []string{
				"• Automatic data expiration",
				"• Useful for volatile/dynamic data",
				"• Cache freshness guarantee",
			},
			cons: []string{
				"• Requires background cleanup",
				"• Doesn't optimize for access patterns",
				"• Expires useful data",
			},
			bestFor: "Session data, temporary caches, expiring content",
		},
	}

	for i, s := range strategies {
		if i > 0 {
			fmt.Println()
		}
		fmt.Printf("  %d. %s\n", i+1, s.name)
		fmt.Printf("     Best for: %s\n", s.bestFor)
		fmt.Println("     Pros:")
		for _, pro := range s.pros {
			fmt.Printf("       %s\n", pro)
		}
		fmt.Println("     Cons:")
		for _, con := range s.cons {
			fmt.Printf("       %s\n", con)
		}
	}

	fmt.Println()
	fmt.Println("REQUEST PATTERNS:")
	fmt.Println()
	fmt.Println("  • Sequential: Cycles through all keys (0, 1, 2, ... 99, 0, 1, ...)")
	fmt.Println("    → Tests behavior with sequential scans → LRU struggles")
	fmt.Println()
	fmt.Println("  • Random: Uniform distribution across all keys")
	fmt.Println("    → Tests general-purpose behavior → Most balanced")
	fmt.Println()
	fmt.Println("  • Hot Spot: 70% access to 10% of keys")
	fmt.Println("    → Tests locality awareness → LFU and LRU excel")
	fmt.Println()
	fmt.Println("  • Skewed: Zipfian-like distribution (power law)")
	fmt.Println("    → Tests real-world patterns → Realistic scenarios")
	fmt.Println()
}

func padString(s string, length int) string {
	if len(s) >= length {
		return s
	}
	return s + strings.Repeat(" ", length-len(s))
}
