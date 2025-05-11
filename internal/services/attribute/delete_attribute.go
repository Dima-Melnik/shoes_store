package attribute

import (
	"github.com/Dima-Melnik/shoes_store/internal/db"
	"github.com/Dima-Melnik/shoes_store/internal/models"
	"github.com/Dima-Melnik/shoes_store/internal/utils"
	"github.com/gin-gonic/gin"
)

func (a *AttributeServices) DeleteAttribute(c *gin.Context, id int) error {
	result := db.DB.Delete(&models.Attribute{}, id)
	if result.Error != nil {
		utils.SendLogrusService("delete attribute", c, result.Error, "DeleteAttribute")
		return result.Error
	}

	return nil
}
