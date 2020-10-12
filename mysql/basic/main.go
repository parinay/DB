package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	// When we import the package prefixed with the blank identifier _ "github.com/go-sql-driver/mysql",
	// this initi() function is called and the driver is available for use.
	_ "github.com/go-sql-driver/mysql"
)

// Users - Schema
type Users struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	City  string `json:"city"`
	State string `json:"state"`
}

const (
	username = "newuser"
	password = "password"
	hostname = "127.0.0.1"
	port     = "3306"
	dbname   = "ecommerce"
)

// dataSourceName - return the data source name
func dataSourceName(dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, dbname)
	//root:Par1nay!@tcp(127.0.0.2:3306)/testDB
}
func main() {

	db, err := sql.Open("mysql", dataSourceName(""))

	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}
	defer db.Close()
	// create
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return
	}
	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return
	}
	log.Printf("No of rows affected are %d\n", no)

	db.Close()

	db, err = sql.Open("mysql", dataSourceName(dbname))
	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return
	}
	defer db.Close()

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB %s", err, dbname)
		return
	}
	log.Printf("Connected to DB %s successfully\n", dbname)
	// create table
	res, err = db.ExecContext(ctx, `CREATE TABLE  users (
		id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100),
		city VARCHAR(20),
		state VARCHAR(12)
		)`)
	if err != nil {
		log.Printf("Error %s when creating tables\n", err)
		return
	}

	// insert
	insert, err := db.Query("INSERT INTO users VALUES (3, 'danny', 'Mumbai','MH')")
	if err != nil {
		log.Printf("Error %s inserting into DB", err)
		return
	}
	defer insert.Close()
	// query
	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Printf("Error %s running SELECT query", err)
		return
	}
	for results.Next() {
		var user Users
		err = results.Scan(&user.ID, &user.Name, &user.City, &user.State)
		if err != nil {
			panic(err.Error())
		}
		log.Println(user.ID, user.Name, user.City, user.State)
	}
}
