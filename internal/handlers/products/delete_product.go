package products

import (
	"net/http"

	"github.com/Dima-Melnik/shoes_store/internal/utils"

	"github.com/gin-gonic/gin"
)

func DeleteProduct(c *gin.Context) {
	id, ok := utils.CheckCorrectID(c)
	if !ok {
		return
	}

	err := productsServices.DeleteProduct(c, id)
	if err != nil {
		utils.SendLoggerHandlers("delete product", err, "DeleteProduct")
		utils.SendError(c, http.StatusBadRequest, "Error in handler DeleteProduct")
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Successfully deleted"})
}
