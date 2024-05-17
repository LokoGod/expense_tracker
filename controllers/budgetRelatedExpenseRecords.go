package controllers

import (
	"github.com/LokoGod/expense_tracker/initializers"
	"github.com/LokoGod/expense_tracker/models"
	"github.com/gin-gonic/gin"
)

func FetchAllBudgetRelatedExpenseRecords(c *gin.Context) {

	id := c.Param("RelatedBudgetID")

	var relatedExpenseRecords []models.ExpenseRelatedBudget
	initializers.DB.Where("Related_Budget_ID = ?", id).Find(&relatedExpenseRecords)

	c.JSON(200, gin.H{
		"All related records": relatedExpenseRecords,
	})
}
