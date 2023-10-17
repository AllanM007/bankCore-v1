package main

import (
	"fmt"
	"time"

	controllers "github.com/AllanM007/bankCore-v1/controllers/users"
	"github.com/AllanM007/bankCore-v1/initializers"
	"github.com/gin-gonic/gin"
)

// func init()  {
// 	initializers.LoadEnvVariables()
// 	initializers.ConnectToDB()
// }

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    }
}

func main() {
  r := gin.Default()

  loc, err := time.LoadLocation("Africa/Nairobi")
	if err != nil {
		fmt.Println(err)
	}
    // handle err
    time.Local = loc // -> this is setting the global timezone

  db := initializers.InitDb()
  r.Use(CORSMiddleware())
  // Use the custom error handler middleware
  // r.Use(middleware.ErrorHandler())
  
  UserRepo := controllers.UserRepo(db)
  //PaymentsRepo := controllers.PaymentRepo(db)

  r.POST("/registerUser", UserRepo.UserRegistration)
  r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}