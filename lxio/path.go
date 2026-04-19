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

// Clean returns the shortest path name equivalent to path by purely lexical processing.
// It removes all `.` and `..` elements, multiple consecutive slashes are replaced by a single slash,
// and any trailing slashes are removed (except for the root path "/").
// If the result is an empty string, Clean returns ".".
// Clean(path) is equivalent to filepath.Clean(path).
//
// Example:
//
//	Clean("/path/to/../file.txt")      // "/path/file.txt"
//	Clean("/path//to///file.txt")      // "/path/to/file.txt"
//	Clean("./file.txt")                // "file.txt"
//	Clean("/path/to/dir/")             // "/path/to/dir"
//	Clean("../../file.txt")            // "../../file.txt"
func Clean(path string) string {
	return filepath.Clean(path)
}

// IsAbs reports whether the path is absolute.
// It returns true if the path begins with a root (e.g., "/" on Unix, "C:\" on Windows).
// The check is purely lexical; no filesystem access is performed.
// IsAbs(path) is equivalent to filepath.IsAbs(path).
//
// Example:
//
//	IsAbs("/path/to/file")   // true
//	IsAbs("./relative")      // false
//	IsAbs("file.txt")        // false
//	IsAbs("C:\\Windows")     // true (Windows)
func IsAbs(path string) bool {
	return filepath.IsAbs(path)
}

// IsRel reports whether the path is relative.
// It returns true if the path is not absolute (i.e., does not begin with a root).
// The check is purely lexical; no filesystem access is performed.
// IsRel(path) is equivalent to !IsAbs(path).
//
// Example:
//
//	IsRel("./file.txt")     // true
//	IsRel("file.txt")       // true
//	IsRel("/path/to/file")  // false
//	IsRel("C:\\Windows")    // false (Windows)
func IsRel(path string) bool {
	return !IsAbs(path)
}

// Rel returns a relative path from basepath to targpath.
// It computes a path relative to basepath that refers to the same location as targpath.
// If the paths cannot be made relative (e.g., on different drives on Windows), it returns an error.
// Rel(basepath, targpath) is equivalent to filepath.Rel(basepath, targpath).
//
// Example:
//
//	Rel("/home/user/docs", "/home/user/docs/file.txt")  // "file.txt"
//	Rel("/home/user/docs", "/home/user/file.txt")       // "../file.txt"
//	Rel("a/b", "a/b/c/d")                               // "c/d"
func Rel(basepath, targpath string) (string, error) {
	return filepath.Rel(basepath, targpath)
}

// SplitExtension splits a path into its base (without extension) and extension parts.
// It returns the path without extension and the extension separately in a single operation.
// More efficient than calling both Extension() and WithoutExtension() separately.
// The extension includes the leading dot; if there is no extension, ext is an empty string.
//
// Example:
//
//	SplitExtension("/path/to/file.txt")    // ("/path/to/file", ".txt")
//	SplitExtension("/path/to/file.tar.gz") // ("/path/to/file.tar", ".gz")
//	SplitExtension("README")               // ("README", "")
//	SplitExtension(".bashrc")              // ("/path/to/", ".bashrc")
//	SplitExtension("")                     // ("", "")
func SplitExtension(path string) (base, ext string) {
	ext = filepath.Ext(path)
	if ext == "" {
		return path, ""
	}
	return strings.TrimSuffix(path, ext), ext
}

// HasExtension reports whether the path has one of the given extensions.
// The comparison is case-sensitive. Extensions should include the leading dot (e.g., ".go", ".txt").
// Returns false if no extensions are provided or if the path has no extension.
// Returns true if the path's extension matches any of the provided extensions.
//
// Example:
//
//	HasExtension("file.txt", ".txt", ".md")       // true
//	HasExtension("file.go", ".js", ".ts")         // false
//	HasExtension("README", ".txt", ".md")         // false
//	HasExtension("archive.tar.gz", ".gz", ".zip") // true
//	HasExtension("file.txt")                      // false (no extensions provided)
func HasExtension(path string, exts ...string) bool {
	if len(exts) == 0 {
		return false
	}

	pathExt := Extension(path)
	if pathExt == "" {
		return false
	}

	for _, ext := range exts {
		if pathExt == ext {
			return true
		}
	}
	return false
}
