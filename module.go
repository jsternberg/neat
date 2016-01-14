package neat

import "fmt"

type ModuleStatus int

const (
	// ModuleOk marks that an action ran, but didn't have to do anything since
	// it was already complete.
	ModuleOk ModuleStatus = iota
	// ModuleChanged marks that an action ran and changed something. Modules
	// that do not trigger or effect other tasks may return ModuleChanged
	// instead of ModuleOk if it is determined that performing the task again
	// would take less time than checking if it should be done.
	ModuleChanged
	// ModuleDeferred marks that an action needed to be deferred for some reason.
	// The most common reason is that the action was unsafe and the module was
	// run in a safe context.
	ModuleDeferred
	// ModuleSkipped marks that an action was skipped and never plans on running.
	ModuleSkipped
	// ModuleFailed marks that an action failed. While this should be returned
	// when a module fails, the marker of a failed action is the error return
	// value associated with this return value.
	// This status is mostly included for completeness and so we don't return
	// a ModuleOk with a failure message.
	ModuleFailed
)

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
	// Execute runs this module. The first return parameter contains any
	// extra information from the result of running this module. The second
	// contains the status. The third contains an error for any failure that
	// may have occurred. If an error is returned, the first two return values
	// may be anything.
	Execute(p *Playbook) (interface{}, ModuleStatus, error)
}

// ModuleE is a convenience struct for methods commonly used when writing
// modules. It is designed as a wrapper around the module interface.
// It is short for ModuleExecutor.
type ModuleE struct {
	Module
}

// Fail is a convenience function for returning a module failure.
// It returns nothing as the result interface,
func (m ModuleE) Fail(v interface{}) (interface{}, ModuleStatus, error) {
	return nil, ModuleFailed, fmt.Errorf("%v", v)
}

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
