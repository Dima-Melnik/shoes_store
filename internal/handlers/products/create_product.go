package products

import (
	"net/http"

	"github.com/Dima-Melnik/shoes_store/internal/models"
	"github.com/Dima-Melnik/shoes_store/internal/utils"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var createdProduct *models.Product

	createProduct, err := productsServices.CreateProduct(c, createdProduct)
	if err != nil {
		utils.SendLoggerHandlers("create product", err, "CreateProduct")
		utils.SendError(c, http.StatusBadRequest, "Error in handler CreateProduct")
		return
	}

	c.JSON(http.StatusCreated, createProduct)
}
