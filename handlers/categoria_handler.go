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
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, categorias)
}

func (h *CategoriaHandler) GetById(c *gin.Context) {
	id, err := utils.GetIDParam(c)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.MsgInvalidID)
		return
	}

	categoria, err := h.service.GetById(id)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, categoria)
}

func (h *CategoriaHandler) Create(c *gin.Context) {
	var categoria models.Categoria
	if err := c.ShouldBindJSON(&categoria); err != nil {
		utils.RespondWithValidationError(c, err)
		return
	}

	if err := h.service.Create(&categoria); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, categoria)
}

func (h *CategoriaHandler) Update(c *gin.Context) {
	id, err := utils.GetIDParam(c)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.MsgInvalidID)
		return
	}

	var categoria models.Categoria
	if err := c.ShouldBindJSON(&categoria); err != nil {
		utils.RespondWithValidationError(c, err)
		return
	}

	if err := h.service.Update(id, &categoria); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, categoria)
}

func (h *CategoriaHandler) Delete(c *gin.Context) {
	id, err := utils.GetIDParam(c)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.MsgInvalidID)
		return
	}

	if err := h.service.Delete(id); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
