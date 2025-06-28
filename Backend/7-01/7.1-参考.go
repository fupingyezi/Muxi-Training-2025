package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Book struct {
	ID     string `json:"id" gorm:"primaryKey;type:varchar(64)"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Stock  int    `json:"stock"`
	Place  string `json:"place"`
}

var db *gorm.DB

func initDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/book_management?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	if err := db.AutoMigrate(&Book{}); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}
}

func AddBook(c *gin.Context) {
	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&newBook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "添加图书失败"})
		return
	}

	c.JSON(http.StatusOK, newBook)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&Book{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var changedBook Book
	if err := c.ShouldBindJSON(&changedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&Book{}).Where("id = ?", id).Updates(changedBook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败"})
		return
	}

	c.JSON(http.StatusOK, changedBook)
}

func SearchAllBook(c *gin.Context) {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}
	c.JSON(http.StatusOK, books)
}

func main() {
	initDB()

	r := gin.Default()

	r.GET("/book/search", SearchAllBook)
	r.POST("/book/add", AddBook)
	r.DELETE("/book/delete/:id", DeleteBook)
	r.PUT("/book/update/:id", UpdateBook)

	if err := r.Run(); err != nil {
		log.Fatalf("Gin 启动失败: %v", err)
	}
}
