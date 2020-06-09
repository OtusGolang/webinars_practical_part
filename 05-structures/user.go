package main

type User struct { // Структура с именованными полями
	Id      int64
	Name    string
	Age     int
	friends []int64 // Приватный элемент
}

func (u *User) HappyBirthday() {
	u.Age++
}
