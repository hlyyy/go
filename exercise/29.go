package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	for {
		 t2 := time.Since(t1)
		 fmt.Printf("t2 type is %T",t2)
		 break
	}
}
