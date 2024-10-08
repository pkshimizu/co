package command

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Executor struct {
	Line string
}

func (e Executor) Run(dir string, args []string, in string, outputStd bool) (out string, err error) {
	values := strings.Split(e.Line, " ")
	name := values[0]
	cmdArgs := []string{}
	for _, arg := range values[1:] {
		if arg == "<args>" {
			cmdArgs = append(cmdArgs, args...)
		} else {
			cmdArgs = append(cmdArgs, arg)
		}
	}
	c := exec.Command(name, cmdArgs...)
	var stdout bytes.Buffer
	c.Stderr = os.Stderr
	if outputStd {
		c.Stdout = os.Stdout
	} else {
		c.Stdout = &stdout
	}

	var si io.WriteCloser
	if si, err = c.StdinPipe(); err != nil {
		return
	}
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

	if !outputStd {
		out = stdout.String()
	}

	return
}

type ExecutorPipeline struct {
	Executors []Executor
}

func (e ExecutorPipeline) String() string {
	lines := []string{}
	for _, ex := range e.Executors {
		lines = append(lines, ex.Line)
	}
	return strings.Join(lines, " | ")
}

func (e ExecutorPipeline) Run(dir string, args []string) (err error) {
	out := ""
	for i, ex := range e.Executors {
		if out, err = ex.Run(dir, args, out, len(e.Executors) == 1 || i == (len(e.Executors)-1)); err != nil {
			break
		}
	}
	return
}

type Command struct {
	Name        string
	Pipelines   []ExecutorPipeline
	WorkingDir  string
	Description string
}

func (cmd Command) Exec(args []string) error {
	for _, pl := range cmd.Pipelines {
		err := pl.Run(cmd.WorkingDir, args)
		if err != nil {
			return err
		}
	}
	return nil
}
