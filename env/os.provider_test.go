package env

import "testing"

func TestOsProvider(t *testing.T) {
	// setup the environment
	t.Setenv("TEST_KEY", "test_value")

	// create a new instance of the OsProvider
	provider := &OsProvider{}

	t.Run("Get", func(t *testing.T) {
		expected := "test_value"

		if actual := provider.Get("TEST_KEY"); actual != expected {
			t.Errorf("Get did not return the expected value, expected: %s, got: %s", expected, actual)
		}
	})

	t.Run("Name", func(t *testing.T) {
		expected := "os"

		if actual := provider.Name(); actual != expected {
			t.Errorf("Name did not return the expected value, expected: %s, got: %s", expected, actual)
		}
	})
}
