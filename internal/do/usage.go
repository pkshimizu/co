package do

import "fmt"

func Print(setting Setting) {
	fmt.Print("Usage: do <command> [<args>]\n\n")

	fmt.Print("Commands:\n")
	ml := getMaxCommandNameLength(setting)
	for _, cmd := range setting.Commands {
		fmt.Printf("  %-*s  %s\n", ml, cmd.Name, cmd.Description)
		fmt.Printf("  %-*s  working dir: %s\n", ml, "", cmd.WorkingDir)
		if len(cmd.ExecList) > 0 {
			fmt.Printf("  %-*s  exec: %s\n", ml, "", cmd.ExecList[0])
		}
		for _, ex := range cmd.ExecList[1:] {
			fmt.Printf("  %-*s        %s\n", ml, "", ex)
		}
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
