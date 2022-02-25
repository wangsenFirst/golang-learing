package main

import (
	"Simple-login-demo/dao"
	"Simple-login-demo/models"
	"Simple-login-demo/routers"
	"Simple-login-demo/setting"
	"fmt"
)

func main() {
	println("Hello WangSen")
	// 加载配置文件
	if err := setting.Init("./conf/config.ini"); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}

	// 连接mysql
	err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer dao.MysqlClose() // 程序退出关闭数据库连接
	// 连接Redis
	err = dao.InitRedis(setting.Conf.RedisConfig)
	if err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer dao.RedisClose()
	// redis 计数器
	//pipe := dao.RedisDB.TxPipeline()
	//incr := pipe.Incr("key")
	//pipe.Expire("key", 10*time.Second)
	//_, exer := pipe.Exec()
	//if exer != nil {
	//	fmt.Println(exer)
	//}
	//fmt.Println(incr.Val())

	// 模型绑定
	dao.MysqlDB.Debug().AutoMigrate(&models.User{})
	// 注册路由
	r := routers.SetupRouter()
	println("================================================")

	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}

}
