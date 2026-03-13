package lxenv

import (
	"os"
	"strconv"
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
