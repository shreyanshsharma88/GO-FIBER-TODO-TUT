package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

var DBPool *pgxpool.Pool

func ConnectDB() {
	var err error
	var dbError error
	dbConnectionString := "postgres://postgres:password@localhost:5432/go_fiber_todo"
	DBPool, dbError = pgxpool.New(context.Background(), dbConnectionString)
	if dbError != nil {
		log.Fatal(err)
	}
	err = DBPool.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")
}
