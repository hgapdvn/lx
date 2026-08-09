package lxio

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Sentinel errors returned by file and directory copy operations.
var (
	// ErrSourceAndDestinationSame is returned when CopyFile receives paths to
	// the same underlying file.
	ErrSourceAndDestinationSame = errors.New("lxio: source and destination are the same file")
	// ErrSourceNotDirectory is returned when the source path is not a directory.
	ErrSourceNotDirectory = errors.New("lxio: source path is not a directory")
	// ErrDestinationExists is returned when the destination path already exists.
	ErrDestinationExists = errors.New("lxio: destination already exists")
	// ErrDestinationWithinSource is returned when CopyDir receives a destination
	// inside the source directory.
	ErrDestinationWithinSource = errors.New("lxio: destination is within source directory")
)

// CopyFile copies the file from src to dst.
// If dst already exists, it is truncated.
// The file permissions and modification time are not preserved.
// Returns ErrSourceAndDestinationSame if src and dst refer to the same file.
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
	defer func() { _ = source.Close() }()

	sourceInfo, err := source.Stat()
	if err != nil {
		return err
	}
	destinationInfo, err := os.Stat(dst)
	if err == nil && os.SameFile(sourceInfo, destinationInfo) {
		return fmt.Errorf("%w: %q and %q", ErrSourceAndDestinationSame, src, dst)
	}
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	// Create destination file
	destination, err := os.Create(dst)
	if err != nil {
		return err
	}

	// Copy contents; close explicitly (not via defer) to capture flush errors
	// (e.g., on NFS or other remote filesystems).
	_, copyErr := io.Copy(destination, source)
	closeErr := destination.Close()

	if copyErr != nil {
		return copyErr
	}
	return closeErr
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
	return Rename(src, dst)
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

	// Create (or truncate) the file with the requested permissions in a single
	// syscall to avoid the TOCTOU race between os.Create and os.Chmod.
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	return f.Close()
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
// Returns ErrDestinationWithinSource if dst is inside src.
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
		return fmt.Errorf("%w: %q", ErrSourceNotDirectory, src)
	}

	// Check if destination already exists
	if _, err := os.Stat(dst); err == nil {
		return fmt.Errorf("%w: %q", ErrDestinationExists, dst)
	} else if !os.IsNotExist(err) {
		return err
	}

	srcAbsPath, err := filepath.Abs(src)
	if err != nil {
		return err
	}
	srcPath, err := filepath.EvalSymlinks(srcAbsPath)
	if err != nil {
		return err
	}
	dstPath, err := resolveDestinationPath(dst)
	if err != nil {
		return err
	}
	inside, err := isWithinDirectory(srcPath, dstPath)
	if err != nil {
		return err
	}
	if inside {
		return fmt.Errorf("%w: %q", ErrDestinationWithinSource, dst)
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
			if err := os.MkdirAll(target, info.Mode().Perm()); err != nil {
				return err
			}
			return os.Chmod(target, info.Mode().Perm())

		case info.Mode()&os.ModeSymlink != 0:
			link, err := os.Readlink(path)
			if err != nil {
				return err
			}
			return os.Symlink(link, target)

		case info.Mode().IsRegular():
			if err := CopyFile(path, target); err != nil {
				return err
			}
			return os.Chmod(target, info.Mode().Perm())

		default:
			// skip unsupported types
			return nil
		}
	})
}

// resolveDestinationPath resolves symlinks in the closest existing ancestor of
// path, then appends the still-missing path components.
func resolveDestinationPath(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	current := absPath
	var suffix []string
	for {
		resolved, err := filepath.EvalSymlinks(current)
		if err == nil {
			return filepath.Join(append([]string{resolved}, suffix...)...), nil
		}
		if !os.IsNotExist(err) {
			return "", err
		}

		parent := filepath.Dir(current)
		if parent == current {
			return "", err
		}
		suffix = append([]string{filepath.Base(current)}, suffix...)
		current = parent
	}
}

func isWithinDirectory(parent, path string) (bool, error) {
	rel, err := filepath.Rel(parent, path)
	if err != nil {
		return false, err
	}
	return rel == "." || (rel != ".." && !strings.HasPrefix(rel, ".."+string(filepath.Separator))), nil
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
	defer func() { _ = f.Close() }()

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
