package attribute

import (
	"github.com/Dima-Melnik/shoes_store/internal/db"
	"github.com/Dima-Melnik/shoes_store/internal/models"
	"github.com/Dima-Melnik/shoes_store/internal/utils"
	"github.com/gin-gonic/gin"
)

func (a *AttributeServices) GetAllAttributes(c *gin.Context) ([]models.Attribute, error) {
	var allAttributes []models.Attribute

	result := db.DB.Find(&allAttributes)
	if result.Error != nil {
		utils.SendLogrusService("all attributes", c, result.Error, "GetAllAttributes")
		return nil, result.Error
	}

	return allAttributes, nil
}
