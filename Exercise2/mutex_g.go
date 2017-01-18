package main

import (
	. "fmt"
	"runtime"
	"time"
)

func goroutine1(ch chan int) {
	for j := 0; j < 1000000; j++ {
		i := <-ch
		i++
		ch <- i
	}
}

func goroutine2(ch chan int) {
	for k := 0; k < 1000000-1; k++ {
		i := <-ch
		i--
		ch <- i
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ch := make(chan int, 1)
	ch <- 0
	go goroutine1(ch)
	go goroutine2(ch)

	time.Sleep(800 * time.Millisecond)
	i := <-ch
	Println(i)

}
