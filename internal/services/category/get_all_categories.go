package category

import (
	"errors"
	"strconv"

	"github.com/Dima-Melnik/shoes_store/internal/db"
	"github.com/Dima-Melnik/shoes_store/internal/models"
	"github.com/Dima-Melnik/shoes_store/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *CategoryServices) GetAllCategories(c *gin.Context) ([]models.Category, error) {
	var allCategories []models.Category

	query := db.DB

	// pagination

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}

	if limit < 10 {
		limit = 10
	}

	offset := (page - 1) * limit

	// search by name

	name := c.Query("name")

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	result := query.Offset(offset).Limit(limit).Find(&allCategories)
	if result.Error != nil {
		utils.SendLogrusService("all categories", c, result.Error, "GetAllCategories")
		return nil, result.Error
	}

	if len(allCategories) == 0 {
		return nil, errors.New("categories is empty")
	}

	return allCategories, nil
}
