package command

import (
	"errors"
	"fmt"
	"os"
)

type Type struct {
	args []string
}

func (t Type) GetType() CommandType {
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
	if (err != nil) {
		return err
	}
	commandType := command.GetType()
	fmt.Fprintf(os.Stdout, "%s is a shell %s", commandName, commandType)
	return nil
}