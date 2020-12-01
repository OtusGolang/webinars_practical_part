package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var mainConfig config

type config struct {
	ServiceName string `yaml:"service_name"`
}

func (m *config) parse() {
	var configPath = flag.String("config", "./config/config.yml", "path to config file")
	flag.Parse()

	configYml, err := ioutil.ReadFile(*configPath)
	if err != nil {
		log.Fatalf("reading config.yml error: %v", err)
	}

	err = yaml.Unmarshal(configYml, m)
	if err != nil {
		log.Fatalf("can't parse config.yml: %v", err)
	}
}

func main() {
	mainConfig.parse()
	fmt.Println(mainConfig.ServiceName)
}
