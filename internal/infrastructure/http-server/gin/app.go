// Package server wires up the Gin framework dependencies
package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type appControllers struct {
}

// AppControllersDependencies lists up the controller dependencies, mostly repositories
type AppControllersDependencies struct {
}

// NewAppControllers is the AppController factory function
func NewAppControllers(deps AppControllersDependencies) appControllers {
	return appControllers{}
}

// NewGinApp sets up the Gin engine
func NewGinApp(controllers appControllers) *gin.Engine {
	app := gin.Default()

	app.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"date": time.Now().String()})
	})

	return app
}
