package handlers

import (
	"net/http"
	"sistema_pos_go/models"
	"sistema_pos_go/services"
	"sistema_pos_go/utils"

	"github.com/gin-gonic/gin"
)

type MesaHandler struct {
	service *services.MesaService
}

func NewMesaHandler(service *services.MesaService) *MesaHandler {
	return &MesaHandler{service: service}
}

func (h *MesaHandler) GetAll(c *gin.Context) {
	mesas, err := h.service.GetAll()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, mesas)
}

func (h *MesaHandler) GetById(c *gin.Context) {
	id, err := utils.GetIDParam(c)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.MsgInvalidID)
		return
	}

	mesa, err := h.service.GetById(id)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, mesa)
}

func (h *MesaHandler) Create(c *gin.Context) {
	var mesa models.Mesa
	if err := c.ShouldBindJSON(&mesa); err != nil {
		utils.RespondWithValidationError(c, err)
		return
	}

	if err := h.service.Create(&mesa); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, mesa)
}

func (h *MesaHandler) Update(c *gin.Context) {
	id, err := utils.GetIDParam(c)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.MsgInvalidID)
		return
	}

	var mesa models.Mesa
	if err := c.ShouldBindJSON(&mesa); err != nil {
		utils.RespondWithValidationError(c, err)
		return
	}

	if err := h.service.Update(id, &mesa); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, mesa)
}

func (h *MesaHandler) PartialUpdate(c *gin.Context) {
	id, err := utils.GetIDParam(c)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.MsgInvalidID)
		return
	}

	var request models.MesaPatch
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondWithValidationError(c, err)
		return
	}

	mesa, err := h.service.PartialUpdate(id, &request)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, mesa)
}

func (h *MesaHandler) Delete(c *gin.Context) {
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
