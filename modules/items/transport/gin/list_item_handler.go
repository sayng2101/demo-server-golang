package ginitem

import (
	"github.com/gin-gonic/gin"
	"github.com/server/common"
	"github.com/server/modules/items/business"
	"github.com/server/modules/items/storage"
	"gorm.io/gorm"
	"net/http"
)

func ListItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		paging.Process()

		store := storage.NewSqlStore(db)
		biz := business.ListItemBusiness(store)
		result, err := biz.ListItem(c.Request.Context(), nil, &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   result,
			"paging": paging,
		})
	}
}
