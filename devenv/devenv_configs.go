package devenv

// valid ide settings for each developer
type IDESettings struct {
	DeveloperName string
	WithTeam      bool
	TeamName      any
	Task          string
	IDE           string
}

// valid tasks for a developer; 			NEED TO EDIT THIS LIST
var validTasks = []string{
	"Application development",
	"Machine learning modelling",
	"Data analysis",
	"Deep learning (model training)",
	"Scripting",
	"Game development",
	"Mobile application development",
}
