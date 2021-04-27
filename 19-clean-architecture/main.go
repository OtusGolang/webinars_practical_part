package main

import (
	"log"

	"github.com/OtusGolang/webinars_practical_part/19-clean-architecture/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
