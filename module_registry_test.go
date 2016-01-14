package neat_test

import (
	"testing"

	"github.com/jsternberg/neat"
)

func TestModuleRegistry_Create(t *testing.T) {
	registry := &neat.ModuleRegistry{}
	registry.Register("mock", &MockFactory{})

	m, err := registry.Create("mock")
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := m.(*Mock); !ok {
		t.Fatalf("expected type *Mock, got %T", m)
	}
}

func TestModuleRegistry_Create_Missing(t *testing.T) {
	registry := &neat.ModuleRegistry{}
	_, err := registry.Create("mock")
	if err == nil {
		t.Fatal("expected an error, got nil")
	} else if err.Error() != "module mock not found" {
		t.Fatalf("expected 'module mock not found' error message, got %s", err)
	}
}

func TestModuleRegistry_Lookup(t *testing.T) {
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

func TestModuleRegistry_Lookup_Missing(t *testing.T) {
	registry := &neat.ModuleRegistry{}
	mf := registry.Lookup("mock")
	if mf != nil {
		t.Fatalf("expected to be unable to find mock factory, got %T", mf)
	}
}
