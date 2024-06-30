package command

import (
	"fmt"
	"os"
	"strings"
)

type Echo struct {
	args []string
}

func (e Echo) GetType() string {
	return BUILTIN
}

func (e Echo) Execute() (error) {
	toEcho := strings.Join(e.args, " ")
	fmt.Fprintf(os.Stdout, "%s", toEcho)
	return nil
}
