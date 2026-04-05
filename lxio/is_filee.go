package lxio

import (
	"errors"
	"os"
)

// IsFileE returns true if the path exists and is a regular file.
// It returns an error for ambiguous failures (like Permission Denied).
func IsFileE(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// It definitely doesn't exist, so it's definitely not a file.
			return false, nil
		}
		// We can't access it to check, bubble up the error.
		return false, err
	}

	// It exists and we can read it. Is it a regular file?
	return info.Mode().IsRegular(), nil
}
