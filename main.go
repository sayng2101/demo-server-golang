package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/server/db"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dns := "root:1234@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	condb, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	if condb != nil {
		fmt.Println("Ket noi thanh cong")
	}

	r := gin.Default()
	// r.MaxMultipartMemory = 8 << 20
	r.POST("/create", db.CreaterItem(condb))
	r.GET("/:title", db.GetItemById(condb))
	r.GET("/", db.GetAllItems(condb))
	r.PATCH("/:id", db.ItemUpdateById(condb))
	r.DELETE("delete/:id", db.DeleteItemById(condb))
	//r.PATCH("/:id/upload", db.UploadImage(condb))
	r.Run(":8080")
}