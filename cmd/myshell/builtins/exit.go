package builtins

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Exit(args []string) (error) {
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