package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindJSON(c *gin.Context, data interface{}) bool {
	if err := c.ShouldBindJSON(&data); err != nil {
		SendLogrusError("json", "binding", err, "Error binding json")
		SendError(c, http.StatusBadRequest, "Error bidning json / invalid body request")
		return false
	}

	return true
}
