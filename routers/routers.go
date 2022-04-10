package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件在哪
	r.Static("/static", "static")
	// 告诉gin框架模板文件在哪
	r.LoadHTMLGlob("templates/*")

	r.GET("/", controller.IndexHandler)

	// v1
	v1Group := r.Group("v1")
	{
		// 代办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看所有代办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
