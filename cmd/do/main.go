package main

import (
	"fmt"
	"os"

	"github.com/pkshimizu/do/internal/do"
)

func main() {
	setting, err := do.Load()
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(os.Args) < 2 {
		do.Print(setting)
		return
	}
	name := os.Args[1]
	cmd := setting.FindCommand(name)
	if cmd == nil {
		do.Print(setting)
		return
	}
	(*cmd).Exec(os.Args[2:])
}
