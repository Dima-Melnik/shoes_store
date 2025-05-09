package category

import (
	"github.com/Dima-Melnik/shoes_store/internal/db"
	"github.com/Dima-Melnik/shoes_store/internal/models"
	"github.com/Dima-Melnik/shoes_store/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *CategoryServices) GetCategoryById(c *gin.Context, id int) (models.Category, error) {
	var categoryByID models.Category

	result := db.DB.First(categoryByID, id)
	if result.Error != nil {
		utils.SendLogrusService("category byID", c, result.Error, "GetCategoryByID")
		return categoryByID, result.Error
	}

	return categoryByID, nil
}
