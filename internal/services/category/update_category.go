package category

import (
	"errors"

	"github.com/Dima-Melnik/shoes_store/internal/db"
	"github.com/Dima-Melnik/shoes_store/internal/models"
	"github.com/Dima-Melnik/shoes_store/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *CategoryServices) UpdateCategory(c *gin.Context, updatedCategory *models.Category, id int) (*models.Category, error) {
	if !utils.BindJSON(c, updatedCategory) {
		return nil, errors.New("error in function BindJSON (UpdateCategory)")
	}

	result := db.DB.Model(&models.Category{}).Where("id = ?", id).Updates(updatedCategory)
	if result.Error != nil {
		utils.SendLogrusService("update category", c, result.Error, "UpdateCategory")
		return nil, result.Error
	}

	return updatedCategory, nil
}
