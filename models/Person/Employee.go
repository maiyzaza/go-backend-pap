package models

import (
	"time"
)

type Employee struct {
	ID        uint       `json:"ID" gorm:"primaryKey"`
	PersonID  uint       `json:"PersonID" gorm:"person_id"`
	RoleID    uint       `json:"RoleID" gorm:"role_id"`
	Salary    float64    `json:"Salary" gorm:"salary"`
	IsActive  bool       `json:"IsActive" gorm:"is_active"`
	CreatedAt time.Time  `json:"CreatedAt" gorm:"created_at"`
	UpdatedAt *time.Time `json:"UpdatedAt" gorm:"updated_at"`

	Person Person `json:"Person" gorm:"foreignKey:PersonID"`
	Role   Role   `json:"Role" gorm:"foreignKey:RoleID"`
}
