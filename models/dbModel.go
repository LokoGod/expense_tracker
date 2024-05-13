package models

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	ExpenseTitle string
	Amount int
	Recurring bool
	ExpenseCategoryID uint
}

type ExpenseCategories struct {
	gorm.Model
	ExpenseCategoryTitle string
	Expense []Expense `gorm:"foreignKey: ExpenseCategoryID"`
}

type Income struct {
	gorm.Model
	IncomeTitle string
	Amount int
	Recurring bool
	IncomeCategoryID uint
}

type IncomeCategories struct {
	gorm.Model
	IncomeCategoryTitle string
	Income []Income `gorm:"foreignKey: IncomeCategoryID"`
}