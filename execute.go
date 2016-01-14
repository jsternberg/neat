package neat

import (
	"bytes"
	"os/exec"
)

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
