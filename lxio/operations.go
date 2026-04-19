package lxio

import (
	"io"
	"os"
	"path/filepath"
)

// CopyFile copies the file from src to dst.
// If dst already exists, it is truncated.
// The file permissions and modification time are not preserved.
// Returns an error if src doesn't exist or if the copy fails.
//
// Example:
//
//	err := lxio.CopyFile("/path/to/source.txt", "/path/to/dest.txt")
//	if err != nil {
//		log.Fatal(err)
//	}
func CopyFile(src, dst string) error {
	// Open source file for reading
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	// Create destination file
	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	// Copy contents
	if _, err = io.Copy(destination, source); err != nil {
		return err
	}

	// Explicitly close to capture any deferred write errors (e.g., on NFS).
	// The deferred Close above acts as a safety net for early returns.
	return destination.Close()
}

// MoveFile moves a file from src to dst.
// If dst already exists, it is replaced.
// This operation uses os.Rename which is atomic on most systems.
// Returns an error if src doesn't exist or if the move fails.
//
// Example:
//
//	err := lxio.MoveFile("/path/to/old.txt", "/path/to/new.txt")
//	if err != nil {
//		log.Fatal(err)
//	}
func MoveFile(src, dst string) error {
	return os.Rename(src, dst)
}

// RemoveFile removes the file at the given path.
// Returns an error if the file cannot be removed.
// It returns an error if the file does not exist.
//
// Example:
//
//	err := lxio.RemoveFile("/path/to/file.txt")
//	if err != nil {
//		log.Fatal(err)
//	}
func RemoveFile(path string) error {
	return os.Remove(path)
}

// CreateDir creates a directory at the given path with the specified permissions.
// If parent directories do not exist, they are created automatically.
// It does not return an error if the directory already exists.
//
// Example:
//
//	err := lxio.CreateDir("/path/to/deeply/nested/dir", 0755)
//	if err != nil {
//		log.Fatal(err)
//	}
func CreateDir(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

// CreateFile creates a file at the given path with the specified permissions.
// If parent directories do not exist, they are created automatically.
// If the file already exists, it is truncated.
// Returns an error if creation fails.
//
// Example:
//
//	err := lxio.CreateFile("/path/to/deeply/nested/file.txt", 0644)
//	if err != nil {
//		log.Fatal(err)
//	}
func CreateFile(path string, perm os.FileMode) error {
	// Create parent directories if they don't exist
	dir := filepath.Dir(path)
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	// Create the file
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	f.Close()

	// Set permissions
	if err := os.Chmod(path, perm); err != nil {
		return err
	}

	return nil
}
