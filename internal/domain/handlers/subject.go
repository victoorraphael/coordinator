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

type subjectHandler struct {
	srv *services.Services
}

func (h *subjectHandler) FindUUID(c *gin.Context) {
	key := c.Param("uuid")
	if key == "" {
		chatty.Infof("falha o buscar subject com uuid = %v: err: uuid vazio", key)
		c.String(http.StatusBadRequest, errs.WrapError(errs.ErrInternalError, "uuid nao pode ser vazio").Error())
		return
	}

	found, err := h.srv.Subject.Find(c, entities.Subject{UUID: key})
	if err != nil {
		chatty.Infof("falha o buscar subject: err: %v", err)
		c.String(http.StatusBadRequest, errs.WrapError(errs.ErrInternalError, "falha ao buscar").Error())
		return
	}

	c.JSON(http.StatusOK, found)
}

func (h *subjectHandler) Create(c *gin.Context) {
	var req entities.Subject
	if err := c.Bind(&req); err != nil {
		chatty.Infof("wrong request payload %v", req)
		c.String(http.StatusBadRequest, errs.WrapError(errs.ErrFieldViolation, "falha ao criar").Error())
		return
	}

	found, err := h.srv.Subject.Find(c, req)
	if err != nil {
		chatty.Infof("falha o buscar subject %v: err: %v", req, err)
		c.String(http.StatusBadRequest, errs.WrapError(errs.ErrInternalError, "falha ao criar").Error())
		return
	}

	if found.UUID != "" {
		chatty.Infof("falha o criar subject: ja existe: err: %v", err)
		c.String(http.StatusBadRequest, errs.WrapError(errs.ErrFieldViolation, "subject ja existe").Error())
		return
	}

	req.UUID = uid.NewUUID().String()
	if err := h.srv.Subject.Create(c, &req); err != nil {
		chatty.Infof("falha o criar subject: err: %v", err)
		c.String(http.StatusBadRequest, errs.WrapError(errs.ErrInternalError, "nao foi possivel criar subject").Error())
		return
	}

	c.JSON(http.StatusOK, req)
}

func RegisterSubjectRoutes(srv *services.Services, router *gin.RouterGroup) {
	hdl := &subjectHandler{srv}
	addressGroup := router.Group("subject")
	addressGroup.
		GET("/:uuid", hdl.FindUUID).
		POST("", hdl.Create)
}
