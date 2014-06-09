package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	flag "github.com/dotcloud/docker/pkg/mflag"
)

func main() {

	var (
		flVersion = flag.Bool([]string{"v", "-version"}, false, "Print version information and quit")
		flHelp    = flag.Bool([]string{"h", "-help"}, false, "Print this message")
		flDebug   = flag.Bool([]string{"-debug"}, false, "Run as DEBUG mode")
		flEditor  = flag.Bool([]string{"e", "-editor"}, false, "Use Editor by default")
		flForce   = flag.Bool([]string{"f", "-force"}, false, "Create README file without prompting")
		flCommand = flag.String([]string{"c", "-command"}, "", "Set Command to open README")
	)

	flag.Parse()

	if *flDebug {
		os.Setenv("DEBUG", "1")
	}

	if *flVersion {
		showVersion()
		os.Exit(0)
	}

	if *flHelp {
		flag.Usage()
		os.Exit(0)
	}

	cmd := retrieveCmd()

	if *flCommand != "" {
		cmd = *flCommand
	}

	if *flEditor {
		cmd = os.Getenv("EDITOR")
	}

	debug("cmd:", cmd)

	path := "."
	if len(os.Args) > 1 {
		path = flag.Arg(0)
	}
	debug("path:", path)

	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	debug("url:", absPath)

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		fmt.Printf("no such file or directory: %s\n", absPath)
		os.Exit(1)
	}

	os.Chdir(absPath)

	url := retrieveURL()
	if cmd == "" && url != "" {
		openByBrowser(url)
	}

	if cmd == "" {
		fmt.Printf("Could not retrieve project url from %s\n", absPath)
		os.Exit(0)
	}

	readmeFile := retrieveReadmeFile(*flForce)
	if readmeFile == "" {
		fmt.Println("Could not retrieve README from %s\n", absPath)
	}

	execOpen(cmd, readmeFile)
}

func showVersion() {
	fmt.Println("docc version 0.1.0")
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func retrieveCmd() string {
	cmd, err := gitConfig("docc.cmd")
	if err != nil {
		panic(err)
	}
	return cmd
}

func retrieveURL() string {

	urlPattern := "remote.origin.url" // Should be configurable
	gitConfigCmd := exec.Command("git", "config", "--local", urlPattern)

	url, err := gitConfigCmd.Output()
	if err != nil || len(url) == 0 {
		return ""
	}

	trimedURL := strings.TrimRight(string(url), "\n")
	return trimedURL
}

func retrieveReadmeFile(forceCreate bool) string {

	readmeFileName := "README.md" // Should be configurable
	readmeFile, err := filepath.Abs(readmeFileName)
	if err != nil {
		panic(err)
	}
	debug("readme:", readmeFile)

	if _, err := os.Stat(readmeFile); os.IsNotExist(err) && !forceCreate {
		fmt.Printf("%s is not found, create it ? [Y/n]: ", readmeFile)

		var ans string
		_, err := fmt.Scanf("%s", &ans)

		if err != nil {
			panic(err)
		}

		if ans != "Y" {
			return ""
		}
	}

	return readmeFile
}

func openByBrowser(url string) {
	openCmd := "open" // Only OSX or Linux
	execOpen(openCmd, url)
}

func execOpen(cmd string, target string) {

	binary, lookErr := exec.LookPath(cmd)
	if lookErr != nil {
		fmt.Printf("open command '%s' not found in $PATH\n", cmd)
		os.Exit(1)
	}

	args := []string{cmd, target}
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

func gitConfig(key string) (string, error) {
	cmd := exec.Command("git", "config", "--path", "--null", "--get", key)
	cmd.Stderr = os.Stderr

	buf, err := cmd.Output()

	if exitError, ok := err.(*exec.ExitError); ok {
		if waitStatus, ok := exitError.Sys().(syscall.WaitStatus); ok {
			if waitStatus.ExitStatus() == 1 {
				return "", nil
			}
		}

		return "", err
	}

	return strings.TrimRight(string(buf), "\000"), nil
}
