package main

import (
	controllers "github.com/AllanM007/bankCore-v1/controllers/users"
	"github.com/AllanM007/bankCore-v1/initializers"
	"github.com/gin-gonic/gin"
)


func init()  {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
  r := gin.Default()
  r.GET("/users", controllers.UserRegistration)
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}