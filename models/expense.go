package models

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	Name string
}