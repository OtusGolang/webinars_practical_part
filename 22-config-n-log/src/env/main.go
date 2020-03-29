package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	httpPort, err := strconv.Atoi(os.Getenv("SHORTENER_PORT"))
	if err != nil {
		panic(fmt.Sprint("SHORTENER_PORT not defined"))
	}
	shortenerHost := os.Getenv("SHORTENER_HOST")
	if shortenerHost == "" {
		panic(fmt.Sprint("SHORTENER_HOST not defined"))
	}
	config := Config{httpPort, shortenerHost}
	fmt.Print(config)
}

type Config struct {
	Port int
	Host string
}
