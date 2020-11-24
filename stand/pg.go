package stand

import "github.com/go-pg/pg/v10"

func ConnectPostgres() *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     DefaultEnv("DB_HOST", "postgres"),
		Database: RequireEnv("DB_DATABASE"),
		Password: RequireEnv("DB_PASSWORD"),
		User:     RequireEnv("DB_USER"),
	})
}
