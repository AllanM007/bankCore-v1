package migrations

import (
	"github.com/AllanM007/bankCore-v1/initializers"
	"github.com/AllanM007/bankCore-v1/models"
)

func init()  {
	initializers.ConnectToDB()
	initializers.LoadEnvVariables()
}

func main()  {
	initializers.DB.AutoMigrate(&models.User{})
}