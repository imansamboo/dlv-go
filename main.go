package main

import (
	"fmt"
	"time"
)

// doWork simulates a task by printing a start message, sleeping for a specified duration,
// and then printing a completion message.
//
// Parameters:
// task (string): The name of the task to be performed.
// seconds (int): The number of seconds to sleep before completing the task.
//
// Return:
// This function does not return any value.
func doWork(task string, seconds int) {
	fmt.Printf("Starting task: %s\n", task)
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Printf("Completed task: %s\n", task)
}

func main() {
	fmt.Println("Program started")

	task1 := "Download files"
	task2 := "Process data"

	// Simulate work
	doWork(task1, 2)
	doWork(task2, 3)
	task3 := task1 + " and " + task2

	doWork(task3, 1)

	fmt.Println("All tasks completed")
}
