package lxio

import (
	"os"
)

// IsDir returns true if the path exists and is a directory.
// It safely returns false if the path doesn't exist, is a file,
// or if there is a permission error.
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return info.IsDir()
}
