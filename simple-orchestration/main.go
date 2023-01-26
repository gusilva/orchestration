package main

import "fmt"

/*
A workflow engine in Golang can be implemented using a combination of goroutines and channels to handle the parallel
execution of tasks.

In this example, the Task struct represents a single unit of work, and has a Run method that simulates the execution of
the task. The main function creates a slice of tasks and runs them concurrently using goroutines. The results of the
tasks are collected using a channel, which allows the main function to wait for all tasks to complete before exiting.
*/

// Task represents a single unit of work
type Task struct {
	Name string
}

// Run runs the task and returns the result
func (t Task) Run() string {
	fmt.Printf("Running task %s...\n", t.Name)
	// simulate work
	return t.Name + " done"
}

func main() {
	// define tasks
	tasks := []Task{
		{Name: "Task 1"},
		{Name: "Task 2"},
		{Name: "Task 3"},
	}

	// create a channel to receive task results
	results := make(chan string)

	// run tasks concurrently
	for _, task := range tasks {
		go func(t Task) {
			results <- t.Run()
		}(task)
	}

	// collect results
	for range tasks {
		fmt.Println(<-results)
	}
}
