package routes

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/hotrungnhan/go-fiber-template/app/v1/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(r fiber.Router) {

	// Routes for GET method:
	r.Get("/books", controllers.GetBooks)   // get list of all books
	r.Get("/book/:id", controllers.GetBook) // get one book by ID

	// Routes for POST method:
	r.Post("/user/sign/up", controllers.UserSignUp) // register a new user
	r.Post("/user/sign/in", controllers.UserSignIn) // auth, return Access & Refresh tokens
}
