package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id                 uint             `gorm:"id"         binding:"required"`
	AccountNo          uint             `gorm:"accountNo"  binding:"required"`
	FirstName          string           `gorm:"firstName"  binding:"required"`
	LastName           string           `gorm:"middleName" binding:"required"`
	IdNo               uint             `gorm:"idNo"       binding:"required"`
}

func CreateUser(db *gorm.DB, user User) (err error) {
	err = db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserById(db *gorm.DB, user *User, userId uint) (err error) {
	err = db.Order("created_at desc").Where("id = ?", userId).First(user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllUsers(db *gorm.DB, user []User) (err error) {
	err = db.Order("created_at desc").Find(user).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(db *gorm.DB, user User) (err error) {
	err = db.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(db *gorm.DB, user User) (err error) {
	err = db.Delete(user).Error
	if err != nil {
		return err
	}
	return nil
}