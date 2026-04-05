package lxio

import (
	"os"
)

// IsFile returns true if the path exists and is a regular file.
// It safely returns false if the path doesn't exist, is a directory,
// or if there is a permission error.
func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	// IsRegular ensures it is an actual file, not a directory,
	// socket, device, or named pipe.
	return info.Mode().IsRegular()
}
