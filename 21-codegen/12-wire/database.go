package main

import "fmt"

// Database represents a simple database connection
type Database struct {
	connectionString string
	logger           Logger
}

// NewDatabase creates a new Database with logger
func NewDatabase(logger Logger) *Database {
	logger.Info("Initializing database connection")
	return &Database{
		connectionString: "localhost:5432",
		logger:           logger,
	}
}

// Query executes a query
func (db *Database) Query(sql string) string {
	db.logger.Info(fmt.Sprintf("Executing query: %s", sql))
	return "Query result from " + sql
}

// UserRepository handles user data access
type UserRepository struct {
	db *Database
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *Database) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// GetUser retrieves a user by ID
func (r *UserRepository) GetUser(id string) string {
	result := r.db.Query("SELECT * FROM users WHERE id = " + id)
	return result
}
