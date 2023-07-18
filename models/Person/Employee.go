package models

import (
	"time"
)

type Employee struct {
	ID        uint       `gorm:"primaryKey"`
	PersonID  uint       `gorm:"person_id"`
	RoleID    uint       `gorm:"role_id"`
	Salary    float64    `gorm:"salary"`
	IsActive  bool       `gorm:"is_active"`
	CreatedAt time.Time  `gorm:"created_at"`
	UpdatedAt *time.Time `gorm:"updated_at"`

	Person Person `gorm:"foreignKey:PersonID"`
	Role   Role   `gorm:"foreignKey:RoleID"`
}
