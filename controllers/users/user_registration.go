package controllers

import (
	"fmt"
	"log"
	"time"

	"github.com/AllanM007/bankCore-v1/initializers"
	"github.com/AllanM007/bankCore-v1/models"
	"github.com/gin-gonic/gin"
)



func UserRegistration( c *gin.Context)  {

	//create a user
	user := models.User{
		Id: 101,
        AccountNo: 18,	
		FirstName: "Jinzhu",
		LastName: "Hawaii",
		IdNo: 293728,
		CreatedOn: time.Now(),
	}
		
    // pass pointer of data to Create
	result := initializers.DB.Create(&user)
	
	if result.Error != nil {
		c.JSON(500, gin.H{"Messge":"Internal Server Error While Adding User"})
		log.Fatal(result.Error)
	}
	
	fmt.Println(user.ID)// returns inserted data's primary key
	
	fmt.Println(result.RowsAffected) // returns inserted records count
	
	c.JSON(200, gin.H{"user": user, "status": "SUCCESS"})
}