package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello, World!")
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go sayHello() // goroutineで実行
	for i := 0; i < 5; i++ {
		fmt.Println("Main Function")
		time.Sleep(300 * time.Millisecond)
	}
}