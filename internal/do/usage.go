package do

import "fmt"

func Print(setting Setting) {
	// ヘルプを表示する
	fmt.Print("Usage: do <command> [<args>]\n\n")

	fmt.Print("Commands\n")
	ml := getMaxCommandNameLength(setting)
	for _, cmd := range setting.Commands {
		fmt.Printf("  %-*s  %s\n", ml, cmd.Name, cmd.Description)
	}
}

func getMaxCommandNameLength(setting Setting) int {
	l := 0
	for _, cmd := range setting.Commands {
		if len(cmd.Name) > l {
			l = len(cmd.Name)
		}
	}
	return l
}
