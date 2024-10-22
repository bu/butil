package env

type DummyProvider struct {
	Env map[string]string
}

func (p *DummyProvider) Get(key string) string {
	v, ok := p.Env[key]
	if !ok {
		return ""
	}
	return v
}

func (p *DummyProvider) Name() string {
	return "dummy"
}
