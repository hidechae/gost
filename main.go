package main

import (
	"github.com/hidechae/gost/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
