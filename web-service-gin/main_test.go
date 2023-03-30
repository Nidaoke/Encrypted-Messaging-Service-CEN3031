package main

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestGetMessages1Count(t *testing.T){
  db, err := sql.Open("sqlite3", "./database.db"); if(err!=nil){panic(err)};
  buffer := "";

  //Get Message id tracker
  tracker := 0;
  row, err := db.Query("SELECT * FROM sqlite_sequence"); if(err!=nil){panic(err)};
  for row.Next(){err = row.Scan(&buffer, &tracker); if(err!=nil){panic(err)}};

  //Add 3 debug messages
  _, err = db.Exec("INSERT INTO messages VALUES(null, \"debug1\", \"debug2\", \"Hi\")"); if(err!=nil){panic(err)};
  _, err = db.Exec("INSERT INTO messages VALUES(null, \"debug1\", \"debug2\", \"Hi2\")"); if(err!=nil){panic(err)};
  _, err = db.Exec("INSERT INTO messages VALUES(null, \"debug2\", \"debug1\", \"Hi back\")"); if(err!=nil){panic(err)};

  //Test
  count := 0;
  rows, err := db.Query("SELECT * FROM messages WHERE (\"from\"==\"debug1\" OR \"to\"==\"debug1\")"); if(err!=nil){panic(err)};
  for rows.Next(){
    count += 1;
  }

  //Go back to before
  _, err = db.Exec("DELETE FROM messages WHERE (\"from\"==\"debug1\" OR \"to\"==\"debug1\")"); if(err!=nil){panic(err)};
  _, err = db.Exec("DELETE FROM sqlite_sequence WHERE \"name\"==\"messages\""); if(err!=nil){panic(err)};
  _, err = db.Exec("INSERT INTO sqlite_sequence VALUES(\"messages\", ?)", tracker); if(err!=nil){panic(err)};

  fmt.Print("Expected Count: 3");
  fmt.Print("\nActual Count: ");
  fmt.Println(count);

  if(count != 3){
    panic("Counts are not equal!");
  }
}

func TestGetMessages1Content(t *testing.T){
  db, err := sql.Open("sqlite3", "./database.db"); if(err!=nil){panic(err)};
  buffer := "";
  buffer2 := 0;

  //Get Message id tracker
  tracker := 0;
  row, err := db.Query("SELECT * FROM sqlite_sequence"); if(err!=nil){panic(err)};
  for row.Next(){err = row.Scan(&buffer, &tracker); if(err!=nil){panic(err)}};

  //Add debug messages
  _, err = db.Exec("INSERT INTO messages VALUES(null, \"debug1\", \"debug2\", \"Hi\")"); if(err!=nil){panic(err)};

  //Test
  mes := "";
  rows, err := db.Query("SELECT * FROM messages WHERE (\"from\"==\"debug1\" OR \"to\"==\"debug1\")"); if(err!=nil){panic(err)};
  for rows.Next(){
    err = rows.Scan(&buffer2, &buffer, &buffer, &mes); if(err!=nil){panic(err)};
  }

  //Go back to before
  _, err = db.Exec("DELETE FROM messages WHERE (\"from\"==\"debug1\" OR \"to\"==\"debug1\")"); if(err!=nil){panic(err)};
  _, err = db.Exec("DELETE FROM sqlite_sequence WHERE \"name\"==\"messages\""); if(err!=nil){panic(err)};
  _, err = db.Exec("INSERT INTO sqlite_sequence VALUES(\"messages\", ?)", tracker); if(err!=nil){panic(err)};

  fmt.Print("Expected Message: Hi");
  fmt.Print("\nActual Message: ");
  fmt.Println(mes);

  if(mes != "Hi"){
    panic("Messages are not equal!");
  }
}

