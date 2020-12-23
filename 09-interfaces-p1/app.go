package main

type App struct {
	Log Logger
}

func (a *App) Run() error {
	return a.Log.LogToFile("Hello")
}
