package help

import (
	"fmt"

	"do/internal/do/setting"
)

func Print(s setting.Setting) {
	fmt.Print("Usage: do <command> [<args>]\n\n")

	fmt.Print("Commands:\n")
	ml := getMaxCommandNameLength(s)
	for _, cmd := range s.Commands {
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

func getMaxCommandNameLength(s setting.Setting) int {
	l := 0
	for _, cmd := range s.Commands {
		if len(cmd.Name) > l {
			l = len(cmd.Name)
		}
	}
	return l
}
