package neat

// ModuleFactory will create a module from the arguments passed into this function.
// If the arguments were wrong or incomplete for some reason, an error will be
// returned to indicate that the Module could not be instantiated.
// Each Module should have a single ModuleFactory that should be registered in
// the appropriate ModuleRegistry.
type ModuleFactory interface {
	New(args ...interface{}) (Module, error)
}

// ModuleFunc is a function that creates a new module.
type ModuleFunc func(args ...interface{}) (Module, error)

func (f ModuleFunc) New(args ...interface{}) (Module, error) {
	return f(args...)
}

// Module encapsulates a single task or unit of behavior. It is the building
// block of the rest of the system.
type Module interface {
	Execute() error
}

// ModuleRegistry keeps a lookup table to access registered ModuleFactory's by
// name.
type ModuleRegistry map[string]ModuleFactory

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

// DefaultRegistry is a global default registry for convenience.
var DefaultRegistry = &ModuleRegistry{}

// RegisterModule calls Register on the DefaultRegistry.
func RegisterModule(name string, factory ModuleFactory) {
	DefaultRegistry.Register(name, factory)
}

// LookupModule calls Lookup on the DefaultRegistry.
func LookupModule(name string) ModuleFactory {
	return DefaultRegistry.Lookup(name)
}
