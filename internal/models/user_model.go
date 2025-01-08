package models

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	ID        uint `gorm:"primaryKey"`
	FirstName *string
	LastName  *string
	Email     string `gorm:"unique"`
	Gender    *string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (reciver UserModel) TableName() string {
	return "users"
}
