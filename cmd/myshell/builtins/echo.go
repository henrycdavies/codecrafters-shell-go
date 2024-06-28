package builtins

import (
	"fmt"
	"os"
	"strings"
)

func Echo(args []string) {
	toEcho := strings.Join(args, " ")
	fmt.Fprintf(os.Stdout, "%s", toEcho)
}