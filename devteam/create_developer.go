package devteam

import (
	"fmt"
	"strings"
)

type Developer struct {
	ID       int
	Name     string
	Level    int
	SkillSet []string
}

var teamDevelopers []Developer // to store all the developers details in the team

func createDeveloper() {
	fmt.Println("Create a new developer profile")

	var id, level int
	var name, skills string

	fmt.Println("Enter developer ID: ")
	fmt.Scan(&id)

	fmt.Println("Enter name: ")
	fmt.Scan(&name)

	fmt.Println("Enter level (integer): ")
	fmt.Scan(&level)

	fmt.Println("Enter the skillset (languages, frameworks, tools) separated by space: ")
	fmt.Scan(&skills)

	skillset := strings.Fields(skills)

	var developer Developer = Developer{id, name, level, skillset}
	teamDevelopers = append(teamDevelopers, developer)

}
