package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "root:password@/testbase")
    if err != nil {
        panic(err)
    }
    defer db.Close()
    fmt.Printf("Hello world")
}
