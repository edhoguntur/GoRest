package models

import "gorm.io/gorm"

// User data model
type User struct {
	gorm.Model
	UserName string
	Password string
	Email    string
}
