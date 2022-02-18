package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	ID       string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

const (
	host     = "118.98.64.74"
	port     = 5432
	user     = "postgres"
	password = "S@y@Br0nt0"
	dbname   = "belajar_golang"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	// db, err := sql.Open("postgres", psqlInfo)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	user := []User{}
	err = db.Select(&user, "SELECT * FROM user_masjit ORDER BY username ASC")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully connected!")
	fmt.Println("data", user)

}
