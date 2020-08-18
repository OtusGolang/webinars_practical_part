package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func main() {
	mainConfig.parse()
	fmt.Println(mainConfig.ServiceName)
}

type config struct {
	ServiceName string `yaml:"service_name"`
}

var mainConfig = config{}

func (m *config) parse() {
	var configPath = flag.String("config", "./config/config.yml", "path to config file")
	flag.Parse()
	configYml, err := ioutil.ReadFile(*configPath)
	if err != nil {
		log.Fatalf("reading config.yml error %v", err)
	}

	err = yaml.Unmarshal(configYml, m)
	if err != nil {
		log.Fatalf("Can't parse config.yml: %v", err)
	}
}
