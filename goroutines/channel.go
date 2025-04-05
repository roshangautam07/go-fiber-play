package main

import (
	"fmt"
)

func producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("Send:", i)

	}
	close(ch)
}
func consumer(ch <-chan int, done chan<- bool) {
	for value := range ch {
		fmt.Println("Received:", value)
	}
	done <- true
}
func main() {
	ch := make(chan int)
	done := make(chan bool)
	go producer(ch)
	go consumer(ch, done)
	<-done
}
