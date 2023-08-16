package models

import (
	"gorm.io/gorm"
)

type Card struct {
	gorm.Model

	Id         int `gorm:"primaryKey;autoIncrement:true;not null;uniqueIndex"`
	UserID     uint
	Name       string `gorm:"type:varchar(60);not null"`
	CardNumber string `gorm:"type:varchar(25);not null"`
}
