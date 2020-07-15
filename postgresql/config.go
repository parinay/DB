package main

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "*******"
	dbname   = "zorbish_demo"
)

type User struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
}
