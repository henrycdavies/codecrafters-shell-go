package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handleCommandNotFound(command string) {
	fmt.Fprintf(os.Stdout, "%s: command not found", command)
}

func exit(args []string) (error) {
	if (len(args) != 1) {
		errMsg := "an exit code arg must be provided to exit command"
		fmt.Fprintf(os.Stdout, "%s", errMsg)
		return errors.New(errMsg)
	}
	code, err := strconv.Atoi(args[0])
	if (err != nil) {
		errMsg := "an exit code arg must be provided to exit command"
		fmt.Fprintf(os.Stdout, "%s", errMsg)
		return errors.New(errMsg)
	}
	os.Exit(code)
	return nil
}

func handleCommand(command string, args []string) {
	switch command {
	case "exit":
		exit(args)
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
	fmt.Fprint(os.Stdout, "\n$ ")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if (err != nil) {
		fmt.Fprintf(os.Stdout, "%s", err)
	}
	command, args, err := parseInput(input)
	if (err == nil) {
		handleCommand(command, args)
	}
	handleInput()
}

func main() {
	handleInput()
}
