package main

import "github.com/gin-gonic/gin"

func main(){
	router := gin.Default()
	router.GET("/hello",func(c *gin.Context) {
		name,ok := c.GetQuery("name") //查询字符串中包含的参数
		if !ok {
			c.JSON(400,gin.H {
				"message":"Bad Request!",
			})
			return //return停止运行
		}
		c.JSON(200,gin.H{
			"message":"Hello,"+name+"!",
		})
	})
	router.Run()
}
