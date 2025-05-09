package category

import (
	"errors"

	"github.com/Dima-Melnik/shoes_store/internal/db"
	"github.com/Dima-Melnik/shoes_store/internal/models"
	"github.com/Dima-Melnik/shoes_store/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *CategoryServices) CreateCategory(c *gin.Context, createdCategory *models.Category) (*models.Category, error) {
	if !utils.BindJSON(c, createdCategory) {
		return nil, errors.New("error in function BindJSON (CreateCategory)")
	}

	result := db.DB.Create(createdCategory)
	if result.Error != nil {
		utils.SendLogrusService("create category", c, result.Error, "CreateCategory")
		return nil, result.Error
	}

	return createdCategory, nil
}
