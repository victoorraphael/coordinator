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
	handlers.RegisterAuthRoutes(s, public)

	private := r.Group("/", AuthMiddleware(), SessionMiddleware())

	// address Routes
	handlers.RegisterAddressRoutes(s, private)

	// student Routes
	handlers.RegisterStudentRoutes(s, private)

	// school Routes
	handlers.RegisterSchoolRoutes(s, private)

	// subject Routes
	handlers.RegisterSubjectRoutes(s, private)

	return r
}
