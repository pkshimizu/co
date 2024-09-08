package setting

import (
	"os"
	"path/filepath"
	"strings"

	"co/internal/co/command"
	yaml "gopkg.in/yaml.v2"
)

func Load() (Setting, error) {
	// 以下の順序で.co.yamlファイルを読み込む
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

	// 3. CO_HOMEディレクトリ
	coHomeDir := os.Getenv("CO_HOME")
	if coHomeDir != "" {
		s, err := loadYaml(coHomeDir)
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
	b, err := os.ReadFile(filepath.Join(dir, ".co.yaml"))
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
		var pipelines []command.ExecutorPipeline
		for _, exec := range cmd.Exec {
			executors := []command.Executor{}
			for _, val := range strings.Split(exec, "|") {
			    executors = append(executors, command.Executor{
			        Line: strings.TrimSpace(val),
			    })
			}
			pipelines = append(pipelines, command.ExecutorPipeline{
				Executors: executors,
			})
		}
		cmds = append(cmds, command.Command{
			Name:          name,
			Pipelines:     pipelines,
			WorkingDir:    wd,
			Description:   cmd.Description,
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
