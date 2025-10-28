//Gin is a lightweight, fast, and easy-to-use HTTP web framework built on top of net/http.
//It helps you build APIs and web applications more easily.

// IT has features like :

// Routing (URL to handler mapping)
// Middleware (logging, authentication, etc.)
// JSON handling
// Error handling
// Query/form parsing
// Rendering templates

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//create a gin router with default middleware (logger and recovery)
	r := gin.Default()

	// define a simple GET endpoint

	r.GET("/ping", func(c *gin.Context) {
		// return json response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//Start the server on port 8080(default)
	//server will listen on 0.0.0.0:8080 (localhost:8080 on windows)

	r.Run()

}

// What this example demonstrates:

// Creating a Gin router with default middleware
// Defining HTTP endpoints with simple handler functions
// Returning JSON responses
// Starting an HTTP server
