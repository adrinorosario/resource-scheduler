package project

import (
	"bufio" // for accepting continuous inputs for developer choice
	"fmt"
	"os"
	"resource_scheduling/devteam"
	"strconv" // convert string to int for index capture
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func composeDevTeam() []devteam.Developer {

	// store the chosen dev team for the project
	// this is the slice that will be returned
	var chosenDevTeam []devteam.Developer

	// capture the indexes (ids) of the chosen devs from the org
	var capturedDevIdx []int

	fmt.Println(strings.Repeat("\n=", 40))
	fmt.Println("Choose the developers for your project")
	fmt.Println(strings.Repeat("=", 40))

	fmt.Println("\nList of organisational developers: ")
	for idx, developer := range devteam.OrgDevelopers {
		fmt.Printf("ID: %d, ", idx+1)
		fmt.Printf("Name: %s\n", developer.Name)
		fmt.Printf("Domain expertise: %s\n", developer.DomainExpertise)
		fmt.Print("Skillset:")
		for _, skill := range developer.SkillSet {
			fmt.Printf("%s, ", skill)
		}
	}

	fmt.Println("Enter the ID of the required developers (press Enter after each, or 'q' to quit once done): ")

	for {
		fmt.Printf("> ")

		// scanner.Scan() waits for the user to press enter
		if !scanner.Scan() {
			break // exit if the EOF or ctrl + z/d occurs
		}

		// trim whitespaces and new lines
		input := strings.TrimSpace(scanner.Text())

		// check if the input is 'q'
		if strings.EqualFold(input, "q") {
			break
		}

		// skip empty presses
		if input == "" {
			continue
		}

		// convert to int
		idx, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Enter a number (int) or 'q' to quit...")
			continue
		}

		if idx > len(devteam.OrgDevelopers) {
			fmt.Println("Invalid input. Enter a valid developer ID from the above list...")
			continue
		}

		capturedDevIdx = append(capturedDevIdx, idx)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input: ", err)
	}

	// append all the developers from devteam.OrgDevelopers
	for _, idx := range capturedDevIdx {
		developer := devteam.OrgDevelopers[idx-1]

		chosenDevTeam = append(chosenDevTeam, developer)
	}

	return chosenDevTeam
}
