package main

import (
	"fmt"

	"github.com/pkshimizu/ca/internal/ca"
)

func main() {
	fmt.Println("start ca")
	setting, err := ca.Load()
	if err != nil {
		fmt.Println(err)
		return
	}
	ca.Print(setting)
}
