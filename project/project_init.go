package project

import (
	"fmt"
	"math/rand/v2"
	"resource_scheduling/devteam"
	"strings"
)

func createProject() {
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("Initialise new project")
	fmt.Println(strings.Repeat("=", 40))

	// record project init inputs
	var projectName, projectType, progressStatus string
	var teamSize, projectId int
	var devTeam []devteam.Developer

	projectId = rand.IntN(100000000)

	fmt.Printf("Enter project name: ")
	fmt.Scan(&projectName)

	projectType = recordProjectType()
	teamSize = devTeamSize()
	progressStatus = "Initiating"

	// choose dev team for the project
	devTeam = composeDevTeam(teamSize) // pass max size for the team

	var project = ProjectSettings{
		ProjectID:      projectId,
		ProjectName:    projectName,
		ProjectType:    ProjectType(projectType),
		ProgressStatus: ProjectStatus(progressStatus),
		TeamSize:       teamSize,
		DeveloperTeam:  devTeam,
	}

	Projects = append(Projects, project)
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

		var selectedProjectType string = AllProjectTypes[projTypeOption-1]

		return selectedProjectType
	}

}

func devTeamSize() int {
	fmt.Println("Required team size: ")
	var teamSize int
	fmt.Scan(&teamSize)

	for {
		if teamSize <= 2 {
			fmt.Println("You need at least 2 people in a team...")
			continue
		}
		return teamSize
	}
}
