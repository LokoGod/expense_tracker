package main

import (
	"github.com/LokoGod/expense_tracker/controllers"
	"github.com/LokoGod/expense_tracker/initializers"
	"github.com/LokoGod/expense_tracker/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	router := gin.Default()
	// middleware
	router.Use(middleware.CORSMiddleware())

	// routing
	router.GET("/api/v1/expense", controllers.FetchAllExpenseRecords)
	router.POST("/api/v1/expense", controllers.AddExpenseRecord)
	router.GET("/api/v1/expense/:id", controllers.FetchSpecificExpenseRecord)
	router.PUT("/api/v1/expense/:id", controllers.UpdateSpecificExpenseRecord)
	router.DELETE("/api/v1/expense/:id", controllers.DeleteSpecificExpenseRecord)

	router.GET("/api/v1/budget", controllers.FetchAllBudgets)
	router.POST("/api/v1/budget", controllers.AddBudget)
	router.DELETE("/api/v1/budget/:id", controllers.DeleteBudget)

	router.GET("/api/v1/relatedExpenseRecord/:BudgetID", controllers.FetchAllBudgetRelatedExpenseRecords)
	//router.GET("/api/v1/totalRelatedRecordAmount/:RelatedBudgetID", controllers.CalTotalBudgetRelatedExpenseRecordAmount)
	//router.GET("/api/v1/calBudgetRemaining/:RelatedBudgetID", controllers.CalBudgetRemaining)

	//router.GET("/api/v1/expenseCategory", controllers.FetchAllExpenseCategories)
	//router.POST("/api/v1/expenseCategory", controllers.AddExpenseCategory)

	if err := router.Run(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
