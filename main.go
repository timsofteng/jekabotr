package main

import (
	"log"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	dbInit()
	dbGetTextMessagesCount()
	tgRun()
}
