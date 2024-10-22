package env

import "testing"

func TestGet(t *testing.T) {
	// Save the original envProvider so we can restore it later
	originalEnvProvider := envProvider

	// Create a new mock provider
	mockProvider := &DummyProvider{
		Env: map[string]string{
			"TEST_KEY": "test_value",
		},
	}

	// Set the mock provider as the envProvider
	SetEnvProvider(mockProvider)

	// Begin Test
	t.Run("GetExistValue", func(t *testing.T) {
		expected := "test_value"

		if actual := Get("TEST_KEY"); actual != expected {
			t.Errorf("Get did not return the expected value, expected: %s, got: %s", expected, actual)
		}
	})

	t.Run("GetExistValueWithDefault", func(t *testing.T) {
		expected := "test_value"

		if actual := Get("TEST_KEY", "default_value"); actual != expected {
			t.Errorf("Get did not return the expected value, expected: %s, got: %s", expected, actual)
		}
	})

	t.Run("GetNotExistValue", func(t *testing.T) {
		if actual := Get("TEST_KEY_NOT_EXIST"); actual != "" {
			t.Error("Get did not return the expected value, expected: empty string, got:", actual)
		}
	})

	t.Run("GetNotExistValueWithDefault", func(t *testing.T) {
		expected := "default_value"

		if actual := Get("TEST_KEY_NOT_EXIST", expected); actual != expected {
			t.Errorf("Get did not return the expected value, expected: %s, got: %s", expected, actual)
		}
	})

	// Restore the original envProvider
	envProvider = originalEnvProvider
}
func TestMustGet(t *testing.T) {
	// Save the original envProvider so we can restore it later
	originalEnvProvider := envProvider

	// Create a new mock provider
	mockProvider := &DummyProvider{
		Env: map[string]string{
			"TEST_KEY": "test_value",
		},
	}

	// Set the mock provider as the envProvider
	SetEnvProvider(mockProvider)

	// Begin Test
	t.Run("MustGetExistValue", func(t *testing.T) {
		// Test the MustGet function
		if MustGet("TEST_KEY") != "test_value" {
			t.Error("MustGet did not return the expected value")
		}
	})

	t.Run("MustGetNotExistValue", func(t *testing.T) {
		// Test the MustGet function
		defer func() {
			if r := recover(); r == nil {
				t.Error("MustGet did not panic")
			}
		}()

		MustGet("TEST_KEY_NOT_EXIST")
	})

	// Restore the original envProvider
	envProvider = originalEnvProvider
}

func TestGetInt(t *testing.T) {
	// Save the original envProvider so we can restore it later
	originalEnvProvider := envProvider

	// Create a new mock provider
	mockProvider := &DummyProvider{
		Env: map[string]string{
			"TEST_KEY":               "123",
			"TEST_KEY_INVALID_VALUE": "invalid_value",
		},
	}

	// Set the mock provider as the envProvider
	SetEnvProvider(mockProvider)

	// Begin Test
	t.Run("GetIntExistValue", func(t *testing.T) {
		expected := int64(123)

		if actual := GetInt("TEST_KEY"); actual != expected {
			t.Errorf("GetInt did not return the expected value, expected: %d, got: %d", expected, actual)
		}
	})

	t.Run("GetIntExistValueWithDefault", func(t *testing.T) {
		expected := int64(123)

		if actual := GetInt("TEST_KEY", 456); actual != expected {
			t.Errorf("GetInt did not return the expected value, expected: %d, got: %d", expected, actual)
		}
	})

	t.Run("GetIntNotExistValue", func(t *testing.T) {
		if actual := GetInt("TEST_KEY_NOT_EXIST"); actual != 0 {
			t.Error("GetInt did not return the expected value, expected: 0, got:", actual)
		}
	})

	t.Run("GetIntNotExistValueWithDefault", func(t *testing.T) {
		expected := int64(456)

		if actual := GetInt("TEST_KEY_NOT_EXIST", expected); actual != expected {
			t.Errorf("GetInt did not return the expected value, expected: %d, got: %d", expected, actual)
		}
	})

	t.Run("GetIntExistValueWithInvalidValue", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("GetInt did not panic")
			}
		}()

		GetInt("TEST_KEY_INVALID_VALUE")
	})

	// Restore the original envProvider
	envProvider = originalEnvProvider
}
