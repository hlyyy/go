package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var i int64 = -1
	var n int64
	var ch = make(chan int64)
	t1 := time.Now()
	runtime.GOMAXPROCS(2)

	go func() {
		for {
			t1 = time.Now()
			for {
				n = <-ch
				if (int64(time.Since(t1)/time.Millisecond))==1 {
					fmt.Println(n)
					break
				}
			} 
		}
	}()

	for {
		i += 1
		ch <- i
	}
}
