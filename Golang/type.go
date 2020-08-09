package main

import (
	"github.com/jinzhu/gorm"
)

type Phone struct {
	gorm.Model
	Name string `gorm:"type:varchar(100)" json:"name" binding:"required"`
	Price string `gorm:"type:varchar(100)" json:"price" binding:"required"` 
}