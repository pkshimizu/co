package setting

import (
	"slices"

	"co/internal/co/command"
)

type Setting struct {
	Commands []command.Command
}

func (s Setting) FindCommand(name string) *command.Command {
	for _, cmd := range s.Commands {
		if cmd.Name == name {
			return &cmd
		}
	}
	return nil

}

func (s *Setting) AddCommands(cmds []command.Command) {
	var names []string
	for _, cmd := range s.Commands {
		names = append(names, cmd.Name)
	}
	for _, cmd := range cmds {
		if slices.Contains(names, cmd.Name) {
			continue
		}
		s.Commands = append(s.Commands, cmd)
		names = append(names, cmd.Name)
	}
}

type YamlCommand struct {
	Exec        []string `yaml:"exec"`
	WorkingDir  string   `yaml:"working_dir"`
	Description string   `yaml:"description"`
}

type YamlSetting struct {
	Commands map[string]YamlCommand `yaml:"commands"`
}
