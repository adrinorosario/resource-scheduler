package devteam

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

// list of approved domains of expertise for a developer
var ApprovedDomainsOfExpertise = []string{
	"Backend",
	"Frontend",
	"Mobile",
	"Infrastructure",
	"AI/ML",
}

// a typical developer profile needed for projects
type Developer struct {
	ID              int
	Name            string
	Level           int
	DomainExpertise string
	Interests       []string
	SkillSet        []string
}

func createDeveloper() {
	fmt.Println("Create a new developer profile")

	var id, level int
	var name, skills, interests, domainExpertise string
	var domainCaputure int

	id = rand.IntN(100000000)

	fmt.Printf("Enter developer name: ")
	fmt.Scan(&name)

	fmt.Println("Enter engineering level (integer): ")
	fmt.Scan(&level)

	fmt.Println("Select domain of expertise from the following list: ")
	for idx, val := range ApprovedDomainsOfExpertise {
		fmt.Printf("%d:\t%d", idx+1, val)
	}

	for {
		fmt.Printf("Enter the correct number: ")
		fmt.Scan(&domainCaputure)

		if domainCaputure > len(ApprovedDomainsOfExpertise) {
			fmt.Println("Wrong choice. Try again...")
			continue
		}
		break
	}

	domainExpertise = ApprovedDomainsOfExpertise[domainCaputure-1]

	fmt.Println("Enter the skillset (languages, frameworks, tools) separated by space: ")
	fmt.Scan(&skills)

	fmt.Println("What are your other interests (can be technical and non-technical): ")
	fmt.Scan(&interests)

	skillset := strings.Fields(skills)
	interestsList := strings.Fields(interests)

	var developer Developer = Developer{id, name, level, domainExpertise, interestsList, skillset}

	// append to global list of org's developers
	OrgDevelopers = append(OrgDevelopers, developer)
}
