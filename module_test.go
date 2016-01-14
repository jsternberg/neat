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

func (*Mock) Execute() error { return nil }

func TestModuleFactory_Lookup(t *testing.T) {
	registry := &neat.ModuleRegistry{}
	registry.Register("mock", &MockFactory{})

	mf := registry.Lookup("mock")
	if mf == nil {
		t.Fatalf("unable to lookup mock module")
	}

	m, err := mf.New()
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := m.(*Mock); !ok {
		t.Fatalf("expected type *Mock, got %T", m)
	}
}

func TestModuleFactory_Lookup_Missing(t *testing.T) {
	registry := &neat.ModuleRegistry{}
	mf := registry.Lookup("mock")
	if mf != nil {
		t.Fatalf("expected to be unable to find mock factory, got %T", mf)
	}
}
