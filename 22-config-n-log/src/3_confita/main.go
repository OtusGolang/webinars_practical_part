package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
)

type Config struct {
	ServiceName string        `config:"serviceName"`
	Port        uint32        `config:"port"`
	Timeout     time.Duration `config:"timeout"`
	Directory   string        `json:"directory"`
}

//serviceName=go-is-go go run ./src/3_confita
func main() {
	cfg := Config{
		ServiceName: "ConfitaTest",
		Port:        5656,
		Timeout:     5 * time.Second,
	}

	loader := confita.NewLoader(
		file.NewBackend("./config/config.json"),
		env.NewBackend(),
	)
	err := loader.Load(context.Background(), &cfg)
	if err != nil {
		log.Fatalf("failed to load: %v", err)
	}

	fmt.Printf("%+v\n", cfg)
}
