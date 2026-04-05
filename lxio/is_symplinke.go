package lxio

import (
	"errors"
	"os"
)

// IsSymlinkE returns true if the path exists and is a symbolic link.
// It returns an error for ambiguous failures (like Permission Denied).
func IsSymlinkE(path string) (bool, error) {
	info, err := os.Lstat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}

	return info.Mode()&os.ModeSymlink != 0, nil
}
