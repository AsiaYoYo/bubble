package setting

import "gopkg.in/ini.v1"

// Cfg 初始化一个全局的AppConf
var Cfg = new(AppConf)

// AppConf 全局配置
type AppConf struct {
	Port       string `ini:"port"`
	*MysqlCfg  `ini:"mysql"`
	*LogConfig `ini:"log"`
}

// MysqlCfg mysql配置
type MysqlCfg struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	DbName   string `ini:"dbname"`
}

// LogConfig 定义log相关配置
type LogConfig struct {
	Level      string `ini:"level"`
	Filename   string `ini:"filename"`
	MaxSize    int    `ini:"maxsize"`
	MaxAge     int    `ini:"max_age"`
	MaxBackups int    `ini:"max_backups"`
}

// Init 加载配置文件
func Init(File string) error {
	return ini.MapTo(Cfg, File)
}
