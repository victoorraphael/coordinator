package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/victoorraphael/coordinator/internal/adapters/repository"
	"time"
)

const (
	DefaultTimeHandler = 5 * time.Second
)

var (
	studentRepository = repository.Person{}
	addressRepository = repository.Address{}
)

type Handlers interface {
	Routes(*echo.Echo)
}

type HandlerAdapter struct {
	handlers []Handlers
}

func NewHandlerAdapter() *HandlerAdapter {
	return &HandlerAdapter{
		handlers: []Handlers{
			&StudentHandler{},
			&AddressHandler{},
		},
	}
}

func (h *HandlerAdapter) AddHandler(handler Handlers) *HandlerAdapter {
	h.handlers = append(h.handlers, handler)
	return h
}

func (h *HandlerAdapter) Connect(e *echo.Echo) {
	for _, handler := range h.handlers {
		handler.Routes(e)
	}
}
