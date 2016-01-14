package neat

import "fmt"

// ModuleRegistry keeps a lookup table to access registered ModuleFactory's by
// name.
type ModuleRegistry map[string]ModuleFactory

// Create creates a new module by name. This is a convenience method for calling
// Lookup and then New. If the lookup doesn't succeed, this returns an error.
func (r *ModuleRegistry) Create(name string, args ...interface{}) (Module, error) {
	factory := r.Lookup(name)
	if factory == nil {
		return nil, fmt.Errorf("module %s not found", name)
	}
	return factory.New(args...)
}

// MustCreate creates a new module and panics if an error occurs.
func (r *ModuleRegistry) MustCreate(name string, args ...interface{}) Module {
	m, err := r.Create(name, args...)
	if err != nil {
		panic(err)
	}
	return m
}

// Register associates name with the passed in factory for future retrieval.
func (r *ModuleRegistry) Register(name string, factory ModuleFactory) {
	(*r)[name] = factory
}

// Lookup finds the ModuleFactory associated with the name. If no factory
// with this name is found, nil is returned.
func (r *ModuleRegistry) Lookup(name string) ModuleFactory {
	factory, _ := (*r)[name]
	return factory
}

func unimplemented(name string) ModuleFunc {
	return ModuleFunc(func(args ...interface{}) (Module, error) {
		return nil, fmt.Errorf("%s unimplemented", name)
	})
}

// DefaultRegistry is a global default registry for convenience.
var DefaultRegistry = &ModuleRegistry{
	"command":  unimplemented("command"),
	"file":     unimplemented("file"),
	"template": unimplemented("template"),
	"service":  unimplemented("service"),
	"shell":    unimplemented("shell"),
}

// CreateModule calls Create on the DefaultRegistry.
func CreateModule(name string, args ...interface{}) (Module, error) {
	return DefaultRegistry.Create(name, args...)
}

// MustCreateModule calls MustCreate on the DefaultRegistry.
func MustCreateModule(name string, args ...interface{}) Module {
	return DefaultRegistry.MustCreate(name, args...)
}

// RegisterModule calls Register on the DefaultRegistry.
func RegisterModule(name string, factory ModuleFactory) {
	DefaultRegistry.Register(name, factory)
}

// LookupModule calls Lookup on the DefaultRegistry.
func LookupModule(name string) ModuleFactory {
	return DefaultRegistry.Lookup(name)
}
