package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"  // 导入驱动
	"github.com/gin-gonic/gin"    
    "fmt"
)

type RegisterPayload struct {
    Username  string  `json:"username"`
    Password  string  `json:"password"`
}

func GetDatabase(username, password, host, port, dbname string) (*sql.DB, error) {
    address := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)
    db, err := sql.Open("mysql", address)
    if err != nil { // 打开失败
        return nil, err
    }
    err = db.Ping()
    if err != nil { // 连接失败
        return nil, err
    }
    return db, nil
}

func main() {
	db, err := GetDatabase("root", "huanglingyun0130", "127.0.0.1", "3306", "testdb")
    if err != nil {
        panic("Link to database failed! Reason: " + err.Error())  // 结束程序并且打印错误原因
    }
	defer db.Close()
	userID := 0
	router := gin.Default()
    router.POST("/register",func (c *gin.Context) {
		var data RegisterPayload
		if err = c.BindJSON(&data); err != nil {
            c.JSON(400, gin.H{
                "message": "Bad Request", // 如果绑定失败，那么我们认定它为一个坏请求，按照规范，状态码应该为400
            })
            return
        }
		if err = db.QueryRow("SELECT id FROM users WHERE username = ?", data.Username).Scan(&userID); err != nil {
			db.Query("INSERT INTO users (username, password) VALUES (?, ?)", data.Username, data.Password)
			c.JSON(200, gin.H{
				"message": data.Username + data.Password,
			})
			return
		} else {
			c.JSON(401, gin.H{
				"message": "User already existed.",
			})
		}
	})
	router.Run()
}
