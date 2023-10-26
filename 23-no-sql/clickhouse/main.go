package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/ClickHouse/clickhouse-go"
)

func main() {
	// Define ClickHouse connection parameters
	dsn := "tcp://localhost:9000?username=default&password="

	// Open a connection to ClickHouse
	conn, err := sql.Open("clickhouse", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Create a table
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS sample (
			id UInt32,
			name String
		) ENGINE = MergeTree()
		ORDER BY id
	`

	_, err = conn.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	// Start a transaction
	tx, err := conn.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback() // Rollback if we encounter an error

	// Prepare the insert statement within the transaction
	insertSQL := `
		INSERT INTO sample (id, name) VALUES (?, ?)
	`

	// Create a prepared statement for the insert
	stmt, err := tx.Prepare(insertSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Define data to be inserted
	data := [][]interface{}{
		{1, "John"},
		{2, "Jane"},
		{3, "Doe"},
	}

	// Batch insert data
	for _, row := range data {
		_, err = stmt.Exec(row...)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	// Query data from the table
	querySQL := `
		SELECT * FROM sample
	`

	rows, err := conn.Query(querySQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate over the result rows
	for rows.Next() {
		var id uint32
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
