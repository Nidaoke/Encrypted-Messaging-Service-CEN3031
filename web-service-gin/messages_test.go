package main

import (
	"database/sql"
	"fmt"
	"testing"
)

type Tests struct {
}

func TestGetMessages(t *testing.T) {
	fmt.Println("START GET MESSAGES TEST")
	//allMsgs, err := getMessages(con *gin.Context)
	const file string = "test.db"

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Open database succuss!")
	//user1 := "tester1"
	//user2 := "tester2"

	rows, err := db.Query("SELECT * FROM messages")

	fmt.Println("Database query succuss!")

	numMsgs := 0

	for rows.Next() {
		var id int64
		var from string
		var to string
		var msg string

		err = rows.Scan(&id, &from, &to, &msg)

		if err != nil {
			//panic(err)
			panic("Failed to iterate through the database! %w")
		} else {
			fmt.Printf("The message is \"%s\" from %s to %s. \n", msg, from, to)
			//fmt.Println("Message from %s to %s is: %s", from, to, msg)
			//t.Logf("Message from %s to %s is: %s", from, to, msg)
		}

		numMsgs++
	}

	fmt.Println("Database search succuss!")
	fmt.Println("Number of messages:", numMsgs)

	rows.Close()

	db.Close()

}

func TestPostMessages(t *testing.T) {
	fmt.Println("START GET MESSAGE 1 TEST")
	const file string = "test.db"

	//i := "null"
	f := "tester3"
	t_ := "tester4"
	m := "This is tester3 messaging tester4"

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Open database succuss!")

	fmt.Println("INSERT CHECK")
	//Insert
	//res, err := db.Exec("INSERT INTO messages(id, from, to, message) values(?, ?, ?, ?)", i, f, t_, m)
	res, err := db.Exec("INSERT INTO messages VALUES(null, ?, ?, ?)", f, t_, m)
	if err != nil {
		panic(err)
		//panic("Failed to insert info into database!")
	}

	var ID int64
	ID, err = res.LastInsertId()
	if err != nil {
		//panic(err)
		panic("Failed to get last insert ID!")
	}

	//fmt.Println("Number of accounts: %w", ID)

	//Query
	rows, err := db.Query("SELECT * FROM messages")
	if err != nil {
		panic("Failed to query into the database!")
	} else {
		t.Logf("Successfully queried from account")
	}

	var scanCheck bool

	for rows.Next() {
		var id int64
		var from string
		var to string
		var message string

		err = rows.Scan(&id, &from, &to, &message)

		if err != nil {
			//panic(err)
			panic("Failed to iterate through the database!")
		}

		if from == f && to == t_ && message == m {
			t.Logf("Message from %s and %s was found!", from, to)
			scanCheck = true
			break
		}
	}
	rows.Close()

	if scanCheck == false {
		panic("The user was not found!")
	} else {
		t.Logf("A new user has successfully posted")
	}

	// delete
	res, err = db.Exec("delete from messages where id=?", ID)
	if err != nil {
		panic(err)
		//panic("Failed to delete!")
	} else {
		fmt.Println("Successfully deleted a message using ID")
		//t.Logf("Successfully deleted a profile using username")
	}

	db.Close()

}

func TestGetMessages1(t *testing.T) {
	fmt.Println("START GET MESSAGE 1 TEST")
	const file string = "test.db"

	name1 := "tester1"
	name2 := "tester2"

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Open database succuss!")

	rows, err := db.Query("SELECT * FROM messages")
	if err != nil {
		panic("Failed to query into the database!")
	} else {
		t.Logf("Successfully queried from messages")
	}

	numMsgs := 0

	for rows.Next() {
		var id int64
		var from string
		var to string
		var msg string

		err = rows.Scan(&id, &from, &to, &msg)

		if err != nil {
			//panic(err)
			panic("Failed to iterate through the database! %w")
		} else if from == name1 && to == name2 {
			fmt.Printf("The message is \"%s\" from %s to %s. \n", msg, from, to)
			//t.Logf("Message from %s to %s is: %s", from, to, msg)
		} else {
			continue
		}

		numMsgs++
	}

	rows.Close()

	db.Close()
}

func TestGetMessages2(t *testing.T) {
	fmt.Println("START GET MESSAGE 2 TEST")
	const file string = "test.db"

	name1 := "tester1"
	name2 := "tester2"

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Open database succuss!")

	rows, err := db.Query("SELECT * FROM messages")
	if err != nil {
		panic("Failed to query into the database!")
	} else {
		t.Logf("Successfully queried from messages")
	}

	numMsgs := 0

	for rows.Next() {
		var id int64
		var from string
		var to string
		var msg string

		err = rows.Scan(&id, &from, &to, &msg)

		if err != nil {
			//panic(err)
			panic("Failed to iterate through the database! %w")
		} else if from == name1 && to == name2 {
			//fmt.Printf("Message from %s to %s is: %s \n", from, to, msg)
			fmt.Printf("The message is \"%s\" from %s to %s. \n", msg, from, to)
			//t.Logf("Message from %s to %s is: %s", from, to, msg)
		} else {
			continue
		}

		numMsgs++
	}

	rows.Close()

	db.Close()
}
