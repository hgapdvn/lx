package lxenv

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	ErrKeyNotFound = errors.New("lxenv: environment variable not found")
)

// Get retrieves the value of an environment variable.
// Returns the value if set, empty string otherwise.
//
// Example:
//
//	value := lxenv.Get("HOME")
//	// value: "/Users/username" (or empty if not set)
func Get(key string) string {
	return os.Getenv(key)
}

// GetOr retrieves the value of an environment variable or returns a default value.
// Returns the environment variable value if set and non-empty, otherwise returns defaultValue.
//
// Example:
//
//	port := lxenv.GetOr("PORT", "8080")
//	// port: "8080" if PORT is not set
func GetOr(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// MustGet retrieves the value of an environment variable.
// Panics if the variable is not set.
//
// Example:
//
//	value := lxenv.MustGet("HOME")
//	// value: "/Users/username"
func MustGet(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic("lxenv: environment variable " + key + " is not set")
	}
	return value
}

// Set sets the value of an environment variable.
// Returns an error if the operation fails.
//
// Example:
//
//	err := lxenv.Set("API_KEY", "secret")
func Set(key, value string) error {
	return os.Setenv(key, value)
}

// Unset removes an environment variable.
// Returns an error if the operation fails.
//
// Example:
//
//	err := lxenv.Unset("TEMP_VAR")
func Unset(key string) error {
	return os.Unsetenv(key)
}

// Has checks if an environment variable is set (even if empty).
// Returns true if the variable exists, false otherwise.
//
// Example:
//
//	if lxenv.Has("DEBUG") {
//	    // DEBUG variable is set
//	}
func Has(key string) bool {
	_, exists := os.LookupEnv(key)
	return exists
}

// NotHas is an alias for !Has.
func NotHas(key string) bool {
	return !Has(key)
}

// Exists is an alias for Has.
func Exists(key string) bool {
	return Has(key)
}

// NotExists is an alias for !Exists.
func NotExists(key string) bool {
	return !Exists(key)
}

// Lookup retrieves an environment variable and reports whether it was set.
// Returns (value, true) if the variable is set, (empty, false) otherwise.
// Unlike Get, this distinguishes between empty and unset variables.
//
// Example:
//
//	if value, exists := lxenv.Lookup("API_KEY"); exists {
//	    // Use value (might be empty string)
//	}
func Lookup(key string) (string, bool) {
	return os.LookupEnv(key)
}

// GetInt retrieves an environment variable as an integer.
// Returns (value, true) if the variable is set and can be parsed as an integer.
// Returns (0, false) if the variable is not set or cannot be parsed.
//
// Example:
//
//	if port, ok := lxenv.GetInt("PORT"); ok {
//	    // Use port as int
//	}
func GetInt(key string) (int, bool) {
	value := os.Getenv(key)
	if value == "" {
		return 0, false
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, false
	}
	return parsed, true
}

// GetIntOr retrieves an environment variable as an integer or returns a default value.
// Returns the parsed integer if the variable is set and valid, otherwise returns defaultValue.
//
// Example:
//
//	port := lxenv.GetIntOr("PORT", 8080)
//	// port: 8080 if PORT is not set or invalid
func GetIntOr(key string, defaultValue int) int {
	if value, ok := GetInt(key); ok {
		return value
	}
	return defaultValue
}

// MustGetInt retrieves an environment variable as an integer.
// Panics if the variable is not set or cannot be parsed as an integer.
//
// Example:
//
//	port := lxenv.MustGetInt("PORT")
//	// port: 8080
func MustGetInt(key string) int {
	value, ok := GetInt(key)
	if !ok {
		panic("lxenv: environment variable " + key + " is not set or not a valid integer")
	}
	return value
}

// GetBool retrieves an environment variable as a boolean.
// Returns (value, true) if the variable is set and can be parsed as a boolean.
// Accepts: "1", "t", "T", "true", "TRUE", "True" for true.
// Accepts: "0", "f", "F", "false", "FALSE", "False" for false.
// Returns (false, false) if the variable is not set or cannot be parsed.
//
// Example:
//
//	if debug, ok := lxenv.GetBool("DEBUG"); ok {
//	    // Use debug as bool
//	}
func GetBool(key string) (bool, bool) {
	value := os.Getenv(key)
	if value == "" {
		return false, false
	}
	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return false, false
	}
	return parsed, true
}

// GetBoolOr retrieves an environment variable as a boolean or returns a default value.
// Returns the parsed boolean if the variable is set and valid, otherwise returns defaultValue.
//
// Example:
//
//	debug := lxenv.GetBoolOr("DEBUG", false)
//	// debug: false if DEBUG is not set or invalid
func GetBoolOr(key string, defaultValue bool) bool {
	if value, ok := GetBool(key); ok {
		return value
	}
	return defaultValue
}

// MustGetBool retrieves an environment variable as a boolean.
// Panics if the variable is not set or cannot be parsed as a boolean.
//
// Example:
//
//	debug := lxenv.MustGetBool("DEBUG")
//	// debug: true
func MustGetBool(key string) bool {
	value, ok := GetBool(key)
	if !ok {
		panic("lxenv: environment variable " + key + " is not set or not a valid boolean")
	}
	return value
}

// GetFloat retrieves an environment variable as a float64.
// Returns (value, true) if the variable is set and can be parsed as a float64.
// Returns (0, false) if the variable is not set or cannot be parsed.
//
// Example:
//
//	if v, ok := lxenv.GetFloat("THRESHOLD"); ok {
//	    // use v (float64)
//	}
func GetFloat(key string) (float64, bool) {
	value := os.Getenv(key)
	if value == "" {
		return 0, false
	}
	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, false
	}
	return parsed, true
}

// GetFloatOr retrieves an environment variable as a float64 or returns a default value.
// Returns the parsed float if the variable is set and valid, otherwise returns defaultValue.
//
// Example:
//
//	threshold := lxenv.GetFloatOr("THRESHOLD", 0.1)
//	// threshold: 0.1 if THRESHOLD is not set or invalid
func GetFloatOr(key string, defaultValue float64) float64 {
	if v, ok := GetFloat(key); ok {
		return v
	}
	return defaultValue
}

// MustGetFloat retrieves the value of an environment variable as a float64.
// Panics if the variable is not set or cannot be parsed as a float64.
//
// Example:
//
//	threshold := lxenv.MustGetFloat("THRESHOLD")
//	// threshold: 0.1
func MustGetFloat(key string) float64 {
	if v, ok := GetFloat(key); ok {
		return v
	}
	if _, exists := os.LookupEnv(key); exists {
		panic("lxenv: environment variable " + key + " is not a valid float")
	}
	panic("lxenv: environment variable " + key + " is not set")
}

// Require ensures the provided keys exist in the environment.
// Returns ErrKeyNotFound wrapped with the missing keys when any key is unset.
func Require(keys ...string) error {
	var missing []string
	for _, key := range keys {
		if NotExists(key) {
			missing = append(missing, key)
		}
	}

	if len(missing) == 0 {
		return nil
	}

	desc := fmt.Sprintf("missing required environment variables: %s", strings.Join(missing, ", "))

	return fmt.Errorf("%w: %s", ErrKeyNotFound, desc)
}
