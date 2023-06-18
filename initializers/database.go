package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/jcarlospontes/go-crud-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to mysql database")
	}

	fmt.Println("Conectado ao banco")

	DB.AutoMigrate(&models.Music{})
}
