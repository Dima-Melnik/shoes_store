package products

import (
	"net/http"

	"github.com/Dima-Melnik/shoes_store/internal/utils"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	createProduct, err := productsServices.CreateProduct(c)
	if err != nil {
		utils.SendLoggerHandlers("create product", err, "CreateProduct")
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, createProduct)
}
