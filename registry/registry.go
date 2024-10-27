package registry

import "sync"

// Registry is the global registry for services
var Registry sync.Map

// Register adds a service to the registry
func Register(name string, service any) {
	if Exist(name) {
		panic("[Registry] " + name + " already exists, to overwrite use Unregister first")
	}

	Registry.Store(name, service)
}

// Unregister removes a service from the registry
func Unregister(name string) {
	Registry.Delete(name)
}

// Get retrieves a service from the registry by name
// this function will panic if the service is not found or the type is incorrect
func Get[T any](name string) T {
	o, ok := Registry.Load(name)

	if !ok {
		panic("[Registry] " + name + " not found")
	}

	return o.(T)
}

// Exist checks if a service exists in the registry
func Exist(name string) bool {
	_, ok := Registry.Load(name)
	return ok
}
