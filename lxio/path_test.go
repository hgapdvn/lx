package lxio_test

import (
	"path/filepath"
	"testing"

	"github.com/hgapdvn/lx/lxio"
	"github.com/hgapdvn/lx/lxslices"
)

func TestJoinPath(t *testing.T) {
	tests := []struct {
		name     string
		parts    []string
		expected string
	}{
		{
			name:     "join multiple path parts",
			parts:    []string{"dir", "subdir", "file.txt"},
			expected: filepath.Join("dir", "subdir", "file.txt"),
		},
		{
			name:     "join single path part",
			parts:    []string{"file.txt"},
			expected: "file.txt",
		},
		{
			name:     "join empty parts returns empty",
			parts:    []string{},
			expected: "",
		},
		{
			name:     "join with trailing separators",
			parts:    []string{"dir", "subdir"},
			expected: filepath.Join("dir", "subdir"),
		},
		{
			name:     "join with absolute path",
			parts:    []string{"/home", "user", "file.txt"},
			expected: filepath.Join("/home", "user", "file.txt"),
		},
		{
			name:     "join with current directory",
			parts:    []string{".", "file.txt"},
			expected: filepath.Join(".", "file.txt"),
		},
		{
			name:     "join with parent directory",
			parts:    []string{"..", "file.txt"},
			expected: filepath.Join("..", "file.txt"),
		},
		{
			name:     "join many parts",
			parts:    []string{"a", "b", "c", "d", "e", "file.txt"},
			expected: filepath.Join("a", "b", "c", "d", "e", "file.txt"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxio.JoinPath(tt.parts...)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestBaseName(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "basename of file path",
			path:     "/path/to/file.txt",
			expected: "file.txt",
		},
		{
			name:     "basename of directory path",
			path:     "/path/to/dir",
			expected: "dir",
		},
		{
			name:     "basename of single file",
			path:     "file.txt",
			expected: "file.txt",
		},
		{
			name:     "basename of empty path",
			path:     "",
			expected: ".",
		},
		{
			name:     "basename with trailing slash",
			path:     "/path/to/dir/",
			expected: "dir",
		},
		{
			name:     "basename of dot",
			path:     ".",
			expected: ".",
		},
		{
			name:     "basename of parent directory",
			path:     "..",
			expected: "..",
		},
		{
			name:     "basename of hidden file",
			path:     "/path/to/.bashrc",
			expected: ".bashrc",
		},
		{
			name:     "basename with unicode",
			path:     "/path/to/文件.txt",
			expected: "文件.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxio.BaseName(tt.path)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestDirName(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
		check    func(result, expected string) bool
	}{
		{
			name:     "dirname of file path",
			path:     "/path/to/file.txt",
			expected: filepath.Join("/path/to"),
			check:    func(result, expected string) bool { return result == expected },
		},
		{
			name:     "dirname of directory path",
			path:     "/path/to/dir",
			expected: filepath.Join("/path/to"),
			check:    func(result, expected string) bool { return result == expected },
		},
		{
			name:     "dirname of single file",
			path:     "file.txt",
			expected: ".",
			check:    func(result, expected string) bool { return result == expected },
		},
		{
			name:     "dirname of empty path",
			path:     "",
			expected: ".",
			check:    func(result, expected string) bool { return result == expected },
		},
		{
			name:     "dirname with trailing slash",
			path:     "/path/to/dir/",
			expected: filepath.Join("/path/to"),
			check: func(result, expected string) bool {
				return result == expected || result == filepath.Dir("/path/to/dir/")
			},
		},
		{
			name:     "dirname of root",
			path:     "/",
			expected: "/",
			check:    func(result, expected string) bool { return result == expected || result == "\\" },
		},
		{
			name:     "dirname of dot",
			path:     ".",
			expected: ".",
			check:    func(result, expected string) bool { return result == expected },
		},
		{
			name:     "dirname of parent",
			path:     "..",
			expected: ".",
			check:    func(result, expected string) bool { return result == expected },
		},
		{
			name:     "dirname of deeply nested path",
			path:     "/a/b/c/d/e/file.txt",
			expected: filepath.Join("/a/b/c/d/e"),
			check:    func(result, expected string) bool { return result == expected },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxio.DirName(tt.path)
			if !tt.check(result, tt.expected) {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestExtension(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "extension of file with extension",
			path:     "/path/to/file.txt",
			expected: ".txt",
		},
		{
			name:     "extension of file without extension",
			path:     "/path/to/README",
			expected: "",
		},
		{
			name:     "extension of hidden file",
			path:     "/path/to/.bashrc",
			expected: ".bashrc",
		},
		{
			name:     "extension of tar.gz file",
			path:     "/path/to/archive.tar.gz",
			expected: ".gz",
		},
		{
			name:     "extension of file with multiple dots",
			path:     "file.backup.2024.bak",
			expected: ".bak",
		},
		{
			name:     "extension of file with numbers in extension",
			path:     "file.7z",
			expected: ".7z",
		},
		{
			name:     "extension of file with uppercase extension",
			path:     "FILE.TXT",
			expected: ".TXT",
		},
		{
			name:     "extension with unicode",
			path:     "文件.txt",
			expected: ".txt",
		},
		{
			name:     "extension of file with no path",
			path:     "file.go",
			expected: ".go",
		},
		{
			name:     "extension of directory with dot",
			path:     "/path/to/dir.bak",
			expected: ".bak",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxio.Extension(tt.path)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestWithoutExtension(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "remove extension from file",
			path:     "/path/to/file.txt",
			expected: "/path/to/file",
		},
		{
			name:     "path without extension unchanged",
			path:     "/path/to/README",
			expected: "/path/to/README",
		},
		{
			name:     "remove extension from hidden file",
			path:     "/path/to/.bashrc",
			expected: "/path/to/",
		},
		{
			name:     "remove extension from tar.gz file",
			path:     "/path/to/archive.tar.gz",
			expected: "/path/to/archive.tar",
		},
		{
			name:     "remove extension from file with multiple dots",
			path:     "file.backup.2024.bak",
			expected: "file.backup.2024",
		},
		{
			name:     "remove extension with single filename",
			path:     "file.go",
			expected: "file",
		},
		{
			name:     "remove uppercase extension",
			path:     "FILE.TXT",
			expected: "FILE",
		},
		{
			name:     "without extension with unicode",
			path:     "文件.txt",
			expected: "文件",
		},
		{
			name:     "without extension empty path",
			path:     "",
			expected: "",
		},
		{
			name:     "without extension preserves full path",
			path:     "/very/long/path/to/file.tar.bz2",
			expected: "/very/long/path/to/file.tar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxio.WithoutExtension(tt.path)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestPathOperationsCombined(t *testing.T) {
	tests := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{
			name: "combine path operations",
			fn: func(t *testing.T) {
				path := lxio.JoinPath("/home", "user", "documents", "file.txt")
				dir := lxio.DirName(path)
				base := lxio.BaseName(path)
				ext := lxio.Extension(base)
				withoutExt := lxio.WithoutExtension(base)

				if !lxslices.Equal([]string{dir, base, ext, withoutExt},
					[]string{filepath.Join("/home", "user", "documents"), "file.txt", ".txt", "file"}) {
					t.Errorf("combined operations failed")
				}
			},
		},
		{
			name: "reconstruct path from components",
			fn: func(t *testing.T) {
				original := lxio.JoinPath("/path", "to", "file.txt")
				dir := lxio.DirName(original)
				base := lxio.BaseName(original)
				reconstructed := lxio.JoinPath(dir, base)

				if reconstructed != original {
					t.Errorf("expected %q, got %q", original, reconstructed)
				}
			},
		},
		{
			name: "change extension",
			fn: func(t *testing.T) {
				original := "/path/to/file.txt"
				withoutExt := lxio.WithoutExtension(original)
				newPath := withoutExt + ".go"

				if newPath != "/path/to/file.go" {
					t.Errorf("expected '/path/to/file.go', got %q", newPath)
				}
			},
		},
		{
			name: "handle files in root directory",
			fn: func(t *testing.T) {
				path := "/file.txt"
				dir := lxio.DirName(path)
				base := lxio.BaseName(path)

				if (dir != "/" && dir != "\\") || base != "file.txt" {
					t.Errorf("expected dir='/' or '\\', base='file.txt', got dir=%q, base=%q", dir, base)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn(t)
		})
	}
}
