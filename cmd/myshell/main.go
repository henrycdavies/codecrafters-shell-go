package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/builtins"
)

func handleCommandNotFound(command string) {
	fmt.Fprintf(os.Stdout, "%s: command not found", command)
}



func handleCommand(command string, args []string) {
	switch command {
	case "exit":
		builtins.Exit(args)
	case "echo":
		builtins.Echo(args)
	default:
		handleCommandNotFound(command)
	}
}

func parseInput(input string) (string, []string, error) {
	trimmed := strings.TrimSpace(input)
	parts := strings.Split(trimmed, " ")
	if (len(parts) == 0) {
		return "", nil, errors.New("no input received")
	}
	command := parts[0]
	var args []string
	if (len(parts) == 1) {
		args = nil
	} else {
		args = parts[1:]
	}
	return command, args, nil
}

func handleInput() {
	fmt.Fprint(os.Stdout, "$ ")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if (err != nil) {
		fmt.Fprintf(os.Stdout, "%s", err)
	}
	command, args, err := parseInput(input)
	if (err == nil) {
		handleCommand(command, args)
	}
	fmt.Fprint(os.Stdout, "\n")
	handleInput()
}

func main() {
	handleInput()
}
