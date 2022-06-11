package controllers

import (
	"fmt"
	"time"

	"github.com/EdricT07/workhours/database"
	"github.com/EdricT07/workhours/models"
	"github.com/gofiber/fiber/v2"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func EnqueteMaandag(c *fiber.Ctx) error {

	var userData models.EnqueteMaandag

	if err := c.BodyParser(&userData); err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var contains bool = contains(models.Partners, userData.Bedrijf)

	fmt.Println(contains)
	// if !contains {
	// 	return c.JSON(fiber.Map{
	// 		"message": "Geen Bedrijf",
	// 	})
	// }

	userData.Jaar, userData.Week = time.Now().UTC().ISOWeek()

	var alreadySubmitted models.EnqueteMaandag

	database.DB.First(&alreadySubmitted, "bedrijf = ?", userData.Bedrijf)
	fmt.Println(alreadySubmitted.Week)
	if alreadySubmitted.Week != 0 {
		return c.JSON(fiber.Map{
			"message": "already submitted",
		})
	}

	database.DB.Create(&userData)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func EnqueteVrijdag(c *fiber.Ctx) error {
	var userData models.EnqueteVrijdag

	if err := c.BodyParser(&userData); err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var contains bool = contains(models.Partners, userData.Bedrijf)

	fmt.Println(contains)
	if !contains {
		return c.JSON(fiber.Map{
			"message": "Geen Bedrijf",
		})
	}

	var alreadySubmitted models.EnqueteVrijdag

	database.DB.First(&alreadySubmitted, "bedrijf = ?", userData.Bedrijf)
	fmt.Println(alreadySubmitted.Week)
	if alreadySubmitted.Week != 0 {
		return c.JSON(fiber.Map{
			"message": "already submitted",
		})
	}
	userData.Jaar, userData.Week = time.Now().UTC().ISOWeek()

	database.DB.Create(&userData)
	return c.JSON(fiber.Map{
		"message": "succes",
	})

}
