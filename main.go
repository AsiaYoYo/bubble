package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
)

func main() {
	// 手动创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	defer dao.Close()

	// 创建表
	models.InitModel()

	r := routers.SetupRouters()

	r.Run(":9090")
}
