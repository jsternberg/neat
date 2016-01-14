package neat_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/jsternberg/neat"
)

func TestModule_Execute_Create(t *testing.T) {
	tmpdir, err := ioutil.TempDir(os.TempDir(), "neat-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	filepath := fmt.Sprintf("%s/test.txt", tmpdir)
	m, err := neat.CreateModule("execute", map[string]interface{}{
		"command": fmt.Sprintf("touch %s", filepath),
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
		t.Fatalf("execute status changed, got %s", status)
	}

	_, err = os.Stat(filepath)
	if err != nil {
		t.Fatal(err)
	}
}
