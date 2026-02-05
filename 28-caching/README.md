# Caching Eviction Mechanisms Demonstration

Interactive benchmark demonstrating 5 different cache eviction strategies with 4 realistic access patterns.

## Aim

Show how different cache eviction algorithms behave with various request patterns, measuring:
- Hit rate (% of requests served from cache)
- Database calls avoided
- Eviction patterns
- Performance trade-offs

**Key Finding**: No universal best strategy - choice depends on access pattern and workload characteristics.

## Architecture

```
Request Generator → [4 Patterns] → API Service → Lazy cache access → Cache → [5 Eviction Strategies]
                                                        ↓
                                                Hit? Return
                                                Miss? DB fetch + cache
```

### Components

| Component | Role | File |
|-----------|------|------|
| **Cache Eviction Strategies** | LRU, LFU, FIFO, Random, TTL | `cache.go` |
| **Request Patterns** | Sequential, Random, Hot Spot, Skewed | `generator.go` |
| **API Service** | Protected API with cache-first lookup | `api.go` |
| **Benchmark** | Runs all 20 combinations (5×4) | `main.go` |

### Cache Strategies

- **LRU** (Least Recently Used) - Remove least recently accessed items
- **LFU** (Least Frequently Used) - Remove least frequently accessed items
- **FIFO** (First In First Out) - Remove oldest inserted items
- **Random** - Remove random items
- **TTL** (Time To Live) - Remove expired items

### Request Patterns

- **Sequential**: Cycles through 100 keys (0-99, 0-99, ...) → Worst case for LRU
- **Random**: Uniform distribution → Baseline behavior
- **Hot Spot**: 70% to 10% of keys → Tests frequency awareness
- **Skewed**: Power-law/Zipfian → Real-world workloads

## Usage

```bash
go run .
```

**Output**: 20 test results with hit rates and metrics for each strategy-pattern combination.

### Metrics Explained

| Metric | Meaning |
|--------|---------|
| **Hits** | Requests served from cache |
| **Misses** | Requests needing DB query |
| **Hit Rate** | (Hits / Total) × 100% |
| **Evictions** | Items removed to make room |
| **DB Calls** | Actual database queries |

## Results Summary

| Pattern | Best | Hit Rate | Why |
|---------|------|----------|-----|
| **Sequential** | All fail | ~0% | Cache thrashing (100 keys, size 20) |
| **Random** | LFU | ~22% | Frequency tracking helps |
| **Hot Spot** | LRU | ~75% | Recency matches frequency |
| **Skewed** | LFU | ~41% | Power-law matches LFU logic |

## When to Use Each

| Strategy | Best For | Weakness |
|----------|----------|----------|
| **LRU** | General purpose, web caching | Fails on sequential access |
| **LFU** | Hot/cold data patterns | Overhead, doesn't adapt quickly |
| **FIFO** | Simple queues, streams | Ignores access patterns |
| **Random** | Debugging, approximation | Unpredictable results |
| **TTL** | Session data, volatile data | Ignores access frequency |

## Code Structure

```
28-caching/
├── cache.go         (5 strategies, 330 lines)
├── generator.go     (4 patterns, 98 lines)
├── api.go           (API service, 60 lines)
├── main.go          (Benchmark, 150 lines)
├── EXAMPLES.go      (10 code templates)
├── go.mod           (No external dependencies)
├── output.txt       (Sample output)
└── README.md        (This file)
```

## Implementation Details

All caches implement a common interface:
```go
type CacheStrategy interface {
    Get(key string) (interface{}, bool)
    Set(key string, value interface{})
    Stats() CacheMetrics
    Name() string
}
```

Thread-safe using RWMutex for concurrent access. Fixed cache size of 20 items.

## Key Insights

1. **LRU is the safe default** - ~20% hit rate on random access, simplicity
2. **LFU wins on patterns** - 75% hit rate with hot spots, 41% on skewed
3. **Pattern matters most** - access pattern choice more important than cache size
4. **FIFO competitive** - surprisingly good (~18% hit rate), simpler than LRU
5. **TTL is different** - time-based, not pattern-aware

## Extensions

See `EXAMPLES.go` for 10 templates:
1. Single strategy benchmark
2. Implement new strategy
3. Add new pattern
4. Concurrency testing
5. Visualization
6. Size benchmarking
7. TTL simulation
8. Hybrid LRU+TTL
9. Cache warming
10. Live metrics

Experiment by modifying cache size, request patterns, or implementing new strategies.
