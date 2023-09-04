package utils

import (
	"github.com/gin-contrib/cors"
)

func DefaultCors() cors.Config {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "id")
	config.AllowHeaders = append(config.AllowHeaders, "Access-Control-Allow-Origin")
	config.AllowAllOrigins = true
	// To be able to send tokens to the server.
	config.AllowCredentials = true

	// OPTIONS method for ReactJS
	config.AddAllowMethods("OPTIONS")
	return config
}
