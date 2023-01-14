package ginitem

import (
	"github.com/gin-gonic/gin"
	"github.com/server/modules/items/business"
	"github.com/server/modules/items/model"
	"github.com/server/modules/items/storage"
	"gorm.io/gorm"
	"net/http"
)

func CreaterItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.ToDoItemCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := storage.NewSqlStore(db)
		biz := business.NewCreateBusiness(store)
		if err := biz.CreateNewItem(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data.Id,
		})
	}
}
