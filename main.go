package main

import (
	"github.com/LokoGod/expense_tracker/controllers"
	"github.com/LokoGod/expense_tracker/initializers"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	router := gin.Default()

	router.GET("/api/v1/expense", controllers.FetchAllExpenseRecords)
	router.POST("/api/v1/expense", controllers.AddExpenseRecord)

	if err := router.Run(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
