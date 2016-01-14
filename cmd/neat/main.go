package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/jsternberg/neat"
)

var (
	flModule = flag.String("m", "execute", "selects the module to run")
	flArgs   = flag.String("a", "", "the arguments to pass to the module")
)

func realMain() int {
	m, err := neat.CreateModule(*flModule)
	if err != nil {
		fmt.Printf("unable to create module: %s\n", err.Error())
		return 1
	}

	playbook := neat.NewPlaybook()
	result, status, err := m.Execute(playbook)
	if err != nil {
		fmt.Printf("module failed: %s\n", err)
	}

	if result != nil {
		fmt.Printf("%s: ", status.String())
		encoder := json.NewEncoder(os.Stdout)
		if err := encoder.Encode(result); err != nil {
			fmt.Println(err)
			return 1
		}
	} else {
		fmt.Println(status.String())
	}
	return 0
}

func main() {
	os.Exit(realMain())
}
