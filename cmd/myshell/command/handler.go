package command

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
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
			path, err := h.searchExecutablePath(command)
			if (err != nil) {
				h.handleCommandNotFound(command)
				return nil, errors.New("Command not found")
			}
			return Executable{ path }, nil
	}
}

func (h Handler) HandleCommand(commandInput string, args []string) {
	command, err := h.GetCommand(commandInput, args)
	if (err == nil) {
		command.Execute()
	}
}

func (h Handler) searchExecutablePath(executableName string) (string, error) {
	path := os.Getenv("PATH")
	if (path == "") {
		return "", errors.New("no PATH environment variable set")
	}
	executableLocations := filepath.SplitList(path)
	for _, loc := range executableLocations {
		exists, path := h.executableExistsInDir(loc, executableName)
		if (exists) {
			return path, nil
		}
	}
	errMsg := fmt.Sprintf("unable to find executable %s in PATH", executableName)
	return "", errors.New(errMsg)
}

func (h Handler) executableExistsInDir(dirName string, executableName string) (bool, string) {
	path := path.Join(dirName, executableName)
	_, err := os.Stat(path)
	if (err != nil) {
		return false, ""
	}
	return true, path
}