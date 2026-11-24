package devenv

import (
	"fmt"
)

func chooseStartingPoint() bool {
	fmt.Println("Choose the application being used (enter the correct number):")

	var currentIDE int
	var currentTask int

	for index, val := range ValidApplicationIdes {
		fmt.Printf("%v\t%v\n", index+1, val)
	}
	fmt.Println("Enter your choice: ")

	_, ideErr := fmt.Scan(&currentIDE)
	if ideErr != nil || currentIDE <= 0 || currentIDE > len(ValidApplicationIdes) {
		fmt.Println("Incorrect input choice. Enter the correct input number and try again...")
		return false
	}

	fmt.Printf("Chosen IDE: %v\n", ValidApplicationIdes[currentIDE-1])

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
	var validIdeName string = ValidApplicationIdes[currentIDE-1]

	return ValidateIDEChoice(validTaskName, validIdeName)
}
