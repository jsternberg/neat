package neat_test

import (
	"testing"

	"github.com/jsternberg/neat"
)

func TestModule_Template_Create(t *testing.T) {
	_, err := neat.CreateModule("template", map[string]interface{}{
		"src":   "foobar.conf.tmpl",
		"dest":  "/etc/foobar.conf",
		"mode":  0644,
		"state": "present",
	})
	if err != nil {
		t.Fatal(err)
	}
}
