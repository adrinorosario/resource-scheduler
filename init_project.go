package main

import (
	"fmt"
	"resource_scheduling/devenv"
)

type IDESettings struct {
	DeveloperName string
	WithTeam      bool
	TeamName      any
	Task          string
	IDE           string
}

var validTasks = []string{
	"Application development",
	"Machine learning modelling",
	"Data analysis",
	"Deep learning (model training)",
	"Scripting",
	"Game development",
	"Mobile application development",
}

func chooseStartingPoint() bool {
	fmt.Println("Choose the application being used (enter the correct number):")

	var currentIDE int
	var currentTask int

	for index, val := range devenv.ValidApplicationIdes {
		fmt.Printf("%v\t%v\n", index+1, val)
	}
	fmt.Println("Enter your choice: ")

	_, ideErr := fmt.Scan(&currentIDE)
	if ideErr != nil || currentIDE <= 0 || currentIDE > len(devenv.ValidApplicationIdes) {
		fmt.Println("Incorrect input choice. Enter the correct input number and try again...")
		return false
	}

	fmt.Printf("Chosen IDE: %v\n", devenv.ValidApplicationIdes[currentIDE-1])

	fmt.Println("Enter the task you are performing (enter the correct number):")
	for index, task := range validTasks {
		fmt.Printf("%v\t%v\n", index+1, task)
	}
	fmt.Println("Enter your choice: ")

	_, taskErr := fmt.Scan(&currentTask)
	if taskErr != nil || currentTask <= 0 || currentTask > len(validTasks) {
		fmt.Println("Incorrect input choice. Enter the correct input number and try again...")
		return false
	}

	fmt.Printf("Chosen task: %v\n", validTasks[currentTask-1])

	var validTaskName string = validTasks[currentTask-1]
	var validIdeName string = devenv.ValidApplicationIdes[currentIDE-1]

	return devenv.ValidateIDEChoice(validTaskName, validIdeName)
}
