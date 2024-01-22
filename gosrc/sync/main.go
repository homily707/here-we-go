package main

import (
	"fmt"
	"time"
)

func main() {
	nobuffer()
}

func nobuffer() {
	// 没有缓冲区的通道
	ch := make(chan int)

	go func() {
		for i := 0; i < 30; i = i + 2 {
			select {
			case ch <- i:
			default:
				fmt.Println("deprecated", i)
			}
			time.Sleep(3*time.Second)
		}
	}()
	go func() {
		for i := 1; i < 30; i = i + 2 {
			select {
			case ch <- i:
			default:
				fmt.Println("deprecated", i)
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		num := <- ch
		fmt.Println("output", num)
	}
}