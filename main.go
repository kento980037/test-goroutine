package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	result1 := func1()

	// channel の初期化
	// 2個のバッファを持った channel を作成
	resChan := make(chan any, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func2(result1, resChan, &wg)
	go func3(result1, resChan, &wg)

	wg.Wait()

	// resChan channel への送信を終了し channel を閉じる
	close(resChan)

	// channel が閉じられるまでループする
	for res := range resChan {
		fmt.Println("res: ", res)
	}

	fmt.Println("took: ", time.Since(start))
}

func func1() string {
	time.Sleep(time.Millisecond * 50)
	return "result1"
}

func func2(s string, reschan chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 50)
	reschan <- "result2" + s
	wg.Done()
}

func func3(s string, reschan chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	reschan <- "result3" + s
	wg.Done()
}
