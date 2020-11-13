package  main

import (
	"fmt"
	"time"
)

var ch =make(chan int)

func Printer(str string) {
	for _,data := range str {
		fmt.Printf("%c",data)
		time.Sleep(time.Second)
	}
}

//person1执行完后才能执行person2
func person1() {
	Printer("hello")
	ch <- 666//给管道传数据
}

func person2() {
	<- ch//从管道取数据，接收，如果通道没有数据他就会堵塞
	Printer("world")
}

func main() {
	go person1()
	go person2()

	for {
	}
}








