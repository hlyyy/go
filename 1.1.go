package main 

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()//获得一个基本的路由变量
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"massage":"hello,gin",
		})
	})
	router.Run()
}
