package lxtypes_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/nthanhhai2909/lx/types"
)

// Test types
type Config struct {
	Host string
	Port int
}

func TestResultSuccess(t *testing.T) {
	t.Run("integer", func(t *testing.T) {
		result := lxtypes.ResultSuccess(42)

		value, err := result.Value()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if value != 42 {
			t.Errorf("Value() = %v, want 42", value)
		}
	})

	t.Run("string", func(t *testing.T) {
		result := lxtypes.ResultSuccess("hello")

		value, err := result.Value()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if value != "hello" {
			t.Errorf("Value() = %v, want hello", value)
		}
	})

	t.Run("struct", func(t *testing.T) {
		config := Config{Host: "localhost", Port: 8080}
		result := lxtypes.ResultSuccess(config)

		value, err := result.Value()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if value.Host != "localhost" || value.Port != 8080 {
			t.Errorf("Value() = %+v, want %+v", value, config)
		}
	})

	t.Run("pointer of struct", func(t *testing.T) {
		config := &Config{Host: "example.com", Port: 3000}
		result := lxtypes.ResultSuccess(config)

		value, err := result.Value()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if value.Host != "example.com" || value.Port != 3000 {
			t.Errorf("Value() = %+v, want %+v", value, config)
		}
	})

	t.Run("zero value", func(t *testing.T) {
		// Even zero values should succeed
		result := lxtypes.ResultSuccess(0)

		value, err := result.Value()
		if err != nil {
			t.Errorf("Expected no error for zero value, got %v", err)
		}
		if value != 0 {
			t.Errorf("Value() = %v, want 0", value)
		}
	})
}

func TestResultFailure(t *testing.T) {
	testErr := errors.New("test error")

	t.Run("integer", func(t *testing.T) {
		result := lxtypes.ResultFailure[int](testErr)

		value, err := result.Value()
		if err == nil {
			t.Error("Expected error, got nil")
		}
		if err.Error() != testErr.Error() {
			t.Errorf("Error = %v, want %v", err, testErr)
		}
		if value != 0 {
			t.Errorf("Value should be zero for failure, got %v", value)
		}
	})

	t.Run("string", func(t *testing.T) {
		result := lxtypes.ResultFailure[string](testErr)

		value, err := result.Value()
		if err == nil {
			t.Error("Expected error, got nil")
		}
		if value != "" {
			t.Errorf("Value should be empty string for failure, got %v", value)
		}
	})

	t.Run("struct", func(t *testing.T) {
		result := lxtypes.ResultFailure[Config](testErr)

		value, err := result.Value()
		if err == nil {
			t.Error("Expected error, got nil")
		}
		if value.Host != "" || value.Port != 0 {
			t.Errorf("Value should be zero Config for failure, got %+v", value)
		}
	})

	t.Run("pointer of struct", func(t *testing.T) {
		result := lxtypes.ResultFailure[*Config](testErr)

		value, err := result.Value()
		if err == nil {
			t.Error("Expected error, got nil")
		}
		if value != nil {
			t.Errorf("Value should be nil for failure, got %+v", value)
		}
	})
}

func TestResultValueOr(t *testing.T) {
	t.Run("success returns original", func(t *testing.T) {
		success := lxtypes.ResultSuccess(42)
		if got := success.ValueOr(0); got != 42 {
			t.Errorf("Success.ValueOr(0) = %v, want 42", got)
		}
	})

	t.Run("failure returns default", func(t *testing.T) {
		failure := lxtypes.ResultFailure[int](errors.New("error"))
		if got := failure.ValueOr(99); got != 99 {
			t.Errorf("Failure.ValueOr(99) = %v, want 99", got)
		}
	})

	t.Run("struct with success", func(t *testing.T) {
		config := Config{Host: "localhost", Port: 8080}
		success := lxtypes.ResultSuccess(config)
		defaultConfig := Config{Host: "default", Port: 80}

		got := success.ValueOr(defaultConfig)
		if got.Host != "localhost" {
			t.Errorf("ValueOr() = %+v, want %+v", got, config)
		}
	})

	t.Run("struct with failure", func(t *testing.T) {
		failure := lxtypes.ResultFailure[Config](errors.New("error"))
		defaultConfig := Config{Host: "default", Port: 80}

		got := failure.ValueOr(defaultConfig)
		if got.Host != "default" {
			t.Errorf("ValueOr() = %+v, want %+v", got, defaultConfig)
		}
	})

	t.Run("pointer of struct with success", func(t *testing.T) {
		config := &Config{Host: "localhost", Port: 8080}
		success := lxtypes.ResultSuccess(config)
		defaultConfig := &Config{Host: "default", Port: 80}

		got := success.ValueOr(defaultConfig)
		if got.Host != "localhost" {
			t.Errorf("ValueOr() = %+v, want %+v", got, config)
		}
	})

	t.Run("pointer of struct with failure", func(t *testing.T) {
		failure := lxtypes.ResultFailure[*Config](errors.New("error"))
		defaultConfig := &Config{Host: "default", Port: 80}

		got := failure.ValueOr(defaultConfig)
		if got.Host != "default" {
			t.Errorf("ValueOr() = %+v, want %+v", got, defaultConfig)
		}
	})
}

