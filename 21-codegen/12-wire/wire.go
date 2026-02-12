//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

// InitializeApp creates and wires all dependencies
func InitializeApp() *UserService {
	// Этот код не будет скомпилирован (см дирктивы вверху), но Wire будет его использовать для генерации реальной функции в wire_gen.go
	// По сути, мы описываем, какие зависимости нам нужны и как их конструкторы, а Wire генерирует код, который это делает
	wire.Build(
		NewLogger,
		NewDatabase,
		NewUserRepository,
		NewPrinterService,
		NewUserService,
	)
	return nil
}
