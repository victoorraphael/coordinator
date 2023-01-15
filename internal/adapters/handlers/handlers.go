package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/victoorraphael/coordinator/internal/services"
)

type Handlers interface {
	Routes(*echo.Echo, *services.Services)
}

type HandlerAdapter struct {
	handlers []Handlers
}

func NewHandlerAdapter() *HandlerAdapter {
	return &HandlerAdapter{
		handlers: []Handlers{
			//&StudentHandler{},
			&AddressHandler{},
		},
	}
}

func (h *HandlerAdapter) AddHandler(handler Handlers) *HandlerAdapter {
	h.handlers = append(h.handlers, handler)
	return h
}

func (h *HandlerAdapter) Connect(e *echo.Echo, services *services.Services) {
	for _, handler := range h.handlers {
		handler.Routes(e, services)
	}
}
