package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id uint
	AccountNo uint
	FirstName string
	LastName string
	IdNo uint
	CreatedOn time.Time
}