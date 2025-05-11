package attribute

import (
	"github.com/Dima-Melnik/shoes_store/internal/db"
	"github.com/Dima-Melnik/shoes_store/internal/models"
	"github.com/Dima-Melnik/shoes_store/internal/utils"
	"github.com/gin-gonic/gin"
)

func (a *AttributeServices) GetAttributeByID(c *gin.Context, id int) (models.Attribute, error) {
	var attributeByID models.Attribute

	result := db.DB.First(&attributeByID, id)
	if result.Error != nil {
		utils.SendLogrusService("attribute byID", c, result.Error, "GetAttributeByID")
		return attributeByID, nil
	}

	return attributeByID, nil
}
