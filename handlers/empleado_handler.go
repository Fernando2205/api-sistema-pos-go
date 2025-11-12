package handlers

import (
	"net/http"
	"sistema_pos_go/models"
	"sistema_pos_go/services"
	"sistema_pos_go/utils"

	"github.com/gin-gonic/gin"
)

type EmpleadoHandler struct {
	service *services.EmpleadoService
}

func NewEmpleadoHandler(service *services.EmpleadoService) *EmpleadoHandler {
	return &EmpleadoHandler{service: service}
}

func (h *EmpleadoHandler) GetAll(c *gin.Context) {
	empleados, err := h.service.GetAll()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, empleados)
}

func (h *EmpleadoHandler) GetById(c *gin.Context) {
	id, err := utils.GetIDParam(c)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.MsgInvalidID)
		return
	}

	empleado, err := h.service.GetById(id)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, empleado)
}

func (h *EmpleadoHandler) Create(c *gin.Context) {
	var empleado models.Empleado
	if err := c.ShouldBindJSON(&empleado); err != nil {
		utils.RespondWithValidationError(c, err)
		return
	}

	if err := h.service.Create(&empleado); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, empleado)
}

func (h *EmpleadoHandler) Update(c *gin.Context) {
	id, err := utils.GetIDParam(c)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.MsgInvalidID)
		return
	}

	var empleado models.Empleado
	if err := c.ShouldBindJSON(&empleado); err != nil {
		utils.RespondWithValidationError(c, err)
		return
	}

	if err := h.service.Update(id, &empleado); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, empleado)
}

func (h *EmpleadoHandler) PartialUpdate(c *gin.Context) {
	id, err := utils.GetIDParam(c)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.MsgInvalidID)
		return
	}

	var request models.EmpleadoPatch
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondWithValidationError(c, err)
		return
	}

	empleado, err := h.service.PartialUpdate(id, &request)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, empleado)
}

func (h *EmpleadoHandler) Delete(c *gin.Context) {
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
