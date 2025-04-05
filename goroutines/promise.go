package main

//js equivalent of promise.then
import (
	"fmt"
	"time"
)

// asyncTask is like a Promise-returning function in JS
func asyncTask(name string, resultChan chan<- string) {
	time.Sleep(2 * time.Second)
	resultChan <- name + " done"
}

func main() {
	result := make(chan string)

	go asyncTask("Task 1", result)

	// This is like `.then()` in JS
	fmt.Println(<-result)
}
