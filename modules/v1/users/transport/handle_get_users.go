package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/server/modules/v1/users/business"
	"github.com/server/modules/v1/users/model"
	"github.com/server/modules/v1/users/storage"
	"gorm.io/gorm"
	"net/http"
)

func GetUsers(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		store := storage.NewSqlStore(db)
		biz := business.GetUsersBusiness(store)
		var result []model.UserResponse
		result, err := biz.GetUsers(c.Request.Context())

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
