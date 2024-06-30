package command

type Executable struct {
	path string
}

func (e Executable) Execute() (error) {
	return nil
}

func (e Executable) GetType() string {
	return e.path
}