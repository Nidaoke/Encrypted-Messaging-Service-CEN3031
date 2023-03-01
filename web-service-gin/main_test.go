package main

import (
	"database/sql"
	"fmt"
	"testing"
)

type Tests struct {
	username string
	password string
	email    string
}

func TestGetAccounts(t *testing.T) {
	const file string = "database.db"

	u := "Marcus"
	p := "Bloop1"
	e := "marcus@gmail.com"

	//db, err := sql.Open("sqlite3", file)
	db, _ := sql.Open("sqlite3", file)

	statement, _ := db.Prepare("CREATE TABLE accounts (username VARCHAR(64) NOT NULL, password VARCHAR(64) NOT NULL, email VARCHAR(64) NOT NULL, PRIMARY KEY (`username`)")

	statement.Exec()

	statement, _ = db.Prepare("INSERT INTO accounts (username, password, email) VALUES (?, ?)")

	statement.Exec(u, p, e)

	rows, _ := db.Query("SELECT username, password, email FROM accounts")

	//var id int
	var username string
	var password string
	var email string

	for rows.Next() {
		rows.Scan(&username, &password, &email)
		if u == username && p == password && e == email {
			fmt.Println("User Found!")
		} else {
			t.Fatalf("No User Found")
		}
	}

	/*if _, err := db.Exec(create); err != nil {
	    return nil, err
	  }

	  var id int64
	  if id, err = res.LastInsertId(); err != nil {
	    fmt.Print("The data base is empty.")
	  }

		fmt.Println(db)
	*/
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
