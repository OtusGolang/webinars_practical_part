package main

import "github.com/jmoiron/sqlx"

type user struct {
	FirstName string `json:"first_name" db:"first_name"`
	Email     string `json:"email" db:"email"`
	Age       uint8  `json:"age" db:"age"`
}

type sqlExecutor interface {
	sqlx.Ext
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
}
