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

func TestAbs(t *testing.T) {
	tests := []struct {
		name        string
		path        string
		expectedErr bool
		expectAbs   bool
	}{
		// Success cases: All return absolute paths
		{
			name:        "absolute path unchanged",
			path:        "/path/to/file.txt",
			expectedErr: false,
			expectAbs:   true,
		},
		{
			name:        "relative path converted to absolute",
			path:        "file.txt",
			expectedErr: false,
			expectAbs:   true,
		},
		{
			name:        "relative path with dot",
			path:        "./file.txt",
			expectedErr: false,
			expectAbs:   true,
		},
		{
			name:        "relative path with dot slash",
			path:        "dir/subdir/file.txt",
			expectedErr: false,
			expectAbs:   true,
		},
		{
			name:        "parent directory reference",
			path:        "../file.txt",
			expectedErr: false,
			expectAbs:   true,
		},
		{
			name:        "current directory",
			path:        ".",
			expectedErr: false,
			expectAbs:   true,
		},
		{
			name:        "empty path returns current directory",
			path:        "",
			expectedErr: false,
			expectAbs:   true,
		},
		{
			name:        "path with unicode characters",
			path:        "文件/路径/file.txt",
			expectedErr: false,
			expectAbs:   true,
		},
		{
			name:        "path with spaces",
			path:        "dir with spaces/file name.txt",
			expectedErr: false,
			expectAbs:   true,
		},
		{
			name:        "deeply nested relative path",
			path:        "a/b/c/d/e/f/g/h/file.txt",
			expectedErr: false,
			expectAbs:   true,
		},
		// Note: expectAbs: false is not applicable for Abs() because:
		// - On success: Abs() always returns an absolute path (that's its purpose)
		// - On error: We don't validate the path format
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := lxio.Abs(tt.path)
			hasErr := err != nil
			if hasErr != tt.expectedErr {
				t.Errorf("Abs(%q) error expectation failed: expected error=%v, got error=%v", tt.path, tt.expectedErr, hasErr)
			}
			isAbs := filepath.IsAbs(result)
			if !hasErr && isAbs != tt.expectAbs {
				t.Errorf("Abs(%q) absolute path expectation failed: expected absolute=%v, got absolute=%v, result=%q", tt.path, tt.expectAbs, isAbs, result)
			}
		})
	}
}

func TestClean(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "empty path returns dot",
			path:     "",
			expected: ".",
		},
		{
			name:     "root path unchanged",
			path:     "/",
			expected: "/",
		},
		{
			name:     "absolute path with trailing slash removed",
			path:     "/path/to/dir/",
			expected: "/path/to/dir",
		},
		{
			name:     "removes dot current directory reference",
			path:     "./file.txt",
			expected: "file.txt",
		},
		{
			name:     "removes multiple consecutive slashes",
			path:     "/path//to///file.txt",
			expected: filepath.Join("/path/to/file.txt"),
		},
		{
			name:     "resolves parent directory reference",
			path:     "/path/to/../file.txt",
			expected: filepath.Join("/path/file.txt"),
		},
		{
			name:     "resolves multiple parent directory references",
			path:     "/path/to/../../file.txt",
			expected: filepath.Join("/file.txt"),
		},
		{
			name:     "relative path with dot",
			path:     "./subdir/file.txt",
			expected: filepath.Join("subdir/file.txt"),
		},
		{
			name:     "relative path with parent directory",
			path:     "dir/../file.txt",
			expected: "file.txt",
		},
		{
			name:     "leading parent directory references preserved",
			path:     "../../file.txt",
			expected: filepath.Join("../../file.txt"),
		},
		{
			name:     "complex path normalization",
			path:     "/path/./to/../to/file.txt",
			expected: filepath.Join("/path/to/file.txt"),
		},
		{
			name:     "path with only dots",
			path:     ".",
			expected: ".",
		},
		{
			name:     "path with only parent references",
			path:     "..",
			expected: "..",
		},
		{
			name:     "multiple slashes at start",
			path:     "///path/to/file.txt",
			expected: filepath.Join("/path/to/file.txt"),
		},
		{
			name:     "trailing slashes removed",
			path:     "/path/to/dir///",
			expected: filepath.Join("/path/to/dir"),
		},
		{
			name:     "unicode path normalized",
			path:     "/路径//to/文件.txt",
			expected: filepath.Join("/路径/to/文件.txt"),
		},
		{
			name:     "path with spaces normalized",
			path:     "/path/to //file name.txt",
			expected: "/path/to /file name.txt",
		},
		{
			name:     "deeply nested path normalized",
			path:     "/a/b/c/d/e/f/g/h/../../../file.txt",
			expected: filepath.Join("/a/b/c/d/e/file.txt"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxio.Clean(tt.path)
			if result != tt.expected {
				t.Errorf("Clean(%q) expected %q, got %q", tt.path, tt.expected, result)
			}
		})
	}
}

