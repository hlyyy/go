package main

import (
	"fmt"
	"time"
	"runtime"
)

func main() {
	var i int64 = -1
	runtime.GOMAXPROCS(2)

	go func() {
		for {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}()

	for {
		i += 1
	//	time.Sleep(time.Second)
	}
}
