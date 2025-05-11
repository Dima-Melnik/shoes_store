package products

import (
	"net/http"

	"github.com/Dima-Melnik/shoes_store/internal/utils"

	"github.com/gin-gonic/gin"
)

func GetProductByID(c *gin.Context) {
	id, ok := utils.CheckCorrectID(c)
	if !ok {
		return
	}

	product, err := productsServices.GetProductById(c, id)
	if err != nil {
		utils.SendLoggerHandlers("product by id", err, "GetProductByID")
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}
