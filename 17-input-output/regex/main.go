package main

import (
	"fmt"
	"regexp"
)

func main() {

	// Определите шаблон регулярного выражения с подвыражениями
	pattern := regexp.MustCompile(`(\w+)-(\d+)`)

	// Входная строка для сопоставления
	input := "example-123"

	// Найти подвыражения во входной строке
	subMatches := pattern.FindStringSubmatch(input)

	// Проверить, есть ли совпадение
	if len(subMatches) > 0 {
		// Первый элемент - это полное совпадение, следующие элементы - подвыражения
		fmt.Println("Полное совпадение:", subMatches[0])
		fmt.Println("Первое подвыражение:", subMatches[1])
		fmt.Println("Второе подвыражение:", subMatches[2])
	} else {
		fmt.Println("Совпадений не найдено")
	}
}
