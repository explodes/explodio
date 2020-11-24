package stand

import "github.com/go-pg/pg/v10"

func ConnectPostgres() *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     DefaultEnv("POSTGRES_HOST", "postgres"),
		Database: RequireEnv("POSTGRES_DATABASE"),
		Password: RequireEnv("POSTGRES_PASSWORD"),
		User:     RequireEnv("POSTGRES_USER"),
	})
}
