package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Book 书籍结构体
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Stock  string `json:"stock"`
}

// Response 通用响应结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var books = make(map[string]Book)

// AddBook 添加书籍
// @Summary 添加书籍
// @Description 传入书籍信息新增书籍
// @Tags 图书
// @Accept json
// @Produce json
// @Param book body Book true "书籍信息"
// @Success 200 {object} Response{data=Book} "书籍信息"
// @Failure 400 {object} Response "参数解析失败"
// @Failure 409 {object} Response "书籍ID已存在"
// @Router /book/add [post]
func AddBook(c *gin.Context) {
	var newBook Book
	if err := c.ShouldBind(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Message: "参数解析失败", Data: err.Error()})
		return
	}

	if _, exists := books[newBook.ID]; exists {
		c.JSON(http.StatusConflict, Response{Code: 409, Message: "书籍ID已存在"})
		return
	}

	books[newBook.ID] = newBook
	c.JSON(http.StatusOK, Response{Code: 200, Message: "添加书籍成功", Data: newBook})
}

// DeleteBook 删除书籍
// @Summary 删除书籍
// @Description 根据书籍 ID 删除书籍
// @Tags 图书
// @Accept json
// @Produce json
// @Param id path string true "书籍ID"
// @Success 200 {object} Response "书籍删除成功"
// @Router /book/delete/{id} [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	delete(books, id)
	c.JSON(http.StatusOK, Response{Code: 200, Message: "书籍删除成功"})
}

// UpdateBook 更新书籍
// @Summary 更新书籍
// @Description 根据 ID 更新书籍信息
// @Tags 图书
// @Accept json
// @Produce json
// @Param id path string true "书籍ID"
// @Param book body Book true "更新后的书籍信息"
// @Success 200 {object} Response{data=Book} "书籍更新成功"
// @Failure 400 {object} Response "参数绑定失败"
// @Failure 400 {object} Response "路径ID与请求体ID不一致"
// @Router /book/update/{id} [put]
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var changedBook Book
	if err := c.ShouldBind(&changedBook); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Message: "参数绑定失败", Data: err.Error()})
		return
	}

	if id != changedBook.ID {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Message: "路径ID与请求体ID不一致"})
		return
	}

	books[id] = changedBook
	c.JSON(http.StatusOK, Response{Code: 200, Message: "书籍修改成功", Data: changedBook})
}

// SearchAllBook 获取所有书籍
// @Summary 获取所有书籍
// @Description 获取所有图书信息
// @Tags 图书
// @Produce json
// @Success 200 {object} map[string]Book
// @Router /book/search [get]
func SearchAllBook(c *gin.Context) {
	c.JSON(http.StatusOK, Response{Code: 200, Message: "书籍获取成功", Data: books})
}

// @title 图书管理系统
// @version 1.0
// @description 实现对图书的增删改查的图书管理系统
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()

	r.GET("/book/search", SearchAllBook)
	r.POST("/book/add", AddBook)
	r.DELETE("/book/delete/:id", DeleteBook)
	r.PUT("/book/update/:id", UpdateBook)

	r.Run()
}
