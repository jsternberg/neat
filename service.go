package neat

type serviceModule struct{}

func (m *serviceModule) Execute(p *Playbook) (interface{}, ModuleStatus, error) {
	return nil, ModuleOk, nil
}
