package plugins

import (
	"fmt"
	"os"
	"os/exec"
)

// must be tagged and you must choose version
func Install(url string) {
	goGetArgs := []string{"mod", "download", url}
	err := execCmd("go", goGetArgs)
	if err != nil {
		fmt.Printf("can't execute 'go mod download %s'\n", url)
		fmt.Println("try again please")
		os.Exit(1)
	}
	err = os.Chdir(getGoPkgRoot() + "/" + url)
	if err != nil {
		fmt.Println("can't change dir:", err)
		os.Exit(1)
	}
	goBuildArgs := []string{"build", "-o", plaginsRoot, "-buildmode", "plugin", getGoPkgRoot() + "/" + url}
	err = execCmd("go", goBuildArgs)
	if err != nil {
		fmt.Println("can't build plugin:", err)
		os.Exit(1)
	}
}

// execute os command
func execCmd(cmd string, args []string) error {
	err := exec.Command(cmd, args...).Run() // #nosec
	if err != nil {
		fmt.Println("can't execute command", cmd, "with args", args)
		fmt.Println("try again please")
		return err
	}
	return nil
}

// get GOPATH or HOME/go if GOPATH is empty
func getGoPkgRoot() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = os.Getenv("HOME") + "/go"
	}
	pth := gopath + "/pkg/mod"
	return pth
}
