package commands

import (
	"fmt"
	"os"
	"os/exec"

	skyclilib "github.com/avran02/sky-cli-lib"
	"github.com/avran02/sky-cli/generator_modules/file"
)

// execute all commands
func ExecAll(commands []skyclilib.OsCommand) error {
	for _, cmd := range commands {
		args := getArgs(cmd)
		err := execCmd(cmd.Name, args) // TODO: add retry
		if err != nil {
			fmt.Println("can't execute command", cmd.Name, args)
			return fmt.Errorf("error: %w", err)
		}
		fmt.Println("command", cmd.Name, "with args", args, "executed")
		fmt.Println("")
	}
	return nil
}

func getArgs(cmd skyclilib.OsCommand) []string {
	commandArgs := make([]string, len(cmd.Args))
	for i, arg := range cmd.Args {
		switch arg.Source.Get() {
		case "FromUser":
			commandArgs[i] = askValue(cmd.Name, arg.Name, commandArgs)
		case "FromUserBool":
			if file.AskIfNeeded(arg.Name) {
				commandArgs[i] = arg.Name
			}
		case "FromPlugin":
			commandArgs[i] = arg.Value
		default:
			fmt.Println("wrong type")
			os.Exit(1)
		}
	}
	return commandArgs
}

// get arguments from user
// func getArgs(cmd skyclilib.OsCommand) []string {
// 	commandArgs := make([]string, 0)
// 	for _, arg := range cmd.Args {
// 		if arg.NeedGetFromUser {
// 			commandArgs = append(commandArgs, askValue(cmd.Name, arg.Name, commandArgs))
// 		} else {
// 			commandArgs = append(commandArgs, arg.Name)
// 		}
// 	}
// 	return commandArgs
// }

// execute os command
func execCmd(cmd string, args []string) error {
	err := exec.Command(cmd, args...).Run() // #nosec
	if err != nil {
		fmt.Println("can't execute command", cmd, args)
		fmt.Println("try again please")
		return err
	}
	return nil
}

// ask user for value and return it or exit with error
func askValue(executingCommand, argName string, previousArgs []string) string {
	fullCmd := "'" + executingCommand
	for _, arg := range previousArgs {
		fullCmd += " " + arg
	}
	fullCmd += "': "
	fmt.Print("Enter ", argName, " for ", fullCmd)
	var answer string
	_, err := fmt.Scanln(&answer)
	if err != nil {
		fmt.Println("can't read answer")
		os.Exit(1)
	}
	return answer
}
