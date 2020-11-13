package main 
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

var userArr = make([]User,0) //userArr用于储存用户的信息，切片
var userId uint = 1
var status = Status{false,""}

type User struct { //创建User的结构体，储存单个用户信息
	Id          uint`json:"-"`
	Username    string
	Password    string
}

type Status struct {
	State   bool
	Detail  string 
}

//用于改密码的结构体
type Newuser struct {
	Username    string
	Oldpassword string
	Newpassword string
}

//判断用户是否存在
func Existed(user User) bool {
	for _,value := range userArr {
		if value.Username == user.Username {
			return true
		}
	}
	return false
}

//验证用户名和密码是否正确
func Verify (user User) bool {
	for _,value := range userArr {
		if value.Username == user.Username && value.Password == user.Password {
			return true
		}
	}
	return false
}

//用于修改密码
func Modify (user Newuser) bool {
	for _,value := range userArr {
		if value.Username == user.Username && value.Password == user.Oldpassword {
			value.Password = user.Oldpassword
			return true
		}
	}
	return false
}

//注册，并验证用户名是否已被注册过
func Register(userInfo []byte) {
	var user User
	json.Unmarshal(userInfo,&user)
	if Existed(user) {
		status = Status{false,"用户已存在"}
		return 
	}
	if user.Username == "" {
		status=Status{false,"用户名为空"}
		return
	}
	if user.Password == "" {
		status=Status{false,"密码为空"}
		return
	}
	user.Id = userId
	userId += 1
	userArr = append(userArr,user)
	status = Status{true,"注册成功"}
}

//登录,验证用户名密码是否正确
func LoginIn(userInfo []byte) {
	var user User
	json.Unmarshal(userInfo,&user)
	if !Existed(user) {
		status = Status{false,"用户名不存在"}
		return 
	}
	if !Verify(user) {
		status = Status{false,"用户名或密码错误"}
		return 
	}
	status = Status{true,"登录成功"}
}

//修改密码
func Update_Password(userInfo []byte) {
	var user Newuser
	json.Unmarshal(userInfo,&user)
	if !Modify(user) {
		status = Status{false,"用户名或旧密码错误"}
		return
	}
	status = Status{true,"修改成功"}
}


//注册的具体交互
func register(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		s,err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err) 
		}
		Register(s)
		res.Write(Feedbook(status))
	} else {
		res.Write([]byte("{\"false\":\"只支持POST方式\"}"))
	}
}

//登录的具体交互
func login(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		s,err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		LoginIn(s)
		res.Write(Feedbook(status))
	} else {
		res.Write([]byte("{\"false\":\"只支持POST方式\"}"))
	}
}

//修改密码的具体交互
func update_password(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		s,err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		Update_Password(s)
		res.Write(Feedbook(status))
	} else {
		res.Write([]byte("{\"false\":\"只支持POST方式\"}"))
	}
}


//将回馈信息转换为json byte
func Feedbook(a Status)[]byte {
	s,err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	return s
}

func main() {
	http.HandleFunc("/login",login)
	http.HandleFunc("/register",register)
	http.HandleFunc("/update_password",update_password)
	err := http.ListenAndServe(":9000",nil)
	if err != nil {
		fmt.Println("监听失败")
	}
}



























