package routes

import (
	"github.com/gofiber/fiber/v2"

	swagger "github.com/arsmn/fiber-swagger/v2"
)

// SwaggerRoute func for describe group of API Docs routes.
func SwaggerRoute(r fiber.Router) {

	// Routes for GET method:
	r.Get("*", swagger.HandlerDefault) // get one user by ID
}
