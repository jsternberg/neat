package neat_test

import (
	"testing"

	"github.com/jsternberg/neat"
)

type PlaybookTest struct {
	*testing.T
	Modules  []neat.Module
	Expected neat.Stats
	Err      string
}

func (t *PlaybookTest) Run() {
	playbook := neat.NewPlaybook()
	for _, module := range t.Modules {
		playbook.Add(module)
	}

	stats, err := playbook.Play()
	if t.Err != "" {
		if err == nil {
			t.Fatalf("expected error, got nil")
		} else if t.Err != err.Error() {
			t.Fatalf("expected error '%s', got '%s'", t.Err, err.Error())
		}
	} else if err != nil {
		t.Fatalf("error running playbook: %s", err)
	}

	if stats.Ok != t.Expected.Ok {
		t.Errorf("ok: expected %d, got %d", t.Expected.Ok, stats.Ok)
	}
	if stats.Changed != t.Expected.Changed {
		t.Errorf("changed: expected %d, got %d", t.Expected.Changed, stats.Changed)
	}
	if stats.Deferred != t.Expected.Deferred {
		t.Errorf("deferred: expected %d, got %d", t.Expected.Deferred, stats.Deferred)
	}
	if stats.Skipped != t.Expected.Skipped {
		t.Errorf("skipped: expected %d, got %d", t.Expected.Skipped, stats.Skipped)
	}
}

func TestPlaybook_Play_Empty(t *testing.T) {
	test := &PlaybookTest{
		T:   t,
		Err: "playbook empty",
	}
	test.Run()
}

func TestPlaybook_Play_SingleModule(t *testing.T) {
	r := &neat.ModuleRegistry{}
	r.Register("mock", &MockFactory{})

	test := &PlaybookTest{
		T: t,
		Modules: []neat.Module{
			r.MustCreate("mock"),
		},
		Expected: neat.Stats{
			Ok:       1,
			Changed:  0,
			Deferred: 0,
			Skipped:  0,
		},
	}
	test.Run()
}

func TestPlaybook_Play_MultiModule(t *testing.T) {
	r := &neat.ModuleRegistry{}
	r.Register("mock", &MockFactory{})

	test := &PlaybookTest{
		T: t,
		Modules: []neat.Module{
			r.MustCreate("mock"),
		},
		Expected: neat.Stats{
			Ok:       2,
			Changed:  0,
			Deferred: 0,
			Skipped:  0,
		},
	}
	test.Run()
}

func TestPlaybook_Play_Noop(t *testing.T) {
	r := &neat.ModuleRegistry{}
	r.Register("mock", &MockFactory{})

	test := &PlaybookTest{
		T: t,
		Modules: []neat.Module{
			r.MustCreate("mock"),
		},
		Expected: neat.Stats{
			Ok:       1,
			Changed:  0,
			Deferred: 0,
			Skipped:  0,
		},
	}
	test.Run()
}

func TestPlaybook_Play_SkippedModule(t *testing.T) {
	r := &neat.ModuleRegistry{}
	r.Register("mock", &MockFactory{})

	test := &PlaybookTest{
		T: t,
		Modules: []neat.Module{
			r.MustCreate("mock"),
		},
		Expected: neat.Stats{
			Ok:       0,
			Changed:  0,
			Deferred: 0,
			Skipped:  1,
		},
	}
	test.Run()
}

func TestPlaybook_Play_DeferredModule(t *testing.T) {
	r := &neat.ModuleRegistry{}
	r.Register("mock", &MockFactory{})

	test := &PlaybookTest{
		T: t,
		Modules: []neat.Module{
			r.MustCreate("mock"),
		},
		Expected: neat.Stats{
			Ok:       0,
			Changed:  0,
			Deferred: 1,
			Skipped:  0,
		},
	}
	test.Run()
}
