package project

import "resource_scheduling/devteam"

// capture the current project stage
type ProjectStatus string

// applicable status for different projects
const (
	Initiating ProjectStatus = "Initiating"
	Planning   ProjectStatus = "Planning"
	Executing  ProjectStatus = "Executing"
	Closing    ProjectStatus = "Closing"
)

// capture the current project type
type ProjectType string

// applicable project types for teams
const (
	Feature        ProjectType = "Create/Modify a feature"
	Product        ProjectType = "New Product"
	Infrastructure ProjectType = "Working on infrastructure"
	Platform       ProjectType = "Platform toolings and pipelines"
	TechDebt       ProjectType = "Rewriting legacy code, refactoring codebases, or migrating infrastructure and projects"
	Refactoring    ProjectType = "Rewriting legacy code, refactoring codebases, or migrating infrastructure and projects"
	Migrations     ProjectType = "Rewriting legacy code, refactoring codebases, or migrating infrastructure and projects"
	KTLO           ProjectType = "Reliability, Scaling, and Maintenance"
	Security       ProjectType = "Security, Privacy, and Compliance"
)

// the structure of a given project
type ProjectSettings struct {
	ProjectID      int
	ProjectName    string
	ProjectType    ProjectType
	TeamSize       int
	ProgressStatus ProjectStatus
	DeveloperTeam  []devteam.Developer
}
