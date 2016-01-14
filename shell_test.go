package neat_test

import (
	"testing"

	"github.com/jsternberg/neat"
)

func TestModule_Shell_Create(t *testing.T) {
	_, err := neat.CreateModule("shell", "(echo hello && echo world) | tail -n 1")
	if err != nil {
		t.Fatal(err)
	}
}
