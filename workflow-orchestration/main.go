package main

import (
	"fmt"
)

/*
	In this example, the Activity struct represents a single unit of work, and has a Run method that simulates the
	execution of the activity. The Workflow struct represents a series of activities and a name, and has a Run method
	that runs the activities in the order they are defined in the Activities slice. The main function creates a slice
	of activities, a workflow and runs the workflow.
*/

// Activity represents a single unit of work
type Activity struct {
	Name string
}

// Run runs the activity and returns the result
func (a Activity) Run() string {
	fmt.Printf("Running activity %s...\n", a.Name)
	// simulate work
	return a.Name + " done"
}

// Workflow represents a series of activities
type Workflow struct {
	Name       string
	Activities []Activity
}

// Run runs the workflow and returns the result
func (w Workflow) Run() string {
	fmt.Printf("Running workflow %s...\n", w.Name)
	var results []string
	for _, activity := range w.Activities {
		results = append(results, activity.Run())
	}
	return "Workflow " + w.Name + " done. Results: " + fmt.Sprint(results)
}

func main() {
	// define activities
	activities := []Activity{
		{Name: "Activity 1"},
		{Name: "Activity 2"},
		{Name: "Activity 3"},
	}

	// define workflow
	workflow := Workflow{
		Name:       "Example Workflow",
		Activities: activities,
	}

	// run workflow
	result := workflow.Run()
	fmt.Println(result)
}
