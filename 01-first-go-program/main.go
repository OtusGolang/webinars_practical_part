package main // Текущий пакет

// Импортируем пакеты
import (
	"os" // Пакет для работы с ОС из стандартной библиотеки Go
	// Пакет "searcher" нашего модуля "github.com/OtusGolang/webinars_practical_part/1.1_first_go_program"
	"github.com/OtusGolang/webinars_practical_part/1.1_first_go_program/searcher"
)

// Константа уровня пакета
const keyword = "чиновник"

// «Точка входа» в программу
func main() {
	// Вызываем функцию "Search" из пакета "searcher"
	searcher.Search(feeds, keyword, os.Stdout)
}
