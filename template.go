package neat

type templateModule struct{}

func (m *templateModule) Execute(p *Playbook) (interface{}, ModuleStatus, error) {
	return nil, ModuleOk, nil
}
