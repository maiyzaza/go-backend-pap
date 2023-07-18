package models

import (
	"time"
)

type Contact struct {
	ID        uint       `gorm:"primaryKey"`
	Type      string     `gorm:"type:varchar(255)"` // line, facebook, email, phone, passport picture, citizen picture
	Value     string     `gorm:"value"`
	PersonID  uint       `gorm:"person_id"`
	IsActive  bool       `gorm:"is_active"`
	CreatedAt time.Time  `gorm:"created_at"`
	UpdatedAt *time.Time `gorm:"updated_at"`

	Person Person `gorm:"foreignKey:PersonID"`
}
