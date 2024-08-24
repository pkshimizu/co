package setting

import (
	"os"
	"path/filepath"

	"do/internal/do/command"
	yaml "gopkg.in/yaml.v2"
)

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

	// 3. DO_HOMEディレクトリ
	doHomeDir := os.Getenv("DO_HOME")
	if doHomeDir != "" {
		s, err := loadYaml(doHomeDir)
		setting.AddCommands(s.Commands)
		if err != nil {
			return setting, err
		}
	}
	// 4. ユーザーホームディレクトリ
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		s, err := loadYaml(userHomeDir)
		setting.AddCommands(s.Commands)
		if err != nil {
			return setting, err
		}
	}
	return setting, nil
}

func loadYaml(dir string) (Setting, error) {
	// 設定を読み込む
	s := YamlSetting{}
	b, err := os.ReadFile(filepath.Join(dir, ".do.yaml"))
	if err != nil {
		return Setting{
			Commands: []command.Command{},
		}, nil
	}
	err = yaml.Unmarshal(b, &s)
	if err != nil {
		return Setting{}, err
	}
	var cmds []command.Command
	for name, cmd := range s.Commands {
		wd, err := getWorkingDir(dir, cmd.WorkingDir)
		if err != nil {
			return Setting{}, err
		}
		cmds = append(cmds, command.Command{
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
	if wd == "" {
		return os.Getwd()
	}
	if filepath.IsAbs(wd) {
		return filepath.Abs(wd)
	} else {
		return filepath.Abs(filepath.Join(base, wd))
	}
}
