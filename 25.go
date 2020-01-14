package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	for {
		go fmt.Print(1)
		fmt.Print(0)
	}
}
