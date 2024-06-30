package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/command"
)

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
	commandInput, args, err := parseInput(input)
	if (err == nil) {
		handler := command.Handler {}
		handler.HandleCommand(commandInput, args)
	}
	handleInput()
}

func main() {
	handleInput()
}
