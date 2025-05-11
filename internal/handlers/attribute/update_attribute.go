package attribute

import (
	"net/http"

	"github.com/Dima-Melnik/shoes_store/internal/models"
	"github.com/Dima-Melnik/shoes_store/internal/utils"
	"github.com/gin-gonic/gin"
)

func UpdateAttribute(c *gin.Context) {
	id, ok := utils.CheckCorrectID(c)
	if !ok {
		return
	}

	var updatedAttribute *models.Attribute

	attribute, err := attributeServices.UpdateAttribute(c, id, updatedAttribute)
	if err != nil {
		utils.SendLoggerHandlers("update attribute", err, "UpdateAttribute")
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, attribute)
}
