package main

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestGetAccounts(t *testing.T) {
	fmt.Println("START GET ACCOUNTS TEST")

	const file string = "test.db"
	const create string = "CREATE TABLE IF NOT EXISTS account (username VARCHAR(264) NOT NULL,password VARCHAR(264) NOT NULL,email VARCHAR(264) NOT NULL, PRIMARY KEY ('username'));"

	//Account Sample
	u := "Marcus"
	p := "Bloop1"
	e := "marcus@gmail.com"

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic("Failed to connect to database!")
	}

	if _, err := db.Exec(create); err != nil {
		//panic(err)
		panic("Failed to create database!")
	}

	fmt.Println("INSERT CHECK")
	//Insert
	res, err := db.Exec("INSERT INTO accounts(username, password, email) values(?, ?, ?)", u, p, e)
	if err != nil {
		//panic(err)
		panic("Failed to insert info into database!")
	}

	fmt.Println("LAST INSERT ID CHECK")
	//Last Insert ID
	var ID int64
	ID, err = res.LastInsertId()
	if err != nil {
		//panic(err)
		panic("Failed to get last insert ID!")
	}

	fmt.Println("Number of accounts: %w", ID)

	// update
	/*newUser := "MarcusLugrand"

	res, err = db.Exec("update account set username=? where id=?", newUser, ID)
	if err != nil {
		//panic(err)
		panic("Failed to insert info into database!")
	} else {
		t.Logf("Updated username successfully from Marcus to MarcusLugrand")
	}*/

	fmt.Println("INSERT ROWS AFFECTED CHECK")
	//Rows
	affect, err := res.RowsAffected()
	if err != nil {
		//panic(err)
		panic("Failed to check rows affected after insert!")
	} else {
		t.Logf("Successfully show what row is affected, which is: %d", affect)
	}

	fmt.Println(affect)

	fmt.Println("QUERY CHECK")
	//Query
	rows, err := db.Query("SELECT * FROM accounts")
	if err != nil {
		panic("Failed to query into the database!")
	} else {
		t.Logf("Successfully queried from account")
	}

	var scanCheck bool

	fmt.Println("FOR LOOP CHECK")
	for rows.Next() {
		//var id int64
		var username string
		var password string
		var email string

		err = rows.Scan(&username, &password, &email)

		if err != nil {
			//panic(err)
			panic("Failed to iterate through the database!")
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

	fmt.Println("DELETE CHECK")
	// delete
	res, err = db.Exec("delete from accounts where username like ?", u)
	if err != nil {
		//panic(err)
		panic("Failed to delete!")
	} else {
		t.Logf("Successfully delete a profile using ID")
	}

	fmt.Println("DELETE ROWS AFFECTED CHECK")
	affect, err = res.RowsAffected()
	if err != nil {
		//panic(err)
		panic("Failed to check rows affected after delete!")
	} else {
		t.Logf("Successfully show what row is affected, which is: %d", affect)
	}

	fmt.Println(affect)

	db.Close()
}

func TestPostAccount(t *testing.T) {
	fmt.Println("START POST ACCOUNTS TEST")
	const file string = "test.db"
	const create string = "CREATE TABLE IF NOT EXISTS accounts (id INT NOT NULL PRIMARY KEY, username VARCHAR(264) NOT NULL,password VARCHAR(264) NOT NULL,email VARCHAR(264) NOT NULL);"

	//Account Sample
	u := "tester9"
	p := "B1914338c1103962d014174aa8266f9430712cae1fa5bebac25d0c60b8a23d75d"
	e := "test9%@gmail.com"

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic("Failed to connect to database!")
	}

	if _, err := db.Exec(create); err != nil {
		//panic(err)
		panic("Failed to create database!")
	}

	//Insert
	res, err := db.Exec("INSERT INTO accounts(username, password, email) values(?, ?, ?)", u, p, e)
	if err != nil {
		//panic(err)
		panic("Failed to insert info into database!")
	} else {
		t.Logf("Successfully inserted account")
	}

	var ID int64
	ID, err = res.LastInsertId()
	if err != nil {
		//panic(err)
		panic("Failed to get last insert ID!")
	}

	fmt.Println("Number of accounts: %w", ID)

	//Query
	rows, err := db.Query("SELECT * FROM accounts")
	if err != nil {
		panic("Failed to query into the database!")
	} else {
		t.Logf("Successfully queried from account")
	}

	var scanCheck bool

	for rows.Next() {
		//var id int64
		var username string
		var password string
		var email string

		err = rows.Scan(&username, &password, &email)

		if err != nil {
			//panic(err)
			panic("Failed to iterate through the database!")
		}

		if username == u && password == p && email == e {
			t.Logf("User %s Found!", username)
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
	res, err = db.Exec("delete from accounts where username like ?", u)
	if err != nil {
		panic(err)
		//panic("Failed to delete!")
	} else {
		t.Logf("Successfully deleted a profile using username")
	}

	db.Close()

}
