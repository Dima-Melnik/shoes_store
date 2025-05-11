package category

import (
	"net/http"

	"github.com/Dima-Melnik/shoes_store/internal/models"
	"github.com/Dima-Melnik/shoes_store/internal/utils"
	"github.com/gin-gonic/gin"
)

func UpdateCategory(c *gin.Context) {
	id, ok := utils.CheckCorrectID(c)
	if !ok {
		return
	}

	var updatedCategory *models.Category

	category, err := categoryServices.UpdateCategory(c, updatedCategory, id)
	if err != nil {
		utils.SendLoggerHandlers("update category", err, "UpdateCategory")
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, category)
}
