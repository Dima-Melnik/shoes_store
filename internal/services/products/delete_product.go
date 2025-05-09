package products

import (
	"github.com/Dima-Melnik/shoes_store/internal/db"
	"github.com/Dima-Melnik/shoes_store/internal/models"
	"github.com/Dima-Melnik/shoes_store/internal/utils"

	"github.com/gin-gonic/gin"
)

func (p *ProductsServices) DeleteProduct(c *gin.Context, id int) error {
	result := db.DB.Delete(&models.Product{}, id)
	if result.Error != nil {
		utils.SendLogrusService("delete product", c, result.Error, "DeleteProduct")
		return result.Error
	}

	return nil
}
