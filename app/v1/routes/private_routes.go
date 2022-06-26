package routes

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/hotrungnhan/go-fiber-template/app/v1/controllers"
	"github.com/hotrungnhan/go-fiber-template/pkg/middleware"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(r fiber.Router) {

	// Routes for POST method:
	r.Post("/book", middleware.JWTProtected(), controllers.CreateBook)           // create a new book
	r.Post("/user/sign/out", middleware.JWTProtected(), controllers.UserSignOut) // de-authorization user
	r.Post("/token/renew", middleware.JWTProtected(), controllers.RenewTokens)   // renew Access & Refresh tokens

	// Routes for PUT method:
	r.Put("/book", middleware.JWTProtected(), controllers.UpdateBook) // update one book by ID

	// Routes for DELETE method:
	r.Delete("/book", middleware.JWTProtected(), controllers.DeleteBook) // delete one book by ID
}
