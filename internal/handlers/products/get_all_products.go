package products

import (
	"net/http"

	"github.com/Dima-Melnik/shoes_store/internal/utils"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products, err := productsServices.GetAllProducts(c)
	if err != nil {
		utils.SendLoggerHandlers("get products", err, "GetProducts")
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}
