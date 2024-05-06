package file

import (
	"fmt"
	"os"
	"strings"
)

// ask user if option is needed and return bool value
func AskIfNeeded(option string) bool {
	fmt.Print("is ", option, " needed? [Y/n]: ")
	var answer string
	fmt.Scanln(&answer)
	answer = strings.ToLower(answer)
	switch answer {
	case "y", "yes", "":
		return true
	case "n", "no":
		return false
	default:
		fmt.Println("Wrong answer")
		os.Exit(1)
	}
	return false
}

// ask user for all custom values and return as map
func GetUserValues(filename string, userValues map[string]string) map[string]string {
	for argName := range userValues {
		userValues[argName] = askValue(filename, argName)
	}
	return userValues
}

// ask user to enter value
func askValue(filename, argName string) string {
	fmt.Print("Enter value for", argName, " in file ", filename, ":")
	var answer string
	_, err := fmt.Scanln(&answer)
	if err != nil {
		fmt.Println("can't read answer")
		os.Exit(1)
	}
	return answer
}
