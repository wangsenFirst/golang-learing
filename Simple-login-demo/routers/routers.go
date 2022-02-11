package routers

import (
	"Simple-login-demo/controller"
	"Simple-login-demo/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("../static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	// 参数校验
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return nil
	}
	// 开放登入界面
	r.GET("/", controller.IndexHandler)
	// 扫码页面
	r.GET("/wx", func(c *gin.Context) {
		c.HTML(http.StatusOK, "weChat.html", nil)
	})
	// v1
	v1Group := r.Group("v1")
	{
		v1Group.POST("/login", controller.Login)
		v1Group.GET("/Wxlogin", controller.WxChat)
	}
	return r
}
