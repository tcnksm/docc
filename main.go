package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func main() {

	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		fmt.Printf("no such file or directory: %s\n", absPath)
	}

	os.Chdir(absPath)

	url := retrieveURL()
	if url != "" {
		openByBrowser(url)
	}

	readmeFileName := "README.md" // Should be configurable
	readmeFile, err := filepath.Abs(readmeFileName)
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(readmeFile); os.IsNotExist(err) {
		fmt.Printf("%s is not found, create it ? [Y/n]: ", readmeFile)

		var ans string
		_, err := fmt.Scanf("%s", &ans)

		if err != nil {
			panic(err)
		}

		if ans != "Y" {
			os.Exit(0)
		}
	}

	openByEdior(readmeFile)
}

func retrieveURL() string {

	urlPattern := "remote.origin.url" // Should be configurable
	gitConfigCmd := exec.Command("git", "config", "--local", urlPattern)

	url, err := gitConfigCmd.Output()
	if err != nil || len(url) == 0 {
		return ""
	}

	trimedURL := strings.TrimRight(string(url), "\n")
	fmt.Println("url: " + trimedURL)
	return trimedURL
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
