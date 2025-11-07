package handlers

import (
	"net/http"
	"sistema_pos_go/models"
	"sistema_pos_go/services"
	"sistema_pos_go/utils"

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

func (h *CategoriaHandler) GetById(c *gin.Context) {
	id, err := utils.GetIDParam(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": utils.MsgInvalidID})
		return
	}

	categoria, err := h.service.GetById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, categoria)
}

func (h *CategoriaHandler) Create(c *gin.Context) {
	var categoria models.Categoria
	if err := c.BindJSON(&categoria); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": utils.MsgInvalidData})
		return
	}

	if err := h.service.Create(&categoria); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, categoria)
}

func (h *CategoriaHandler) Update(c *gin.Context) {
	id, err := utils.GetIDParam(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": utils.MsgInvalidID})
		return
	}

	var categoria models.Categoria
	if err := c.BindJSON(&categoria); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": utils.MsgInvalidData})
		return
	}

	if err := h.service.Update(id, &categoria); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, categoria)
}

/*
func (h *CategoriaHandler) PartialUpdate(c *gin.Context) {
	id, err := utils.GetIDParam(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": utils.MsgInvalidID})
		return
	}

	var campos map[string]interface{}
	if err := c.BindJSON(&campos); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": utils.MsgInvalidData})
		return
	}

	categoria, err := h.service.PartialUpdate(id, campos)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, categoria)
}
*/

func (h *CategoriaHandler) Delete(c *gin.Context) {
	id, err := utils.GetIDParam(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": utils.MsgInvalidID})
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"mensaje": utils.MsgDeleted})
}
