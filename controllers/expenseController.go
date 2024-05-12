package controllers

import (
	"github.com/LokoGod/expense_tracker/initializers"
	"github.com/LokoGod/expense_tracker/models"
	"github.com/gin-gonic/gin"
	"log"
)

func FetchAllExpenseRecords(c *gin.Context) {
	// Get expenseRecords
	var expenseRecords []models.Expense
	initializers.DB.Find(&expenseRecords)

	// Respond with the data
	c.JSON(200, gin.H{
		"All Records": expenseRecords,
	})
}

func AddExpenseRecord(c *gin.Context) {

	// Get data from req.body
	var body struct {
		ExpenseTitle string
		Amount       int
		Recurring    bool
	}

	if err := c.Bind(&body); err != nil {
		log.Fatalf("Error binding data: %v", err)
	}

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

func FetchSpecificExpenseRecord(c *gin.Context) {
	//	Get ID from Endpoint
	id := c.Param("id")

	//	Get singleExpenseRecord
	var specificExpenseRecord models.Expense
	initializers.DB.First(&specificExpenseRecord, id)

	//	Respond with data
	c.JSON(200, gin.H{
		"Requested Record": specificExpenseRecord,
	})
}

func UpdateSpecificExpenseRecord(c *gin.Context) {
	//	Get ID from Endpoint
	id := c.Param("id")

	// Get data from req.body
	var body struct {
		ExpenseTitle string
		Amount       int
		Recurring    bool
	}

	if err := c.Bind(&body); err != nil {
		log.Fatalf("Error binding data: %v", err)
	}

	//	Find the record to be updated
	var specificExpenseRecord models.Expense
	initializers.DB.First(&specificExpenseRecord, id)

	//	Update the record
	initializers.DB.Model(&specificExpenseRecord).Updates(models.Expense{
		ExpenseTitle: body.ExpenseTitle,
		Amount:       body.Amount,
		Recurring:    body.Recurring,
	})

	//	Respond the data
	c.JSON(200, gin.H{
		"Updated Record": specificExpenseRecord,
	})

}

func DeleteSpecificExpenseRecord(c *gin.Context) {

}
