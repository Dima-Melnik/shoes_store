package attribute

import (
	"net/http"

	"github.com/Dima-Melnik/shoes_store/internal/utils"
	"github.com/gin-gonic/gin"
)

func DeleteAttribute(c *gin.Context) {
	id, ok := utils.CheckCorrectID(c)
	if !ok {
		return
	}

	err := attributeServices.DeleteAttribute(c, id)
	if err != nil {
		utils.SendLoggerHandlers("delete attribute", err, "DeleteAttribute")
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "suuccessfully deleted"})
}
