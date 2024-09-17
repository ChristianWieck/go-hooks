package main

import (
	"os"
	"os/exec"
)

func main() {
	// First we need to make sure golines and gofumpt are installed
	exec.Command("go", "install", "github.com/segmentio/golines@latest")
	exec.Command("go", "install", "mvdan.cc/gofumpt@latest")

	// Then run golines
	exec.Command(
		"golines",
		"-w",
		"-m",
		"90",
		"--base-formatter",
		"gofumpt",
		".",
	)
	os.Exit(0)
}
