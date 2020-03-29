package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"time"
)

type config struct {
	Home         string        `env:"HOME"`
	Port         int           `env:"PORT" envDefault:"3000"`
	IsProduction bool          `env:"PRODUCTION,required"`
	Hosts        []string      `env:"HOSTS" envSeparator:":"`
	Duration     time.Duration `env:"DURATION"`
	TempFolder   string        `env:"TEMP_FOLDER" envDefault:"${HOME}/tmp"`
}

/*
PRODUCTION=true HOSTS="host1:host2:host3" DURATION=1s go run ./caarlos-env/main.go
{Home:/your/home Port:3000 IsProduction:true Hosts:[host1 host2 host3] Duration:1s}
*/
func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	fmt.Printf("%+v\n", cfg)
}
