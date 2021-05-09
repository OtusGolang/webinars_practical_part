package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	open()
	readPartOf()
	readAll()
	readAt()
}

func open() {
	var file *os.File

	file, err := os.OpenFile("no such file", os.O_RDWR, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("not found")
			return
		}
		panic(err)
	}
	defer file.Close()

	// Some file logic
}

func readPartOf() {
	// Вычитываем мегабайт данных с помощью килобайтного буфера.
	const n = 1 << 10
	buf := make([]byte, 1024)

	file, err := os.Open("/etc/hosts")
	if err != nil {
		panic(err)
	}

	var offset int
	for offset < n {
		read, err := file.Read(buf[offset:])
		offset += read
		if err == io.EOF {
			fmt.Println("!!! EOF !!!")
			break
		}
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(string(buf))
}

func readAll() {
	file, err := os.Open("/etc/hosts")
	if err != nil {
		panic(err)
	}

	// 1
	b1 := make([]byte, 1<<20)
	n, err := io.ReadFull(file, b1)
	if err == io.ErrUnexpectedEOF {
		fmt.Println("buffer is larger than file")
	} else if err != nil {
		fmt.Println("read", n, "bytes")
		panic(err)
	}

	// 2
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		panic(err)
	}

	b2, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	if string(b1) != string(b2) {
		panic("data is not equal")
	}
}

func readAt() {
	file, err := os.Open("/etc/hosts")
	if err != nil {
		panic(err)
	}

	// tail -c 10 /etc/hosts
	offset, err := file.Seek(-10, io.SeekEnd)
	if err != nil {
		panic(err)
	}

	b := make([]byte, 20)
	if _, err := file.ReadAt(b, offset); err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
