package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	fmt.Println("docc")
	// 1) git project and it has github page?
	//    -> open it browser
	configValue := "remote.origin.url"
	gitConfigCmd := exec.Command("git", "config", "--local", configValue)

	url, err := gitConfigCmd.Output()
	if err == nil && len(url) != 0 {
		u := strings.TrimRight(string(url), "\n")
		openByBrowser(u)
	}

	// 2) There is not github page but it has README(.md*)
	//    -> opent it $EDITOR
	readmeFile := "README.md"
	openByEdior(readmeFile)
}

func openByBrowser(url string) {
	openCmd := "open" // Only OSX or Linux
	execOpen(openCmd, url)
}

func openByEdior(filename string) {
	openCmd := os.Getenv("EDITOR")
	execOpen(openCmd, filename)
}

func execOpen(cmd string, target string) {

	binary, lookErr := exec.LookPath(cmd)
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{cmd, target}
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
