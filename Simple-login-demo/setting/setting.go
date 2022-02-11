package setting

import (
	"gopkg.in/ini.v1"
)

var Conf = new(AppConfig)

// AppConfig 应用程序配置
type AppConfig struct {
	Release       bool `ini:"release"`
	Port          int  `ini:"port"`
	*MySQLConfig  `ini:"mysql"`
	*RedisConfig  `ini:"redis"`
	*WeChatConfig `ini:"wechat"`
}

// MySQLConfig 数据库配置
type MySQLConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	DB       string `ini:"db"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
}

// RedisConfig 数据库配置
type RedisConfig struct {
	Addr     string `ini:"addr"`
	Password string `ini:"password"`
	DB       int    `ini:"db"`
}

type WeChatConfig struct {
	Corpid      string `ini:"corpid"`
	Corpsecret  string `ini:"corpsecret"`
	Appid       string `ini:"appid"`
	Agentid     string `ini:"agentid"`
	RedirectUri string `ini:"redirect_uri"`
}

func Init(file string) error {
	return ini.MapTo(Conf, file)
}
