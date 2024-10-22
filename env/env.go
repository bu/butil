package env

import (
	"fmt"
	"strconv"
)

// envProvider is the EnvProvider implementation used by the package.
var envProvider Provider = &OsProvider{}

// SetEnvProvider sets the EnvProvider implementation used by the package.
func SetEnvProvider(p Provider) {
	envProvider = p
}

// Getenv returns the value of the environment variable named by the key.
func Get(key string, defaults ...string) string {
	val := envProvider.Get(key)
	if val == "" && len(defaults) > 0 {
		return defaults[0]
	}
	return val
}

// MustGet returns the value of the environment variable named by the key.
// If the key is not found, it panics.
func MustGet(key string) string {
	val := envProvider.Get(key)
	if val == "" {
		panic(fmt.Sprintf("[%s] Cannot get environment variable %s", envProvider.Name(), key))
	}
	return val
}

// GetInt returns the value of the environment variable named by the key as an int.
func GetInt(key string, defaults ...int64) int64 {
	val := envProvider.Get(key)
	hasDefault := len(defaults) > 0

	if val == "" {
		if hasDefault {
			return defaults[0]
		}

		return 0
	}

	// Parse the value as an int
	vInt, err := strconv.ParseInt(val, 10, 64)
	if err != nil && !hasDefault {
		panic(fmt.Sprintf("[%s] Cannot parse environment variable %s as int, and there's no default value", envProvider.Name(), key))
	}

	// Return the value
	return vInt
}
