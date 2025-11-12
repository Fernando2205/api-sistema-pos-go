package middleware

import (
	"fmt"
	"net/http"
	"sistema_pos_go/utils"

	"github.com/gin-gonic/gin"
)

// ErrorHandler middleware maneja errores globales y rutas no encontradas
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Si ya se escribió una respuesta, no hacer nada
		if c.Writer.Written() {
			return
		}

		// Manejar error 404 (ruta no encontrada) - Reutiliza NoRouteHandler
		if c.Writer.Status() == http.StatusNotFound {
			NoRouteHandler(c)
			return
		}

		// Verificar si hay errores en el contexto
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
}

// RecoveryHandler maneja panics y los convierte en respuestas de error estructuradas
func RecoveryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("Error interno del servidor: %v", err)
				utils.RespondWithError(c, http.StatusInternalServerError, message)
				c.Abort()
			}
		}()
		c.Next()
	}
}

// NoRouteHandler maneja las rutas que no existen (404)
func NoRouteHandler(c *gin.Context) {
	message := fmt.Sprintf("No se encontró el recurso: %s", c.Request.URL.Path)
	utils.RespondWithError(c, http.StatusNotFound, message)
}
