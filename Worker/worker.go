package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // ジョブ処理に1秒
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// ワーカーを3つ作成
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// ジョブを送信
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 結果を受信
	for a := 1; a <= numJobs; a++ {
		fmt.Println("Result:", <-results)
	}
}