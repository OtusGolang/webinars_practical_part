package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	lsCmd := exec.Command("ls")
	wcCmd := exec.Command("wc", "-l")

	pipe, _ := lsCmd.StdoutPipe()
	wcCmd.Stdin = pipe
	wcCmd.Stdout = os.Stdout

	log.Println("ls start err:", lsCmd.Start())
	log.Println("wc start err:", wcCmd.Start())
	log.Println("ls wait err:", lsCmd.Wait())
	log.Println("wc wait err:", wcCmd.Wait())
}
