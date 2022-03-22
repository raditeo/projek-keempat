package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// go goroutine()

	fmt.Println("main execution started")

	go firstProcess(8)

	secondProcess(8)

	fmt.Println("No. of goroutines:", runtime.NumGoroutine())

	time.Sleep(time.Second * 1)

	fmt.Println("main execution ended")
}

// func goroutine() {
// 	fmt.Println("Yahoo")
// }

func firstProcess(index int) {
	fmt.Println("first process func started")
	for i := 1; i <= index; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("first process func ended")
}

func secondProcess(index int) {
	fmt.Println("second process func started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("second process func ended")
}
