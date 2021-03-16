package routers

import (
	"bubble/controller"
	"bubble/logger"

	"github.com/gin-gonic/gin"
)

// SetupRouters 定义路由
func SetupRouters() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.LoadHTMLGlob("templates/*")
	r.Static("./static", "static")

	r.GET("/", controller.IndexHandler)

	// 定义一个路由组
	v1Group := r.Group("/v1")
	{
		// 添加一个todo
		v1Group.POST("/todo", controller.CreateATodo)

		// 查看所有todo
		v1Group.GET("/todo", controller.GetTodoList)
		// 更新一个todo
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除一个todo
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
