package command

import (
	"errors"
	"fmt"
	"os"
)

type Type struct {
	args []string
}

func (t Type) GetType() string {
	return BUILTIN
}

func (t Type) Execute() (error) {
	if (len(t.args) != 1) {
		return errors.New("type command only takes one argument")
	}
	commandName := t.args[0]
	handler := Handler{}
	var arr []string
	command, err := handler.GetCommand(commandName, arr)

	if (err == nil ){
		commandType := command.GetType()
		fmt.Fprintf(os.Stdout, "%s is %s\n", commandName, commandType)
		return nil
	}
	return err
}