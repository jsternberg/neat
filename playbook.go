package neat

import "errors"

var (
	errPlaybookEmpty = errors.New("playbook empty")
)

// Playbook keeps the underlying plan and execution order.
type Playbook struct {
	modules []Module
	safe    bool
}

// NewPlaybook creates a new playbook. The playbook defaults to running
// in safe mode and with no modules.
func NewPlaybook() *Playbook {
	return &Playbook{
		safe: true,
	}
}

// Add adds a Module to the list of modules to run in the playbook.
func (p *Playbook) Add(m Module) {
	p.modules = append(p.modules, m)
}

// Play runs the playbook.
func (p *Playbook) Play() (Stats, error) {
	stats := Stats{}
	if len(p.modules) == 0 {
		return stats, errPlaybookEmpty
	}
	return stats, nil
}
