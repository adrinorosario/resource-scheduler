package devenv

import (
	"fmt"
	"slices"
	"strings"
)

var ValidApplicationIdes = []string{
	"VS Code",
	"BlueJ",
	"IntelliJ",
	"Pycharm",
	"Jupyter Notebook",
	"Google Colab local runtime",
	"Android Studio",
	"XCode",
	"UnityIDE",
	"Terminal",
}

var IdeTaskMapped = map[string][]string{
	"Application development":        {"VS Code", "BlueJ", "IntelliJ", "Pycharm", "Android Studio", "XCode"},
	"Machine learning modelling":     {"VS Code", "Pycharm", "Jupyter Notebook", "Google Colab local runtime"},
	"Data analysis":                  {"Pycharm", "Jupyter Notebook", "Google Colab local runtime", "VS Code"},
	"Deep learning (model training)": {"Jupyter Notebook", "Google Colab local runtime", "VS Code"},
	"Scripting":                      {"VS Code", "Terminal"},
	"Game development":               {"UnityIDE", "IntelliJ"},
	"Mobile application development": {"XCode", "Android Studio"},
}

func ValidateIDEChoice(taskName string, ideName string) bool {
	if strings.EqualFold(taskName, "Mobile application development") {
		fmt.Println("Enter the platform you are developing for (iOS or Android):")

		var platform string
		_, err := fmt.Scan(&platform)

		if err != nil {
			return verifyTargetMobilePlatform(ideName, platform)
		} else {
			return false
		}
	} else if slices.Contains(IdeTaskMapped[taskName], ideName) {
		return true
	} else {
		return false
	}
}
