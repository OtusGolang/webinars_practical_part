// FOR *nix: //go:generate ./command.sh
// FOR WINDOWS: //go:generate cmd /C command.win.bat
//go:generate ./command.sh

package main

import "fmt"

func main() {
	fmt.Println("run any unix command in go:generate")
}

//go:generate -command list ls -l
//go:generate -command bye echo "Goodbye, world!"

//go:generate bye
//go:generate list
//go:generate go run generate.go

//go:generate echo f=$GOFILE p=$GOPACKAGE r=$GOROOT a=$GOARCH o=$GOOS d=$DOLLAR l=$GOLINE
// Список: https://pkg.go.dev/cmd/go#hdr-Generate_Go_files_by_processing_source

//go:generate pwd

// go generate
// go generate -v
// go generate -x
// go generate -n
// go generate -run bye

// to make go install work:
// export PATH=$PATH:$(go env GOPATH)/bin
