package main

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
	//"github.com/techschool/simplebank/util"
)

const (
	dbDriver = "sqlite3"
	dbSource = "test.db"

	//dbDriver = "postgres"
	//dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslm"
)

var testQueries *Queries

//var testDB *sql.DB

func TestMain(m *testing.M) {
	/*conn, err := run(m)
	  //conn, err := sql.Open("sqlite3", file)
		if err != nil {
			fmt.Errorf("Cannot connect to database: ", err)
		}*/

	fmt.Println("START MAIN TEST")

	/*config, err := util.LoadConfig("../..")
		if err != nil {
			panic(err)
			//panic("cannot load config:")
		}

		testDB, err = sql.Open(config.DBDriver, config.DBSource)
		if err != nil {
			//panic(err)
			panic("cannot connect to db:")
		}
	  testQueries = New(testDB)*/

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		//panic(err)
		panic("cannot connect to db:")
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}

/*func TestGetAccounts(t *testing.T) {

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
		//panic(err)
		panic("Failed to create database!")
	}

	//Insert
	res, err := db.Exec("INSERT INTO account(id, username, password, email) values(?, ?, ?, ?)", 1, u, p, e)
	if err != nil {
		//panic(err)
		panic("Failed to insert info into database!")
	}

	var ID int64
	ID, err = res.LastInsertId()
	if err != nil {
		//panic(err)
		panic("Failed to insert info into database!")
	}

	fmt.Println("Number of accounts: %w", ID)

	// update

	newUser := "MarcusLugrand"

	res, err = db.Exec("update account set username=? where id=?", newUser, ID)
	if err != nil {
		//panic(err)
		panic("Failed to insert info into database!")
	} else {
		t.Logf("Updated username successfully from Marcus to MarcusLugrand")
	}

	affect, err := res.RowsAffected()
	if err != nil {
		//panic(err)
		panic("Failed to check rows affected after delete!")
	} else {
		t.Logf("Successfully show what row is affected, which is: ")
	}

	fmt.Println(affect)

	//Query
	rows, err := db.Query("SELECT * FROM account")
	if err != nil {
		panic("Failed to query into the database!")
	} else {
		t.Logf("Successfully queried from account")
	}

	var scanCheck bool

	for rows.Next() {
		var id int64
		var username string
		var password string
		var email string

		err = rows.Scan(&id, &username, &password, &email)

		if err != nil {
			//panic(err)
			panic("Failed to iterate through the database!")
		}

		if username == newUser && p == password && e == email {
			t.Logf("User %s Found!", username)
			scanCheck = true
			break
		}
	}
	rows.Close()

	if scanCheck == false {
		panic("The user was not found!")
	}

	// delete
	res, err = db.Exec("delete from account where id=?", ID)
	if err != nil {
		//panic(err)
		panic("Failed to delete!")
	} else {
		t.Logf("Successfully delete a profile using ID")
	}

	affect, err = res.RowsAffected()
	if err != nil {
		//panic(err)
		panic("Failed to check rows affected after delete!")
	} else {
		t.Logf("Successfully show what row is affected, which is: ")
	}

	fmt.Println(affect)

	db.Close()
}

func TestPostAccount(t *testing.T) {

	//in :=

	//TRIAL
	t.Fatal("not implemented")

	/*tests := []Tests {
	  name: "Username",
	  pass: "Password"
	}
}*/
