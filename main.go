package main

import (
	shared "github.com/hotrungnhan/go-fiber-template/app/shared/routes"
	r1 "github.com/hotrungnhan/go-fiber-template/app/v1/routes"
	"github.com/hotrungnhan/go-fiber-template/pkg/configs"
	m "github.com/hotrungnhan/go-fiber-template/pkg/middleware"
	"github.com/hotrungnhan/go-fiber-template/pkg/utils"

	"github.com/gofiber/fiber/v2"

	_ "github.com/hotrungnhan/go-fiber-template/docs"
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	m.FiberMiddleware(app)   // Register Fiber's middleware for app.
	shared.SwaggerRoute(app) // Register a route for API Docs (Swagger).
	// Create routes group.A
	v1 := app.Group("/api/v1")
	{
		// Routes.
		r1.PublicRoutes(v1)  // Register a public routes for app.
		r1.PrivateRoutes(v1) // Register a private routes for app.
		r1.NotFoundRoute(v1) // Register route for 404 Error.
	}
	// Start server (with or without graceful shutdown).
	if configs.Get().STAGE == configs.DEVELOPMENT {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
