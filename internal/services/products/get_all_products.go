package products

import (
	"errors"
	"strconv"

	"github.com/Dima-Melnik/shoes_store/internal/db"
	"github.com/Dima-Melnik/shoes_store/internal/models"
	"github.com/Dima-Melnik/shoes_store/internal/utils"

	"github.com/gin-gonic/gin"
)

func (p *ProductsServices) GetAllProducts(c *gin.Context) ([]models.Product, error) {
	var allProducts []models.Product

	query := db.DB

	// Pagination

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}

	if limit < 10 {
		limit = 10
	}

	offset := (page - 1) * limit

	// Filter

	name := c.Query("name")

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	minPrice := c.Query("min_price")
	if minPrice != "" {
		minPriceInt, err := strconv.Atoi(minPrice)
		if err == nil {
			query = query.Where("price >= ?", minPriceInt)
		}
	}

	maxPrice := c.Query("max_price")
	if maxPrice != "" {
		maxPriceInt, err := strconv.Atoi(maxPrice)
		if err == nil {
			query = query.Where("price <= ?", maxPriceInt)
		}
	}

	result := query.Limit(limit).Offset(offset).Order("create_at DESC").Preload("Category").Preload("Attributes").Find(&allProducts)
	if result.Error != nil {
		utils.SendLogrusService("getting all products", c, result.Error, "GetAllProducts")
		return nil, result.Error
	}

	if len(allProducts) == 0 {
		return nil, errors.New("products is empty")
	}

	return allProducts, nil
}
