package main

import (
	"bubble/dao"
	"bubble/logger"
	"bubble/models"
	"bubble/routers"
	"bubble/setting"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./bubble conf/config.ini")
		return
	}
	// 加载配置文件
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("init config failed, err:%v\n", err)
		return
	}

	// 初始化logger
	if err := logger.InitLogger(setting.Cfg.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	// 手动创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := dao.InitMysql(setting.Cfg.MysqlCfg)
	if err != nil {
		panic(err)
	}
	defer dao.Close()

	// 创建表
	models.InitModel()

	// 注册路由
	r := routers.SetupRouters()

	if err := r.Run(fmt.Sprintf(":%s", setting.Cfg.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
