package httpHandlers

import (
	"LineGPT/pkg/controllers"
	"net/http"
)

type Handler struct {
	Controller *controllers.Controller
}

func NewHandler(c *controllers.Controller) *Handler {
	return &Handler{Controller: c}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	//health check
}
