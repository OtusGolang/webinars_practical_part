package main

import (
	"encoding/json"
	"testing"

	"github.com/OtusGolang/webinars_practical_part/21-codegen/easyjson/student"
)

// go test -bench=. -benchmem

type NotEasyStudent struct {
	FirstName  string
	SecondName string
	Age        int
	Marks      map[student.Discipline]int
}

func BenchmarkEasyStudent(b *testing.B) {
	s := student.Student{
		FirstName:  "Otus",
		SecondName: "Otusov",
		Age:        25,
		Marks: map[student.Discipline]int{
			"Golang":     5,
			"JavaScript": 3,
		},
	}
	for i := 0; i < b.N; i++ {
		data, err := json.Marshal(s)
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal(data, &student.Student{}); err != nil {
			panic(err)
		}
	}
}

func BenchmarkNotEasyStudent(b *testing.B) {
	s := NotEasyStudent{
		FirstName:  "Otus",
		SecondName: "Otusov",
		Age:        25,
		Marks: map[student.Discipline]int{
			"Golang":     5,
			"JavaScript": 3,
		},
	}
	for i := 0; i < b.N; i++ {
		data, err := json.Marshal(s)
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal(data, &NotEasyStudent{}); err != nil {
			panic(err)
		}
	}
}
