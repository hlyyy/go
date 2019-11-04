package main
import "fmt"

func bijiaodaxiao (a,b int) (i,j int) {
	if a<b {
		i,j=b,a
	} else {
		i,j=a,b
	}
	return 
}

func main() {
	var i,j,a int
	fmt.Println("让我们来画一个长方形")
	fmt.Printf("一边:")
	fmt.Scan(&i)
	fmt.Printf("另一边:")
	fmt.Scan(&j)
	i,j=bijiaodaxiao(i,j)
	for ;j>=1;j-- {
		for a=i;a>=1;a-- {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

