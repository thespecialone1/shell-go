package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

var history []string
var historyIndex int

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		// Prompt the current directory, hostname and username
		dir, _ := os.Getwd()
		// hostname, _ := os.Hostname()
		currentUser, _ := user.Current()
		fmt.Printf("%s:%s $ ", currentUser.Username, dir)
		// fmt.Print("$")
		// Read the keyboard input
		input, err := reader.ReadString('\n')
		if err != nil {	
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		// Clean input (trim spaces and newline)
		input = strings.TrimSpace(input)

		if len(input) > 0 {
			history = append(history, input)
			historyIndex = len(history) // Reset the history index
		}
		// Handle the execution of the command/input
		if err := execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

// Executing Commands
func execInput(input string) error {
	// handle spaces in the input
	args := strings.Fields(input)
	if len(args) == 0 {
		return nil // No command provided, return nil
	}
	// Split the input to get the command and the arguments
	// args := strings.Split(input, " ")
	// check for built-in commands
	switch args[0] {
	case "cd":
		// Check if a path is provided (change directory(cd) to home directory if no path is provided)
		if len(args) < 2 {
			return errors.New("path required")
		}
		// Handele path with spaces by joining the arguments
		dir := strings.Join(args[1:], " ")
		
		// Change the directory and return the error
		return os.Chdir(dir)
	case "exit":
		os.Exit(0)
		// Add the command to the history
	case "history":
		// Print the history
		for i, cmd := range history {
			fmt.Printf("%d %s\n", i+1, cmd)
		}
		return nil
	}

	// Prepare the command to execute (pass the command and the arguments separately)
	cmd := exec.Command(args[0], args[1:]...)
	// Set the output to the terminal
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	//Execute the command and return the error
	return cmd.Run()
}

// History {Still in progress had to use third party library for this}
func chkHistory(input string) string {
	// Navigate up
	if input == "up" {
		if historyIndex > 0 {
			historyIndex--
		}
	}
	// Navigate down
	if input == "down" {
		if historyIndex < len(history)-1 {
			historyIndex++
		}
	}
	if len(history) > 0 && historyIndex >= len(history) {
		historyIndex = len(history) - 1
		return history[historyIndex]
	}
	return ""
}
