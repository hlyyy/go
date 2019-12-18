package main

import "github.com/gin-gonic/gin"

type RegisterPayload struct {
	Username  string `json:"username"` //以字典形式储存
	Password  string `json:"password"`
}

func main() {
	router := gin.Default()
	router.POST("/register",func(c *gin.Context){
		var data RegisterPayload
		//返回一个error类型的变量,不为nil则未绑定成功
		if err := c.BindJSON(&data); err != nil {
			c.JSON(400,gin.H{
				"massage": "Bad Request",
			})
			return 
		}
		c.JSON(200,gin.H{
			"massage": data.Username + data.Password,
		})
	})
	router.Run()
}

