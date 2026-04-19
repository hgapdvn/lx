package lxio

import (
	"path/filepath"
	"strings"
)

// JoinPath joins path elements into a single path.
// It handles the platform-specific path separators correctly.
// The returned path is empty if no arguments are provided.
//
// Example:
//
//	JoinPath("dir", "subdir", "file.txt")
//	// Output: "dir/subdir/file.txt" (Unix) or "dir\subdir\file.txt" (Windows)
func JoinPath(parts ...string) string {
	if len(parts) == 0 {
		return ""
	}
	return filepath.Join(parts...)
}

// BaseName returns the last element of the path.
// If the path is empty, BaseName returns ".".
// BaseName(path) is equivalent to filepath.Base(path).
//
// Example:
//
//	BaseName("/path/to/file.txt") // "file.txt"
//	BaseName("/path/to/dir/")     // "dir"
//	BaseName("file.txt")          // "file.txt"
//	BaseName("")                  // "."
func BaseName(path string) string {
	return filepath.Base(path)
}

// DirName returns all but the last element of the path, typically the path's directory.
// If the path is empty, DirName returns ".".
// DirName(path) is equivalent to filepath.Dir(path).
//
// Example:
//
//	DirName("/path/to/file.txt") // "/path/to"
//	DirName("/path/to/dir/")     // "/path/to"
//	DirName("file.txt")          // "."
//	DirName("")                  // "."
func DirName(path string) string {
	return filepath.Dir(path)
}

// Extension returns the file extension of the path.
// The extension is the suffix starting with the final dot.
// If there is no extension, Extension returns an empty string.
// Extension(path) is equivalent to filepath.Ext(path).
//
// Example:
//
//	Extension("/path/to/file.txt")    // ".txt"
//	Extension("/path/to/file.tar.gz") // ".gz"
//	Extension("README")               // ""
//	Extension(".bashrc")              // ".bashrc"
func Extension(path string) string {
	return filepath.Ext(path)
}

// WithoutExtension returns the path without its file extension.
// If the path has no extension, it is returned unchanged.
//
// Example:
//
//	WithoutExtension("/path/to/file.txt")    // "/path/to/file"
//	WithoutExtension("/path/to/file.tar.gz") // "/path/to/file.tar"
//	WithoutExtension("README")               // "README"
//	WithoutExtension(".bashrc")              // ".bashrc"
func WithoutExtension(path string) string {
	ext := filepath.Ext(path)
	if ext == "" {
		return path
	}
	return strings.TrimSuffix(path, ext)
}

// Abs returns an absolute representation of the path.
// If the path is not absolute, it is joined with the current working directory.
// If there is an error retrieving the working directory, the error is returned.
// Abs(path) is equivalent to filepath.Abs(path).
//
// Example:
//
//	Abs("file.txt")     // "/current/working/directory/file.txt" (Unix)
//	Abs("/path/to/file") // "/path/to/file"
//	Abs("./dir/file")   // "/current/working/directory/dir/file" (Unix)
func Abs(path string) (string, error) {
	return filepath.Abs(path)
}
