package category

import (
	"net/http"

	"github.com/Dima-Melnik/shoes_store/internal/utils"
	"github.com/gin-gonic/gin"
)

func GetCategoryByID(c *gin.Context) {
	id, ok := utils.CheckCorrectID(c)
	if !ok {
		return
	}

	categoryByID, err := categoryServices.GetCategoryById(c, id)
	if err != nil {
		utils.SendLoggerHandlers("category byID", err, "GetCategoryByID")
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, categoryByID)
}
