package command

import (
	"fmt"
	"os"
)

type Cd struct {
	args []string
}


func (c Cd) GetType() string {
	return BUILTIN
}

func (c Cd) Execute() (error) {
	switch len(c.args) {
	case 0:
		return nil
	case 1:
		path := c.args[0]
		c.changeDirectory(path)
	}
	return nil
}

func (c Cd) changeDirectory(dir string) (error) {
	err := os.Chdir(dir)
	if (err != nil) {
		fmt.Fprintf(os.Stderr, "cd: %s: No such file or directory\n", dir)
		return err
	}
	return nil
}