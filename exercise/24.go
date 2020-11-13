package main

import (
	"fmt"
)

func main() {

	go func () {
		for i := 0;i<5;i++ {
			fmt.Println("go")
		}
	} ()

	for i := 0;i<2;i++ {
		//runtime.Gosched()
		fmt.Println("main")
	}
}
