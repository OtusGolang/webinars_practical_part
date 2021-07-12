package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("../env/env")
	cmd.Env = append(os.Environ(),
		"USER=petya",
		"CITY=Msk",
	)

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
