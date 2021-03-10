package setting

import "gopkg.in/ini.v1"

// Cfg 初始化一个全局的AppConf
var Cfg = new(AppConf)

// AppConf 全局配置
type AppConf struct {
	Port      string `ini:"port"`
	*MysqlCfg `ini:"mysql"`
}

// MysqlCfg mysql配置
type MysqlCfg struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	DbName   string `ini:"dbname"`
}

// Init 加载配置文件
func Init(File string) error {
	return ini.MapTo(Cfg, File)
}
