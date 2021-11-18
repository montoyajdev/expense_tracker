package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user_dev"
	password = "a123"
	dbname   = "bizzytrack_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var id int
	err = db.QueryRow(`
	INSERT INTO users(name,email)
	VALUES($1,$2)
	RETURNING ID`,
		"Jon Calhoun", "jon@jon.com").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("ID is ....", id)
}
