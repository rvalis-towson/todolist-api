package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

/**
 * Todo struct
 */
type Todo struct {
	ID        uuid.UUID `json:"id"`
	Completed bool      `json:"completed"`
	Contents  string    `json:"contents"`
	Title     string    `json:"title"`
	ImageURL  string    `json:"image_url"`
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
	r.PUT("/todo/:id", updateTodo)
	r.GET("/todo/:id", getTodo)
	r.GET("/todos", getAllTodos)
	r.DELETE("/todo/:id", delTodo)
	r.Run(":8888")
}

func createTodo(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var todo Todo
	c.BindJSON(&todo)
	todo.ID = uuid.New()
	db.Create(&todo)
	c.JSON(200, todo)
}

func updateTodo(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var todo Todo
	var todoBind Todo

	if c.ShouldBind(&todoBind) == nil {
		id := c.Params.ByName("id")
		if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			todo.Completed = todoBind.Completed
			if todoBind.Contents != "" {
				todo.Contents = todoBind.Contents
			}
			if todoBind.Title != "" {
				todo.Title = todoBind.Title
			}
			if todoBind.ImageURL != "" {
				todo.ImageURL = todoBind.ImageURL
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
