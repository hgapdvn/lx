package lxio

import (
	"errors"
	"os"
)

// ----------------------------------------------- Exists Stats -------------------------------------------------------

// Exists returns true if the path exists.
// It returns an error for ambiguous failures (e.g. permission denied).
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// ExistsOK returns true if the path exists.
// It ignores any errors and safely returns false.
func ExistsOK(path string) bool {
	ok, _ := Exists(path)
	return ok
}

// MustExist returns true if the path exists.
// It panics if an error for ambiguous failures (e.g. permission denied).
func MustExist(path string) bool {
	ok, err := Exists(path)
	if err != nil {
		panic(err)
	}
	return ok
}

// NotExists returns true if the file does not exist, false if it explicitly does.
// It returns an error for ambiguous failures (like Permission Denied).
// When an error occurs, returns (false, error)—conservative assumption that existence cannot be determined.
func NotExists(path string) (bool, error) {
	exists, err := Exists(path)
	if err != nil {
		return false, err
	}
	return !exists, nil
}

// NotExistsOK returns true if the file does not exist, false if it explicitly does.
// It swallows any errors and safely defaults to false (conservative: assume file exists or is inaccessible).
func NotExistsOK(path string) bool {
	ok, _ := NotExists(path)
	return ok
}

// MustNotExist return true if the file does not exist, false if it explicitly does.
// It panics if an error for ambiguous failures (like Permission Denied).
//
// WARNING: When MustNotExist panics, the file existence state is ambiguous.
// This function is suitable for scenarios where you expect either the file to definitely not exist,
// or to panic on any access issues. Use NotExists() if you need explicit error handling.
func MustNotExist(path string) bool {
	exists, err := NotExists(path)
	if err != nil {
		panic(err)
	}
	return exists
}

// ---------------------------------------------- Dir Stats  -----------------------------------------------------------

// IsDir returns true if the path exists and is a directory.
// It returns an error for ambiguous failures (like Permission Denied).
func IsDir(path string) (bool, error) {
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

// IsDirOK returns true if the path exists and is a directory.
// It ignores any errors and safely returns false.
func IsDirOK(path string) bool {
	ok, _ := IsDir(path)
	return ok
}

// MustBeDir returns true if the path exists and is a directory.
// It panics if an error for ambiguous failures (like Permission Denied).
func MustBeDir(path string) bool {
	isDir, err := IsDir(path)
	if err != nil {
		panic(err)
	}
	return isDir
}

// ---------------------------------------------- File Stats  ----------------------------------------------------------

// IsFile returns true if the path exists and is a regular file.
// It returns an error for ambiguous failures (like Permission Denied).
func IsFile(path string) (bool, error) {
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

// IsFileOK returns true if the path exists and is a regular file.
// It ignores any errors and safely returns false.
func IsFileOK(path string) bool {
	isFile, _ := IsFile(path)
	return isFile
}

// MustBeFile returns true if the path exists and is a regular file.
// It panics if an error for ambiguous failures (like Permission Denied).
func MustBeFile(path string) bool {
	isFile, err := IsFile(path)
	if err != nil {
		panic(err)
	}
	return isFile
}

// ---------------------------------------------- Symlink Stats  -------------------------------------------------------

// IsSymlink returns true if the path exists and is a symbolic link.
// It returns an error for ambiguous failures (like Permission Denied).
func IsSymlink(path string) (bool, error) {
	info, err := os.Lstat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}

	return info.Mode()&os.ModeSymlink != 0, nil
}

// IsSymlinkOK returns true if the path exists and is a symbolic link.
// It ignores any errors and safely returns false.
func IsSymlinkOK(path string) bool {
	isSymlink, _ := IsSymlink(path)
	return isSymlink
}

// MustBeSymlink returns true if the path exists and is a symbolic link.
// It panics if an error for ambiguous failures (like Permission Denied).
func MustBeSymlink(path string) bool {
	isSymlink, err := IsSymlink(path)
	if err != nil {
		panic(err)
	}
	return isSymlink
}
