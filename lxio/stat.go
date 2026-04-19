package lxio

import (
	"errors"
	"os"
	"time"
)

// Permission mode constants for Unix file permissions (informational only).
// These represent standard Unix permission bits and are provided for reference,
// but should not be relied upon to determine actual file access permissions.
//
// ⚠️  IMPORTANT: These constants reflect only permission bits, not actual capabilities.
// Actual file access depends on:
// - Process effective UID/GID and group membership
// - ACLs (Linux/macOS)
// - Filesystem type and mount options (NFS, Docker volumes, etc.)
// - SELinux or AppArmor policies
// - Process capabilities (Linux)
// - Platform-specific security models (Windows uses ACLs, not bits)
//
// For checking actual access capability, use the IsReadable/IsWritable functions
// which attempt actual file operations instead of checking bits.
const (
	PermExec  = 1 // Unix execute permission bit (--x) — informational only
	PermWrite = 2 // Unix write permission bit (-w-) — informational only
	PermRead  = 4 // Unix read permission bit (r--) — informational only
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

// ---------------------------------------------- Empty Stats  ----------------------------------------------------------

// IsEmpty returns true if the path is empty (file is 0 bytes or directory has no entries).
// Returns an error for nonexistent paths or ambiguous failures (like Permission Denied).
// For files: returns true if size is 0.
// For directories: returns true if there are no entries.
func IsEmpty(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// Path doesn't exist
			return false, err
		}
		// We can't access it, bubble up the error
		return false, err
	}

	// Check if it's a directory
	if info.IsDir() {
		// For directories, check if it has any entries
		entries, err := os.ReadDir(path)
		if err != nil {
			return false, err
		}
		return len(entries) == 0, nil
	}

	// For files, check if size is 0
	return info.Size() == 0, nil
}

// IsEmptyOK returns true if the path is empty (file is 0 bytes or directory has no entries).
// It ignores any errors and safely returns false.
func IsEmptyOK(path string) bool {
	ok, _ := IsEmpty(path)
	return ok
}

// ---------------------------------------------- Size Stats  ----------------------------------------------------------

// Size returns the size of the file in bytes.
// It returns an error for nonexistent paths or ambiguous failures (like Permission Denied).
// For directories, it returns the size of the directory itself, not the size of its contents.
func Size(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// Path doesn't exist, return error
			return 0, err
		}
		// We can't access it, bubble up the error
		return 0, err
	}

	return info.Size(), nil
}

// SizeOK returns the size of the file in bytes.
// It ignores any errors and safely returns 0.
// For directories, it returns the size of the directory itself, not the size of its contents.
func SizeOK(path string) int64 {
	size, _ := Size(path)
	return size
}

// ---------------------------------------------- ModTime Stats  -------------------------------------------------------

// ModTime returns the last modification time of the file or directory.
// It returns an error for nonexistent paths or ambiguous failures (like Permission Denied).
func ModTime(path string) (time.Time, error) {
	info, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// Path doesn't exist, return error
			return time.Time{}, err
		}
		// We can't access it, bubble up the error
		return time.Time{}, err
	}

	return info.ModTime(), nil
}

// ---------------------------------------------- Permissions Stats  ---------------------------------------------------

// IsReadable attempts to open the path for reading.
// Returns true if the operation succeeds, false otherwise.
// This reflects actual read permission taking into account effective UID/GID,
// group membership, ACLs, filesystem type, and platform-specific access controls.
// This is more reliable than checking permission bits directly.
//
// Note: This performs an actual file operation and is subject to TOCTOU races.
// For security-critical operations, re-check or use atomic operations.
func IsReadable(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	f.Close()
	return true
}

// IsWritable attempts to determine if the path is writable.
// For files: attempts to open for writing without truncating.
// For directories: checks if we can create a temporary file in the directory.
// Returns true if the operation succeeds, false otherwise.
// This reflects actual write permission taking into account effective UID/GID,
// group membership, ACLs, filesystem type, and platform-specific access controls.
// This is more reliable than checking permission bits directly.
//
// Note: This performs actual file operations and is subject to TOCTOU races.
// For security-critical operations, re-check or use atomic operations.
func IsWritable(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	// For directories, try to create a temporary file
	if info.IsDir() {
		tempFile, err := os.CreateTemp(path, ".write-check-")
		if err != nil {
			return false
		}
		tempFile.Close()
		os.Remove(tempFile.Name())
		return true
	}

	// For files, try to open for writing
	f, err := os.OpenFile(path, os.O_WRONLY, 0)
	if err != nil {
		return false
	}
	f.Close()
	return true
}

// IsExecutable checks if the path exists and the executable bit is set in owner permissions.
// Warning: This only checks permission bits, not actual execute capability.
// On Unix: checks owner execute bit
// On Windows: returns true if path exists and is a regular file (Windows doesn't use execute bits)
//
// This is a limited check. In real systems, execute permission depends on:
// - Filesystem capabilities and mount options
// - SELinux/AppArmor policies
// - Process capabilities and sandboxing
// The only reliable way to check if something is executable is to attempt execution.
func IsExecutable(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	// On Windows, execute permission doesn't apply; just check if it's a regular file
	// os.Stat returns a mode that reflects file type but not Windows ACLs
	if os.PathSeparator == '\\' {
		// Windows: no execute bit model, just check file exists
		return info.Mode().IsRegular()
	}

	// Unix: check if owner has execute permission (simplified check)
	// This is just a best-guess based on permission bits
	ownerPerm := (info.Mode() >> 6) & 07
	return (ownerPerm & 1) != 0
}
