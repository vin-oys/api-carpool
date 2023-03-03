package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID            uint
	Username      uint
	Password      string
	ContactNumber string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	RoleID        string
}
