package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	UserId                uint64             `gorm:"userId"`
	Reference             string             `gorm:"reference"  binding:"required"`
	Amount                float64            `gorm:"amount"     binding:"required"`
	Channel               uint64             `gorm:"channel"    binding:"required"`
	Item                  uint64             `gorm:"item"       binding:"required"`
	Status                string             `gorm:"status"     binding:"required"`
	Active                bool               `gorm:"active"`
}

type PaymentItem struct {
	gorm.Model
	Name                  string             `gorm:"name"        binding:"required"`
	Active                bool               `gorm:"active"`
}

type PaymentChannel struct {
	gorm.Model
	Name                  string             `gorm:"name"        binding:"required"`
	Active                bool               `gorm:"active"`
}

func CreatePayment(db *gorm.DB, newPayment *Payment) (err error) {
	err = db.Create(newPayment).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPaymentsById(db *gorm.DB, payments *Payment, paymentId uint64) (err error) {
	err = db.Order("created_at desc").First(payments).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllPayments(db *gorm.DB, payments *[]Payment) (err error) {
	err = db.Order("created_at desc").Find(payments).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllActivePayments(db *gorm.DB, payments *[]Payment) (err error) {
	err = db.Order("created_at desc").Where("active = ?", true).Find(payments).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserPayments(db *gorm.DB, payments *[]Payment, userId uint) (err error) {
	err = db.Order("created_at desc").Where("user_id = ? AND active = ?", userId, true).Find(payments).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdatePayment(db *gorm.DB, payment *Payment) (err error) {
	err = db.Save(payment).Error
	if err != nil {
		return err
	}
	return nil
}

func DeletePayment(db *gorm.DB, payment *Payment, paymentId uint) (err error) {
	err = db.Where("id = ?", paymentId).Delete(payment).Error
	if err != nil {
		return err
	}
	return nil
}