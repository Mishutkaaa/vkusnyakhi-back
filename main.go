package main

import (
	"database/sql"
	"log"
	"os"
)

type Drinks struct {
	id       int
	name     string
	category []string
}

type Food struct {
	id       int
	name     string
	category []string
}

var (
	conn = os.Getenv("CONN")
)

func main() {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Println("err conn db", err)
		return
	}
	if err := db.Ping(); err != nil {
		log.Println("cannot conn db")
		return
	}
}
