package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID        int    `gorm:"primaryKey;autoIncrement;not null;unique"`
	Nickname  string `gorm:"type:varchar(60);not null;unique"`
	FirstName *string
	LastName  *string
	Email     string `gorm:"type:varchar(80);not null;unique_index"`
	Password  string `gorm:"type:varchar(160);not null"`
}
