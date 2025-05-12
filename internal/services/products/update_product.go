package products

import (
	"errors"
	"net/http"

	"github.com/Dima-Melnik/shoes_store/internal/db"
	"github.com/Dima-Melnik/shoes_store/internal/models"
	"github.com/Dima-Melnik/shoes_store/internal/utils"

	"github.com/gin-gonic/gin"
)

func (p *ProductsServices) UpdateProduct(c *gin.Context, id int, updatedProd *models.Product) (*models.Product, error) {
	if !utils.BindJSON(c, updatedProd) {
		return nil, errors.New("error binding json (updateProduct)")
	}

	var existingProduct models.Product
	foundResult := db.DB.First(&updatedProd, id)
	if foundResult.Error != nil {
		utils.SendLogrusService("found product", c, foundResult.Error, "UpdateProduct")
		return nil, foundResult.Error
	}

	var input models.CreateProductInput
	if !utils.BindJSON(c, input) {
		utils.SendError(c, http.StatusBadRequest, "Error bidning input")
		return nil, errors.New("error binding input")
	}

	existingProduct.Name = input.Name
	existingProduct.ImageURL = input.ImageURL
	existingProduct.Price = input.Price
	existingProduct.CategoryID = input.CategoryID

	if len(input.AttributeIDs) > 0 {
		var attributes []models.Attribute
		loadResult := db.DB.Where("id IN ?", input.AttributeIDs).Find(&attributes)
		if loadResult.Error != nil {
			utils.SendLogrusService("load attributes", c, loadResult.Error, "UpdateProduct")
			return nil, loadResult.Error
		}
		updatedProd.Attributes = attributes
	}

	result := db.DB.Save(&updatedProd)
	if result.Error != nil {
		utils.SendLogrusService("save updates", c, result.Error, "UpdateProduct")
		return nil, result.Error
	}

	return updatedProd, nil
}
