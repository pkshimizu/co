package do

import (
	"os"
	"path/filepath"

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

type YamlCommand struct {
	Exec        []string `yaml:"exec"`
	WorkingDir  string   `yaml:"working_dir"`
	Description string   `yaml:"description"`
}

type YamlSetting struct {
	Commands map[string]YamlCommand `yaml:"commands"`
}

func Load() (Setting, error) {
	// 以下の順序で.ca.yamlファイルを読み込む
	// 1. カレントディレクトリ
	setting, err := loadYaml(".")
	if err != nil {
		return setting, err
	}
	// 2. 上位ディレクトリ
	// 3. CA_HOMEディレクトリ
	// 4. ホームディレクトリ
	return setting, nil
}

func loadYaml(dir string) (Setting, error) {

	// 設定を読み込む
	setting := YamlSetting{}
	b, err := os.ReadFile(dir + "/.do.yaml")
	if err != nil {
		return Setting{}, err
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
			Execs:       cmd.Exec,
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
