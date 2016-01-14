package neat_test

import (
	"testing"

	"github.com/jsternberg/neat"
)

type MockFactory struct{}

func (*MockFactory) New(args ...interface{}) (neat.Module, error) {
	return &Mock{}, nil
}

type Mock struct{}

func (*Mock) Execute(p *neat.Playbook) (interface{}, neat.ModuleStatus, error) {
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
