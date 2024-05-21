package models

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	ExpenseTitle  string
	Amount        int
	Recurring     bool
	RelatedBudget []ExpenseRelatedBudget `gorm:"foreignKey: RelatedBudgetID"`
}

type ExpenseRelatedBudget struct {
	ID               uint `gorm:"primarykey"`
	RelatedBudgetID  uint
	RelatedExpenseID uint
}

type Budget struct {
	gorm.Model
	BudgetTitle    string
	BudgetAmount   int
	BudgetDetail   string
	RelatedExpense []ExpenseRelatedBudget `gorm:"foreignKey: RelatedExpenseID"`
}

//type ExpenseCategories struct {
//	ID                   uint `gorm:"primarykey"`
//	ExpenseCategoryTitle string
//	Expense              []Expense `gorm:"foreignKey: ExpenseCategoryID"`
//}

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
