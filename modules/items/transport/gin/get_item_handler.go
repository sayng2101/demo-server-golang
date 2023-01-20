package ginitem

import (
	"github.com/gin-gonic/gin"
	"github.com/server/modules/items/business"
	"github.com/server/modules/items/storage"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetItemById(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := storage.NewSqlStore(db)
		biz := business.GetItemBusiness(store)
		data, err := biz.GetItemById(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   data,
		})
	}
}