func TestResultFromError(t *testing.T) {
	t.Run("nil error creates success", func(t *testing.T) {
		value, err := strconv.Atoi("42")
		result := convertToResult(value, err)

		v, e := result.Value()
		if e != nil {
			t.Errorf("Expected no error, got %v", e)
		}
		if v != 42 {
			t.Errorf("Value = %v, want 42", v)
		}
	})

	t.Run("non-nil error creates failure", func(t *testing.T) {
		value, err := strconv.Atoi("invalid")
		result := convertToResult(value, err)

		_, e := result.Value()
		if e == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("struct with nil error", func(t *testing.T) {
		config := Config{Host: "localhost", Port: 8080}
		result := convertToResult(config, nil)

		v, e := result.Value()
		if e != nil {
			t.Errorf("Expected no error, got %v", e)
		}
		if v.Host != "localhost" {
			t.Errorf("Value = %+v, want %+v", v, config)
		}
	})

	t.Run("struct with error", func(t *testing.T) {
		result := convertToResult(Config{}, errors.New("config error"))

		_, e := result.Value()
		if e == nil {
			t.Error("Expected error, got nil")
		}
	})
}

// Helper function to convert (value, error) to Result
func convertToResult[T any](value T, err error) lxtypes.Result[T] {
	if err != nil {
		return lxtypes.ResultFailure[T](err)
	}
	return lxtypes.ResultSuccess(value)
}

func TestResultChaining(t *testing.T) {
	t.Run("chain with success", func(t *testing.T) {
		result := lxtypes.ResultSuccess(42)
		value, err := result.Value()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		finalValue := value + 10
		if finalValue != 52 {
			t.Errorf("Expected 52, got %v", finalValue)
		}
	})

	t.Run("chain with failure", func(t *testing.T) {
		result := lxtypes.ResultFailure[int](errors.New("error"))
		value := result.ValueOr(0) + 10
		if value != 10 {
			t.Errorf("Expected 10, got %v", value)
		}
	})

	t.Run("struct modification", func(t *testing.T) {
		result := lxtypes.ResultSuccess(Config{Host: "localhost", Port: 8080})
		config := result.ValueOr(Config{Host: "default", Port: 80})
		config.Port += 1

		if config.Port != 8081 {
			t.Errorf("Expected Port 8081, got %v", config.Port)
		}
	})

	t.Run("pointer struct modification", func(t *testing.T) {
		original := &Config{Host: "localhost", Port: 8080}
		result := lxtypes.ResultSuccess(original)
		config := result.ValueOr(&Config{Host: "default", Port: 80})
		config.Port += 1

		// Should modify the original
		if original.Port != 8081 {
			t.Errorf("Expected Port 8081, got %v", original.Port)
		}
	})
}

func TestResultValueErrorPattern(t *testing.T) {
	t.Run("success with value-error pattern", func(t *testing.T) {
		result := lxtypes.ResultSuccess(42)

		if value, err := result.Value(); err == nil {
			if value != 42 {
				t.Errorf("Expected value 42, got %v", value)
			}
		} else {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("failure with value-error pattern", func(t *testing.T) {
		result := lxtypes.ResultFailure[int](errors.New("test error"))

		if value, err := result.Value(); err != nil {
			if value != 0 {
				t.Errorf("Expected zero value, got %v", value)
			}
			if err.Error() != "test error" {
				t.Errorf("Expected 'test error', got %v", err)
			}
		} else {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("struct with value-error pattern", func(t *testing.T) {
		config := Config{Host: "localhost", Port: 8080}
		result := lxtypes.ResultSuccess(config)

		if value, err := result.Value(); err == nil {
			if value.Host != "localhost" {
				t.Errorf("Expected localhost, got %v", value.Host)
			}
		} else {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}

func TestResultRealWorldScenarios(t *testing.T) {
	t.Run("database query simulation", func(t *testing.T) {
		// Simulate successful database query
		findUser := func(id int) lxtypes.Result[Config] {
			if id > 0 {
				return lxtypes.ResultSuccess(Config{Host: "user-service", Port: 5000})
			}
			return lxtypes.ResultFailure[Config](errors.New("user not found"))
		}

		result := findUser(123)
		config := result.ValueOr(Config{Host: "default", Port: 80})
		if config.Host != "user-service" {
			t.Errorf("Expected user-service, got %v", config.Host)
		}

		result2 := findUser(-1)
		config2 := result2.ValueOr(Config{Host: "default", Port: 80})
		if config2.Host != "default" {
			t.Errorf("Expected default, got %v", config2.Host)
		}
	})

	t.Run("API call simulation", func(t *testing.T) {
		// Simulate API call
		callAPI := func(endpoint string) lxtypes.Result[string] {
			if endpoint != "" {
				return lxtypes.ResultSuccess(fmt.Sprintf("Response from %s", endpoint))
			}
			return lxtypes.ResultFailure[string](errors.New("invalid endpoint"))
		}

		result := callAPI("/users")
		if response, err := result.Value(); err == nil {
			if response != "Response from /users" {
				t.Errorf("Expected 'Response from /users', got %v", response)
			}
		} else {
			t.Errorf("Expected success, got error: %v", err)
		}
	})

	t.Run("file operation simulation", func(t *testing.T) {
		// Simulate file read
		readConfig := func(path string) lxtypes.Result[*Config] {
			if path == "config.json" {
				return lxtypes.ResultSuccess(&Config{Host: "file-host", Port: 9000})
			}
			return lxtypes.ResultFailure[*Config](errors.New("file not found"))
		}

		result := readConfig("config.json")
		if config, err := result.Value(); err == nil {
			if config.Host != "file-host" {
				t.Errorf("Expected file-host, got %v", config.Host)
			}
		} else {
			t.Errorf("Expected success, got error: %v", err)
		}

		result2 := readConfig("missing.json")
		defaultConfig := &Config{Host: "default", Port: 80}
		config2 := result2.ValueOr(defaultConfig)
		if config2.Host != "default" {
			t.Errorf("Expected default, got %v", config2.Host)
		}
	})
}
