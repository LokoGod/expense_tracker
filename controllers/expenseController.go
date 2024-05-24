package controllers

import (
	"github.com/LokoGod/expense_tracker/initializers"
	"github.com/LokoGod/expense_tracker/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func FetchAllExpenseRecords(c *gin.Context) {
	// Get expenseRecords with associated Budget loaded
	var expenseRecords []models.Expense
	initializers.DB.Preload("Budget").Find(&expenseRecords)

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
		BudgetID     uint
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the expense record
	expenseRecord := models.Expense{
		ExpenseTitle: body.ExpenseTitle,
		Amount:       body.Amount,
		Recurring:    body.Recurring,
		BudgetID:     body.BudgetID,
	}

	// Add the expense record to the database
	result := initializers.DB.Create(&expenseRecord)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Return the data
	c.JSON(http.StatusCreated, gin.H{"added": expenseRecord})
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

// UpdateSpecificExpenseRecord Need to update this function
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

// DeleteSpecificExpenseRecord This deleteFunc kinda does a soft-delete & does not show in fetch req
func DeleteSpecificExpenseRecord(c *gin.Context) {
	//	Get ID from Endpoint
	id := c.Param("id")

	//	Delete record
	initializers.DB.Delete(&models.Expense{}, id)

	//	Respond data
	c.Status(200)
}
