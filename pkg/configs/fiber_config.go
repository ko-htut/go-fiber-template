package configs

import (
	"github.com/gofiber/fiber/v2"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig() fiber.Config {
	// Define server settings.

	// Return Fiber configuration.
	return fiber.Config{
		Prefork:     Get().SERVER.PREFORK,
		ReadTimeout: Get().SERVER.READ_TIMEOUT,
	}
}
