package category

import (
	"github.com/Dima-Melnik/shoes_store/internal/db"
	"github.com/Dima-Melnik/shoes_store/internal/models"
	"github.com/Dima-Melnik/shoes_store/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *CategoryServices) DeleteCategory(c *gin.Context, id int) error {
	result := db.DB.Delete(&models.Category{}, id)
	if result.Error != nil {
		utils.SendLogrusService("delete category", c, result.Error, "DeleteCategory")
		return result.Error
	}

	return nil
}
