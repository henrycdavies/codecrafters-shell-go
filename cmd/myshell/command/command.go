package command

const BUILTIN = "a shell builtin"

type Command interface {
	GetType() string
    Execute() (error)
}