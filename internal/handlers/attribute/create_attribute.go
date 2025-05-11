package attribute

import (
	"net/http"

	"github.com/Dima-Melnik/shoes_store/internal/models"
	"github.com/Dima-Melnik/shoes_store/internal/utils"
	"github.com/gin-gonic/gin"
)

func CreateAttribute(c *gin.Context) {
	var createdAttribute *models.Attribute

	attribute, err := attributeServices.CreateAttribute(c, createdAttribute)
	if err != nil {
		utils.SendLoggerHandlers("create attribute", err, "CreateAttribute")
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, attribute)
}
