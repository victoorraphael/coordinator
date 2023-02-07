package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/golangsugar/chatty"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/internal/domain/services"
	"github.com/victoorraphael/coordinator/pkg/errs"
	"github.com/victoorraphael/coordinator/pkg/uid"
	"net/http"
)

type schoolHandler struct {
	srv *services.Services
}

func RegisterSchoolRoutes(s *services.Services, router *gin.RouterGroup) {
	hdl := &schoolHandler{s}
	schoolGroup := router.Group("school")
	schoolGroup.POST("", hdl.Create)
}

func (s *schoolHandler) Create(c *gin.Context) {
	var req entities.CreateSchool
	if err := c.Bind(&req); err != nil {
		chatty.Errorf("falha ao criar escola: err: %v", err)
		c.String(http.StatusInternalServerError, "não foi possível criar escola")
		return
	}

	if err := req.Validate(); err != nil {
		chatty.Errorf("error ao criar escola: %v", err.Error())
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if req.AddressUUID == "" {
		chatty.Error("error ao criar escola: endereco vazio")
		c.String(http.StatusBadRequest, errs.WrapError(errs.ErrFieldViolation, "endereco id vazio").Error())
		return
	}

	addr, err := s.srv.Address.Find(c, entities.Address{UUID: req.AddressUUID})
	if err != nil {
		chatty.Errorf("error ao criar escola: %v", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	school := entities.School{
		UUID:      uid.NewUUID().String(),
		Name:      req.Name,
		AddressID: addr.ID,
	}

	_, errc := s.srv.School.Create(c, school)
	if errc != nil {
		chatty.Errorf("error ao criar escola: err: %v", err)
		c.String(http.StatusBadRequest, "não foi possível criar escola")
		return
	}
}
