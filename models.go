package main

import (
	"context"
	"fmt"
	"os"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

func initDB() {
	err := godotenv.Load(".env")
	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	ctx := context.Background()

	var bla []map[string]int

	query := `SELECT count(*) AS text FROM text_messages`

	err = pgxscan.Select(ctx, db, &bla, query)

	fmt.Print(bla)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

}
