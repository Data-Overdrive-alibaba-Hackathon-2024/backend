package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var (
	dbName = DBName()
	dbUser = DBUser()
	dbPass = DBPass()
	dbHost = DBHost()
	dbPort = DBPort()
)

func NewDBPool() *sql.DB {
	conn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalln(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	log.Println("Connected to database")

	return db
}
