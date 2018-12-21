package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

/**
 * Todo struct
 */
type Todo struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}

var db *gorm.DB
var err error

func main() {

	db, _ = gorm.Open("sqlite3", "./todo.db")
	defer db.Close()

	r := gin.Default()
	db.AutoMigrate(&Todo{})

	r.POST("/todo/create", createTodo)
	r.GET("/todo/:id", getTodo)
	r.GET("/todos", getAllTodos)
	r.DELETE("/todo/:id", delTodo)
	r.Run(":8888")
}

func createTodo(c *gin.Context) {
	var todo Todo
	c.BindJSON(&todo)
	db.Create(&todo)
	c.JSON(200, todo)
}

func getTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo Todo
	if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, todo)
	}
}

func delTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo Todo
	db.Where("id = ?", id).Delete(&todo)
	c.JSON(200, gin.H{id: "deleted"})
}

func getAllTodos(c *gin.Context) {
	var todos []Todo
	if err := db.Find(&todos).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, todos)
	}
}
