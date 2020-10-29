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
func dbConnection() (*sql.DB, error) {

	db, err := sql.Open("mysql", dataSourceName(""))

	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return nil, err
	}
	// defer db.Close()
	// create
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return nil, err
	}
	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return nil, err
	}
	log.Printf("No of rows affected are %d\n", no)

	db.Close()

	db, err = sql.Open("mysql", dataSourceName(dbname))
	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return nil, err
	}
	// defer db.Close()

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB %s", err, dbname)
		return nil, err
	}
	log.Printf("Connected to DB %s successfully\n", dbname)

	return db, nil
}
func createTable(db *sql.DB) error {
	// create table
	query := `CREATE TABLE  IF NOT EXISTS users (
		id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100),
		city VARCHAR(20),
		state VARCHAR(12)
		)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating the table\n", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return err
	}
	log.Printf("Rows affected when creating the table:%d\n", rows)
	return nil
}
func insert(db *sql.DB, u Users) error {

	// insert
	query := "INSERT INTO users VALUES (?, ?, ?, ?)"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	insert, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s inserting into DB", err)
		return err
	}
	defer insert.Close()

	res, err := insert.ExecContext(ctx, u.ID, u.Name, u.City, u.State)
	if err != nil {
		log.Printf("Error %s when inseting row into users table", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d users created ", rows)
	return nil
}
func main() {
	db, err := dbConnection()
	if err != nil {
		log.Printf("Error %s when setting up the db connection", err)
		return
	}
	defer db.Close()

	err = createTable(db)
	if err != nil {
		log.Printf("Error %s when creating the table", err)
		return
	}

	u := Users{
		ID:    1,
		Name:  "Danny",
		City:  "Mumbai",
		State: "MH",
	}
	err = insert(db, u)
	if err != nil {
		log.Printf("Error %s when inserting a row in the table", err)
		return
	}
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
