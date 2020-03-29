package main

import (
	"context"
	"fmt"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"time"
)
import "github.com/heetch/confita"

func main() {

	// serviceName=go-is-go go run confita/main.go
	cfg := Config{
		ServiceName: "ConfitaTest",
		Port:        5656,
		Timeout:     5 * time.Second,
	}
	loader := confita.NewLoader(
		file.NewBackend("/Users/a.zheltak/GolandProjects/awesomeProject/config/config.json"),
		env.NewBackend(),
	)
	err := loader.Load(context.Background(), &cfg)
	if err != nil {
		panic(err)
	}
	fmt.Print(cfg)
}

type Config struct {
	ServiceName string        `config:"serviceName"`
	Port        uint32        `config:"port"`
	Timeout     time.Duration `config:"timeout"`
}
