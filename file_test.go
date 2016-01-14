package neat_test

import (
	"testing"

	"github.com/jsternberg/neat"
)

func TestModule_File_Create(t *testing.T) {
	_, err := neat.CreateModule("file", map[string]interface{}{
		"path":  "/path/to/foobar",
		"mode":  0644,
		"state": "present",
	})
	if err != nil {
		t.Fatal(err)
	}
}
