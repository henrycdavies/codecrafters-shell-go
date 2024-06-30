package command

import (
	"os"
	"os/exec"
)

type Executable struct {
	path string
	args []string
}

func (e Executable) Execute() (error) {
	cmd := exec.Command(e.path, e.args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}

func (e Executable) GetType() string {
	return e.path
}