func TestGetMessages2Count(t *testing.T){
  db, err := sql.Open("sqlite3", "./database.db"); if(err!=nil){panic(err)};
  buffer := "";

  //Get Message id tracker
  tracker := 0;
  row, err := db.Query("SELECT * FROM sqlite_sequence"); if(err!=nil){panic(err)};
  for row.Next(){err = row.Scan(&buffer, &tracker); if(err!=nil){panic(err)}};

  //Add 3 debug messages
  _, err = db.Exec("INSERT INTO messages VALUES(null, \"debug1\", \"debug2\", \"Hi\")"); if(err!=nil){panic(err)};
  _, err = db.Exec("INSERT INTO messages VALUES(null, \"debug1\", \"debug2\", \"Hi2\")"); if(err!=nil){panic(err)};
  _, err = db.Exec("INSERT INTO messages VALUES(null, \"debug2\", \"debug1\", \"Hi back\")"); if(err!=nil){panic(err)};

  //Test
  count := 0;
  rows, err := db.Query("SELECT * FROM messages WHERE ((\"from\"==\"debug1\" AND \"to\"==\"debug2\") OR (\"from\"==\"debug2\" AND \"to\"==\"debug1\"))"); if(err!=nil){panic(err)};
  for rows.Next(){
    count += 1;
  }

  //Go back to before
  _, err = db.Exec("DELETE FROM messages WHERE ((\"from\"==\"debug1\" AND \"to\"==\"debug2\") OR (\"from\"==\"debug2\" AND \"to\"==\"debug1\"))"); if(err!=nil){panic(err)};
  _, err = db.Exec("DELETE FROM sqlite_sequence WHERE \"name\"==\"messages\""); if(err!=nil){panic(err)};
  _, err = db.Exec("INSERT INTO sqlite_sequence VALUES(\"messages\", ?)", tracker); if(err!=nil){panic(err)};

  fmt.Print("Expected Count: 3");
  fmt.Print("\nActual Count: ");
  fmt.Println(count);

  if(count != 3){
    panic("Counts are not equal!");
  }
}

func TestGetMessages2Content(t *testing.T){
  db, err := sql.Open("sqlite3", "./database.db"); if(err!=nil){panic(err)};
  buffer := "";
  buffer2 := 0;

  //Get Message id tracker
  tracker := 0;
  row, err := db.Query("SELECT * FROM sqlite_sequence"); if(err!=nil){panic(err)};
  for row.Next(){err = row.Scan(&buffer, &tracker); if(err!=nil){panic(err)}};

  //Add debug messages
  _, err = db.Exec("INSERT INTO messages VALUES(null, \"debug1\", \"debug2\", \"Hi\")"); if(err!=nil){panic(err)};

  //Test
  mes := "";
  rows, err := db.Query("SELECT * FROM messages WHERE ((\"from\"==\"debug1\" AND \"to\"==\"debug2\") OR (\"from\"==\"debug2\" AND \"to\"==\"debug1\"))"); if(err!=nil){panic(err)};
  for rows.Next(){
    err = rows.Scan(&buffer2, &buffer, &buffer, &mes); if(err!=nil){panic(err)};
  }

  //Go back to before
  _, err = db.Exec("DELETE FROM messages WHERE ((\"from\"==\"debug1\" AND \"to\"==\"debug2\") OR (\"from\"==\"debug2\" AND \"to\"==\"debug1\"))"); if(err!=nil){panic(err)};
  _, err = db.Exec("DELETE FROM sqlite_sequence WHERE \"name\"==\"messages\""); if(err!=nil){panic(err)};
  _, err = db.Exec("INSERT INTO sqlite_sequence VALUES(\"messages\", ?)", tracker); if(err!=nil){panic(err)};

  fmt.Print("Expected Message: Hi");
  fmt.Print("\nActual Message: ");
  fmt.Println(mes);

  if(mes != "Hi"){
    panic("Messages are not equal!");
  }
}

func TestGetMessagesCount(t *testing.T){
  db, err := sql.Open("sqlite3", "./database.db");
  if(err!=nil) { panic(err) };
  row, err := db.Query("SELECT * FROM messages ORDER BY id DESC LIMIT 1");
  if(err!=nil) { panic(err) };
  //fmt.Println(row);
  expectedCount := 0;
  buffer := "";
  for row.Next(){
    err = row.Scan(&expectedCount, &buffer, &buffer, &buffer);
    if(err!=nil) { panic(err)};
  };
  expectedCount += 1;

  actualCount := 0;
  rows, err := db.Query("SELECT * FROM messages");
  if(err!=nil) { panic(err) };
  for rows.Next(){
    actualCount += 1;
  }

  fmt.Print("Expected Count: ");
  fmt.Print(expectedCount);
  fmt.Print("\nActual Count: ");
  fmt.Println(actualCount);

  if(actualCount != expectedCount){
    panic("Counts are not equal");
  }
}

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

/*func TestPostAccount(t *testing.T) {

	//in :=

	//TRIAL
	t.Fatal("not implemented")

	/*tests := []Tests {
	  name: "Username",
	  pass: "Password"
	}
}*/
