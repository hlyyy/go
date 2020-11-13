package main

import(
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		defer fmt.Println("子协程调用完毕")
		for i := 0;i<2;i++ {
			fmt.Println("i=",i)
			time.Sleep(time.Second)
		}
		ch <- 666
	}()
	s := <- ch
	fmt.Println("s=",s)
}


