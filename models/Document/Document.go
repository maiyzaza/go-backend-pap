package models

import (
	"time"
)

type Document struct {
	ID          uint       `gorm:"primaryKey"`
	Type        string     `gorm:"type:varchar(255)"` //RENT PERSON, OWNER, AGENCY, BUYER, SELLER
	DocumentUrl string     `gorm:"document_url"`
	IsActive    bool       `gorm:"is_active"`
	CreatedAt   time.Time  `gorm:"created_at"`
	UpdatedAt   *time.Time `gorm:"updated_at"`
}
