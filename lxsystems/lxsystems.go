package lxsystems

import (
	"os"
	"runtime"
)

const (

	// OSName stores the operating system name on which the program is being executed, derived from runtime.GOOS.
	OSName = runtime.GOOS

	// OSArch stores the architecture on which the program is being executed, derived from runtime.GOARCH.
	OSArch = runtime.GOARCH
)

// GoVersion returns the current version of the Go runtime as a string.
func GoVersion() string {
	return runtime.Version()
}

// UserHomeDir returns the user's home directory, or an error if it cannot be determined.
func UserHomeDir() (string, error) {
	return os.UserHomeDir()
}

// UserHomeDirOrEmpty returns the user's home directory as a string.
// If the home directory cannot be determined, it returns an empty string.
func UserHomeDirOrEmpty() string {
	dir, _ := os.UserHomeDir()
	return dir
}

// UserHomeDirOr returns the user's home directory, or defaultValue if it cannot be determined.
func UserHomeDirOr(defaultValue string) string {
	dir, err := os.UserHomeDir()
	if err != nil {
		return defaultValue
	}
	return dir
}

// TempDir returns the system's default temporary directory as a string.'
func TempDir() string {
	return os.TempDir()
}

// WorkingDir returns the current working directory, or an error if it cannot be determined.
func WorkingDir() (string, error) {
	return os.Getwd()
}

// WorkingDirOrEmpty returns the current working directory as a string.
func WorkingDirOrEmpty() string {
	dir, _ := os.Getwd()
	return dir
}

// WorkingDirOr returns the current working directory, or defaultValue if it cannot be determined.
func WorkingDirOr(defaultValue string) string {
	dir, err := os.Getwd()
	if err != nil {
		return defaultValue
	}
	return dir
}

// NumCPU returns the number of logical CPUs available to the current process.
func NumCPU() int {
	return runtime.NumCPU()
}

// IsWindows returns true if the operating system is Windows, false otherwise.
func IsWindows() bool {
	return OSName == "windows"
}

// IsLinux returns true if the operating system is Linux, false otherwise.
func IsLinux() bool {
	return OSName == "linux"
}

// IsMacOS returns true if the operating system is MacOS, false otherwise.
func IsMacOS() bool {
	return OSName == "darwin"
}

// HostName returns the host name of the current machine.
// It returns an error if the host name cannot be determined.
func HostName() (string, error) {
	return os.Hostname()
}

func HostNameOrEmpty() string {
	name, _ := os.Hostname()
	return name
}

func HostNameOr(defaultValue string) string {
	name, err := os.Hostname()
	if err != nil {
		return defaultValue
	}
	return name
}

// PID returns the process id of the caller.
func PID() int {
	return os.Getpid()
}
