package neat

import "os"

type fileModule struct {
	Path  string
	Mode  os.FileMode
	State string
}

func (m *fileModule) Execute(p *Playbook) (interface{}, ModuleStatus, error) {
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
