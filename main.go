package main

import (
	"fmt"

	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	databaseConnectionURL := "postgres://postgres:postgres@localhost:5432/voyager"
	connection, err := InitializeDatabase(databaseConnectionURL)

	if err != nil {
		panic("unable to connect to database")
	}

	fmt.Println("Sucesfully connected to database:", connection)
}

//InitializeDatabase Connect to the database and apply the migrations. Return the database connection handle.
func InitializeDatabase(url string) (*gorm.DB, error) {

	// Extract the database configuration using the DATABASE_URL environment variable.
	parsedURL, err := pq.ParseURL(url)

	if err != nil {
		return nil, err
	}

	// Open a connection to postgres
	database, err := gorm.Open(postgres.Open(parsedURL))

	// Check if there was an error while opening a connection to the database
	if err != nil {
		return nil, err
	}

	// Return the database object
	return database, nil
}
