package httphdl

import (
	"github.com/gin-gonic/gin"
	"github.com/victoorraphael/coordinator/internal/domain/handlers"
	"github.com/victoorraphael/coordinator/internal/domain/services"
	"net/http"
)

func Routes(s *services.Services, test ...bool) *gin.Engine {
	r := gin.Default()
	public := r.Group("")
	public.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": true})
	})

	// auth Routes
	{
		authGroup := public.Group("/auth")
		hdl := handlers.NewAuthHandler(s)
		authGroup.
			POST("/login", hdl.Login)
	}

	var private *gin.RouterGroup
	if len(test) > 0 {
		private = r.Group("/")
	} else {
		private = r.Group("/", AuthMiddleware())
	}

	// address Routes
	{
		addressGroup := private.Group("address")
		hdl := handlers.NewAddressHandler(s)
		addressGroup.
			GET("", hdl.Find).
			POST("", hdl.Create)
	}

	// student Routes
	{
		studentGroup := private.Group("students")
		hdl := handlers.NewStudentHandler(s)
		studentGroup.
			GET("", hdl.Find).
			POST("", hdl.Create)
	}

	return r
}
