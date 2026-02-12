package main

func main() {
	// Тут - вызов сгенерироанной функции, которая создаёт все зависимости
	app := InitializeApp()

	// Use the app
	app.GetUserInfo("123")
	app.GetUserInfo("456")
}
