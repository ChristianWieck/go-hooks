package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// First we need to make sure golines and gofumpt are installed
	exec.Command("go", "install", "github.com/segmentio/golines@latest")
	exec.Command("go", "install", "mvdan.cc/gofumpt@latest")

	// First check IF golines would reformat stuff
	cmd := exec.Command(
		"golines",
		"-l",
		"-m",
		"90",
		".",
	)

	stdout, _ := cmd.Output()

	if len(stdout) > 0 {
		formattedFiles := strings.Split(string(stdout), "\n")
		fmt.Println("Formatted files:")
		for _, file := range formattedFiles {
			fmt.Printf("\t%s\n", file)
		}
		// Formatting happened, run golines
		// again to apply the formatting to the code base
		cmd := exec.Command(
			"golines",
			"-w",
			"-m",
			"90",
			"--base-formatter",
			"gofumpt",
			".",
		)

		cmd.Output()
		os.Exit(1) // Run failed
	}
	os.Exit(0)
}
