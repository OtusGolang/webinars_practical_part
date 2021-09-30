package main

import (
	"log"
)

func main() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	// Println writes to the standard logger.
	log.Println("main started")
	// Fatalln is Println() followed by a call to os.Exit(1)
	log.Fatalln("fatal message")
	// Panicln is Println() followed by a call to panic()
	log.Panicln("panic message")
}
