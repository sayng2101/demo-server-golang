package db

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/server/model"
	"gorm.io/gorm"
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
		if err := db.Create(&data).Error; err != nil {
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

func GetItemById(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data []model.ToDoItem
		// id, err := strconv.Atoi(c.Param("id"))
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"error": err.Error(),
		// 	})
		// 	return
		// }
		title := c.Param("title")
		if err := db.Where("title LIKE ?", "%"+title+"%").Find(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

func GetAllItems(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var result []model.ToDoItem
		if err := db.Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}

func ItemUpdateById(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.ToDoItemUpdate
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}

func DeleteItemById(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.ToDoItem
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = db.Where("id = ?", id).Delete(data).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})

	}
}

// func UploadImage(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var data model.UploadImage
// 		id, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}
// 		file, err := c.FormFile("file")
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		if err := db.Where("id = ?", id).UpdateColumn(data , file).Error; err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{
// 			"data": true,
// 		})
// 	}
// }
