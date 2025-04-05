package main

import (
	"errors"
	"fmt"
	"time"
)

func task(name string, duration time.Duration, shouldFail bool, successChan chan string, errChan chan error) {
	time.Sleep(duration)
	if shouldFail {
		errChan <- errors.New(name + " failed")
	} else {
		successChan <- name + " success"
	}
}

func main() {
	successChan := make(chan string)
	errChan := make(chan error)
	totalTasks := 3
	failureCount := 0

	go task("Task 1", 500*time.Millisecond, true, successChan, errChan)
	go task("Task 2", 1*time.Second, false, successChan, errChan)
	go task("Task 3", 2*time.Second, false, successChan, errChan)

	for {
		select {
		case result := <-successChan:
			fmt.Println("✅ First success:", result)
			return
		case err := <-errChan:
			fmt.Println("❌", err)
			failureCount++
			if failureCount == totalTasks {
				fmt.Println("❌ All tasks failed")
				return
			}
		}
	}
}
