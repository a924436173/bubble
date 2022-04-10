package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
)

func main() {
	// 创建数据库
	// sql : create database bubble;
	// 连接数据库
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	defer dao.Close() // 程序退出关闭数据库连接
	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()

	r.Run(":9000")
}
