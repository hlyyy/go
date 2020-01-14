package main

import (
    "fmt"
    "time"
)

func StartCac() {
    t1 := time.Now() // get current time
    fmt.Println(t1)
	//logic handlers
    for i := 0; i < 1000; i++ {
        fmt.Print("*")
    }
    elapsed := time.Since(t1)
    fmt.Println("App elapsed: ", elapsed)
}

func main(){
    StartCac()
}
