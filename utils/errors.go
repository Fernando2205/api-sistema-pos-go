package utils

import "errors"

// Errores comunes de la aplicación
var (
	ErrInvalidID         = errors.New("ID inválido")
	ErrNotFound          = errors.New("recurso no encontrado")
	ErrInvalidData       = errors.New("datos inválidos")
	ErrDuplicateEntry    = errors.New("entrada duplicada")
	ErrEmptyField        = errors.New("campo vacío")
	ErrMaxLengthExceeded = errors.New("longitud máxima excedida")
	ErrNilPointer        = errors.New("valor nulo no permitido")
)
