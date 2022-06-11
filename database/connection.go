package database

import (
	"github.com/EdricT07/workhours/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("partner:v~Smb0231@tcp(127.0.0.1:3306)/enquete?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

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
