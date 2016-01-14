package neat_test

import (
	"testing"

	"github.com/jsternberg/neat"
)

func TestModule_Service_Create(t *testing.T) {
	_, err := neat.CreateModule("service", map[string]interface{}{
		"path":  "neatd",
		"state": "restarted",
	})
	if err != nil {
		t.Fatal(err)
	}
}
