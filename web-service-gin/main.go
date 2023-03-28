package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"crypto/sha256"
	"encoding/hex"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/thinkerou/favicon"
)

var db *sql.DB
var hasher = sha256.New()

type account struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type message struct {
	Id      int    `json:"id"`
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
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

	//Marcus Added Begin
	//DefaultConfig as a start point
	config := cors.DefaultConfig()
	config.AllowHeaders = []string{"Authorization", "content-type"}
	config.AllowOrigins = []string{"http://localhost:4200"} //Change to * at some point
	router.Use(cors.New(config))
	//End

	router.Use(favicon.New("./favicon.ico"))

	router.GET("/", getAccounts)
	router.POST("/", postAccount)
	router.GET("/messages", getMessages)
  router.GET("/messages/:name", getMessage1);
  router.GET("/messages/:name/:name2", getMessage2)
	router.POST("/messages", postMessage)
	router.Run("localhost:9000")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getAccounts(con *gin.Context) {
	var accounts []account

	rows, err := db.Query("SELECT * FROM accounts") //Query the accounts table
	checkErr(err)

	for rows.Next() {
		var acc account
		err := rows.Scan(&acc.Username, &acc.Password, &acc.Email)
		checkErr(err)
		accounts = append(accounts, acc)
	}
	rows.Close()

	con.JSON(http.StatusOK, accounts)
}

func postAccount(con *gin.Context) {
	var acc account

	err := con.BindJSON(&acc)
	checkErr(err)

	hasher.Reset()
	hasher.Write([]byte(acc.Password))
	acc.Password = hex.EncodeToString(hasher.Sum(nil))

	_, err = db.Exec("INSERT INTO accounts (username, password, email) VALUES (?, ?, ?)", acc.Username, acc.Password, acc.Email)
	checkErr(err)

	con.JSON(http.StatusCreated, acc)
}

func getMessages(con *gin.Context) {
	var messages []message

	rows, err := db.Query("SELECT * FROM messages")
	checkErr(err)

	for rows.Next() {
		var mes message
		err := rows.Scan(&mes.Id, &mes.From, &mes.To, &mes.Message)
		checkErr(err)
		messages = append(messages, mes)
	}
	rows.Close()

	con.JSON(http.StatusOK, messages)
}

func getMessage1(con *gin.Context) {
	var messages []message
  name := con.Param("name");

	rows, err := db.Query("SELECT * FROM messages WHERE (\"from\"==? OR \"to\"==?)", name, name);
	checkErr(err)

	for rows.Next() {
		var mes message
		err := rows.Scan(&mes.Id, &mes.From, &mes.To, &mes.Message)
		checkErr(err)
		messages = append(messages, mes)
	}
	rows.Close()

	con.JSON(http.StatusOK, messages)
}

func getMessage2(con *gin.Context) {
	var messages []message
  name1 := con.Param("name");
  name2 := con.Param("name2");

	rows, err := db.Query("SELECT * FROM messages WHERE ((\"from\"==? AND \"to\"==?) OR (\"from\"==? AND \"to\"==?))", name1, name2, name2, name1);
	checkErr(err)

	for rows.Next() {
		var mes message
		err := rows.Scan(&mes.Id, &mes.From, &mes.To, &mes.Message)
		checkErr(err)
		messages = append(messages, mes)
	}
	rows.Close()

	con.JSON(http.StatusOK, messages)
}

func postMessage(con *gin.Context) {
	var mes message

	err := con.BindJSON(&mes)
	checkErr(err)

	_, err = db.Exec("INSERT INTO messages VALUES (?, ?, ?, ?)", nil, mes.From, mes.To, mes.Message)
	checkErr(err)

	con.JSON(http.StatusCreated, mes)
}
