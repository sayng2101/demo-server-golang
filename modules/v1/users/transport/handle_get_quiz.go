package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/server/modules/v1/users/business"
	"github.com/server/modules/v1/users/model"
	"github.com/server/modules/v1/users/storage"
	"gorm.io/gorm"
	"net/http"
)

func GetQuiz(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var filter model.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSqlStore(db)
		biz := business.GetQuizBusiness(store)
		var result []model.Quiz
		result, err := biz.GetQuiz(c.Request.Context(), &filter)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
