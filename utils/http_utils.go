package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetIDParam extrae y convierte el parámetro "id" de la URL a int
func GetIDParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.Param("id"))
}
