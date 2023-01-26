package main

import (
	"fmt"
	"time"
)

/*
	A workflow engine based on Golang with the features like workflows, activities, workflow states, error handling,
	retries, and input-output parameters can be implemented using structs to represent the different components of
	the workflow, and a combination of goroutines, channels, and error handling to handle the parallel execution of
	tasks, and the flow of data between them.

	In this example, the Activity struct represents a single unit of work, and has a Run method that simulates the
	execution of the activity, it also has retries, input-output parameters, and a state, the Run method will check
	if the activity state is not completed
*/

// Activity represents a single unit of work
type Activity struct {
	Name       string
	Retries    int
	Input      interface{}
	Output     interface{}
	State      string
	lastError  error
	runCounter int
}

// Run runs the activity and returns the result
func (a *Activity) Run() error {
	a.runCounter++
	fmt.Printf("Running activity %s...\n", a.Name)
	// simulate work
	time.Sleep(time.Second)
	if a.runCounter < a.Retries {
		a.lastError = fmt.Errorf("error running activity %s", a.Name)
		a.State = "Error"
		return a.lastError
	}
	a.State = "Completed"
	fmt.Printf("Activity %s done\n", a.Name)
	return nil
}

// Workflow represents a series of activities
type Workflow struct {
	Name       string
	Activities []*Activity
}

// Run runs the workflow and returns the result
func (w *Workflow) Run() error {
	fmt.Printf("Running workflow %s...\n", w.Name)
	for _, activity := range w.Activities {
		for activity.State != "Completed" {
			err := activity.Run()
			if err != nil {
				fmt.Printf("Error running activity %s: %v\n", activity.Name, err)
				time.Sleep(time.Second * 2)
				continue
			}
			break
		}
	}
	fmt.Println("Workflow completed")
	return nil
}

func main() {
	// define activities
	activity1 := &Activity{Name: "Activity 1", Retries: 3}
	activity2 := &Activity{Name: "Activity 2", Retries: 3}
	activity3 := &Activity{Name: "Activity 3", Retries: 3}

	// define workflow
	workflow := &Workflow{
		Name:       "Example Workflow",
		Activities: []*Activity{activity1, activity2, activity3},
	}

	// run workflow
	err := workflow.Run()
	if err != nil {
		fmt.Printf("Error running workflow: %v\n", err)
	}
}
