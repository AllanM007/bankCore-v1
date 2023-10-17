package controllers

import (
	"net/http"
	"strings"

	"github.com/AllanM007/bankCore-v1/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func UserRepo(db *gorm.DB) *User {
	db.AutoMigrate()

	return &User{DB : db}
}

type Member struct {
	Id                uint64         `json:"id"         binding:"required"`
	FirstName         string         `json:"firstName"  binding:"required"`
	MiddleName        string         `json:"middleName" binding:"required"`
	LastName          string         `json:"lastName"   binding:"required"`
	Email             string         `json:"email"      binding:"required"`
	KraPin            string         `json:"kraPin"     binding:"required"` 
	IdNo              uint64         `json:"idNo"       binding:"required"`
	Gender            string         `json:"gender"     binding:"required"`
}

func (repository *User)UserRegistration( c *gin.Context)  {

	var member Member

	if err := c.BindJSON(&member) ; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status":"BAD_REQUEST", "message":"Invalid Request!!", "error": err.Error()})
		return
	}

	newMember := &models.User{
		
	}
		
    // pass pointer of data to Create
	result := models.CreateUser(repository.DB, *newMember)
	if result != nil {

		if strings.Contains(result.(*pgconn.PgError).Message, "duplicate key value violates unique constraint") {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "Duplicate conflict while creating parent account","status":"DUPLICATE_ENTITY"})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error()})
			return
		}
	}
	
	// if result.Error != nil {
	// 	c.JSON(500, gin.H{"Messge":"Internal Server Error While Adding User"})
	// 	log.Fatal(result.Error)
	// }
		
	
	c.JSON(http.StatusCreated, gin.H{"status": "SUCCESS", "message": "User created succesfully!!"})
}