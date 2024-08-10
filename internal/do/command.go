package do

import (
	"os/exec"
)

type Command struct {
	Name        string
	Execs       []string
	WorkingDir  string
	Description string
}

func (cmd Command) Exec(args []string) error {
	for _, ex := range cmd.Execs {
		c := exec.Command(ex, args...)
		c.Dir = cmd.WorkingDir
		out, err := c.Output()
		if err != nil {
			return err
		}
		println(string(out))
	}
	return nil
}
