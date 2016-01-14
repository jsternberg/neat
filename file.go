package neat

import (
	"errors"
	"fmt"
	"os"
)

func File(args ...interface{}) (Module, error) {
	m := &fileModule{}
	if err := decodeArgs(m, args...); err != nil {
		return nil, err
	}

	if m.Path == "" {
		return nil, errors.New("path required")
	}
	switch m.State {
	case "present", "absent":
	case "":
		m.State = "present"
	default:
		return nil, fmt.Errorf("invalid state: %s", m.State)
	}
	return m, nil
}

type fileModule struct {
	Path  string
	Mode  os.FileMode
	State string
}

func (m *fileModule) Execute(p *Playbook) (interface{}, ModuleStatus, error) {
	switch m.State {
	case "present":
		return m.ensurePresent()
	case "absent":
		return m.ensureAbsent()
	default:
		panic(fmt.Sprintf("unknown state: %s", m.State))
	}
}

func (m *fileModule) ensurePresent() (interface{}, ModuleStatus, error) {
	e := ModuleE{Module: m}
	f, err := os.Open(m.Path)
	if err != nil {
		e.Changed = true
		f, err = os.Create(m.Path)
		if err != nil {
			return e.Fail(err)
		}
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		e.Fail(err)
	}

	if stat.Mode() != m.Mode {
		e.Changed = true
		if err := f.Chmod(m.Mode); err != nil {
			return e.Fail(err)
		}
	}
	return e.Ok(nil)
}

func (m *fileModule) ensureAbsent() (interface{}, ModuleStatus, error) {
	e := ModuleE{Module: m}
	_, err := os.Stat(m.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return e.Ok(nil)
		}
		return e.Fail(err)
	}

	e.Changed = true
	if err := os.Remove(m.Path); err != nil {
		e.Fail(err)
	}
	return e.Ok(nil)
}
