package models

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	ExpenseTitle string
	Amount int
	Recurring bool
}