package do

import (
	"bytes"
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
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		c.Stdout = &stdout
		c.Stderr = &stderr
		c.Dir = cmd.WorkingDir
		err := c.Run()
		if err != nil {
			println(stdout.String())
			println(stderr.String())
		} else {
			println(stdout.String())
		}
	}
	return nil
}
