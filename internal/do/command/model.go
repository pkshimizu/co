package command

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Executors struct {
	ExecList []string
}

func (e Executors) String() string {
	return strings.Join(e.ExecList, " | ")
}

func (e Executors) Run(dir string, args []string) (err error) {
	out := ""
	errMsg := ""
	for _, ex := range e.ExecList {
		values := strings.Split(ex, " ")
		cmdName := values[0]
		cmdArgs := append(values[1:])
		out, errMsg, err = runCommand(dir, cmdName, cmdArgs, out)
	}
	if len(out) > 0 {
		_, err = fmt.Fprint(os.Stdin, out)
	}
	if len(errMsg) > 0 {
		_, err = fmt.Fprint(os.Stderr, errMsg)
	}
	return
}

func runCommand(dir string, name string, args []string, in string) (out string, errMsg string, err error) {
	c := exec.Command(name, args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr
	var si io.WriteCloser
	if si, err = c.StdinPipe(); err != nil {
		return
	}
	defer func() {
		err = si.Close()
	}()
	c.Dir = dir
	if err = c.Start(); err != nil {
		return
	}
	if _, err = fmt.Fprint(si, in); err != nil {
		return
	}
	if err = si.Close(); err != nil {
		return
	}
	if err = c.Wait(); err != nil {
		return
	}
	out = stdout.String()
	errMsg = stderr.String()
	return
}

type Command struct {
	Name          string
	ExecutorsList []Executors
	WorkingDir    string
	Description   string
}

func (cmd Command) Exec(args []string) error {
	for _, ex := range cmd.ExecutorsList {
		err := ex.Run(cmd.WorkingDir, args)
		if err != nil {
			return err
		}
	}
	return nil
}
