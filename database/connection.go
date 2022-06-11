package database

import (
	"os"

	"github.com/EdricT07/workhours/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var dbconn string = os.Getenv("DATABASE_URL")
	connection, err := gorm.Open(postgres.Open(dbconn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection
	connection.AutoMigrate(&models.Company{})
	connection.AutoMigrate(&models.EnqueteMaandag{})
	connection.AutoMigrate(&models.EnqueteVrijdag{})
	connection.AutoMigrate(&models.User{})
	connection.AutoMigrate(&models.Partner{})

}
