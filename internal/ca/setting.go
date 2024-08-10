package ca

import (
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Setting struct {
	Commands []Command
}

type YamlCommand struct {
	Exec        string `yaml:"exec"`
	WorkingDir  string `yaml:"working_dir"`
	Description string `yaml:"description"`
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
	b, err := os.ReadFile(dir + "/.ca.yaml")
	if err != nil {
		return Setting{}, err
	}
	err = yaml.Unmarshal(b, &setting)
	if err != nil {
		return Setting{}, err
	}
	var cmds []Command
	for name, cmd := range setting.Commands {
		cmds = append(cmds, Command{
			Name:        name,
			WorkingDir:  cmd.WorkingDir,
			Description: cmd.Description,
		})
	}
	return Setting{
		Commands: cmds,
	}, nil
}
