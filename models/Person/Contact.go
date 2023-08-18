package models

import (
	"time"
)

type Contact struct {
	ID        uint       `json:"ID" gorm:"primaryKey"`
	Type      string     `json:"Type" gorm:"type:varchar(255)"`
	Value     string     `json:"Value" gorm:"value"`
	PersonID  uint       `json:"PersonID" gorm:"person_id"`
	IsActive  bool       `json:"IsActive" gorm:"is_active"`
	CreatedAt time.Time  `json:"CreatedAt" gorm:"created_at"`
	UpdatedAt *time.Time `json:"UpdatedAt" gorm:"updated_at"`

	Person Person `json:"Person" gorm:"foreignKey:PersonID"`
}
