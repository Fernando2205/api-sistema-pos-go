package utils

import (
	"errors"
	"time"
)

// Errores comunes de la aplicación
var (
	ErrInvalidID         = errors.New("ID inválido")
	ErrNotFound          = errors.New("recurso no encontrado")
	ErrInvalidData       = errors.New("datos inválidos")
	ErrEmptyField        = errors.New("campo vacío")
	ErrMaxLengthExceeded = errors.New("longitud máxima excedida")
	ErrAlreadyExists     = errors.New("recurso ya existe")
	ErrDuplicateEntry    = errors.New("entrada duplicada")
)

// Códigos de error personalizados
const (
	ErrCodeBadRequest          = "BAD_REQUEST"
	ErrCodeNotFound            = "NOT_FOUND"
	ErrCodeConflict            = "CONFLICT"
	ErrCodeInternalServerError = "INTERNAL_SERVER_ERROR"
	ErrCodeValidationFailed    = "VALIDATION_FAILED"
)

// ErrorResponse es la estructura estándar para respuestas de error
type ErrorResponse struct {
	ErrorCode string    `json:"errorCode"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Path      string    `json:"path"`
}

// NewErrorResponse crea una nueva respuesta de error
func NewErrorResponse(errorCode, message, path string) ErrorResponse {
	return ErrorResponse{
		ErrorCode: errorCode,
		Message:   message,
		Timestamp: time.Now(),
		Path:      path,
	}
}
