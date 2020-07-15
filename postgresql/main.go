package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// errorF() - error function
func errorF(err error) {
	if err != nil {
		panic(err)
	}
}

func connectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	// connect to the DB - validate the args provided
	db, err := sql.Open("postgres", psqlInfo)
	/*if err != nil {
		panic(err)
	}*/
	errorF(err)

	// connect to the DB - open a connection to the DB
	err = db.Ping()
	errorF(err)
	fmt.Println("Successfully connected!")
	return db
}
func main() {

	// open and connect to DB
	db := connectDB()
	defer db.Close()

	//
	sqlStatement := `
	DELETE FROM users
	WHERE id = $1;`

	_, err := db.Exec(sqlStatement, 1)
	errorF(err)
	fmt.Printf("Delete :: Successful\n")

	// create
	id := 0
	sqlStatement = `
	INSERT INTO users (age, email, first_name, last_name)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	err = db.QueryRow(sqlStatement, 5, "swara@P.com", "swara", "P").Scan(&id)
	errorF(err)

	fmt.Println("Create :: New record ID is:\n", id)

	// update
	sqlStatement = `UPDATE users
	SET age = $2, first_name = $3, last_name = $4, email = $5
	WHERE id =$1
	RETURNING id, email;`

	var email string
	/*
		res, err := db.Exec(sqlStatement, id, "spruha", "K", "spruha@K.com")
		errorF(err)
		count, err := res.RowsAffected()
		errorF(err)
		fmt.Println(count)
	*/
	err = db.QueryRow(sqlStatement, id, 5, "SpruhA", "K", "SpruhA@K.com").Scan(&id, &email)
	errorF(err)

	fmt.Printf("Update :: ID=%d email=%s\n", id, email)

	// query - single record
	var user User
	sqlStatement = `SELECT * FROM users WHERE id=$1;`

	// QueryRow() error handling is deferred untill Scan()
	// And error is returned when 0 records are found.
	row := db.QueryRow(sqlStatement, id)
	err = row.Scan(&user.ID, &user.Age, &user.FirstName, &user.LastName, &user.Email)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Printf("Query :: %v\n", user)
	default:
		errorF(err)
	}

	// query - multiple records
	fmt.Println("Multi row Query ::")
	sqlStatement = `SELECT id, first_name FROM users LIMIT $1;`

	// Query() No error is returned if no records are found.
	// err could be returned even before you iterate over rows.
	rows, err := db.Query(sqlStatement, 3)
	errorF(err)

	// Defer Close()
	// Next() method closed the *Rows, if there are no more rows left.
	// In case of error you will have to manially call close.
	defer rows.Close()

	// Next() will return an error or nil
	for rows.Next() {
		var id int
		var firstname string

		err = rows.Scan(&id, &firstname)
		errorF(err)
		fmt.Printf("ID = %d First Name = %s\n", id, firstname)
	}
	// Verify there was no error when calling rows.Next()
	err = rows.Err()
	errorF(err)
}
