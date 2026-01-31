package models

import (
	"gorm.io/gorm"
	"time"
)

type EventCategory struct {
	gorm.Model
	Category string `gorm:"type:varchar(100);not null"`
}

type Event struct {
	gorm.Model
	Title           string `gorm:"type:varchar(200);not null"`
	Description     string `gorm:"type:text"`
	When            string `gorm:"type:varchar(200);not null"`
	Where           string `gorm:"type:varchar(200);not null"`
	AlwaysVisible   bool
	OnAir           time.Time
	OffAir          time.Time
	EventCategoryID uint
	EventCategory   EventCategory `gorm:"foreignKey:EventCategoryID"`
}
