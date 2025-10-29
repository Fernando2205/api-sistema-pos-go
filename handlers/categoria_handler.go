package handlers

import (
	"net/http"
	"sistema_pos_go/services"

	"github.com/gin-gonic/gin"
)

type CategoriaHandler struct {
	service *services.CategoriaService
}

func NewCategoriaHandler(service *services.CategoriaService) *CategoriaHandler {
	return &CategoriaHandler{service: service}
}

func (h *CategoriaHandler) GetAll(c *gin.Context) {
	categorias, err := h.service.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, categorias)
}
