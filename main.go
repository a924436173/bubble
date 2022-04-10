package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
	"bubble/settings"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage：./bubble conf/config.ini")
		return
	}
	// 加载配置文件
	if err := settings.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	// 创建数据库
	// sql : create database bubble;
	// 连接数据库
	err := dao.InitMysql(settings.Conf.MySQLConfig)
	if err != nil {
		panic(err)
	}
	defer dao.Close() // 程序退出关闭数据库连接
	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()

	r.Run(":9000")
}
