package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// タイムアウト付きのコンテキストを作成
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // 必ずキャンセル関数を呼ぶ

	// goroutineで処理を実行
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // タイムアウト時の処理
				fmt.Println("Context canceled:", ctx.Err())
				return
			default:
				fmt.Println("Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	// メイン関数で待機
	time.Sleep(3 * time.Second)
	fmt.Println("Main function finished")
}
