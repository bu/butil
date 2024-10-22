package env

// Provider is an interface for getting environment variables.
type Provider interface {
	Get(key string) string // returns the value of the environment variable named by the key
	Name() string          // returns the name of the provider
}
