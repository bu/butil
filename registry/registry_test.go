package registry

import (
	"reflect"
	"sync"
	"testing"
)

type TestInterface interface {
	Test()
}

type TestStruct struct {
}

func (t *TestStruct) Test() {
	return
}

// TestRegister
func TestRegister(t *testing.T) {
	// clear the registry
	Registry = sync.Map{}

	// Try to register a service
	t.Run("Register", func(t *testing.T) {
		Register("test", &TestStruct{})

		// Check if the service is registered
		_, ok := Registry.Load("test")
		if !ok {
			t.Error("Register did not register the service")
		}
	})

	// Try to register the same service again
	// it should panic because the service is already registered
	t.Run("RegisterAgain", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Register did not panic when registering the same service again")
			}
		}()

		Register("test", &TestStruct{})
	})
}

// TestUnregister
func TestUnregister(t *testing.T) {
	// clear the registry
	Registry = sync.Map{}

	// Register a service
	Register("test", &TestStruct{})

	// Try to unregister the service
	t.Run("Unregister", func(t *testing.T) {
		Unregister("test")

		// Check if the service is unregistered
		_, ok := Registry.Load("test")
		if ok {
			t.Error("Unregister did not unregister the service")
		}
	})

	// Try to unregister the service again
	// it should just work
	t.Run("UnregisterAgain", func(t *testing.T) {
		Unregister("test")
	})

	// Try to unregister a service that does not exist
	// it should not panic
	t.Run("UnregisterNotExist", func(t *testing.T) {
		Unregister("test_not_exist")
	})
}

// TestExist
func TestExist(t *testing.T) {
	// clear the registry
	Registry = sync.Map{}

	// Register a service
	Register("test", &TestStruct{})

	// Check if the service exists
	t.Run("Exist", func(t *testing.T) {
		if !Exist("test") {
			t.Error("Exist did not find the service")
		}
	})

	// Check if the service does not exist
	t.Run("NotExist", func(t *testing.T) {
		if Exist("test_not_exist") {
			t.Error("Exist found a service that does not exist")
		}
	})
}

// TestGet
func TestGet(t *testing.T) {
	// clear the registry
	Registry = sync.Map{}

	// Register a service
	Register("test", TestStruct{})

	// Get the service
	t.Run("Get", func(t *testing.T) {
		if actual := Get[TestStruct]("test"); reflect.TypeOf(actual) != reflect.TypeOf(TestStruct{}) {
			t.Error("Get did not return the service")
		}
	})

	// Get a service that does not exist
	// it should panic
	t.Run("GetNotExist", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Get did not panic when the service does not exist")
			}
		}()

		Get[TestStruct]("test_not_exist")
	})

	// Clear the registry, and test for interface type
	Registry = sync.Map{}

	// Register a service
	Register("test", &TestStruct{})

	// Get the service
	t.Run("GetInterface", func(t *testing.T) {
		if actual := Get[TestInterface]("test"); reflect.TypeOf(actual) != reflect.TypeOf(&TestStruct{}) {
			t.Error("Get did not return the service")
		}
	})

	// if we try to get pointer to interface, it should panic
	t.Run("GetInterfacePointer", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Get did not panic when the service does not exist")
			}
		}()

		Get[*TestInterface]("test")
	})
}
