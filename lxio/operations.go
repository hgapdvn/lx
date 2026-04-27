package lxio

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"
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

// RemoveAll removes the file or directory at the given path and any children it contains.
// It removes everything recursively. It does not return an error if the path does not exist.
//
// Example:
//
//	err := lxio.RemoveAll("/path/to/dir")
//	if err != nil {
//		log.Fatal(err)
//	}
func RemoveAll(path string) error {
	return os.RemoveAll(path)
}

// RemoveIfExists removes the file or directory at the given path if it exists.
// It does not return an error if the path does not exist.
// If the path is a directory, it and all its contents are removed recursively.
//
// Example:
//
//	err := lxio.RemoveIfExists("/path/to/file.txt")
//	if err != nil {
//		log.Fatal(err)
//	}
func RemoveIfExists(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // Path doesn't exist, no error
		}
		return err // Some other error occurred
	}
	// Path exists, remove it
	return os.RemoveAll(path)
}

// CopyDir recursively copies the directory from src to dst.
// If dst already exists, it is not overwritten (error is returned).
// The directory structure and file permissions are preserved.
// Returns an error if src doesn't exist or if the copy fails.
//
// Example:
//
//	err := lxio.CopyDir("/path/to/source/dir", "/path/to/dest/dir")
//	if err != nil {
//		log.Fatal(err)
//	}
func CopyDir(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !srcInfo.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	// Check if destination already exists
	if _, err := os.Stat(dst); err == nil {
		return fmt.Errorf("destination %q already exists", dst)
	}

	if err := os.MkdirAll(dst, srcInfo.Mode().Perm()); err != nil {
		return err
	}

	return filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		target := filepath.Join(dst, rel)

		info, err := d.Info()
		if err != nil {
			return err
		}

		switch {
		case d.IsDir():
			return os.MkdirAll(target, info.Mode().Perm())

		case info.Mode()&os.ModeSymlink != 0:
			link, err := os.Readlink(path)
			if err != nil {
				return err
			}
			return os.Symlink(link, target)

		case info.Mode().IsRegular():
			return CopyFile(path, target)

		default:
			// skip unsupported types
			return nil
		}
	})
}

// Touch creates a file at the given path if it doesn't exist.
// If the file already exists, its access and modification times are updated to the current time.
// The file is created with default permissions (0644).
//
// Example:
//
//	err := lxio.Touch("/path/to/file.txt")
//	if err != nil {
//		log.Fatal(err)
//	}
func Touch(path string) error {
	// Try to open the file; if it doesn't exist, create it
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// Update the access and modification time to current time
	now := time.Now()
	return os.Chtimes(path, now, now)
}

// Rename renames (moves) the file or directory from oldpath to newpath.
// If newpath already exists, it will be replaced on Unix systems.
// This is an alias for os.Rename provided for code clarity.
//
// Example:
//
//	err := lxio.Rename("/path/to/old.txt", "/path/to/new.txt")
//	if err != nil {
//		log.Fatal(err)
//	}
func Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}
