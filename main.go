package main

import (
	"fmt"
	"os"
	"os/exec"
)

var Context = "default"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: docker <command> [args]")
		os.Exit(1)
	}
	var args []string
	var env []string
	if os.Args[1] == "buildx" {
		fmt.Printf("Running docker buildx with context default\n")
		args = append([]string{fmt.Sprintf("--context=%s", Context)}, os.Args[1:]...)
		env = append(os.Environ(), fmt.Sprintf("BUILDX_BUILDER=%s", Context))
		fmt.Println(args)
	} else {
		args = os.Args[1:]
		env = os.Environ()
	}
	// Run passed in arguments as subprocess
	cmd := exec.Command("docker", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = env
	err := cmd.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			os.Exit(exitError.ExitCode())
		} else {
			os.Exit(1)
		}
	}
	os.Exit(0)
}
