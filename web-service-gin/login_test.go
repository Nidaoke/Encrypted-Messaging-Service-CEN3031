package main

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestCheckLogin(t *testing.T) {
	fmt.Println("START GET REQUEST TEST")
	const file string = "test.db"

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Open database succuss!")

	rows, err := db.Query("SELECT * FROM accounts")

	fmt.Println("Database query succuss!")

	numRequests := 0

	expectedUser := "Tester1"
	expectedPass := "1914338c1103962d014174aa8266f9430712cae1fa5bebac25d0c60b8a23d75d"

	var user string
	var pass string
	var email string

	for rows.Next() {

		err = rows.Scan(&user, &pass, &email)

		if err != nil {
			panic("Failed to iterate through the database! %w")
		}

		if user == expectedUser && expectedPass == pass {
			fmt.Printf("The Login Check is successful for %s and %s. \n", user, pass)
			break
		}

		numRequests++
	}

	fmt.Println("Number of requests in database:", numRequests)

	rows.Close()

	db.Close()

}
