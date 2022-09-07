package repository

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var ctx context.Context

func NewDB(databaseUrl string) *pgxpool.Pool {
	ctx = context.Background()

	db, err := pgxpool.Connect(ctx, databaseUrl)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return db
	// defer db.Close(context.Background())
}
