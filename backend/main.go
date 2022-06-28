package main

import (
	"database/sql"

	"github.com/rg-km/final-project-engineering-66/backend/api"
	"github.com/rg-km/final-project-engineering-66/backend/repository"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}
	usersRepo := repository.NewUserRepository(db)
	kampusRepo := repository.NewkampusRepository(db)
	jurusanRepo := repository.NewJurusanRepository(db)
	reviewRepo := repository.NewReviewRepository(db)
	api := api.NewAPI(*usersRepo, *kampusRepo, *jurusanRepo, *reviewRepo)
	api.Start()
}
