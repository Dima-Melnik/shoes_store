package category

import (
	"net/http"

	"github.com/Dima-Melnik/shoes_store/internal/utils"
	"github.com/gin-gonic/gin"
)

func DeleteCategory(c *gin.Context) {
	id, ok := utils.CheckCorrectID(c)
	if !ok {
		return
	}

	err := categoryServices.DeleteCategory(c, id)
	if err != nil {
		utils.SendLoggerHandlers("delete category", err, "DeleteCategory")
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "successfully deleted"})
}