func TestIsAbs(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{
			name:     "absolute path with leading slash",
			path:     "/path/to/file.txt",
			expected: true,
		},
		{
			name:     "relative path with filename",
			path:     "file.txt",
			expected: false,
		},
		{
			name:     "relative path with current directory",
			path:     "./file.txt",
			expected: false,
		},
		{
			name:     "relative path with parent directory",
			path:     "../file.txt",
			expected: false,
		},
		{
			name:     "empty path",
			path:     "",
			expected: false,
		},
		{
			name:     "root path",
			path:     "/",
			expected: true,
		},
		{
			name:     "absolute path with trailing slash",
			path:     "/path/to/dir/",
			expected: true,
		},
		{
			name:     "relative path with multiple levels",
			path:     "dir/subdir/file.txt",
			expected: false,
		},
		{
			name:     "absolute path with multiple slashes",
			path:     "//path//to//file.txt",
			expected: true,
		},
		{
			name:     "dot only",
			path:     ".",
			expected: false,
		},
		{
			name:     "double dot only",
			path:     "..",
			expected: false,
		},
		{
			name:     "absolute path with unicode",
			path:     "/路径/到/文件.txt",
			expected: true,
		},
		{
			name:     "relative path with unicode",
			path:     "路径/到/文件.txt",
			expected: false,
		},
		{
			name:     "absolute path with spaces",
			path:     "/path with spaces/file.txt",
			expected: true,
		},
		{
			name:     "relative path with spaces",
			path:     "path with spaces/file.txt",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxio.IsAbs(tt.path)
			if result != tt.expected {
				t.Errorf("IsAbs(%q) expected %v, got %v", tt.path, tt.expected, result)
			}
		})
	}
}

func TestIsRel(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{
			name:     "relative path with filename",
			path:     "file.txt",
			expected: true,
		},
		{
			name:     "relative path with current directory",
			path:     "./file.txt",
			expected: true,
		},
		{
			name:     "relative path with parent directory",
			path:     "../file.txt",
			expected: true,
		},
		{
			name:     "relative path with multiple levels",
			path:     "dir/subdir/file.txt",
			expected: true,
		},
		{
			name:     "absolute path with leading slash",
			path:     "/path/to/file.txt",
			expected: false,
		},
		{
			name:     "root path",
			path:     "/",
			expected: false,
		},
		{
			name:     "empty path",
			path:     "",
			expected: true,
		},
		{
			name:     "dot only",
			path:     ".",
			expected: true,
		},
		{
			name:     "double dot only",
			path:     "..",
			expected: true,
		},
		{
			name:     "relative path with unicode",
			path:     "路径/到/文件.txt",
			expected: true,
		},
		{
			name:     "absolute path with unicode",
			path:     "/路径/到/文件.txt",
			expected: false,
		},
		{
			name:     "relative path with spaces",
			path:     "path with spaces/file.txt",
			expected: true,
		},
		{
			name:     "absolute path with spaces",
			path:     "/path with spaces/file.txt",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxio.IsRel(tt.path)
			if result != tt.expected {
				t.Errorf("IsRel(%q) expected %v, got %v", tt.path, tt.expected, result)
			}
		})
	}
}

