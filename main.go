package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$")
		// Read the keyboard input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		// Handle the execution of the command/input
		if err := execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

// Executing Commands
func execInput(input string) error {
	//Remove the newline character
	input = strings.TrimSuffix(input, "\n")
	// Split the input to get the command and the arguments
	args := strings.Split(input, " ")
	// check for built-in commands
	switch args[0] {
	case "cd":
		// Check if a path is provided (change directory(cd) to home directory if no path is provided)
		if len(args) < 2 {
			return errors.New("path required")
		}
		// Change the directory and return the error
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}
	// Prepare the command to execute (pass the command and the arguments separately)
	cmd := exec.Command(args[0], args[1:]...)
	// Set the output to the terminal
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	//Execute the command and return the error
	return cmd.Run()
}
