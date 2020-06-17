package main

import (
	"bytes"
	"os"
)

type Logger struct {
	Level int
	File *os.File
	Buf *bytes.Buffer
}

func (l Logger) LogToFile(msg string) error {
	_, err := l.File.Write([]byte(msg))
	return err
}

func (l Logger) LogToBuffer(msg string) error {
	l.Buf.Write([]byte(msg))
	return nil
}
