package neat_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/jsternberg/neat"
)

func TestModule_File_Create(t *testing.T) {
	tmpdir, err := ioutil.TempDir(os.TempDir(), "neat-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	m, err := neat.CreateModule("file", map[string]interface{}{
		"path":  fmt.Sprintf("%s/test.txt", tmpdir),
		"mode":  0644,
		"state": "present",
	})
	if err != nil {
		t.Fatal(err)
	}

	playbook := neat.NewPlaybook()
	_, status, err := m.Execute(playbook)
	if err != nil {
		t.Fatal(err)
	}
	if status != neat.ModuleChanged {
		t.Fatalf("expected status changed, got %s", status)
	}
}
