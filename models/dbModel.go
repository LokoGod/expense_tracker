package models

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	ExpenseTitle string
	Amount       int
	Recurring    bool
	BudgetID     uint
	Budget       Budget `gorm:"foreignKey:BudgetID"`
}

type Budget struct {
	gorm.Model
	BudgetTitle    string
	BudgetAmount   int
	BudgetDetail   string
	ExpenseRecords []Expense
}

type Income struct {
	gorm.Model
	IncomeTitle      string
	Amount           int
	Recurring        bool
	IncomeCategoryID uint
}

type IncomeCategories struct {
	gorm.Model
	IncomeCategoryTitle string
	Income              []Income `gorm:"foreignKey: IncomeCategoryID"`
}