func TestRel(t *testing.T) {
	tests := []struct {
		name        string
		basepath    string
		targpath    string
		expected    string
		expectedErr bool
	}{
		{
			name:        "same directory",
			basepath:    "/home/user",
			targpath:    "/home/user/file.txt",
			expected:    "file.txt",
			expectedErr: false,
		},
		{
			name:        "nested target",
			basepath:    "/home/user",
			targpath:    "/home/user/docs/file.txt",
			expected:    filepath.Join("docs/file.txt"),
			expectedErr: false,
		},
		{
			name:        "parent directory",
			basepath:    "/home/user/docs",
			targpath:    "/home/user/file.txt",
			expected:    filepath.Join("../file.txt"),
			expectedErr: false,
		},
		{
			name:        "sibling directories",
			basepath:    "/home/user/docs",
			targpath:    "/home/user/downloads/file.txt",
			expected:    filepath.Join("../downloads/file.txt"),
			expectedErr: false,
		},
		{
			name:        "relative to relative path",
			basepath:    "a/b",
			targpath:    "a/b/c/d",
			expected:    filepath.Join("c/d"),
			expectedErr: false,
		},
		{
			name:        "relative base with parent reference",
			basepath:    "a/b/c",
			targpath:    "a/b/file.txt",
			expected:    filepath.Join("../file.txt"),
			expectedErr: false,
		},
		{
			name:        "same path",
			basepath:    "/home/user/file.txt",
			targpath:    "/home/user/file.txt",
			expected:    ".",
			expectedErr: false,
		},
		{
			name:        "root to absolute path",
			basepath:    "/",
			targpath:    "/home/user/file.txt",
			expected:    filepath.Join("home/user/file.txt"),
			expectedErr: false,
		},
		{
			name:        "deep nesting",
			basepath:    "/a/b/c/d/e",
			targpath:    "/a/b/c/x/y/z",
			expected:    filepath.Join("../../x/y/z"),
			expectedErr: false,
		},
		{
			name:        "with unicode paths",
			basepath:    "/home/user/文件",
			targpath:    "/home/user/文件/file.txt",
			expected:    "file.txt",
			expectedErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := lxio.Rel(tt.basepath, tt.targpath)
			hasErr := err != nil
			if hasErr != tt.expectedErr {
				t.Errorf("Rel(%q, %q) error expectation failed: expected error=%v, got error=%v (%v)", tt.basepath, tt.targpath, tt.expectedErr, hasErr, err)
			}
			if !hasErr && result != tt.expected {
				t.Errorf("Rel(%q, %q) expected %q, got %q", tt.basepath, tt.targpath, tt.expected, result)
			}
		})
	}
}

func TestSplitExtension(t *testing.T) {
	tests := []struct {
		name         string
		path         string
		expectedBase string
		expectedExt  string
	}{
		{
			name:         "file with single extension",
			path:         "/path/to/file.txt",
			expectedBase: "/path/to/file",
			expectedExt:  ".txt",
		},
		{
			name:         "file with no extension",
			path:         "/path/to/README",
			expectedBase: "/path/to/README",
			expectedExt:  "",
		},
		{
			name:         "hidden file with extension",
			path:         "/path/to/.bashrc",
			expectedBase: "/path/to/",
			expectedExt:  ".bashrc",
		},
		{
			name:         "file with double extension",
			path:         "/path/to/archive.tar.gz",
			expectedBase: "/path/to/archive.tar",
			expectedExt:  ".gz",
		},
		{
			name:         "file with multiple dots",
			path:         "file.backup.2024.bak",
			expectedBase: "file.backup.2024",
			expectedExt:  ".bak",
		},
		{
			name:         "file with numeric extension",
			path:         "file.7z",
			expectedBase: "file",
			expectedExt:  ".7z",
		},
		{
			name:         "uppercase extension",
			path:         "FILE.TXT",
			expectedBase: "FILE",
			expectedExt:  ".TXT",
		},
		{
			name:         "file in current directory with extension",
			path:         "file.go",
			expectedBase: "file",
			expectedExt:  ".go",
		},
		{
			name:         "empty path",
			path:         "",
			expectedBase: "",
			expectedExt:  "",
		},
		{
			name:         "path with only dot",
			path:         ".",
			expectedBase: "",
			expectedExt:  ".",
		},
		{
			name:         "path with only double dot",
			path:         "..",
			expectedBase: ".",
			expectedExt:  ".",
		},
		{
			name:         "unicode filename with extension",
			path:         "/path/to/文件.txt",
			expectedBase: "/path/to/文件",
			expectedExt:  ".txt",
		},
		{
			name:         "unicode filename without extension",
			path:         "/path/to/文件",
			expectedBase: "/path/to/文件",
			expectedExt:  "",
		},
		{
			name:         "filename with spaces and extension",
			path:         "/path/to/my file.txt",
			expectedBase: "/path/to/my file",
			expectedExt:  ".txt",
		},
		{
			name:         "filename with spaces no extension",
			path:         "/path/to/my file",
			expectedBase: "/path/to/my file",
			expectedExt:  "",
		},
		{
			name:         "deeply nested path with extension",
			path:         "/a/b/c/d/e/f/file.tar.bz2",
			expectedBase: "/a/b/c/d/e/f/file.tar",
			expectedExt:  ".bz2",
		},
		{
			name:         "relative path with extension",
			path:         "./subdir/file.txt",
			expectedBase: "./subdir/file",
			expectedExt:  ".txt",
		},
		{
			name:         "parent reference path with extension",
			path:         "../file.txt",
			expectedBase: "../file",
			expectedExt:  ".txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			base, ext := lxio.SplitExtension(tt.path)
			if base != tt.expectedBase {
				t.Errorf("SplitExtension(%q) base expected %q, got %q", tt.path, tt.expectedBase, base)
			}
			if ext != tt.expectedExt {
				t.Errorf("SplitExtension(%q) ext expected %q, got %q", tt.path, tt.expectedExt, ext)
			}
		})
	}
}

