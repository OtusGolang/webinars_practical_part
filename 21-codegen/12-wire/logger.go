package main

import (
	"fmt"
	"log"
)

// Logger is a simple logger interface
type Logger interface {
	Info(msg string)
	Error(msg string)
}

// SimpleLogger is a concrete implementation of Logger
type SimpleLogger struct {
	prefix string
}

// NewLogger creates a new SimpleLogger
func NewLogger() Logger {
	return &SimpleLogger{
		prefix: "[APP]",
	}
}

// Info logs an info message
func (l *SimpleLogger) Info(msg string) {
	log.Println(l.prefix, "INFO:", msg)
}

// Error logs an error message
func (l *SimpleLogger) Error(msg string) {
	log.Println(l.prefix, "ERROR:", msg)
}

// PrinterService prints messages using logger
type PrinterService struct {
	logger Logger
}

// NewPrinterService creates a new PrinterService
func NewPrinterService(logger Logger) *PrinterService {
	return &PrinterService{
		logger: logger,
	}
}

// Print prints a message
func (p *PrinterService) Print(msg string) {
	p.logger.Info(fmt.Sprintf("Printing: %s", msg))
	fmt.Println(">> " + msg)
}
