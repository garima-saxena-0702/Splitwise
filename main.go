package main

import (	
	"github.com/gin-gonic/gin"
	"splitwise/handlers"
)

func main() {
	//router handlers
	router := gin.Default()
	
	router.POST("/user", handlers.CreateUser)
	router.GET("/users", handlers.GetUsers)
	router.POST("/group", handlers.CreateGroup)
	router.GET("/transactions", handlers.GetTransactions)
	router.POST("/transaction", handlers.CreateTransaction)

	router.Run() // listen and serve on 0.0.0.0:8080
}