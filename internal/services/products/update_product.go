package products

import (
	"errors"

	"github.com/Dima-Melnik/shoes_store/internal/db"
	"github.com/Dima-Melnik/shoes_store/internal/models"
	"github.com/Dima-Melnik/shoes_store/internal/utils"

	"github.com/gin-gonic/gin"
)

func (p *ProductsServices) UpdateProduct(c *gin.Context, id int, updatedProd *models.Product) (*models.Product, error) {
	if !utils.BindJSON(c, updatedProd) {
		return nil, errors.New("error in function BindJSON (UpdateProduct)")
	}

	result := db.DB.Model(&models.Product{}).Where("id = ?", id).Updates(updatedProd)
	if result.Error != nil {
		utils.SendLogrusService("update product", c, result.Error, "UpdateProduct")
		return nil, result.Error
	}

	return updatedProd, nil
}
