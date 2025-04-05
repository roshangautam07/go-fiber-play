// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// // Simulate fetching user data from an API
// func fetchUserData(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Fetching user data...")
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("User data fetched ✅")
// }

// // Simulate image resizing
// func resizeImage(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Resizing image...")
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("Image resized ✅")
// }

// // Simulate saving logs to a file
// func saveLogs(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Saving logs to file...")
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Logs saved ✅")
// }

// func main() {
// 	var wg sync.WaitGroup

// 	wg.Add(3) // 3 tasks

// 	go fetchUserData(&wg)
// 	go resizeImage(&wg)
// 	go saveLogs(&wg)

// 	wg.Wait() // Wait for all tasks to finish
// 	fmt.Println("✅ All tasks completed!")
// }

package main

import (
	"fmt"
	"sync"
	"time"
)

// Define task type
type taskFunc func(wg *sync.WaitGroup)

// Simulate fetching user data from an API
func fetchUserData(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Fetching user data...")
	time.Sleep(2 * time.Second)
	fmt.Println("User data fetched ✅")
}

// Simulate image resizing
func resizeImage(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Resizing image...")
	time.Sleep(3 * time.Second)
	fmt.Println("Image resized ✅")
}

// Simulate saving logs to a file
func saveLogs(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Saving logs to file...")
	time.Sleep(1 * time.Second)
	fmt.Println("Logs saved ✅")
}

func main() {
	var wg sync.WaitGroup

	// List of tasks
	tasks := []taskFunc{
		fetchUserData,
		resizeImage,
		saveLogs,
	}

	for _, task := range tasks {
		wg.Add(1)
		go task(&wg)
	}

	wg.Wait()
	fmt.Println("✅ All dynamic tasks completed!")
}
