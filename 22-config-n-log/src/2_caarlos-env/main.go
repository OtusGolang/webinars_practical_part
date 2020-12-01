package main

import (
	"fmt"
	"log"
	"time"

	"github.com/caarlos0/env"
)

type config struct {
	Home         string        `env:"HOME"`
	Port         int           `env:"PORT" envDefault:"3000"`
	IsProduction bool          `env:"PRODUCTION,required"`
	Hosts        []string      `env:"HOSTS" envSeparator:":"`
	Duration     time.Duration `env:"DURATION"`
	TempFolder   string        `env:"TEMP_FOLDER" envExpand:"true"`
}

/*
PRODUCTION=true HOSTS="host1:host2:host3" DURATION=1s go run ./src/caarlos-env
{Home:/your/home Port:3000 IsProduction:true Hosts:[host1 host2 host3] Duration:1s}
*/
func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("%+v\n", cfg)
}
