package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB  *gorm.DB

func ConnectToDB()  {


	var err error;

	dsn := os.Getenv("DATABASE_URL")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	fmt.Println("database connection succesful", DB)

	if err != nil {
		log.Fatal(err)
	}
	
	    // defer db.Close(c.Background())

}