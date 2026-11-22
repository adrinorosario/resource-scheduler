package devenv

import (
	"fmt"
	"strings"
)

func verifyTargetMobilePlatform(ideName string, platformName string) bool {
	switch platformName {
	case "iOS":
		if strings.EqualFold(ideName, "Xcode") {
			return true
		}
		fmt.Println("Cannot target Apple platforms without using XCode.")
		return false

	case "Android":

		if strings.EqualFold(ideName, "Android studio") || strings.EqualFold(ideName, "IntelliJ") {
			return true
		}
		fmt.Printf("Cannot develop for Android on %v. Choose a valid platform like Android studio or IntelliJ", ideName)
		return false

	case "Cross-platform":

		if strings.EqualFold(ideName, "Google colab local runtime") || strings.EqualFold(ideName, "Jupyter notebook") || strings.EqualFold(ideName, "BlueJ") || strings.EqualFold(ideName, "Pycharm") || strings.EqualFold(ideName, "UnityIDE") || strings.EqualFold(ideName, "Terminal") {
			fmt.Printf("Cannot perform cross-platform development using %v. Choose another IDE.", ideName)
			return false
		}
		return true

	default:
		fmt.Printf("Cannot match platform %v", platformName)
		return false
	}
}
