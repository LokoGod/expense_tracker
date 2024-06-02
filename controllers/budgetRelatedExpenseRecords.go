package controllers

import (
	"github.com/LokoGod/expense_tracker/initializers"
	"github.com/LokoGod/expense_tracker/models"
	"github.com/gin-gonic/gin"
)

func FetchAllBudgetRelatedExpenseRecords(c *gin.Context) {

	id := c.Param("BudgetID")

	var relatedExpenseRecords []models.Expense
	initializers.DB.Where("Budget_ID = ?", id).Find(&relatedExpenseRecords)

	c.JSON(200, gin.H{
		"All related records": relatedExpenseRecords,
	})
}

func CalTotalBudgetRelatedExpenseRecordAmount(c *gin.Context) {
	id := c.Param("BudgetID")

	var relatedExpenseRecords []models.Expense
	initializers.DB.Where("Budget_ID = ?", id).Find(&relatedExpenseRecords)

	var totalExpenseRecordAmount int
	for _, record := range relatedExpenseRecords {
		totalExpenseRecordAmount += record.Amount
	}

	c.JSON(200, gin.H{
		"Total": totalExpenseRecordAmount,
	})
}

//func CalBudgetRemaining(c *gin.Context) {
//	id := c.Param("RelatedBudgetID")
//
//	var relatedExpenseRecords []models.ExpenseRelatedBudget
//	initializers.DB.Where("Related_Budget_ID = ?", id).Find(&relatedExpenseRecords)
//
//	var budget models.Budget
//	initializers.DB.First(&budget, id)
//
//	budgetAmount := budget.BudgetAmount
//
//	var totalExpenseRecordAmount int
//	for _, record := range relatedExpenseRecords {
//		var expense models.Expense
//		initializers.DB.First(&expense, record.RelatedExpenseID)
//		totalExpenseRecordAmount += expense.Amount
//	}
//
//	remainingBudget := budgetAmount - totalExpenseRecordAmount
//
//	c.JSON(200, gin.H{
//		"All related records": relatedExpenseRecords,
//		"Related budget":      budget,
//		"Total spent amount":  totalExpenseRecordAmount,
//		"Remaining amount":    remainingBudget,
//	})
//}
