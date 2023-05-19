package initializers

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func InitializeDB(c *gin.Context)  {

	dsn := os.Getenv("DATABASE_URL")
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	fmt.Println("database connection succesful", db)

	
}