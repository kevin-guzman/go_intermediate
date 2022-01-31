package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go doSome(c)
	g := 25
	h := &g
	fmt.Println("g", g, "h", h, "*h", *h)
}

func doSome(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("done")
	c <- 1
}
