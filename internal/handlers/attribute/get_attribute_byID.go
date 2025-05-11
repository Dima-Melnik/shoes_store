package attribute

import (
	"net/http"

	"github.com/Dima-Melnik/shoes_store/internal/utils"
	"github.com/gin-gonic/gin"
)

func GetAttributeByID(c *gin.Context) {
	id, ok := utils.CheckCorrectID(c)
	if !ok {
		return
	}

	attribute, err := attributeServices.GetAttributeByID(c, id)
	if err != nil {
		utils.SendLoggerHandlers("attribute byID", err, "GetAttributeByID")
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, attribute)
}
