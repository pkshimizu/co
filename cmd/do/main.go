package main

import (
	"fmt"

	"github.com/pkshimizu/do/internal/do"
)

func main() {
	setting, err := do.Load()
	if err != nil {
		fmt.Println(err)
		return
	}
	do.Print(setting)
}
