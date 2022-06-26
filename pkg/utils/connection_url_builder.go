package utils

import (
	"fmt"
	. "github.com/hotrungnhan/go-fiber-template/pkg/configs"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n string) (string, error) {
	// Define URL to connection.
	var url string

	// Switch given names.
	switch n {
	case "postgres":
		// URL for PostgreSQL connection.
		url = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			Get().DB.HOST,
			Get().DB.PORT,
			Get().DB.USERNAME,
			Get().DB.PASSWORD,
			Get().DB.NAME,
			Get().DB.SSL,
		)
	case "redis":
		// URL for Redis connection.
		url = fmt.Sprintf(
			"%s:%s",
			Get().CACHE.HOST,
			Get().CACHE.PORT,
		)
	case "fiber":
		// URL for Fiber connection.
		url = fmt.Sprintf(
			"%s:%d",
			Get().SERVER.HOST,
			Get().SERVER.PORT,
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return url, nil
}
