package neat_test

import (
	"testing"

	"github.com/jsternberg/neat"
)

type MockFactory struct{}

func (*MockFactory) New(args ...interface{}) (neat.Module, error) {
	var v interface{}
	if len(args) >= 1 {
		v = args[0]
	}
	return &Mock{v}, nil
}

type Mock struct {
	v interface{}
}

func (m *Mock) Execute(p *neat.Playbook) (interface{}, neat.ModuleStatus, error) {
	if m.v != nil {
		switch v := m.v.(type) {
		case error:
			me := neat.ModuleE{Module: m}
			return me.Fail(v.Error())
		case neat.ModuleStatus:
			return nil, v, nil
		default:
			panic(v)
		}
	}
	return nil, neat.ModuleOk, nil
}

func TestModuleFunc(t *testing.T) {
	registry := &neat.ModuleRegistry{}
	registry.Register("mock", neat.ModuleFunc(
		func(args ...interface{}) (neat.Module, error) {
			return &Mock{}, nil
		},
	))

	mf := registry.Lookup("mock")
	_, err := mf.New()
	if err != nil {
		t.Fatalf("unable to lookup mock module")
	}
}

func TestModuleE_Ok(t *testing.T) {
	registry := &neat.ModuleRegistry{}
	registry.Register("mock", &MockFactory{})

	m := neat.ModuleE{Module: registry.MustCreate("mock")}
	result, status, err := m.Ok("module succeeded")
	actual, ok := result.(string)
	if !ok || actual != "module succeeded" {
		t.Errorf("expected result to be 'module succeded', got '%v'", result)
	}
	if status != neat.ModuleOk {
		t.Errorf("expected status to be %s, got %s", neat.ModuleOk, status)
	}
	if err != nil {
		t.Error(err)
	}
}

func TestModuleE_Changed(t *testing.T) {
	registry := &neat.ModuleRegistry{}
	registry.Register("mock", &MockFactory{})

	m := neat.ModuleE{Module: registry.MustCreate("mock")}
	m.Changed = true
	result, status, err := m.Ok("module succeeded")
	actual, ok := result.(string)
	if !ok || actual != "module succeeded" {
		t.Errorf("expected result to be 'module succeded', got '%v'", result)
	}
	if status != neat.ModuleChanged {
		t.Errorf("expected status to be %s, got %s", neat.ModuleChanged, status)
	}
	if err != nil {
		t.Error(err)
	}
}

func TestModuleE_Fail(t *testing.T) {
	registry := &neat.ModuleRegistry{}
	registry.Register("mock", &MockFactory{})

	m := neat.ModuleE{Module: registry.MustCreate("mock")}
	result, status, err := m.Fail("module failed to run")
	if result != nil {
		t.Errorf("expected result to be nil, got %v", result)
	}
	if status != neat.ModuleFailed {
		t.Errorf("expected status to be %s, got %s", neat.ModuleFailed, status)
	}
	if err == nil || err.Error() != "module failed to run" {
		t.Errorf("expected error to be 'module failed to run', got '%s'", err)
	}
}

func TestModuleStatus_String(t *testing.T) {
	tests := []struct {
		Expected string
		Status   neat.ModuleStatus
	}{
		{Expected: "ok", Status: neat.ModuleOk},
		{Expected: "changed", Status: neat.ModuleChanged},
		{Expected: "deferred", Status: neat.ModuleDeferred},
		{Expected: "skipped", Status: neat.ModuleSkipped},
		{Expected: "failed", Status: neat.ModuleFailed},
	}

	for _, test := range tests {
		actual := test.Status.String()
		if actual != test.Expected {
			t.Errorf("expected '%s', got '%s'", test.Expected, actual)
		}
	}
}
