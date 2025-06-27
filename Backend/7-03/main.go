package main

import (
	"github.com/cqhasy/muxisdk/dao"
	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Stock  string `json:"stock"`
	Place  string `json:"place"`
}

// var books = make(map[string]Book)
var books *dao.OrmClient

// AddBook @Summary 添加图书
// @Description 创建一本新图书
// @Tags Book
// @Accept json
// @Produce json
// @Param book body Book true "要添加的图书"
// @Success 200 {object} Book
// @Failure 400 {object} map[string]interface{}
// @Router /book/add [post]
func AddBook(c *gin.Context) {
	var newBook Book
	if er := c.ShouldBind(&newBook); er != nil {
		c.JSON(400, gin.H{"error": er.Error()})
		return
	}

	err := books.Create(c, &newBook)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, newBook)
}

// DeleteBook @Summary 删除图书
// @Description 根据 ID 删除图书
// @Tags Book
// @Param id path int true "图书 ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /book/delete/{id} [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var k = dao.Key{"id", id}
	err := books.Delete(c, &Book{}, k)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "删除成功"})
}

// UpdateBook @Summary 更新图书
// @Description 根据 ID 更新图书信息
// @Tags Book
// @Accept json
// @Produce json
// @Param id path int true "图书 ID"
// @Param book body Book true "更新后的图书数据"
// @Success 200 {object} Book
// @Failure 400 {object} map[string]interface{}
// @Router /book/update/{id} [put]
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var changedBook Book
	if err := c.ShouldBind(&changedBook); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if id != changedBook.ID {
		c.JSON(400, gin.H{"error": "路径 ID 与请求体 ID 不一致"})
		return
	}
	var k = dao.Key{"id", id}
	err := books.Update(c, &changedBook, k)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, changedBook)
}

// SearchBook @Summary 查找图书
// @Description 查询一本图书（参数未指定）
// @Tags Book
// @Produce json
// @Success 200 {object} Book
// @Failure 400 {object} map[string]interface{}
// @Router /book/search [get]
func SearchBook(c *gin.Context) {
	var re Book
	err := books.Get(c, &re)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, re)
}

func main() {
	db := dao.NewOrmClient("root:chenhaoqi318912@tcp(my-mysql:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local")
	db.InitTable(&Book{})
	books = db
	r := gin.Default()

	r.GET("/book/search", SearchBook)
	r.POST("/book/add", AddBook)
	r.DELETE("/book/delete/:id", DeleteBook)
	r.PUT("/book/update/:id", UpdateBook)
	r.Run()
}
