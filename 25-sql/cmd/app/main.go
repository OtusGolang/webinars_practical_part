package main

import (
	"context"
	"log"

	"github.com/OtusGolang/webinars_practical_part/25-sql/internal/app"
	"github.com/OtusGolang/webinars_practical_part/25-sql/internal/config"
	"github.com/OtusGolang/webinars_practical_part/25-sql/internal/repository/psql"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c, err := config.Read("configs/local.toml")
	if err != nil {
		log.Fatal(err)
	}

	r := new(psql.Repo)
	if err := r.Connect(ctx, c.PSQL.DSN); err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	a, err := app.New(r)
	if err != nil {
		log.Fatal(err)
	}

	if err := a.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
