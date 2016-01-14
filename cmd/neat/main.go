package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jsternberg/neat"
)

var (
	flModule = flag.String("m", "command", "selects the module to run")
	flArgs   = flag.String("a", "", "the arguments to pass to the module")
)

func realMain() int {
	m, err := neat.CreateModule(*flModule, *flArgs)
	if err != nil {
		fmt.Printf("unable to create module: %s\n", err.Error())
		return 1
	}
	m.Execute()
	return 0
}

func main() {
	os.Exit(realMain())
}
