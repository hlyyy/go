package main

import "github.com/gin-gonic/gin"

func main(){
	router := gin.Default()//获得一个基本的路由变量
	router.GET("/hello",func(c *gin.Context) {//请求与相应信息
		c.JSON(200, gin.H{//gin.H()简化JSON的生成
			"message": "Hello,gin!",
		})
	})
	router.Run()//启动路由
}

