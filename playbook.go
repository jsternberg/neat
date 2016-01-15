package neat

import (
	"errors"
	"fmt"
)

var (
	errPlaybookEmpty = errors.New("playbook empty")
)

// Playbook keeps the underlying plan and execution order.
type Playbook struct {
	modules   []Module
	safe      bool
	checkMode bool
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

// SetSafe changes if the playbook is being run in safe mode.
func (p *Playbook) SetSafe(safe bool) {
	p.safe = safe
}

// Safe returns if this playbook is in safe mode.
func (p *Playbook) Safe() bool {
	return p.safe
}

// SetCheckMode changes if the playbook is being run in check mode.
func (p *Playbook) SetCheckMode(checkMode bool) {
	p.checkMode = checkMode
}

// CheckMode returns if this playbook is being run in check mode.
func (p *Playbook) CheckMode() bool {
	return p.checkMode
}

// Play runs the playbook.
func (p *Playbook) Play() (Stats, error) {
	stats := Stats{}
	if len(p.modules) == 0 {
		return stats, errPlaybookEmpty
	}

	for _, m := range p.modules {
		_, status, err := m.Execute(p)
		if err != nil {
			return stats, fmt.Errorf("playbook execution error: %s", err)
		}
		stats.Record(status)
	}
	return stats, nil
}
