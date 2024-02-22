package main

import (
	"fmt"
	"os"

	"github.com/c-128/paper.gen/cli"
)

func main() {
	err := cli.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
