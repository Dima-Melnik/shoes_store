package products

import (
	"net/http"

	"github.com/Dima-Melnik/shoes_store/internal/models"
	"github.com/Dima-Melnik/shoes_store/internal/utils"

	"github.com/gin-gonic/gin"
)

func UpdateProduct(c *gin.Context) {
	id, ok := utils.CheckCorrectID(c)
	if !ok {
		return
	}

	var updatedProduct *models.Product

	product, err := productsServices.UpdateProduct(c, id, updatedProduct)
	if err != nil {
		utils.SendLoggerHandlers("update product", err, "UpdateProduct")
		utils.SendError(c, http.StatusBadRequest, "Error in handler UpdateProduct")
		return
	}

	c.JSON(http.StatusOK, product)
}
