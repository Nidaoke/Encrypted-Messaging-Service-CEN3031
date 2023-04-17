package main

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestGetRequests(t *testing.T) {
	fmt.Println("START GET REQUEST TEST")
	const file string = "test.db"

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Open database succuss!")

	rows, err := db.Query("SELECT * FROM requests")

	fmt.Println("Database query succuss!")

	numRequests := 0

	var expectedId int64 = 0

	var id int64
	var sentFrom string
	var sentTo string

	for rows.Next() {

		err = rows.Scan(&id, &sentFrom, &sentTo)

		if err != nil {
			panic("Failed to iterate through the database! %w")
		} else {
			fmt.Printf("The request is from %s to %s. \n", sentFrom, sentTo)
		}

		numRequests++
	}

	fmt.Print("Expected Request Id: ")
	fmt.Print(expectedId)
	fmt.Print("\nActual Request Id: ")
	fmt.Println(id)

	if expectedId != id {
		panic("Id's do not match!")
	}

	fmt.Println("Number of requests in database:", numRequests)

	rows.Close()

	db.Close()

}

// NOT DONE
func TestGetRequestFrom(t *testing.T) {
	fmt.Println("START GET REQUEST FROM TEST")
	const file string = "test.db"

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Open database succuss!")

	rows, err := db.Query("SELECT * FROM requests")

	fmt.Println("Database query succuss!")

	numRequests := 0

	expectedSentFrom := "tester1"

	var id int64
	var sentFrom string
	var sentTo string

	for rows.Next() {

		err = rows.Scan(&id, &sentFrom, &sentTo)

		if err != nil {
			panic("Failed to iterate through the database! %w")
		}

		numRequests++
	}

	fmt.Print("Expected User sent from: ")
	fmt.Print(expectedSentFrom)
	fmt.Print("\nActual User sent from: ")
	fmt.Println(sentFrom)

	if expectedSentFrom != sentFrom {
		panic("Names do not match!")
	}

	fmt.Println("Number of requests in database:", numRequests)

	rows.Close()

	db.Close()

}

func TestGetRequestTo(t *testing.T) {
	fmt.Println("START GET REQUEST TO TEST")
	const file string = "test.db"

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Open database succuss!")

	rows, err := db.Query("SELECT * FROM requests")

	fmt.Println("Database query succuss!")

	numRequests := 0

	expectedSentTo := "tester2"

	var id int64
	var sentFrom string
	var sentTo string

	for rows.Next() {

		err = rows.Scan(&id, &sentFrom, &sentTo)

		if err != nil {
			panic("Failed to iterate through the database! %w")
		}
		numRequests++
	}

	fmt.Print("Expected User sent to: ")
	fmt.Print(expectedSentTo)
	fmt.Print("\nActual User sent to: ")
	fmt.Println(sentTo)

	if expectedSentTo != sentTo {
		panic("Names do not match!")
	}

	fmt.Println("Number of requests in database:", numRequests)

	rows.Close()

	db.Close()
}

func TestPostRequest(t *testing.T) {
	fmt.Println("START POST REQUEST TEST")
	const file string = "test.db"

	var id int64
	var sentFrom string
	var sentTo string

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Open database succuss!")

	_, err = db.Exec("INSERT INTO requests VALUES(1, \"Marcus\", \"Jon\")")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM requests")
	if err != nil {
		panic("Failed to query into the database!")
	} else {
		t.Logf("Successfully queried from requests")
	}

	numReqs := 0

	for rows.Next() {
		err = rows.Scan(&id, &sentFrom, &sentTo)
		if err != nil {
			panic(err)
		}
		numReqs++
	}

	fmt.Println("Number of requests in database:", numReqs)

	_, err = db.Exec("DELETE FROM requests WHERE \"sentfrom\"==\"Marcus\"")
	if err != nil {
		panic(err)
	}

	fmt.Print("Expected Request ID: ")
	fmt.Print("1")
	fmt.Print("\nActual Request ID: ")
	fmt.Println(id)
	fmt.Print("Expected User sent from: ")
	fmt.Print("Marcus")
	fmt.Print("\nActual User sent from: ")
	fmt.Println(sentFrom)
	fmt.Print("Expected User sent to: ")
	fmt.Print("Jon")
	fmt.Print("\nActual User sent to: ")
	fmt.Println(sentTo)

	if id != 1 {
		panic("Request Id's do not match!")
	}
	if sentFrom != "Marcus" {
		panic("Sent From User does not match!")
	}
	if sentTo != "Jon" {
		panic("Sent To User does not match!")
	}

	rows.Close()

	db.Close()
}

