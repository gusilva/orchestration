package main

import (
	"fmt"
	"time"
)

// Workflow struct represents a workflow
type Workflow struct {
	Name       string
	Activities []*Activity
	State      string
	Retries    int
}

// Activity struct represents an activity in a workflow
type Activity struct {
	Name         string
	InputParams  map[string]interface{}
	OutputParams map[string]interface{}
	Execute      func() error
	Retries      int
}

// Start starts the workflow and its activities
func (w *Workflow) Start() error {
	w.State = "Running"
	for _, activity := range w.Activities {
		err := activity.Execute()
		if err != nil {
			w.State = "Error"
			return err
		}
		// Check if the activity has any retries, if so, execute the activity again
		for i := 0; i < activity.Retries; i++ {
			err = activity.Execute()
			if err == nil {
				break
			}
			time.Sleep(time.Second * 2)
		}
		if err != nil {
			w.State = "Error"
			return err
		}
	}
	w.State = "Completed"
	return nil
}

func main() {
	// Create a new workflow
	myWorkflow := &Workflow{
		Name: "My Workflow",
		Activities: []*Activity{
			{
				Name: "Activity 1",
				InputParams: map[string]interface{}{
					"param1": "value1",
				},
				Retries: 3,
				Execute: func() error {
					fmt.Println("Executing Activity 1")
					return nil
				},
			},
			{
				Name: "Activity 2",
				InputParams: map[string]interface{}{
					"param1": "value1",
					"param2": "value2",
				},
				Retries: 3,
				Execute: func() error {
					fmt.Println("Executing Activity 2")
					return nil
				},
			},
		},
	}

	// Start the workflow
	err := myWorkflow.Start()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Workflow completed successfully")
	}
}

//// Workflow struct represents a workflow
//type Workflow struct {
//	Name       string
//	Activities []*Activity
//	State      string
//	Retries    int
//}
//
//// Activity struct represents an activity in a workflow
//type Activity struct {
//	Name         string
//	InputParams  map[string]interface{}
//	OutputParams map[string]interface{}
//	Execute      func() error
//	Retries      int
//	Retried      int
//}
//
//// Start starts the workflow and its activities
//func (w *Workflow) Start() error {
//	w.State = "Running"
//	for _, activity := range w.Activities {
//		for activity.Retried <= activity.Retries {
//			err := activity.Execute()
//			if err != nil {
//				activity.Retried++
//				time.Sleep(time.Second)
//				continue
//			}
//			break
//		}
//		if activity.Retried > activity.Retries {
//			w.State = "Error"
//			return fmt.Errorf("activity %s failed after %d retries", activity.Name, activity.Retries)
//		}
//	}
//	w.State = "Completed"
//	return nil
//}
//
//func main() {
//	// Create a new workflow
//	myWorkflow := &Workflow{
//		Name: "My Workflow",
//		Activities: []*Activity{
//			{
//				Name: "Activity 1",
//				Execute: func() error {
//					fmt.Println("Executing Activity 1")
//					// Simulating an error
//					return fmt.Errorf("Error in Activity 1")
//				},
//				Retries: 2,
//			},
//			{
//				Name: "Activity 2",
//				Execute: func() error {
//					fmt.Println("Executing Activity 2")
//					return nil
//				},
//				Retries: 3,
//			},
//		},
//		Retries: 3,
//	}
//
//	// Start the workflow
//	err := myWorkflow.Start()
//	if err != nil {
//		fmt.Println("Error:", err)
//	} else {
//		fmt.Println("Workflow completed successfully")
//	}
//}

//import (
//	"fmt"
//	"time"
//)
//
//package main
//
//import (
//"fmt"
//"time"
//)
//
//// Workflow struct represents a workflow
//type Workflow struct {
//	Name       string
//	Activities []Activity
//	State      string
//}
//
//// Activity struct represents an activity in a workflow
//type Activity struct {
//	Name        string
//	InputParams map[string]interface{}
//	OutputParams map[string]interface{}
//	Execute     func() error
//	Retry       int
//}
//
//// Start starts the workflow and its activities
//func (w *Workflow) Start() error {
//	w.State = "Running"
//	for i, activity := range w.Activities {
//		err := activity.Execute()
//		for j := 0; j < activity.Retry; j++ {
//			if err != nil {
//				fmt.Printf("Error executing activity %s, retrying...\n", activity.Name)
//				time.Sleep(2 * time.Second)
//				err = activity.Execute()
//			} else {
//				break
//			}
//		}
//		if err != nil {
//			w.State = "Error"
//			return err
//		}
//		// Pass output parameters of the current activity to the next activity
//		if i < len(w.Activities)-1 {
//			for key, value := range activity.OutputParams {
//				w.Activities[i+1].InputParams[key] = value
//			}
//		}
//	}
//	w.State = "Completed"
//	return nil
//}
//
//func main() {
//	// Create a new workflow
//	myWorkflow := &Workflow{
//		Name: "My Workflow",
//		Activities: []Activity{
//			{
//				Name: "Activity 1",
//				InputParams: map[string]interface{}{
//					"param1": "value1",
//				},
//				OutputParams: map[string]interface{}{
//					"param2": "value2",
//				},
//				Execute: func() error {
//					fmt.Println("Executing Activity 1")
//					// check for errors
//					if val, ok := InputParams["param1"]; ok {
//						if val != "value1" {
//							return fmt.Errorf("Invalid input param1 value: %v", val)
//						}
//					}
//					return nil
//				},
//				Retry: 2,
//			},
//			{
//				Name: "Activity 2",
//				InputParams: map[string]interface{}{
//					"param3": "value3",
//				},
//				OutputParams: map[string]interface{}{
//					"param4": "value4",
//				},
//				Execute:
