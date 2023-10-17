package payments

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/AllanM007/bankCore-v1/models"
	"github.com/AllanM007/bankCore-v1/utilities"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type Payments struct {
	DB *gorm.DB
}

func PaymentRepo(db *gorm.DB) *Payments {
 	db.AutoMigrate(
		&models.Payment{},
		&models.PaymentChannel{},
		&models.PaymentItem{},
	)
	return &Payments{DB : db}
}


type NewPayment struct {
	UserId              uint64           `json:"userId"         binding:"required"`
	Amount              float64          `json:"amount"         binding:"required"`
	PaymentItem         uint64           `json:"paymentItem"    binding:"required"`
	PaymentChannel      uint64           `json:"paymentChannel" binding:"required"`
}

func (repository *Payments)AddPayment(c *gin.Context)  {

	var newPayment NewPayment

	if err := c.BindJSON(&newPayment); err != nil  {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status":"INVALID_REQUEST", "message":"Invalid Request!!", "error": err.Error()})
		return
	}

	paymentRef := utilities.GenerateRandomString(10)

	payment := &models.Payment{
		UserId: newPayment.UserId,
		Reference: paymentRef,
	}

	err := models.CreatePayment(repository.DB, payment)
	if err != nil {
		if strings.Contains(err.(*pgconn.PgError).Message, "duplicate key value violates unique constraint") {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "Duplicate conflict while creating payment","status":"DUPLICATE_ENTITY"})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{"status":"SUCCESS", "message":"Payment created succesfully!!"})
	
}

func (repository *Payments)GetPayments(c *gin.Context)  {
	var payments []models.Payment
	var data = make([]map[string]interface{}, 0)
	
	for item := 0; item < len(payments); item++ {
		paymentItem := map[string]interface{}{
			"userId": payments[item].UserId,
			"reference": payments[item].Reference,
			"amount": payments[item].Amount,
			"createdOn": payments[item].CreatedAt,
		}	
		data = append(data, paymentItem)
	}

	c.JSON(http.StatusOK, gin.H{"status":"SUCCESS", "payload": data})
}

func (repository *Payments)GetUserPayments(c *gin.Context)  {
	userId,_ := strconv.Atoi(c.Param("id"))
	var payments []models.Payment
	var data = make([]map[string]interface{}, 0)

	err := models.GetUserPayments(repository.DB, &payments, uint(userId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message":"User payments not found", "status":"NOT_FOUND"})
			return
		}
	
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for item := 0; item < len(payments); item++ {
		paymentItem := map[string]interface{}{
			"userId": payments[item].UserId,
			"reference": payments[item].Reference,
			"amount": payments[item].Amount,
		}

		data = append(data, paymentItem)
	}
	c.JSON(http.StatusOK, gin.H{"status":"SUCCESS", "payload":data})
}

type UpdatePayment struct {
	PaymentId             uint64              `json:"paymentId"    binding:"required"`
	PaymentItem           uint64              `json:"paymentItem"`
	PaymentChannel        uint64              `json:"paymentChannel"`
	Status                string              `json:"status"`
}

func (repository *Payments)UpdatePayment(c *gin.Context)  {

	var updatePayment UpdatePayment

	if err := c.BindJSON(&updatePayment); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status":"BAD_REQUEST", "message":"Invalid request!!", "error": err.Error()})
		return
	}

	var payment models.Payment
	err := models.GetPaymentsById(repository.DB, &payment, updatePayment.PaymentId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message":"Payment not found!!", "status":"NOT_FOUND"})
			return
		}
	
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	payment.Item = updatePayment.PaymentItem
	payment.Channel = updatePayment.PaymentChannel 

	err = models.UpdatePayment(repository.DB, &payment)
	if err != nil {
		if strings.Contains(err.(*pgconn.PgError).Message, "duplicate key value violates unique constraint") {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "Duplicate conflict while updating payment!!","status":"DUPLICATE_ENTITY"})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"status":"SUCCESS", "message":"Payment updated succesfully!!"})
	
}

func (repository *Payments)DeletePayment(c *gin.Context)  {
	paymentId,_ := strconv.Atoi(c.Param("id"))

	var payment models.Payment
	err := models.DeletePayment(repository.DB, &payment, uint(paymentId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message":"Payment not found!!", "status":"NOT_FOUND"})
			return
		}
	
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status":"SUCCESS", "message":"Payment deleted succesfully!!"})
}