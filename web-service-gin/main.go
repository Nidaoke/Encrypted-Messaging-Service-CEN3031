package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"crypto/sha256"
	"encoding/hex"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var hasher = sha256.New()

type account struct{
    Username string `json:"username"`
	  Password string `json:"password"`
}

func main() {
    var err error
    db, err = sql.Open("sqlite3", "./database.db")
    checkErr(err)

    //insert
    //inserter, err := db.Prepare("INSERT INTO accounts(username, password) values (?,?)")
    //checkErr(err)

    //_, err = inserter.Exec("Nido", "1234")
    //checkErr(err)

    fmt.Println("Connected!")

    router := gin.Default()
    router.GET("/", getAccounts)
    router.POST("/", postAccount)
    router.Run("localhost:9000")
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func getAccounts(con *gin.Context) {
    var accounts []account

    rows, err := db.Query("SELECT * FROM accounts") 		//Query the accounts table
    checkErr(err)

    for rows.Next(){
        var acc account
        err := rows.Scan(&acc.Username, &acc.Password)
        checkErr(err)
        accounts = append(accounts, acc)
    }
    rows.Close()

    con.JSON(http.StatusOK, accounts)
}

func postAccount(con *gin.Context){
    var acc account

    err := con.BindJSON(&acc)
    checkErr(err)

    hasher.Reset()
    hasher.Write([]byte(acc.Password))
    acc.Password = hex.EncodeToString(hasher.Sum(nil))

    _, err = db.Exec("INSERT INTO accounts (username, password) VALUES (?, ?)", acc.Username, acc.Password)
    checkErr(err)

    con.JSON(http.StatusCreated, acc)
}

//OLD POST ACCOUNTS, NEED TO WRITE NEW ONE. USE SOMETHING LIKE //insert ABOVE

/*func postAccounts(con *gin.Context){
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
  }*/
