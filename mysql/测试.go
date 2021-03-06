package main
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

type Lable struct {
									//基本结构体
	ID    int	`gorm:"AUTO_INCREMENT"`     //自动递增，默认主键
	Name  string `gorm:"type:varchar(100)"`
}

//手动设置数据表名
func (Lable) TableName() string {
	return "lable1"
}

func main() {
	db,err := gorm.Open("mysql","root:huanglingyun0130@/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	defer db.Close()
	//lable := Lable{Name:"Helen"}
	//db.Create(&lable)
	//db.Create(&Lable{Name:"huanglingyun"})
	var l Lable
	res := db.Where("name=?","Helen").Find(&l)
	fmt.Println(l)
	//db.Model(&Lable{}).Where(&Lable{Name:"Helen"}).Find(&l)
	//fmt.Println(l)
	if res.RecordNotFound() {
		fmt.Println("Record Not Found")
	} else {
		fmt.Println("Record Found")
	}
}

