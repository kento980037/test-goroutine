package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	result1 := func1()
	result2 := func2(result1)
	result3 := func3(result1)

	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println("took: ", time.Since(start))
}

func func1() string {
	time.Sleep(time.Millisecond * 50)
	return "result1"
}

func func2(s string) string {
	time.Sleep(time.Millisecond * 50)
	return "result2" + s
}

func func3(s string) string {
	time.Sleep(time.Millisecond * 100)
	return "result3" + s
}
