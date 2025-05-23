package main

import (
	"certdeck/config"
	"certdeck/middleware"
	v1 "certdeck/routes/v1"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config.LoadConfig()

	Application := gin.New()

	// 加载中间件
	Application.Use(middleware.Cors())                    // 跨域
	Application.Use(middleware.Jwt())                     // Jwt
	Application.Use(middleware.Logger(middleware.HttpIn)) // 日志
	Application.Use(middleware.RecoveryWithLogger())      // Panic

	baseRouter := Application.Group(config.GlobalConfig.Application.Name)

	v1.BasicRegister(baseRouter)
	v1.EmailRegister(baseRouter)

	err := Application.Run(":" + config.GlobalConfig.Application.Port)
	if err != nil {
		log.Fatal(config.GlobalConfig.Application.Name+" 启动失败 : ", err)
	}
}
