package main

import (
	"context"
	"fmt"
	"log"

	"github.com/OtusGolang/webinars_practical_part/25-sql/internal/app"
	"github.com/OtusGolang/webinars_practical_part/25-sql/internal/config"
	"github.com/OtusGolang/webinars_practical_part/25-sql/internal/repository/psql"
)

func main() {
	if err := mainImpl(); err != nil {
		log.Fatal(err)
	}
}

func mainImpl() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c, err := config.Read("configs/local.toml")
	if err != nil {
		return fmt.Errorf("cannot read config: %v", err)
	}

	r := new(psql.Repo)
	if err := r.Connect(ctx, c.PSQL.DSN); err != nil {
		return fmt.Errorf("cannot connect to psql: %v", err)
	}
	defer func() {
		if err := r.Close(); err != nil {
			log.Println("cannot close psql connection", err)
		}
	}()

	a, err := app.New(r)
	if err != nil {
		return fmt.Errorf("cannot create app: %v", err)
	}

	if err := a.Run(ctx); err != nil {
		return fmt.Errorf("cannot run app: %v", err)
	}

	return nil
}
