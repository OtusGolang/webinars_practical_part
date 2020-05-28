package app

import (
	"context"
	"log"

	"github.com/OtusGolang/webinars_practical_part/25-sql/internal/repository"
)

type App struct {
	r repository.BaseRepo
}

func New(r repository.BaseRepo) (*App, error) {
	return &App{r: r}, nil
}

func (a *App) Run(ctx context.Context) error {
	books, err := a.r.GetBooks(ctx)
	if err != nil {
		return err
	}

	log.Println("books:")
	for _, b := range books {
		log.Printf("\t %+v", b)
	}

	return nil
}
