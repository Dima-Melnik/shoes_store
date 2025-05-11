package attribute

import (
	"errors"

	"github.com/Dima-Melnik/shoes_store/internal/db"
	"github.com/Dima-Melnik/shoes_store/internal/models"
	"github.com/Dima-Melnik/shoes_store/internal/utils"
	"github.com/gin-gonic/gin"
)

func (a *AttributeServices) CreateAttribute(c *gin.Context, createAtribute *models.Attribute) (*models.Attribute, error) {
	if !utils.BindJSON(c, createAtribute) {
		return nil, errors.New("error in function BindJSON (CreateAttribute)")
	}

	result := db.DB.Create(&createAtribute)
	if result.Error != nil {
		utils.SendLogrusService("create attribute", c, result.Error, "CreateAttribute")
		return nil, result.Error
	}

	return createAtribute, nil
}
