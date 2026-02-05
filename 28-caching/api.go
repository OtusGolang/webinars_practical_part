package main

import (
	"fmt"
	"sync"
	"time"
)

// User represents a user entity
type User struct {
	ID      string
	Name    string
	Email   string
	Fetched time.Time
}

// APIService represents the protected API
type APIService struct {
	cache CacheStrategy
	mu    sync.RWMutex
	calls int
}

// NewAPIService creates a new API service with a cache
func NewAPIService(cache CacheStrategy) *APIService {
	return &APIService{
		cache: cache,
		calls: 0,
	}
}

// GetUser retrieves a user (checks cache first)
func (api *APIService) GetUser(userID string) *User {
	// Check cache first
	if cached, found := api.cache.Get(userID); found {
		return cached.(*User)
	}

	// Cache miss - fetch from "database"
	user := api.fetchUserFromDB(userID)

	// Store in cache
	api.cache.Set(userID, user)

	return user
}

// fetchUserFromDB simulates a database fetch
func (api *APIService) fetchUserFromDB(userID string) *User {
	api.mu.Lock()
	api.calls++
	api.mu.Unlock()

	// Simulate database latency
	time.Sleep(5 * time.Millisecond)

	return &User{
		ID:      userID,
		Name:    fmt.Sprintf("User %s", userID),
		Email:   fmt.Sprintf("%s@example.com", userID),
		Fetched: time.Now(),
	}
}

// GetCallCount returns the number of actual DB calls made
func (api *APIService) GetCallCount() int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	return api.calls
}

// ResetCallCount resets the counter
func (api *APIService) ResetCallCount() {
	api.mu.Lock()
	defer api.mu.Unlock()
	api.calls = 0
}
