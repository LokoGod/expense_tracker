package main

import (
	"github.com/LokoGod/expense_tracker/initializers"
	"github.com/LokoGod/expense_tracker/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// Automigration
	initializers.DB.AutoMigrate(&models.Expense{})
	initializers.DB.AutoMigrate(&models.ExpenseRelatedBudget{})
	initializers.DB.AutoMigrate(&models.Budget{})
	initializers.DB.AutoMigrate(&models.Income{})
	initializers.DB.AutoMigrate(&models.IncomeCategories{})
}
