package routes

import (
	"github.com/EdricT07/workhours/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// Auth
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Get("/user", controllers.User)
	app.Post("/logout", controllers.Logout)
	app.Get("/isloggedin", controllers.IsLoggedInRoute)

	app.Post("/enquete-maandag", controllers.EnqueteMaandag)
	app.Post("/enquete-vrijdag", controllers.EnqueteVrijdag)

	app.Get("/partners", controllers.GetPartners)
	app.Get("/getpartnerinfo/:bedrijf", controllers.GetPartnerInfo)
	app.Post("/createpartner", controllers.CreatePartner)
}
