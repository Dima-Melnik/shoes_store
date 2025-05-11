package attribute

import (
	"net/http"

	"github.com/Dima-Melnik/shoes_store/internal/utils"
	"github.com/gin-gonic/gin"
)

func GetAllAttributes(c *gin.Context) {
	attributes, err := attributeServices.GetAllAttributes(c)
	if err != nil {
		utils.SendLoggerHandlers("all attributes", err, "GetAllAttributes")
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, attributes)
}