func TestGetFriends(t *testing.T) {
	fmt.Println("START GET FRIENDS TEST")
	const file string = "test.db"

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Open database succuss!")

	rows, err := db.Query("SELECT * FROM friends")
	if err != nil {
		panic("Failed to query into the database!")
	} else {
		t.Logf("Successfully queried from friends")
	}

	//Number of friend connections
	numConns := 0

	expectedId := 0
	expectedUser1 := "tester1"
	expectedUser2 := "tester2"

	var id int64
	var user1 string
	var user2 string

	for rows.Next() {

		err = rows.Scan(&id, &user1, &user2)

		if err != nil {
			panic("Failed to iterate through the database! %w")
		}

		numConns++
	}

	fmt.Println("Number of friend connections in database:", numConns)

	fmt.Print("Expected ID: ")
	fmt.Print(expectedId)
	fmt.Print("\nActual ID: ")
	fmt.Println(id)
	fmt.Print("Expected User Name 1: ")
	fmt.Print(expectedUser1)
	fmt.Print("\nActual User Name 1: ")
	fmt.Println(user1)
	fmt.Print("Expected User Name 2: ")
	fmt.Print(expectedUser2)
	fmt.Print("\nActual User Name 2: ")
	fmt.Println(user2)

	if id != 0 {
		panic("Request Id's do not match!")
	}
	if expectedUser1 != "tester1" {
		panic("User 1 does not match!")
	}
	if expectedUser2 != "tester2" {
		panic("User 2 does not match!")
	}

	rows.Close()

	db.Close()
}

func TestPostFriend(t *testing.T) {
	fmt.Println("START POST FRIENDS TEST")
	const file string = "test.db"

	var expectedId int64 = 1
	var expectedUser1 string = "Marcus"
	var expectedUser2 string = "Jon"

	var id int64
	var user1 string
	var user2 string

	//Number of friend connections in database
	numConns := 0

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Open database succuss!")

	_, err = db.Exec("INSERT INTO friends VALUES(1, \"Marcus\", \"Jon\")")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM friends")
	if err != nil {
		panic("Failed to query into the database!")
	} else {
		t.Logf("Successfully queried from requests")
	}

	for rows.Next() {
		err = rows.Scan(&id, &user1, &user2)
		if err != nil {
			panic(err)
		}
		numConns++
	}

	fmt.Println("Number of requests in database:", numConns)

	_, err = db.Exec("DELETE FROM friends WHERE \"user1\"==\"Marcus\"")
	if err != nil {
		panic(err)
	}

	fmt.Print("Expected Request ID: ")
	fmt.Print("1")
	fmt.Print("\nActual Request ID: ")
	fmt.Println(id)
	fmt.Print("Expected User 1: ")
	fmt.Print("Marcus")
	fmt.Print("\nActual User 1: ")
	fmt.Println(user1)
	fmt.Print("Expected User 2: ")
	fmt.Print("Jon")
	fmt.Print("\nActual User 2: ")
	fmt.Println(user2)

	if expectedId != 1 {
		panic("Request Id's do not match!")
	}
	if expectedUser1 != "Marcus" {
		panic("Sent From User does not match!")
	}
	if expectedUser2 != "Jon" {
		panic("Sent To User does not match!")
	}

	rows.Close()

	db.Close()
}
