package main

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestGetAccounts(t *testing.T) {

	const file string = "test.db"
	const create string = "CREATE TABLE IF NOT EXISTS account (id INT NOT NULL PRIMARY KEY, username VARCHAR(264) NOT NULL,password VARCHAR(264) NOT NULL,email VARCHAR(264) NOT NULL);"

	//Account Sample
	u := "Marcus"
	p := "Bloop1"
	e := "marcus@gmail.com"

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic("Failed to connect to database!")
	}

	if _, err := db.Exec(create); err != nil {
		panic(err)
	}

	res, err := db.Exec("INSERT INTO account(id, username, password, email) values(?, ?, ?, ?)", 1, u, p, e)
	if err != nil {
		panic(err)
		//panic("Failed to insert info into database!")
	}

	var ID int64
	ID, err = res.LastInsertId()
	if err != nil {
		panic(err)
		//panic("Failed to insert info into database!")
	}

	fmt.Println("Number of accounts: %w", ID)

	rows, err := db.Query("SELECT * FROM account")
	//rows, err := db.Query("SELECT * username, password, email FROM account")
	if err != nil {
		panic("Failed to query into the database!")
	}

	var scanCheck bool

	for rows.Next() {
		var id int64
		var username string
		var password string
		var email string

		err = rows.Scan(&id, &username, &password, &email)

		if err != nil {
			panic(err)
			//panic("Failed to iterate through the database!")
		}

		if username == u && p == password && e == email {
			t.Logf("User %s Found!", username)
			scanCheck = true
			break
		}
	}
	rows.Close()

	if scanCheck == false {
		panic("The user was not found!")
	}
}

/*func TestPostAccount(t *testing.T) {

	//in :=

	//TRIAL
	t.Fatal("not implemented")

	/*tests := []Tests {
	  name: "Username",
	  pass: "Password"
	}
}*/
