package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title    string `gorm:"uniqueIndex"`
	WikiID   int    `gorm:"uniqueIndex"`
	IsSexual bool
}

type Link struct {
	gorm.Model
	FromID int
	ToID   int
	From       Article `gorm:"foreignKey:FromID;references:WikiID"`
	To         Article `gorm:"foreignKey:ToID;references:WikiID"`
}

type Cache struct {
	gorm.Model
	StartID int
	EndID int
	Data string
}
