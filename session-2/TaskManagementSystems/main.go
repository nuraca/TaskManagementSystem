package main

import (
	"fmt"
	"time"
)

func main() {
	// Constants
	const TotalTasks = 100
	const (
		Low = iota
		Medium
		High
	)

	// Variables
	projectName := "Task Management System"
	projectStatus := "IN PROGRESS"
	tasksCreated := 25
	projectCompleted := false
	startDate := time.Date(2024, 9, 18, 0, 0, 0, 0, time.UTC)

	// Convert the boolean to a string for output
	projectCompletedStr := fmt.Sprintf("%t", projectCompleted)

	// Output message
	fmt.Println("Welcome to the Task Management System!")
	fmt.Printf("Project start date is: %s\n", startDate.Format("2006-01-02 15:04:05"))
	fmt.Printf("Project: %s\n\n", projectName)

	fmt.Printf("Current project status: %s\n", projectStatus)
	fmt.Printf("Tasks completed: %d out of %d\n", tasksCreated, TotalTasks)
	fmt.Printf("Task priorities: %d-Low, %d-Medium, %d-High\n", Low, Medium, High)
	fmt.Printf("Is the project completed? %s\n", projectCompletedStr)
}
