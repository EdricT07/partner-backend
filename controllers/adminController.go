package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/EdricT07/workhours/database"
	"github.com/EdricT07/workhours/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func GetPartnersDB() map[string]string {
	var partners []models.Partner
	var partnersMap map[string]string = make(map[string]string)
	database.DB.Find(&partners)

	for _, value := range partners {
		partnersMap[value.Partner] = value.Partner
	}
	return partnersMap
}

func isLoggedIn(c *fiber.Ctx) bool {
	cookie := c.Cookies("jwt")

	_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	return err == nil

}

func GetPartners(c *fiber.Ctx) error {
	var isLoggedInVar bool = isLoggedIn(c)

	if isLoggedInVar {
		return c.JSON(models.Partners)
	} else {
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}
}

func GetPartnerInfo(c *fiber.Ctx) error {
	var partner string = c.Params("bedrijf")
	partner = strings.ReplaceAll(partner, "%20", " ")

	var isLoggedInVar bool = isLoggedIn(c)

	if !isLoggedInVar {
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}
	var maandag models.EnqueteMaandag
	var vrijdag models.EnqueteVrijdag

	var week int
	var jaar int

	jaar, week = time.Now().UTC().ISOWeek()

	database.DB.Where("bedrijf = ? AND week = ? AND jaar = ?", partner, week, jaar).First(&maandag)
	database.DB.Where("bedrijf = ? AND week = ? AND jaar = ?", partner, week, jaar).First(&vrijdag)

	return c.JSON([]interface{}{maandag, vrijdag})
}

func checkForValue(userValue string, students map[string]string) bool {

	//traverse through the map
	for _, value := range students {

		//check if present value is equals to userValue
		if value == userValue {

			//if same return true
			return true
		}
	}

	//if value not found return false
	return false
}

func CreatePartner(c *fiber.Ctx) error {
	var isLoggedInVar bool = isLoggedIn(c)

	if isLoggedInVar {
		var partner models.Partner

		if err := c.BodyParser(&partner); err != nil {
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		var PartnersMap map[string]string = GetPartnersDB()
		fmt.Println(PartnersMap)
		if checkForValue(partner.Partner, PartnersMap) {
			c.JSON(fiber.Map{
				"message": "Partner bestaat al",
			})
		}

		database.DB.Create(&partner)
		return c.JSON(fiber.Map{
			"message": "success",
		})
	} else {
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

}

func IsLoggedInRoute(c *fiber.Ctx) error {
	var isLoggedIn bool = isLoggedIn(c)

	if isLoggedIn {
		return c.JSON(fiber.Map{
			"message": "success",
		})
	} else {
		return c.JSON(fiber.Map{
			"message": "failed",
		})
	}
}
