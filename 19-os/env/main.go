package main

import (
	"fmt"
	"os"
)

func main() {
	env := os.Environ() // слайс строк "key=value"
	fmt.Println(env[0]) // USER=rob

	user, ok := os.LookupEnv("USER")
	fmt.Println(user, ok) // rob

	os.Setenv("PASSWORD", "qwe123")                     // установить
	os.Unsetenv("PASSWORD")                             // удалить
	fmt.Println(os.ExpandEnv("$USER lives in ${CITY}")) // "шаблонизация"
}
