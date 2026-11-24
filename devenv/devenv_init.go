package devenv

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// define command line flags
var (
	projectName      = flag.String("name", "", "Name of the project")
	projectDirectory = flag.String("dir", "", "Directory of the project")
)

func Init() {
	// parse the flags if they haven't been parsed
	if !flag.Parsed() {
		flag.Parse()
	}

	// check if the parsed inputs are empty
	if *projectName == "" || *projectDirectory == "" {
		log.Fatalf("Usage: -%s -name=<project_name> -dir=<project_directory>", os.Args[0])
	}

	fmt.Println("Initlialising project development environment...")
	fmt.Printf("Initialising Project %s in %s", *projectName, *projectDirectory)

	// cd into the project directory
	// check pwd
	pwd, pwdErr := os.Getwd()
	if pwdErr != nil {
		log.Fatal(pwdErr)
	}

	// initialise project directory
	projectInitDirectory := filepath.Join(pwd, *projectDirectory)

	// creates the directory if it does not exist
	// 0755 gives read/write/execute to owner, and read/execute to group/others
	projectDirErr := os.Mkdir(projectInitDirectory, 0755)

	if projectDirErr != nil {
		log.Fatalf("Failed to create directory: %v", projectDirErr)
	}

	fmt.Printf("%v present/created", projectInitDirectory)

	// cd into the project directory
	changeDirToProjectDirCmd := exec.Command("cd", projectInitDirectory)

	// execute cd command and check for error
	cdCheckOutput, cdCheckErr := changeDirToProjectDirCmd.CombinedOutput()
	if cdCheckErr != nil {
		log.Fatalf("cd command failed: %v\nOutput: %s", cdCheckErr, cdCheckOutput)
	} else {
		fmt.Printf("Changed directory to %v", projectInitDirectory)
		// CHECK PWD HERE AND CONFIRM IF DIR WAS CHANGED
	}

	// check if git is present on the system
	gitCheckCmd := exec.Command("git", "--version")

	// run the command and capture the standard out and error
	gitCheckOutput, gitCheckError := gitCheckCmd.CombinedOutput()
	if gitCheckError != nil {
		log.Fatalf("Command failed: %v\nOutput: %s", gitCheckError, gitCheckOutput)
	} else {
		// if you enter this block, it means git is present in the system
		// here, programmatically create a new project directory in the name
		// of the initialised project as well as the project id
		// cd into the directory, and intitialise a git repository and make the first
		// initial commit
		fmt.Printf("Output: %s", gitCheckOutput)
	}
}
