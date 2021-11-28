package main

import (
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

	/*psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	 db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} */
}
