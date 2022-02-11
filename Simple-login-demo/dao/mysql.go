package dao

import (
	"Simple-login-demo/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	MysqlDB *gorm.DB
)

func InitMySQL(cfg *setting.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	MysqlDB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return MysqlDB.DB().Ping()
}

func MysqlClose() {
	MysqlDB.Close()
}
