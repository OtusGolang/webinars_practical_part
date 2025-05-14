package main

// Processor — базовая функция‑обработчик.
// В реальном приложении это может быть что угодно: работа с файлом,
// запрос к БД, вычисления, отправка данных и т.д.
// Принимает строку и возвращает строку.
type Processor func(string) string

// Middleware — функция‑обёртка над Processor.
// Принимает "следующий" Processor и возвращает новый Processor,
// выполняя дополнительные действия до и/или после вызова next.
type Middleware func(Processor) Processor

// Chain строит окончательный Processor, последовательно оборачивая исходный p
// всеми middleware из списка. Middleware применяются в порядке передачи
// (m1, m2, m3) → m1(next(m2(next(m3(p)))))
func Chain(p Processor, mws ...Middleware) Processor {
	// Оборачиваем в обратном порядке, чтобы первый в списке выполнился первым
	for i := len(mws) - 1; i >= 0; i-- {
		p = mws[i](p)
	}
	return p
}
