package command

import (
	"errors"
	"fmt"
	"os"
)
type Handler struct {}

func (h Handler) handleCommandNotFound(command string) {
	fmt.Fprintf(os.Stdout, "%s: not found", command)
}

func (h Handler) GetCommand(command string, args []string) (Command, error) {
	switch command {
		case "exit":
			return Exit{ args }, nil
		case "echo":
			return Echo{ args }, nil
		case "type":
			return Type{ args }, nil
		default:
			h.handleCommandNotFound(command)
			return nil, errors.New("Command not found")
	}
}

func (h Handler) HandleCommand(commandInput string, args []string) {
	command, err := h.GetCommand(commandInput, args)
	if (err == nil) {
		command.Execute()
	}
}