package main
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db,err := gorm.Open("mysql","user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	defer db.Close()
}

