package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func main() {
	mainConfig.parse()
	fmt.Println(mainConfig.ServiceName)
}

var mainConfig = Config{}

func (m *Config) parse() {
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

type Config struct {
	ServiceName string `yaml:"service_name"`
}
