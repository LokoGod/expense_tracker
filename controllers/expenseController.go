package controllers

import (
	"github.com/LokoGod/expense_tracker/initializers"
	"github.com/LokoGod/expense_tracker/models"
	"github.com/gin-gonic/gin"
)

func AddExpenseRecord(c *gin.Context) {

	// Get data from req.body
	var body struct {
		ExpenseTitle string
		Amount       int
		Recurring    bool
	}

	c.Bind(&body)

	// Add the expense
	expenseRecord := models.Expense{ExpenseTitle: body.ExpenseTitle, Amount: body.Amount, Recurring: body.Recurring}

	result := initializers.DB.Create(&expenseRecord)

	if result.Error != nil {
		c.Status(500)
		return
	}

	// return the data
	c.JSON(201, gin.H{
		"Added": expenseRecord,
	})
}

func FetchAllExpenseRecords(c *gin.Context) {
	// Get expenseRecords
	var expenseRecords []models.Expense
	initializers.DB.Find(&expenseRecords)

	// Respond with the data
	c.JSON(200, gin.H{
		"All Records": expenseRecords,
	})

}
