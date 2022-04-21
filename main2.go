package main

import (
	"fmt"
	"time"
)

var c1 = make(chan int, 1)
var c2 = make(chan int, 1)
var c3 = make(chan int, 1)

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c2 <- i
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c3 <- i
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case i1 := <-c1:
			fmt.Println("c1 ", i1)
		case i2 := <-c2:
			fmt.Println("c2 ", i2)
		case i3 := <-c3:
			fmt.Println("c3 ", i3)
		default:
			break
		}
	}
}
