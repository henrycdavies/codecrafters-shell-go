package command

import (
	"fmt"
	"os"
)

const HOME_DIRECTORY = "~"

type Cd struct {
	args []string
}


func (c Cd) GetType() string {
	return BUILTIN
}

func (c Cd) Execute() (error) {
	switch len(c.args) {
	case 0:
		c.goToHomeDirectory()
		return nil
	case 1:
		path := c.args[0]
		if (path == HOME_DIRECTORY) {
			c.goToHomeDirectory()
			return nil
		}
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

func (c Cd) goToHomeDirectory() (error) {
	homeDir, err := os.UserHomeDir()
	if (err != nil) {
		return err
	}
	c.changeDirectory(homeDir)
	return nil
}