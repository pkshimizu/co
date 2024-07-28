package main

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	fmt.Println("start ca")
	loadCommands()
}

func loadCommands() Setting {
	// 以下の順序で.ca.yamlファイルを読み込む
	// 1. カレントディレクトリ
	setting := loadYaml(".")
	return setting
	// 2. 上位ディレクトリ
	// 3. CA_HOMEディレクトリ
	// 4. ホームディレクトリ
}

type SettingCommand struct {
	Exec        string `yaml:"exec"`
	WorkingDir  string `yaml:"working_dir"`
	Description string `yaml:"description"`
}

type Setting struct {
	Commands map[string]SettingCommand `yaml:"commands"`
}

func loadYaml(dir string) Setting {
	// yamlファイルを読み込む
	setting := Setting{}
	b, err := os.ReadFile(dir + "/.ca.yaml")
	if err != nil {
		return setting
	}
	err = yaml.Unmarshal(b, &setting)
	if err != nil {
		return setting
	}
	return setting
}
