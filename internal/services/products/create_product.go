package products

import (
	"errors"

	"github.com/Dima-Melnik/shoes_store/internal/db"
	"github.com/Dima-Melnik/shoes_store/internal/models"
	"github.com/Dima-Melnik/shoes_store/internal/utils"

	"github.com/gin-gonic/gin"
)

func (p *ProductsServices) CreateProduct(c *gin.Context, createdProd *models.Product) (*models.Product, error) {
	if !utils.BindJSON(c, createdProd) {
		return nil, errors.New("error in function BindJSON (CreateProduct)")
	}

	result := db.DB.Create(createdProd)
	if result.Error != nil {
		utils.SendLogrusService("Creating product", c, result.Error, "CreateProduct")
		return nil, result.Error
	}

	return createdProd, nil
}
