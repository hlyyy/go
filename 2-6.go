package main
import "fmt"
func main() {
	var a int
	fmt.Println("请输入您的身高:")
	fmt.Scan(&a)
	fmt.Printf("您的标准体重是%f\n",float64(a-100)*0.9)
}
