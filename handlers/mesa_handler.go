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
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, mesas)
}

func (h *MesaHandler) GetById(c *gin.Context) {
	id, err := utils.GetIDParam(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": utils.MsgInvalidID})
		return
	}

	mesa, err := h.service.GetById(id)
	if err != nil {
		status := utils.GetHTTPStatusFromError(err)
		c.IndentedJSON(status, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, mesa)
}

func (h *MesaHandler) Create(c *gin.Context) {
	var mesa models.Mesa
	if err := c.BindJSON(&mesa); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": utils.MsgInvalidData})
		return
	}

	if err := h.service.Create(&mesa); err != nil {
		status := utils.GetHTTPStatusFromError(err)
		c.IndentedJSON(status, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, mesa)
}

func (h *MesaHandler) Update(c *gin.Context) {
	id, err := utils.GetIDParam(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": utils.MsgInvalidID})
		return
	}

	var mesa models.Mesa
	if err := c.BindJSON(&mesa); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": utils.MsgInvalidData})
		return
	}

	if err := h.service.Update(id, &mesa); err != nil {
		status := utils.GetHTTPStatusFromError(err)
		c.IndentedJSON(status, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, mesa)
}

func (h *MesaHandler) PartialUpdate(c *gin.Context) {
	id, err := utils.GetIDParam(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": utils.MsgInvalidID})
		return
	}

	var request models.MesaPatch
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": utils.MsgInvalidData})
		return
	}

	mesa, err := h.service.PartialUpdate(id, &request)
	if err != nil {
		status := utils.GetHTTPStatusFromError(err)
		c.IndentedJSON(status, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, mesa)
}

func (h *MesaHandler) Delete(c *gin.Context) {
	id, err := utils.GetIDParam(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": utils.MsgInvalidID})
		return
	}

	if err := h.service.Delete(id); err != nil {
		status := utils.GetHTTPStatusFromError(err)
		c.IndentedJSON(status, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
