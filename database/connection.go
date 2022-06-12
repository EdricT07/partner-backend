package database

import (
	"github.com/EdricT07/workhours/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var dbconn string = "tester:secret@tcp(db:3306)/test"
	connection, err := gorm.Open(mysql.Open(dbconn), &gorm.Config{})

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
