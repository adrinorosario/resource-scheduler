package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting application...")

	for {
		// choose the starting point of your development journey
		var startingPointStatus bool = chooseStartingPoint()
		if !startingPointStatus {
			continue
		} else {
			fmt.Println("Starting point of application development environment ready...(exit for now)")
			break
		}
	}
}
