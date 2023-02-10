package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/server/modules/v1/users/business"
	"github.com/server/modules/v1/users/model"
	"github.com/server/modules/v1/users/storage"
	"gorm.io/gorm"
)

func LoginUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var user model.LoginUser
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		// user := c.PostForm("account")
		// pass := c.PostForm("password")
		// if err := c.ShouldBindQuery(&user); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"error": err.Error(),
		// 	})
		// 	return
		// }
		// if err := c.ShouldBindQuery(&pass); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"error": err.Error(),
		// 	})
		// 	return
		// }
		store := storage.NewSqlStore(db)
		biz := business.LoginUserBusiness(store)
		data, err := biz.LoginUser(c.Request.Context(), user.Account, user.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": data,
		})
	}
}
