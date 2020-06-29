package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {

	db, err := sql.Open("mysql", "root:passowrd@tcp(127.0.0.1:3306)/testDB")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// insert
	insert, err := db.Query("INSERT INTO users VALUES (3, 'danny', 'Mumbai','MH')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	// query
	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var user Users
		err = results.Scan(&user.Id, &user.Name, &user.City, &user.State)
		if err != nil {
			panic(err.Error())
		}
		log.Println(user.Id, user.Name, user.City, user.State)
	}

}
