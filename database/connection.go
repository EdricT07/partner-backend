package database

import (
	"os"

	"github.com/EdricT07/workhours/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var dbuser string = os.Getenv("DBUSER")
	var dbpassword string = os.Getenv("DBPASSWORD")
	var ip string = os.Getenv("IP")
	var port string = os.Getenv("PORT")
	var dbname string = os.Getenv("DBNAME")

	dsn := "host=" + ip + " user=" + dbuser + " password=" + dbpassword + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Europe/Amsterdam"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

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
