package utils

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetIDParam extrae y convierte el parámetro "id" de la URL a int
func GetIDParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.Param("id"))
}

// ValidateID verifica que el ID sea válido (mayor o igual a MinId)
func ValidateID(id int) error {
	if id < MinId {
		return ErrInvalidData
	}
	return nil
}

// GetHTTPStatusFromError mapea errores de dominio a códigos HTTP apropiados
func GetHTTPStatusFromError(err error) int {
	switch {
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, ErrInvalidID):
		return http.StatusBadRequest
	case errors.Is(err, ErrInvalidData):
		return http.StatusBadRequest
	case errors.Is(err, ErrEmptyField):
		return http.StatusBadRequest
	case errors.Is(err, ErrMaxLengthExceeded):
		return http.StatusBadRequest
	case errors.Is(err, ErrDuplicateEntry):
		return http.StatusConflict
	case errors.Is(err, ErrAlreadyExists):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
