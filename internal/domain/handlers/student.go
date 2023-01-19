package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/internal/domain/services"
	"log"
	"net/http"
	"time"
)

type StudentHandler struct {
	personService services.IPersonService
}

func NewStudentHandler(s *services.Services) *StudentHandler {
	return &StudentHandler{personService: s.Person}
}

func (s *StudentHandler) Find(c *gin.Context) {
	list, err := s.personService.FetchAll(entities.PersonStudent)
	if err != nil {
		log.Println("falha ao buscar estudantes: err:", err)
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
		log.Println("error ao fazer unmarshal de estudante: err:", err)
		c.String(http.StatusBadRequest, "campos inválidos")
		return
	}

	p := entities.Person{
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		Birthdate: req.Birthdate,
		AddressID: req.Address.ID,
		SchoolID:  req.SchoolID,
	}
	uid, err := s.personService.Create(p)
	if err != nil {
		log.Println("error ao criar estudante: err:", err)
		c.String(http.StatusBadRequest, "não foi possível criar estudante")
		return
	}

	p.UUID = uid

	c.JSON(http.StatusCreated, p)
}
