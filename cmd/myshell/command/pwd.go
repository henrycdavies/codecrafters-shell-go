package command

import (
	"errors"
	"fmt"
	"os"
)

type Pwd struct {}


func (p Pwd) GetType() string {
	return BUILTIN
}

func (p Pwd) Execute() (error) {
	currDir, err := os.Getwd()
	if (err != nil) {
		return errors.New("error getting current path")
	}
	fmt.Fprintf(os.Stdout, "%s\n", currDir)
	return nil
}
