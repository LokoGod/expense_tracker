package controllers

import (
	"github.com/LokoGod/expense_tracker/initializers"
	"github.com/LokoGod/expense_tracker/models"
	"github.com/gin-gonic/gin"
	"log"
)

func FetchAllBudgets(c *gin.Context) {
	var budgets []models.Budget
	initializers.DB.Find(&budgets)

	c.JSON(200, gin.H{
		"All Budgets": budgets,
	})
}

func AddBudget(c *gin.Context) {
	var body struct {
		BudgetTitle  string
		BudgetAmount int
		BudgetDetail string
	}

	if err := c.Bind(&body); err != nil {
		log.Fatalf("Error binding data: %v", err)
	}
	budget := models.Budget{BudgetTitle: body.BudgetTitle, BudgetAmount: body.BudgetAmount, BudgetDetail: body.BudgetDetail}

	result := initializers.DB.Create(&budget)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(201, gin.H{
		"Added": budget,
	})
}

func DeleteBudget(c *gin.Context) {
	//	Get ID from Endpoint
	id := c.Param("id")

	//	Delete record
	initializers.DB.Delete(&models.Budget{}, id)

	//	Respond data
	c.Status(200)
}
