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

	filepath := fmt.Sprintf("%s/test.txt", tmpdir)
	m, err := neat.CreateModule("file", map[string]interface{}{
		"path":  filepath,
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

	stat, err := os.Stat(filepath)
	if err != nil {
		t.Fatal(err)
	} else if stat.Mode() != 0644 {
		t.Errorf("expected file to be mode 0644, got %v", stat.Mode())
	}
}

func TestModule_File_Remove(t *testing.T) {
	tmpdir, err := ioutil.TempDir(os.TempDir(), "neat-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	filepath := fmt.Sprintf("%s/test.txt", tmpdir)
	f, err := os.Create(filepath)
	if err != nil {
		t.Fatal(err)
	}
	f.Close()

	m, err := neat.CreateModule("file", map[string]interface{}{
		"path":  filepath,
		"state": "absent",
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

	_, err = os.Stat(filepath)
	if err == nil {
		t.Fatal("file present, but was supposed to be removed")
	} else if !os.IsNotExist(err) {
		t.Fatal(err)
	}
}
