package router

import (
	"github.com/gin-gonic/gin"
	"github.com/server/db"
	"gorm.io/gorm"
)

func RouterInit(r *gin.RouterGroup, condb *gorm.DB) {

	r.POST("/create", db.CreaterItem(condb))
	r.GET("/:id", db.GetItemById(condb))
	r.GET("/", db.GetAllItems(condb))
}
