package project

import "resource_scheduling/devteam"

type ProjectSettings struct {
	ProjectName string
	ProjectLead string
	ProjectID   int
	Developers  []devteam.Developer
}
