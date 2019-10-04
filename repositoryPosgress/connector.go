package repositoryPosgress

import "github.com/go-pg/pg"

func buildPgConnection() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "rootpass",
		Database: "minesweeper",
	})
	return db
}
