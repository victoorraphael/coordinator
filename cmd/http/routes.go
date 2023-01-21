package main

import (
	"github.com/gin-gonic/gin"
	"github.com/victoorraphael/coordinator/internal/domain/handlers"
	"github.com/victoorraphael/coordinator/internal/domain/services"
	"net/http"
)

func routes(s *services.Services) *gin.Engine {
	r := gin.Default()
	public := r.Group("")
	public.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": true})
	})

	// auth routes
	{
		authGroup := public.Group("/auth")
		hdl := handlers.NewAuthHandler(s)
		authGroup.
			POST("/login", hdl.Login)
	}

	private := r.Group("/", AuthMiddleware())

	// address routes
	{
		addressGroup := private.Group("address")
		hdl := handlers.NewAddressHandler(s)
		addressGroup.
			GET("", hdl.Find).
			POST("", hdl.Create)
	}

	// student routes
	{
		studentGroup := private.Group("students")
		hdl := handlers.NewStudentHandler(s)
		studentGroup.
			GET("", hdl.Find).
			POST("", hdl.Create)
	}

	return r
}
