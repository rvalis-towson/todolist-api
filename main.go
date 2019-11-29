package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

/**
 * Todo struct
 */
type Todo struct {
	ID      uint   `json:"id"`
	Done    bool   `json:"done"`
	Content string `json:"content"`
}

var db *gorm.DB
var err error

func main() {

	db, _ = gorm.Open("sqlite3", "./todo.db")
	defer db.Close()

	r := gin.Default()
	db.AutoMigrate(&Todo{})
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"POST", "GET", "DELETE"},
		AllowHeaders:  []string{"Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	r.POST("/todo/create", createTodo)
	r.POST("/todo/update", updateTodo)
	r.GET("/todo/:id", getTodo)
	r.GET("/todos", getAllTodos)
	r.DELETE("/todo/:id", delTodo)
	r.Run(":8888")
}

func createTodo(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var todo Todo
	c.BindJSON(&todo)
	db.Create(&todo)
	c.JSON(200, todo)
}

func updateTodo(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var todo Todo
	var todoBind Todo

	if c.ShouldBind(&todoBind) == nil {
		id := todoBind.ID
		if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
			c.AbortWithStatus(404)
		} else {
			todo.Done = todoBind.Done
			if todoBind.Content != "" {
				todo.Content = todoBind.Content
			}
			db.Save(&todo)
			c.JSON(200, todo)
		}
	} else {
		c.AbortWithStatus(404)
	}
}

func getTodo(c *gin.Context) {
	c.Header("Content-Type", "application/json")
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
	c.Header("Content-Type", "application/json")
	id := c.Params.ByName("id")
	var todo Todo
	db.Where("id = ?", id).Delete(&todo)
	c.JSON(200, gin.H{id: "deleted"})
}

func getAllTodos(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var todos []Todo
	if err := db.Find(&todos).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, todos)
	}
}
