package models

import (
	"gorm.io/gorm"
)

type EventCategory struct {
	gorm.Model
	Category string `gorm:"type:varchar(100);not null"`
}

// TODO: Event struct