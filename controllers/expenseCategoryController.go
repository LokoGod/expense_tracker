package controllers

import (
	"github.com/LokoGod/expense_tracker/initializers"
	"github.com/LokoGod/expense_tracker/models"
	"github.com/gin-gonic/gin"
	"log"
)

func FetchAllExpenseCategories(c *gin.Context) {
	var expenseCategories []models.ExpenseCategories
	initializers.DB.Find(&expenseCategories)

	c.JSON(200, gin.H{
		"All Categories": expenseCategories,
	})
}

func AddExpenseCategory(c *gin.Context) {
	var body struct {
		ExpenseCategoryTitle string
	}

	if err := c.Bind(&body); err != nil {
		log.Fatalf("Error binding data: %v", err)
	}
	expenseCategories := models.ExpenseCategories{ExpenseCategoryTitle: body.ExpenseCategoryTitle}

	result := initializers.DB.Create(&expenseCategories)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(201, gin.H{
		"Added": expenseCategories,
	})
}
