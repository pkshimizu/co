package do

import (
	"os"
	"path/filepath"
	"slices"

	yaml "gopkg.in/yaml.v2"
)

type Setting struct {
	Commands []Command
}

func (s Setting) FindCommand(name string) *Command {
	for _, cmd := range s.Commands {
		if cmd.Name == name {
			return &cmd
		}
	}
	return nil

}

func (s *Setting) AddCommands(cmds []Command) {
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

func Load() (Setting, error) {
	// 以下の順序で.do.yamlファイルを読み込む
	// 1. カレントディレクトリ
	dir, err := os.Getwd()
	if err != nil {
		return Setting{}, err
	}
	setting, err := loadYaml(dir)
	if err != nil {
		return setting, err
	}

	// 2. 上位ディレクトリ
	for {
		parentDir := filepath.Join(dir, "..")
		if parentDir == dir {
			break
		}
		dir = parentDir
		s, err := loadYaml(dir)
		setting.AddCommands(s.Commands)
		if err != nil {
			return setting, err
		}
	}

	// 3. CA_HOMEディレクトリ
	// 4. ホームディレクトリ
	return setting, nil
}

func loadYaml(dir string) (Setting, error) {
	// 設定を読み込む
	setting := YamlSetting{}
	b, err := os.ReadFile(filepath.Join(dir, ".do.yaml"))
	if err != nil {
		return Setting{
			Commands: []Command{},
		}, nil
	}
	err = yaml.Unmarshal(b, &setting)
	if err != nil {
		return Setting{}, err
	}
	var cmds []Command
	for name, cmd := range setting.Commands {
		wd, err := getWorkingDir(dir, cmd.WorkingDir)
		if err != nil {
			return Setting{}, err
		}
		cmds = append(cmds, Command{
			Name:        name,
			ExecList:    cmd.Exec,
			WorkingDir:  wd,
			Description: cmd.Description,
		})
	}
	return Setting{
		Commands: cmds,
	}, nil
}

func getWorkingDir(base string, wd string) (string, error) {
	if filepath.IsAbs(wd) {
		return filepath.Abs(wd)
	} else {
		return filepath.Abs(filepath.Join(base, wd))
	}
}
