package command

type CommandType string

const (
    EchoCommand CommandType = "echo"
    ExitCommand CommandType = "exit"
)

const BUILTIN = "builtin"

type Command interface {
	GetType() CommandType
    Execute() (error)
}