func TestHasExtension(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		exts     []string
		expected bool
	}{
		{
			name:     "single matching extension",
			path:     "file.txt",
			exts:     []string{".txt"},
			expected: true,
		},
		{
			name:     "multiple extensions with match",
			path:     "file.txt",
			exts:     []string{".txt", ".md", ".go"},
			expected: true,
		},
		{
			name:     "multiple extensions no match",
			path:     "file.go",
			exts:     []string{".txt", ".md", ".js"},
			expected: false,
		},
		{
			name:     "no extensions provided",
			path:     "file.txt",
			exts:     []string{},
			expected: false,
		},
		{
			name:     "file with no extension",
			path:     "README",
			exts:     []string{".txt", ".md"},
			expected: false,
		},
		{
			name:     "double extension match last part",
			path:     "archive.tar.gz",
			exts:     []string{".gz", ".zip"},
			expected: true,
		},
		{
			name:     "double extension no match",
			path:     "archive.tar.gz",
			exts:     []string{".tar", ".zip"},
			expected: false,
		},
		{
			name:     "case sensitive match",
			path:     "FILE.TXT",
			exts:     []string{".txt"},
			expected: false,
		},
		{
			name:     "case sensitive uppercase extension",
			path:     "FILE.TXT",
			exts:     []string{".TXT"},
			expected: true,
		},
		{
			name:     "hidden file match",
			path:     ".bashrc",
			exts:     []string{".bashrc"},
			expected: true,
		},
		{
			name:     "hidden file no match",
			path:     ".bashrc",
			exts:     []string{".txt"},
			expected: false,
		},
		{
			name:     "numeric extension match",
			path:     "file.7z",
			exts:     []string{".7z", ".zip"},
			expected: true,
		},
		{
			name:     "with full path",
			path:     "/path/to/file.txt",
			exts:     []string{".txt"},
			expected: true,
		},
		{
			name:     "with relative path",
			path:     "./dir/file.go",
			exts:     []string{".go", ".js"},
			expected: true,
		},
		{
			name:     "unicode extension",
			path:     "文件.txt",
			exts:     []string{".txt"},
			expected: true,
		},
		{
			name:     "path with spaces",
			path:     "/path/to/my file.txt",
			exts:     []string{".txt"},
			expected: true,
		},
		{
			name:     "empty path",
			path:     "",
			exts:     []string{".txt"},
			expected: false,
		},
		{
			name:     "multiple matches returns true on first match",
			path:     "file.md",
			exts:     []string{".txt", ".md", ".go", ".js"},
			expected: true,
		},
		{
			name:     "match at end of list",
			path:     "file.go",
			exts:     []string{".txt", ".md", ".js", ".go"},
			expected: true,
		},
		{
			name:     "many extensions no match",
			path:     "file.rs",
			exts:     []string{".go", ".js", ".ts", ".py", ".java", ".cpp"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxio.HasExtension(tt.path, tt.exts...)
			if result != tt.expected {
				t.Errorf("HasExtension(%q, %v) expected %v, got %v", tt.path, tt.exts, tt.expected, result)
			}
		})
	}
}
