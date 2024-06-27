package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func handleCommandNotFound(command string) {
	fmt.Fprintf(os.Stdout, "%s: command not found", command)
}

func main() {
	fmt.Fprint(os.Stdout, "$ ")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if (err != nil) {
		fmt.Fprintf(os.Stdout, "%s", err)
	}
	command := strings.TrimSpace(input)
	handleCommandNotFound(command)

}
