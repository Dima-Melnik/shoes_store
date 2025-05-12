package products

import (
	"errors"
	"net/http"

	"github.com/Dima-Melnik/shoes_store/internal/db"
	"github.com/Dima-Melnik/shoes_store/internal/models"
	"github.com/Dima-Melnik/shoes_store/internal/utils"

	"github.com/gin-gonic/gin"
)

func (p *ProductsServices) CreateProduct(c *gin.Context) (*models.Product, error) {
	var input models.CreateProductInput

	if !utils.BindJSON(c, input) {
		return nil, errors.New("error in function BindJSON(CreateProductInput)")
	}

	product := models.Product{
		Name:       input.Name,
		ImageURL:   input.ImageURL,
		Price:      input.Price,
		CategoryID: input.CategoryID,
	}

	var attributes []models.Attribute
	if len(input.AttributeIDs) > 0 {
		loadResult := db.DB.Where("id IN ?", input.AttributeIDs).Find(&attributes)
		if loadResult.Error != nil {
			utils.SendError(c, http.StatusInternalServerError, "Failed to load attributes")
			return nil, loadResult.Error
		}
		product.Attributes = attributes
	}

	result := db.DB.Create(&product)
	if result.Error != nil {
		utils.SendLogrusService("Creating product", c, result.Error, "CreateProduct")
		return nil, result.Error
	}

	return &product, nil
}
