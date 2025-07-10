package main

import (
	"PersonBlog/config"
	"PersonBlog/logger"
	"PersonBlog/router"
	"fmt"
)

func main() {
	//初始化配置
	config.InitConfig()
	logger.InitLog()
	//config.DB.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	//初始化路由器
	r := router.InitRouter()

	//启动服务
	addr := fmt.Sprintf(":%d", config.Conf.App.Port)
	err := r.Run(addr)
	if err != nil {
		return
	}
}
