package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/pressly/goose"
)

var (
	dir           = "./db/migrations"
	driver        = "mysql"
	createCommand = "create"
	upCommand     = "up"
	downCommand   = "down"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide command")
	}
	command := os.Args[1]
	switch command {
	case createCommand:
		name := os.Args[2]
		if name == "" {
			log.Fatal("Please provide migration name")
		}
		if err := goose.Run("create", nil, dir, os.Args[1], "sql"); err != nil {
			log.Fatalf("goose run: %v", err)
		}
		return
	case upCommand:
	case downCommand:
		if len(os.Args) != 2 {
			log.Fatal("Invalid args length")
		}

		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASS")
		dbHost := os.Getenv("DB_HOST")
		dbName := os.Getenv("DB_NAME")

		dbstring := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbHost, dbName)
		db, err := sql.Open(driver, dbstring)
		if err != nil {
			log.Fatalf("Failed open connction: %s", err)
		}
		if err := goose.Run(command, db, dir); err != nil {
			log.Fatalf("Failed goose run: %s", err)
		}
	default:
		log.Fatal("Invalid command")
	}

}
