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

type studentHandler struct {
	srv *services.Services
}

func RegisterStudentRoutes(srv *services.Services, router *gin.RouterGroup) {
	hdl := &studentHandler{srv}
	studentGroup := router.Group("students")
	studentGroup.
		GET("", hdl.Find).
		POST("", hdl.Create)
}

func (s *studentHandler) Find(c *gin.Context) {
	list, err := s.srv.Person.FetchAll(c, entities.PersonStudent)
	if err != nil {
		chatty.Errorf("falha ao buscar estudantes: err: %v", err)
		c.String(http.StatusInternalServerError, "error ao buscar estudantes")
		return
	}

	c.JSON(http.StatusOK, list)
}

func (s *studentHandler) Create(c *gin.Context) {
	var req entities.CreateStudent
	if err := c.Bind(&req); err != nil {
		chatty.Errorf("error ao fazer unmarshal de estudante: err: %v", err)
		c.String(http.StatusBadRequest, "campos inválidos")
		return
	}

	if err := req.Validate(); err != nil {
		chatty.Errorf("error ao criar estudante: %v", err.Error())
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	per, err := s.srv.Person.Search(c, entities.Person{Email: req.Email})
	if err != nil {
		chatty.Errorf("error ao criar estudante: %v", err.Error())
		c.String(http.StatusBadRequest, errs.WrapError(errs.ErrInternalError, "nao foi possivel verificar se email ja existe").Error())
		return
	}

	if per.UUID != "" {
		chatty.Error("error ao criar estudante: email ja cadastrado")
		c.String(http.StatusBadRequest, errs.WrapError(errs.ErrFieldViolation, "email ja cadastrado ").Error())
		return
	}

	if req.AddressID == "" {
		chatty.Error("error ao criar estudante: endereco vazio")
		c.String(http.StatusBadRequest, errs.WrapError(errs.ErrFieldViolation, "endereco id vazio").Error())
		return
	}

	addr, err := s.srv.Address.Find(c, entities.Address{UUID: req.AddressID})
	if err != nil {
		chatty.Errorf("error ao criar estudante: %v", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	p := entities.Person{
		UUID:      uid.NewUUID().String(),
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		Birthdate: req.Birthdate,
		AddressID: addr.ID,
	}
	_, err = s.srv.Person.Create(c, p)
	if err != nil {
		chatty.Errorf("error ao criar estudante: err: %v", err)
		c.String(http.StatusBadRequest, "não foi possível criar estudante")
		return
	}

	c.JSON(http.StatusCreated, p)
}
