package main

import "fmt"

func sendData(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i // データを送信
	}
	close(ch) // チャネルを閉じる
}

func main() {
	ch := make(chan int) // チャネルを作成
	go sendData(ch)      // goroutineでデータを送信

	for data := range ch { // チャネルからデータを受信
		fmt.Println(data)
	}
}