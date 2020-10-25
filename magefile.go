// +build mage

package main

import (
	"fmt"
	"os/exec"
	// mg contains helpful utility functions, like Deps
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	fmt.Println("Tidying...")
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Run()

	fmt.Println("Formatting...")
	cmd = exec.Command("go", "fmt", "-x")
	cmd.Run()

	fmt.Println("Installing...")
	cmd = exec.Command("go", "install")
	return cmd.Run()
}
