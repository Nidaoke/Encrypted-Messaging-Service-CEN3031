package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type account struct{ 										//Account type, holds a username and password string
	//ID uint `json:"id"` 									//ig we can use username for id? soemthing about relational algebra?
	Username string `json:"username"`
	Password string `json:"password"`
}

var accounts = []account{ 									//Accounts container, an array of type 'account's
	{Username: "user1", Password: "password123"},
	{Username: "user2", Password: "password321"},
}

func main(){ 												//Main function called with 'go run .'
	router := gin.Default() 								//Basic default router
	
	router.GET("/", getAccounts) 							//When getting main page, call getAccounts
	router.POST("/", postAccounts) 							//When posting to main page, call postAccounts

	router.Run("localhost:8080") 							//Run this on port 8080 locally
}

func getAccounts(con *gin.Context){
	con.JSON(http.StatusOK, accounts) 						//Return list of accounts as JSON data
}

func postAccounts(con *gin.Context){
	var newAccount account 									//Create a new account

	if e := con.BindJSON(&newAccount); e != nil{ 			//Bind the JSON parameter to our new account, return if failed
		return
	}

	accounts = append(accounts, newAccount) 				//Append the new account to our list
	con.JSON(http.StatusCreated, newAccount) 				//Return the http post response with our new account
}