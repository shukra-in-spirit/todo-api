package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-api/internal/pkg/constants"
)

func main() {
	// router := gin.Default()
	// This will create an engine (router) with default middleware
	// logger and recovery (crash-free) middleware enabled
	// fastest and easiest way to get running

	// using gin.New() allows us to create a blank router
	router := gin.New()
	// default configuration for the router is:
	// RedirectTrailingSlash: true, RedirectFixedPath: false
	// HandleMethodNotAllowed: false, ForwardedByClientIP: true
	// UseRawPath: false, UnescapePathValues: true

	// We will be creating the endpoints next
	// The initial endpoints are status and version
	router.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, constants.TodoStatusHealthy)
		// returns the status healthy string
		// "Status: healthy"
	})
	router.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, "Version: %s \n", constants.TodoApiVersion)
	})
	// in the GET method, we are passing a Method Expression

	// The following command will run the server
	// default port value is 8080
	router.Run(constants.TodoPortNumber)
}
