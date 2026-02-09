package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("../env/env", "-v")
	cmd.Env = append(os.Environ(),
		`USER="petya"`,
		//"CITY=Msk",
	)
	myRighter := &myWriter{}
	cmd.Stdout = myRighter

	// output, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//log.Printf("output: %s", output)
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	err := cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

// custom writer
type myWriter struct{}

func (w *myWriter) Write(p []byte) (n int, err error) {

	// split by lines and output with timestamp
	//time.Sleep(time.Second * 3)
	lines := strings.Split(strings.TrimSpace(string(p)), "\n")
	for _, line := range lines {
		log.Printf("From app: %s", line)
	}
	return len(p), nil
}
