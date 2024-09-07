package help

import (
	"fmt"
	"slices"
	"sort"

	"do/internal/do/setting"
)

func Print(version string, s setting.Setting) {
	fmt.Printf("do version: %s\n", version)
	fmt.Print("Usage: do <command> [<args>]\n\n")

	fmt.Print("Commands:\n")
	ml := getMaxCommandNameLength(s)
	cmds := slices.Clone(s.Commands)
	sort.SliceStable(cmds, func(i, j int) bool {
		return cmds[i].Name < cmds[j].Name
	})
	for _, cmd := range cmds {
		fmt.Printf("  %-*s  %s\n", ml, cmd.Name, cmd.Description)
		fmt.Printf("  %-*s  working dir: %s\n", ml, "", cmd.WorkingDir)
		if len(cmd.Pipelines) > 0 {
			fmt.Printf("  %-*s  exec: %s\n", ml, "", cmd.Pipelines[0].String())
		}
		for _, pl := range cmd.Pipelines[1:] {
			fmt.Printf("  %-*s        %s\n", ml, "", pl.String())
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
