package env

import "os"

type OsProvider struct{}

func (p *OsProvider) Get(key string) string {
	return os.Getenv(key)
}

func (p *OsProvider) Name() string {
	return "os"
}
