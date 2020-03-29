package main

import (
	"github.com/otusteam/go/cleancalendar/cmd"
	"log"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
