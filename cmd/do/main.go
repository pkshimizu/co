package main

import (
	"fmt"
	"os"

	"github.com/pkshimizu/do/internal/do/help"
	"github.com/pkshimizu/do/internal/do/setting"
)

func main() {
	s, err := setting.Load()
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(os.Args) < 2 {
		help.Print(s)
		return
	}
	name := os.Args[1]
	cmd := s.FindCommand(name)
	if cmd == nil {
		help.Print(s)
		return
	}
	cmd.Exec(os.Args[2:])
}
