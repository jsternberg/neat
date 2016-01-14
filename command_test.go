package neat_test

import (
	"testing"

	"github.com/jsternberg/neat"
)

func TestModule_Command_Create(t *testing.T) {
	_, err := neat.CreateModule("command", "touch /tmp")
	if err != nil {
		t.Fatal(err)
	}
}
