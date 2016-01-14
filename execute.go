package neat

import (
	"bytes"
	"errors"
	"os/exec"
)

func Execute(args ...interface{}) (Module, error) {
	m := &executeModule{}
	if err := decodeArgs(m, args...); err != nil {
		return nil, err
	}

	if m.Command == "" {
		return nil, errors.New("command required")
	}
	return m, nil
}

type executeModule struct {
	Command string
}

func (m *executeModule) Execute(p *Playbook) (interface{}, ModuleStatus, error) {
	e := ModuleE{Module: m}
	e.Changed = true

	var stdout, stderr bytes.Buffer
	cmd := exec.Command("/bin/sh", "-c", m.Command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return e.Fail(err)
	}

	return e.Ok(map[string]string{
		"stdout": stdout.String(),
		"stderr": stderr.String(),
	})
}
