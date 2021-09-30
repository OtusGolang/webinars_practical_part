package main

import (
	"fmt"

	"github.com/OtusGolang/webinars_practical_part/21-codegen/05-easyjson/student"
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

	//data, err := json.Marshal(s)
	data, err := s.MarshalJSON()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
