package main

import (
	"fmt"
	"math/rand"
)

// RequestGenerator generates requests with different patterns
type RequestGenerator struct {
	profile    string
	maxID      int
	sequential int
}

// RequestProfile represents a single request
type RequestProfile struct {
	ID    string
	Count int
}

const (
	ProfileSequential = "sequential"
	ProfileRandom     = "random"
	ProfileHotSpot    = "hotspot"
	ProfileSkewed     = "skewed"
)

// NewRequestGenerator creates a new request generator
func NewRequestGenerator(profile string, maxID int) *RequestGenerator {
	return &RequestGenerator{
		profile:    profile,
		maxID:      maxID,
		sequential: 0,
	}
}

// Next generates the next request
func (g *RequestGenerator) Next() RequestProfile {
	switch g.profile {
	case ProfileSequential:
		id := fmt.Sprintf("user_%d", g.sequential%g.maxID)
		g.sequential++
		return RequestProfile{ID: id, Count: 1}

	case ProfileRandom:
		id := fmt.Sprintf("user_%d", rand.Intn(g.maxID))
		return RequestProfile{ID: id, Count: 1}

	case ProfileHotSpot:
		// 70% of requests to 10% of keys
		if rand.Float64() < 0.7 {
			id := fmt.Sprintf("user_%d", rand.Intn(g.maxID/10))
			return RequestProfile{ID: id, Count: 1}
		}
		id := fmt.Sprintf("user_%d", rand.Intn(g.maxID))
		return RequestProfile{ID: id, Count: 1}

	case ProfileSkewed:
		// Zipfian-like distribution
		userID := rand.Intn(g.maxID)
		// Make smaller IDs more likely
		if rand.Float64() < 0.5 {
			userID = rand.Intn(g.maxID / 4)
		}
		id := fmt.Sprintf("user_%d", userID)
		return RequestProfile{ID: id, Count: 1}

	default:
		id := fmt.Sprintf("user_%d", rand.Intn(g.maxID))
		return RequestProfile{ID: id, Count: 1}
	}
}

// GetProfileName returns a human-readable name
func GetProfileName(profile string) string {
	switch profile {
	case ProfileSequential:
		return "Sequential Access (cycling through all keys)"
	case ProfileRandom:
		return "Uniform Random (all keys equally likely)"
	case ProfileHotSpot:
		return "Hot Spot (70% access to 10% of keys)"
	case ProfileSkewed:
		return "Skewed Distribution (smaller keys more frequent)"
	default:
		return "Unknown"
	}
}
