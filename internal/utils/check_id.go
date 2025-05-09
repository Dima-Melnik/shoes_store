package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CheckCorrectID(c *gin.Context) (int, bool) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendLogrusError("id", "getting", err, "Error getting id / invalid id value")
		SendError(c, http.StatusBadRequest, "Invalid id")
		return id, false
	}

	return id, true
}
