package category

import (
	"net/http"

	"github.com/Dima-Melnik/shoes_store/internal/models"
	"github.com/Dima-Melnik/shoes_store/internal/utils"
	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var createdCategory *models.Category

	category, err := categoryServices.CreateCategory(c, createdCategory)
	if err != nil {
		utils.SendLoggerHandlers("create category", err, "CreateCategory")
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, category)
}
