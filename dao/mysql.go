package dao

import (
	"bubble/setting"
	"fmt"

	"github.com/jinzhu/gorm"
	// mysql驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 定义一个全局的db
var DB *gorm.DB

// InitMysql 初始化数据库
func InitMysql(Cfg *setting.MysqlCfg) (err error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		Cfg.User, Cfg.Password, Cfg.Host, Cfg.Port, Cfg.DbName)
	fmt.Println(dsn)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

// Close 向外暴露一个close函数
func Close() {
	DB.Close()
}
