package main

import (
	. "fmt"
	"runtime"
	"time"
)

var i = 0

func goroutine1() {
	for j := 0; j < 1000000; j++ {
		i++
	}
}

func goroutine2() {
	for k := 0; k < 1000000; k++ {
		i--
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go goroutine1()
	go goroutine2()

	time.Sleep(100 * time.Millisecond)
	Println(i)

}
