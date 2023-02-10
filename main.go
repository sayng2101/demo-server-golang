package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	transport2 "github.com/server/modules/v1/users/transport"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dns := "root:1234@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	if db != nil {
		fmt.Println("Ket noi thanh cong")
	}
	r := gin.Default()

	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.MaxMultipartMemory = 8 << 20
	v1 := r.Group("/v1")

	//quiz app
	{
		user := v1.Group("/user")
		{
			user.POST("/create", transport2.CreateUser(db))
			user.POST("/login", transport2.LoginUser(db))
		}
	}

	//item
	//r.Static("/file", "")

	r.Run(":8080")

}
