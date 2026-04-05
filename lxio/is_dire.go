package lxio

import (
	"errors"
	"os"
)

// IsDirE returns true if the path exists and is a directory.
// It returns an error for ambiguous failures (like Permission Denied).
func IsDirE(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// It definitely doesn't exist, so it's definitely not a directory.
			return false, nil
		}
		// We can't access it to check, bubble up the error.
		return false, err
	}

	// It exists and we can read it. Is it a directory?
	return info.IsDir(), nil
}
