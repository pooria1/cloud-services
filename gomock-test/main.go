package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("goroutine job finished!")
	}()

	defer printHello()
	time.Sleep(time.Second * 6)
}

func printHello() {
	fmt.Println("hello boy!!")
}
