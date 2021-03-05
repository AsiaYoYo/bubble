package dao

import (
	"github.com/jinzhu/gorm"
	// mysql驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 定义一个全局的db
var DB *gorm.DB

// InitMysql 初始化数据库
func InitMysql() (err error) {
	dsn := "root:root@(127.0.0.1:3306)/bubble?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	return
}

// Close 向外暴露一个close函数
func Close() {
	DB.Close()
}
