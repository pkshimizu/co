package do

import (
	"bytes"
	"os/exec"
	"strings"
)

type Command struct {
	Name        string
	Execs       []string
	WorkingDir  string
	Description string
}

func (cmd Command) Exec(args []string) error {
	for _, ex := range cmd.Execs {
		values := strings.Split(ex, " ")
		cmdName := values[0]
		cmdArgs := append(values[1:], args...)
		c := exec.Command(cmdName, cmdArgs...)
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		c.Stdout = &stdout
		c.Stderr = &stderr
		c.Dir = cmd.WorkingDir
		err := c.Run()
		stdMsg := stdout.String()
		if len(stdMsg) > 0 {
			println(stdMsg)
		}
		if err != nil {
			errMsg := stderr.String()
			if len(errMsg) > 0 {
				println(errMsg)
			}
			return err
		}
	}
	return nil
}
