package lib

import (
	"context"
	"fmt"
	"jekabot/models"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var ctx context.Context

func GetDatabaseInstance(cDB models.DatabaseConfig) *pgxpool.Pool {
	databaseUrl := fmt.Sprintf("%s://%s:%s@%s/%s",
		cDB.Type,
		cDB.User,
		cDB.Password,
		cDB.Addr,
		cDB.DBName)

	log.Println(databaseUrl)

	ctx = context.Background()

	db, err := pgxpool.Connect(ctx, databaseUrl)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return db
}
