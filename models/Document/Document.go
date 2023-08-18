package models

import (
	"time"
)

type Document struct {
	ID          uint       `json:"ID" gorm:"primaryKey"`
	Type        string     `json:"Type" gorm:"type:varchar(255)"`
	DocumentUrl string     `json:"DocumentUrl" gorm:"document_url"`
	IsActive    bool       `json:"IsActive" gorm:"is_active"`
	CreatedAt   time.Time  `json:"CreatedAt" gorm:"created_at"`
	UpdatedAt   *time.Time `json:"UpdatedAt" gorm:"updated_at"`
}
