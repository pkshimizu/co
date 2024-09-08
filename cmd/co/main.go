package main

import (
	"fmt"
	"os"

	"co/internal/co/help"
	"co/internal/co/setting"
)

const Version = "0.1.0"

func main() {
	s, err := setting.Load()
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(os.Args) < 2 {
		help.Print(Version, s)
		return
	}
	name := os.Args[1]
	cmd := s.FindCommand(name)
	if cmd == nil {
		help.Print(Version, s)
		return
	}
	cmd.Exec(os.Args[2:])
}
