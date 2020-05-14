package main

import (
	"encoding/json"
	"fmt"

	"github.com/OtusGolang/webinars_practical_part/21-codegen/easyjson/student"
)

func main() {
	s := student.Student{
		FirstName:  "Otus",
		SecondName: "Otusov",
		Age:        25,
		Marks: map[student.Discipline]int{
			"Golang":     5,
			"JavaScript": 3,
		},
	}
	data, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
