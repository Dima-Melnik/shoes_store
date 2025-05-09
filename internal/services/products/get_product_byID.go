package products

import (
	"github.com/Dima-Melnik/shoes_store/internal/db"
	"github.com/Dima-Melnik/shoes_store/internal/models"
	"github.com/Dima-Melnik/shoes_store/internal/utils"

	"github.com/gin-gonic/gin"
)

func (p *ProductsServices) GetProductById(c *gin.Context, id int) (models.Product, error) {
	var productByID models.Product

	result := db.DB.First(&productByID, id)
	if result.Error != nil {
		utils.SendLogrusService("product by id", c, result.Error, "GetProductByID")
		return productByID, nil
	}

	return productByID, nil
}
