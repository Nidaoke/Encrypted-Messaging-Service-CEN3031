package main

import (
	"crypto/sha256"											//Used for encrypting
	"encoding/hex"											//Encodes bytes to hex string
	"database/sql"											//SQL database
	"fmt"													//IO stuff
	"log"													//Logging
	"net/http"												//HTTP
	"os"													//Get info from machine

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB												//Reference to database, stored globally
var hasher = sha256.New()									//Hashing function, stored globally

type account struct{ 										//Account type, holds a username and password string
	//ID uint `json:"id"` 									//ig we can use username for id? soemthing about relational algebra?
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	//Database stuff
    cfg := mysql.Config{									//Configure database connection properties
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "softengproj",
    }
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())			//Open Database
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()									//Try to connect to database
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")

	//HTTP Stuff
	router := gin.Default() 								//Basic default router
	
	router.GET("/", getAccounts) 							//When getting main page, call getAccounts
	router.POST("/", postAccounts) 							//When posting to main page, call postAccounts

	router.Run("localhost:8080") 							//Run this on port 8080 locally
}

func getAccounts(con *gin.Context) {
	var accounts []account

	rows, err := db.Query("SELECT * FROM accounts") 		//Query the accounts table
	if err != nil{
		log.Fatal(fmt.Errorf("getAccounts: %v", err))
		return
	}
	
	defer rows.Close() 										//IDK what this does tbh
	for rows.Next(){										//Go through the rows, scan them into our account struct, append to list
		var acc account
		if err := rows.Scan(&acc.Username, &acc.Password); err != nil{
			log.Fatal(fmt.Errorf("getAccounts: %v", err))
			return
		}
		accounts = append(accounts, acc)
	}

	if err := rows.Err(); err != nil{
		log.Fatal(fmt.Errorf("getAccounts: %v", err))
		return
	}
	con.JSON(http.StatusOK, accounts)						//Send list as json to context
}

func postAccounts(con *gin.Context){
	var acc account

	if err := con.BindJSON(&acc); err != nil{				//Bind the curl info to an account
		log.Fatal(fmt.Errorf("postAccounts: %v", err))
		return
	}

	hasher.Reset()											//Clear hashing function
	hasher.Write([]byte(acc.Username))						//Write info to it
	acc.Username = hex.EncodeToString(hasher.Sum(nil))		//Get the encrypted value and overwrite (stored as string)

	hasher.Reset()
	hasher.Write([]byte(acc.Password))
	acc.Password = hex.EncodeToString(hasher.Sum(nil))

															//Try to insert this into our database
	result, err := db.Exec("INSERT INTO accounts (username, password) VALUES (?, ?)", acc.Username, acc.Password)
	if err != nil{
		log.Fatal(fmt.Errorf("postAccounts: %v", err))
		return
	}

	username, err := result.LastInsertId()					//This can probably be removed
	if err != nil{
		log.Fatal(fmt.Errorf("postAccounts: %v", err))
	}

	fmt.Printf("Response: %v\n", username)
	con.JSON(http.StatusCreated, acc)						//Send out a context message
}