package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	numberOfGoroutine = 4
	taskLoad          = 10
	wg_10             sync.WaitGroup
)

func main() {
	wg_10.Add(numberOfGoroutine)
	tasks := make(chan int, taskLoad)
	for i := 0; i < numberOfGoroutine; i++ {
		go Worker(tasks, i)
	}
	for i := 0; i < taskLoad; i++ {
		tasks <- i
	}
	close(tasks)
	wg_10.Wait()
}
func Worker(tasks chan int, workerID int) {
	defer wg_10.Done()
	for {
		taskID, ok := <-tasks
		if !ok {
			fmt.Println("tasks channel closed,workID: ", workerID, " shutting down.")
			return
		}
		fmt.Println("I'm worker", workerID, "taskID:", taskID)
		sleep := rand.Intn(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
	}
}
