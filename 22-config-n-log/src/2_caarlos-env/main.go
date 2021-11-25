package main

import (
	"fmt"
	"log"
	"time"

   "github.com/caarlos0/env/v6"
)

type config struct {
	Home         string        `env:"HOME"`
	Port         int           `env:"PORT" envDefault:"3000"`
	IsProduction bool          `env:"PRODUCTION,required"`
	Hosts        []string      `env:"HOSTS" envSeparator:":"`
	Duration     time.Duration `env:"DURATION"`
	TempFolder   string        `env:"TEMP_FOLDER" envDefault:"${HOME}/tmp" envExpand:"true"`
}

/*
PRODUCTION=true HOSTS="host1:host2:host3" DURATION=1s go run ./src/2_caarlos-env
{Home:/your/home Port:3000 IsProduction:true Hosts:[host1 host2 host3] Duration:1s}
*/
func main() {
	cfg := config{}

	//opts := env.Options{Environment: map[string]string{
	//	"PRODUCTION": "true",
	//}}

	if err := env.Parse(&cfg); err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("%+v\n", cfg)
}
