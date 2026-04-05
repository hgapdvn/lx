package lxio

import (
	"errors"
	"os"
)

// ExistsE returns true if the file exists, false if it explicitly does not.
// It returns an error for ambiguous failures (like Permission Denied).
func ExistsE(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
