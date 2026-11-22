package project

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

func createProject() {
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("Initialise new project")
	fmt.Println(strings.Repeat("=", 40))

	// record project init inputs
	var projectName, projectType, progressStatus string
	var teamSize, projectId int

	projectId := rand.IntN(100000000)

	fmt.Printf("Enter project name: ")
	fmt.Scan(&projectName)

	projectType = recordProjectType()
	teamSize = devTeamSize()
	projectStatus = "Initiating"


}

func recordProjectType() string {
	
	for {
		fmt.Println("Choose the project type from the following (input the correct number):")

		for idx, projType := range AllProjectTypes {
			fmt.Printf("%d:\t%d", idx+1, projType)
		}

		var projTypeOption int
		fmt.Printf("Enter project type: ")
		fmt.Scan(&projTypeOption)

		if projTypeOption <= 0 || projTypeOption > len(AllProjectTypes) {
			fmt.Print("Invalid choice. Select the correct option...")
			continue
		}

		var selectedProjectType := AllProjectTypes[projTypeOption-1]

		return selectedProjectType
	}

}

func devTeamSize() int {
	fmt.Println("Required team size: ")
	var teamSize
	fmt.Scan(&teamSize)

	for {
		if teamSize <= 2 {
			fmt.Println("You need at least 2 people in a team...")
			continue
		}
		return teamSize
	}
}