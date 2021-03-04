package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// DB 定义一个全局的db
var DB *gorm.DB

// InitMysql 初始化数据库
func InitMysql() (err error) {
	dsn := "root:root@(127.0.0.1:3306)/bubble?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	return
}

func main() {
	// 手动创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := InitMysql()
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("./static", "static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 创建表
	DB.AutoMigrate(&Todo{})

	// 定义一个路由组
	v1Group := r.Group("/v1")
	{
		// 添加一个todo
		v1Group.POST("/todo", func(c *gin.Context) {
			// 从请求中获取数据
			var todo Todo
			c.BindJSON(&todo)
			// 将数据存入数据库
			if err := DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})

		// 查看所有todo
		v1Group.GET("/todo", func(c *gin.Context) {
			var todoList []Todo
			if err := DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})
		// 更新一个todo
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}
			var todo Todo
			if err := DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			c.BindJSON(&todo)
			if err := DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		// 删除一个todo
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}
			if err := DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{id: "deleted"})
			}
		})
	}
	r.Run(":9090")
}
