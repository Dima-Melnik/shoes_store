package category

import (
	"net/http"

	"github.com/Dima-Melnik/shoes_store/internal/utils"
	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	categories, err := categoryServices.GetAllCategories(c)
	if err != nil {
		utils.SendLoggerHandlers("getting categories", err, "GetCategories")
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, categories)
}
