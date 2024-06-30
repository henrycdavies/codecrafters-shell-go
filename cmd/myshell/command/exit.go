package command

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Exit struct {
	args[] string
}

func (e Exit) GetType() string {
	return BUILTIN
}

func (e Exit) Execute() (error) {
	if (len(e.args) != 1) {
		errMsg := "an exit code arg must be provided to exit command"
		fmt.Fprintf(os.Stderr, "%s\n", errMsg)
		return errors.New(errMsg)
	}
	code, err := strconv.Atoi(e.args[0])
	if (err != nil) {
		errMsg := "an exit code arg must be provided to exit command"
		fmt.Fprintf(os.Stderr, "%s\n", errMsg)
		return errors.New(errMsg)
	}
	os.Exit(code)
	return nil
}

