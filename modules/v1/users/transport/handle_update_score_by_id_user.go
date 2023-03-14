package transport

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"net/http"
	"strconv"

	"github.com/server/modules/v1/users/business"
	"github.com/server/modules/v1/users/model"
	"github.com/server/modules/v1/users/storage"
)

func UpdateScoreByIdUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var dataUser model.UpdateScoreUser
		if err := c.ShouldBind(&dataUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := storage.NewSqlStore(db)
		biz := business.UpdateScoreByIdUserBusiness(store)
		if err := biz.UpdateScoreUser(c.Request.Context(), id, &dataUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": true,
		})
	}
}
