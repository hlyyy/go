package main
import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

//有json转换的函数，在对待struct时，需要结构体中的属性首字母大写
type User struct {
	Username  string
	Password  string
}

type Status struct {
	Status  bool
	Detail  string
}

var UserArr=make([]User,1)
var status=Status{false,""}

func Existed(user User) bool {
	for _,value :=range UserArr {
		if user.Username==value.Username {
			return true
		} 
	}
	return false
}

func Verify(user User) bool {
	for _,value :=range UserArr {
		if user.Username==value.Username && user.Password==value.Password {
			return true
		}
	}
	return false
}


func Register (userInfo []byte) {
	var user User 
	json.Unmarshal(userInfo,&user)//反序列化，解析json编码的数据并将结果存入user中
	if Existed(user){
		status=Status{false,"用户已存在"}
		return 
	} else {
		UserArr=append(UserArr,user)
		status=Status{true,"注册成功"}
	}
}

func Login (userInfo []byte) {
	var user User
	json.Unmarshal(userInfo,&user)
	if !Existed(user) {
		status=Status{false,"用户不存在"}
		return 
	}
	if !Verify(user) {
		status=Status{false,"用户名或密码错误"}
		return 
	}
	status=Status{true,"登录成功"}
}

func register(res http.ResponseWriter,req *http.Request) {
	if req.Method=="POST" {
		s,_:=ioutil.ReadAll(req.Body)
		Register(s)
		res.Write(Feedbook(status))
	} else {
		res.Write([]byte("{\"false\":\"只支持POST方式\"}"))
	}
}

func login (res http.ResponseWriter,req *http.Request) {
	if req.Method=="POST" {
		s,_:=ioutil.ReadAll(req.Body)
		Login(s)
		res.Write(Feedbook(status))
	}else {
		res.Write([]byte("{\"false\":\"只支持POST方式\"}"))
	}
}

//返回statua的json编码
func Feedbook(a Status)[]byte {
	s,_:=json.Marshal(a)
	return s
}

func main() {
	http.HandleFunc("/register",register)
	http.HandleFunc("/login",login)
	if err :=http.ListenAndServe(":9000",nil);err!=nil {
		fmt.Println("监听失败")
	}
}





































