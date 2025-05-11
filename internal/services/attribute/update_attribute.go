package attribute

import (
	"errors"

	"github.com/Dima-Melnik/shoes_store/internal/db"
	"github.com/Dima-Melnik/shoes_store/internal/models"
	"github.com/Dima-Melnik/shoes_store/internal/utils"
	"github.com/gin-gonic/gin"
)

func (a *AttributeServices) UpdateAttribute(c *gin.Context, id int, updateAttribute *models.Attribute) (*models.Attribute, error) {
	if !utils.BindJSON(c, updateAttribute) {
		return nil, errors.New("error in function BindJSON (UreateAttribute)")
	}

	result := db.DB.Model(&models.Attribute{}).Where("id = ?", id).Updates(updateAttribute)
	if result.Error != nil {
		utils.SendLogrusService("update attribute", c, result.Error, "UpdateAttribute")
		return nil, result.Error
	}

	return updateAttribute, nil
}
