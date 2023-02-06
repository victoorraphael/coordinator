package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/golangsugar/chatty"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/internal/domain/services"
	"github.com/victoorraphael/coordinator/pkg/uid"
	"net/http"
	"time"
)

type StudentHandler struct {
	srv *services.Services
}

func NewStudentHandler(s *services.Services) *StudentHandler {
	return &StudentHandler{s}
}

func (s *StudentHandler) Find(c *gin.Context) {
	list, err := s.srv.Person.FetchAll(c, entities.PersonStudent)
	if err != nil {
		chatty.Errorf("falha ao buscar estudantes: err: %v", err)
		c.String(http.StatusInternalServerError, "error ao buscar estudantes")
		return
	}

	c.JSON(http.StatusOK, list)
}

type CreateStudentRequest struct {
	Name      string           `json:"name"`
	Email     string           `json:"email"`
	Phone     string           `json:"phone"`
	Birthdate time.Time        `json:"birthdate"`
	Address   entities.Address `json:"address"`
	SchoolID  int64            `json:"school_id"`
}

func (s *StudentHandler) Create(c *gin.Context) {
	var req CreateStudentRequest
	if err := c.Bind(&req); err != nil {
		chatty.Errorf("error ao fazer unmarshal de estudante: err: %v", err)
		c.String(http.StatusBadRequest, "campos inválidos")
		return
	}

	if req.Address.UUID != "" {
		s.srv.Address
	}

	p := entities.Person{
		UUID:      uid.NewUUID().String(),
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		Birthdate: req.Birthdate,
		AddressID: req.Address.ID,
	}
	_, err := s.srv.Student.Create(ctx, p)
	if err != nil {
		chatty.Errorf("error ao criar estudante: err: %v", err)
		c.String(http.StatusBadRequest, "não foi possível criar estudante")
		return
	}

	c.JSON(http.StatusCreated, p)
}
