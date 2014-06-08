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
	)

	flag.Parse()

	if *flDebug {
		os.Setenv("DEBUG", "1")
	}

	if *flVersion {
		showVersion()
		fmt.Println(os.Args[1])
		os.Exit(0)
	}

	if *flHelp {
		flag.Usage()
		os.Exit(0)
	}

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
	if !*flEditor && url != "" {
		openByBrowser(url)
	}

	readmeFile := retrieveReadmeFile(*flForce)
	if readmeFile != "" {
		openByEdior(readmeFile)
	}

	os.Exit(0)
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
