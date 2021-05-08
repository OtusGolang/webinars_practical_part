package student

type Discipline = string

//go:generate easyjson -all student.go
type Student struct {
	FirstName  string `json:"first_name"`
	SecondName string
	Age        int
	Marks      map[Discipline]int
}
