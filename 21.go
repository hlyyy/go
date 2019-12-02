package main 
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

//定义全局变量db，连接数据库mysql
var db,_ = gorm.Open("mysql","root:@/usersinfo?charset=utf8&parseTime=True&loc=Local")

var status = Status{false,""}

type User struct { //创建User的结构体，储存单个用户信息
	Id          int    `gorm:"AUTO_INCREMENT";json:"-"`
	Username    string `gorm:"type:varchar(100)"`
	Password    string `gorm:"type:varchar(100)"`
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
	var l User
	res := db.Where("username= ? ",user.Username).Find(&l)
	//判断用户数据是否被查找到
	if res.RecordNotFound() {
			return  false
	}
	return true
}

//验证用户名和密码是否正确
func Verify (user User) bool {
	var l User
	res := db.Where("username = ? AND password = ?",user.Username,user.Password).Find(&l)
	if res.RecordNotFound() {
			return false
	}
	return true
}

//用于修改密码
func Modify (user Newuser) bool {
	var l User
	res := db.Where("username = ? AND password = ?",user.Username,user.Oldpassword).Find(&l)
	if res.RecordNotFound() {
		//将Newuser结构体中的信息转换到User中
		l.Username = user.Username
		l.Password = user.Oldpassword
		return false
	}
	//更新
	db.Model(&l).Where("username = ?",user.Username).Update("password",user.Newpassword)
	return true
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
		status = Status{false,"用户名为空"}
		return
	}
	if user.Password == "" {
		status = Status{false,"密码为空"}
		return
	}
	db.Create(&user)//将用户信息储存到数据表中
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
	db,err := gorm.Open("mysql","root:huanglingyun0130@/usersinfo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	http.HandleFunc("/login",login)
	http.HandleFunc("/register",register)
	http.HandleFunc("/update_password",update_password)
	err = http.ListenAndServe(":9000",nil)
	if err != nil {
		fmt.Println("监听失败")
	}
}



























