package lxio_test

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/hgapdvn/lx/lxio"
)

// setupTestEnvironment creates a temporary folder with a file, a directory,
// and a symlink to use in our tests. It returns the paths to these items.
func setupTestEnvironment(t *testing.T) (testDir, testFile, testSymlink, testMissing string) {
	// Create a temporary directory that cleans itself up after the test
	testDir = t.TempDir()

	// 1. Create a regular file
	testFile = filepath.Join(testDir, "regular_file.txt")
	err := os.WriteFile(testFile, []byte("hello world"), 0644)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	// 2. Create a sub-directory
	testSubDir := filepath.Join(testDir, "sub_dir")
	err = os.Mkdir(testSubDir, 0755)
	if err != nil {
		t.Fatalf("failed to create test directory: %v", err)
	}

	// 3. Create a symlink pointing to the regular file
	testSymlink = filepath.Join(testDir, "symlink_file.txt")
	err = os.Symlink(testFile, testSymlink)
	if err != nil {
		// Symlinks can sometimes fail on Windows without admin rights.
		t.Logf("Warning: failed to create symlink (normal on Windows without elevated privileges): %v", err)
	}

	// 4. Define a path that definitely doesn't exist
	testMissing = filepath.Join(testDir, "does_not_exist.md")

	return testSubDir, testFile, testSymlink, testMissing
}

// ======================================== Exists Tests ========================================

func TestExists(t *testing.T) {
	testDir, testFile, testSymlink, missingPath := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	tests := []struct {
		name              string
		pathSetup         func() string
		skipOnWindows     bool
		expected          bool
		expectedErr       bool
		shouldHavePermErr bool
	}{
		{
			name:        "file exists",
			pathSetup:   func() string { return testFile },
			expected:    true,
			expectedErr: false,
		},
		{
			name:        "directory exists",
			pathSetup:   func() string { return testDir },
			expected:    true,
			expectedErr: false,
		},
		{
			name:        "path does not exist",
			pathSetup:   func() string { return missingPath },
			expected:    false,
			expectedErr: false,
		},
		{
			name:        "symlink exists",
			pathSetup:   func() string { return testSymlink },
			expected:    true,
			expectedErr: false,
		},
		{
			name:              "permission error",
			skipOnWindows:     true,
			expected:          false,
			expectedErr:       true,
			shouldHavePermErr: true,
			pathSetup: func() string {
				secureDir := t.TempDir()
				secretFile := filepath.Join(secureDir, "secret.txt")
				if err := os.WriteFile(secretFile, []byte("secret"), 0644); err != nil {
					t.Fatalf("failed to create secret file: %v", err)
				}
				if err := os.Chmod(secureDir, 0000); err != nil {
					t.Fatalf("failed to change permissions: %v", err)
				}
				t.Cleanup(func() {
					_ = os.Chmod(secureDir, 0755)
				})
				return secretFile
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "symlink exists" && runtime.GOOS == "windows" {
				if _, err := os.Lstat(testSymlink); err != nil {
					t.Skip("Symlink not created (normal on Windows)")
				}
			}

			if tt.skipOnWindows && runtime.GOOS == "windows" {
				t.Skip("Skipping permission tests on Windows")
			}

			path := tt.pathSetup()
			result, err := lxio.Exists(path)
			hasErr := err != nil

			if hasErr != tt.expectedErr {
				t.Errorf("Exists(%q) error expectation failed: expected error=%v, got error=%v (%v)", path, tt.expectedErr, hasErr, err)
			}

			if tt.shouldHavePermErr && (err == nil || !errors.Is(err, os.ErrPermission)) {
				t.Errorf("Exists(%q) expected permission error, got: %v", path, err)
			}

			if !hasErr && result != tt.expected {
				t.Errorf("Exists(%q) expected %v, got %v", path, tt.expected, result)
			}
		})
	}
}

func TestExistsOK(t *testing.T) {
	testDir, testFile, _, missingPath := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	tests := []struct {
		name          string
		pathSetup     func() string
		skipOnWindows bool
		expected      bool
	}{
		{
			name:      "file exists",
			pathSetup: func() string { return testFile },
			expected:  true,
		},
		{
			name:      "path does not exist",
			pathSetup: func() string { return missingPath },
			expected:  false,
		},
		{
			name:          "permission error is swallowed",
			skipOnWindows: true,
			expected:      false,
			pathSetup: func() string {
				secureDir := t.TempDir()
				secretFile := filepath.Join(secureDir, "secret.txt")
				if err := os.WriteFile(secretFile, []byte("secret"), 0644); err != nil {
					t.Fatalf("failed to create secret file: %v", err)
				}
				if err := os.Chmod(secureDir, 0000); err != nil {
					t.Fatalf("failed to change permissions: %v", err)
				}
				t.Cleanup(func() {
					_ = os.Chmod(secureDir, 0755)
				})
				return secretFile
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skipOnWindows && runtime.GOOS == "windows" {
				t.Skip("Skipping permission tests on Windows")
			}

			path := tt.pathSetup()
			result := lxio.ExistsOK(path)
			if result != tt.expected {
				t.Errorf("ExistsOK(%q) expected %v, got %v", path, tt.expected, result)
			}
		})
	}
}

func TestMustExist(t *testing.T) {
	testDir, testFile, _, _ := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	tests := []struct {
		name          string
		pathSetup     func() string
		skipOnWindows bool
		expected      bool
		shouldPanic   bool
	}{
		{
			name:        "file exists",
			pathSetup:   func() string { return testFile },
			expected:    true,
			shouldPanic: false,
		},
		{
			name:        "path does not exist",
			pathSetup:   func() string { return filepath.Join(testDir, "nonexistent.txt") },
			expected:    false,
			shouldPanic: false,
		},
		{
			name:          "panics on permission error",
			skipOnWindows: true,
			shouldPanic:   true,
			pathSetup: func() string {
				secureDir := t.TempDir()
				secretFile := filepath.Join(secureDir, "secret.txt")
				if err := os.WriteFile(secretFile, []byte("secret"), 0644); err != nil {
					t.Fatalf("failed to create secret file: %v", err)
				}
				if err := os.Chmod(secureDir, 0000); err != nil {
					t.Fatalf("failed to change permissions: %v", err)
				}
				t.Cleanup(func() {
					_ = os.Chmod(secureDir, 0755)
				})
				return secretFile
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skipOnWindows && runtime.GOOS == "windows" {
				t.Skip("Skipping permission tests on Windows")
			}

			path := tt.pathSetup()

			if tt.shouldPanic {
				defer func() {
					if recover() == nil {
						t.Errorf("MustExist(%q) should panic", path)
					}
				}()
			}

			result := lxio.MustExist(path)
			if !tt.shouldPanic && result != tt.expected {
				t.Errorf("MustExist(%q) expected %v, got %v", path, tt.expected, result)
			}
		})
	}
}

func TestNotExists(t *testing.T) {
	testDir, testFile, _, missingPath := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	tests := []struct {
		name              string
		pathSetup         func() string
		skipOnWindows     bool
		expected          bool
		expectedErr       bool
		shouldHavePermErr bool
	}{
		{
			name:        "file exists",
			pathSetup:   func() string { return testFile },
			expected:    false,
			expectedErr: false,
		},
		{
			name:        "path does not exist",
			pathSetup:   func() string { return missingPath },
			expected:    true,
			expectedErr: false,
		},
		{
			name:              "permission error",
			skipOnWindows:     true,
			expected:          false,
			expectedErr:       true,
			shouldHavePermErr: true,
			pathSetup: func() string {
				secureDir := t.TempDir()
				secretFile := filepath.Join(secureDir, "secret.txt")
				if err := os.WriteFile(secretFile, []byte("secret"), 0644); err != nil {
					t.Fatalf("failed to create secret file: %v", err)
				}
				if err := os.Chmod(secureDir, 0000); err != nil {
					t.Fatalf("failed to change permissions: %v", err)
				}
				t.Cleanup(func() {
					_ = os.Chmod(secureDir, 0755)
				})
				return secretFile
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skipOnWindows && runtime.GOOS == "windows" {
				t.Skip("Skipping permission tests on Windows")
			}

			path := tt.pathSetup()
			result, err := lxio.NotExists(path)
			hasErr := err != nil

			if hasErr != tt.expectedErr {
				t.Errorf("NotExists(%q) error expectation failed: expected error=%v, got error=%v (%v)", path, tt.expectedErr, hasErr, err)
			}

			if tt.shouldHavePermErr && (err == nil || !errors.Is(err, os.ErrPermission)) {
				t.Errorf("NotExists(%q) expected permission error, got: %v", path, err)
			}

			if !hasErr && result != tt.expected {
				t.Errorf("NotExists(%q) expected %v, got %v", path, tt.expected, result)
			}
		})
	}
}

func TestNotExistsOK(t *testing.T) {
	testDir, testFile, _, missingPath := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	tests := []struct {
		name          string
		pathSetup     func() string
		skipOnWindows bool
		expected      bool
	}{
		{
			name:      "file exists",
			pathSetup: func() string { return testFile },
			expected:  false,
		},
		{
			name:      "path does not exist",
			pathSetup: func() string { return missingPath },
			expected:  true,
		},
		{
			name:          "permission error is swallowed",
			skipOnWindows: true,
			expected:      false,
			pathSetup: func() string {
				secureDir := t.TempDir()
				secretFile := filepath.Join(secureDir, "secret.txt")
				if err := os.WriteFile(secretFile, []byte("secret"), 0644); err != nil {
					t.Fatalf("failed to create secret file: %v", err)
				}
				if err := os.Chmod(secureDir, 0000); err != nil {
					t.Fatalf("failed to change permissions: %v", err)
				}
				t.Cleanup(func() {
					_ = os.Chmod(secureDir, 0755)
				})
				return secretFile
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skipOnWindows && runtime.GOOS == "windows" {
				t.Skip("Skipping permission tests on Windows")
			}

			path := tt.pathSetup()
			result := lxio.NotExistsOK(path)
			if result != tt.expected {
				t.Errorf("NotExistsOK(%q) expected %v, got %v", path, tt.expected, result)
			}
		})
	}
}

func TestMustNotExist(t *testing.T) {
	testDir, testFile, _, missingPath := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	tests := []struct {
		name          string
		pathSetup     func() string
		skipOnWindows bool
		expected      bool
		shouldPanic   bool
	}{
		{
			name:        "path does not exist",
			pathSetup:   func() string { return missingPath },
			expected:    true,
			shouldPanic: false,
		},
		{
			name:        "file exists",
			pathSetup:   func() string { return testFile },
			expected:    false,
			shouldPanic: false,
		},
		{
			name:        "directory exists",
			pathSetup:   func() string { return testDir },
			expected:    false,
			shouldPanic: false,
		},
		{
			name:          "panics on permission error",
			skipOnWindows: true,
			shouldPanic:   true,
			pathSetup: func() string {
				secureDir := t.TempDir()
				secretFile := filepath.Join(secureDir, "secret.txt")
				if err := os.WriteFile(secretFile, []byte("secret"), 0644); err != nil {
					t.Fatalf("failed to create secret file: %v", err)
				}
				if err := os.Chmod(secureDir, 0000); err != nil {
					t.Fatalf("failed to change permissions: %v", err)
				}
				t.Cleanup(func() {
					_ = os.Chmod(secureDir, 0755)
				})
				return secretFile
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skipOnWindows && runtime.GOOS == "windows" {
				t.Skip("Skipping permission tests on Windows")
			}

			path := tt.pathSetup()

			if tt.shouldPanic {
				defer func() {
					if recover() == nil {
						t.Errorf("MustNotExist(%q) should panic", path)
					}
				}()
			}

			result := lxio.MustNotExist(path)
			if !tt.shouldPanic && result != tt.expected {
				t.Errorf("MustNotExist(%q) expected %v, got %v", path, tt.expected, result)
			}
		})
	}
}

// ======================================== IsDir Tests ========================================

func TestIsDir(t *testing.T) {
	testDir, testFile, _, missingPath := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	tests := []struct {
		name              string
		pathSetup         func() string
		skipOnWindows     bool
		expected          bool
		expectedErr       bool
		shouldHavePermErr bool
	}{
		{
			name:        "directory returns true",
			pathSetup:   func() string { return testDir },
			expected:    true,
			expectedErr: false,
		},
		{
			name:        "file returns false",
			pathSetup:   func() string { return testFile },
			expected:    false,
			expectedErr: false,
		},
		{
			name:        "missing path returns false",
			pathSetup:   func() string { return missingPath },
			expected:    false,
			expectedErr: false,
		},
		{
			name:              "permission error",
			skipOnWindows:     true,
			expected:          false,
			expectedErr:       true,
			shouldHavePermErr: true,
			pathSetup: func() string {
				secureDir := t.TempDir()
				secretFile := filepath.Join(secureDir, "secret.txt")
				if err := os.WriteFile(secretFile, []byte("secret"), 0644); err != nil {
					t.Fatalf("failed to create secret file: %v", err)
				}
				if err := os.Chmod(secureDir, 0000); err != nil {
					t.Fatalf("failed to change permissions: %v", err)
				}
				t.Cleanup(func() {
					_ = os.Chmod(secureDir, 0755)
				})
				return secretFile
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skipOnWindows && runtime.GOOS == "windows" {
				t.Skip("Skipping permission tests on Windows")
			}

			path := tt.pathSetup()
			result, err := lxio.IsDir(path)
			hasErr := err != nil

			if hasErr != tt.expectedErr {
				t.Errorf("IsDir(%q) error expectation failed: expected error=%v, got error=%v (%v)", path, tt.expectedErr, hasErr, err)
			}

			if tt.shouldHavePermErr && (err == nil || !errors.Is(err, os.ErrPermission)) {
				t.Errorf("IsDir(%q) expected permission error, got: %v", path, err)
			}

			if !hasErr && result != tt.expected {
				t.Errorf("IsDir(%q) expected %v, got %v", path, tt.expected, result)
			}
		})
	}
}

func TestIsDirOK(t *testing.T) {
	testDir, testFile, _, _ := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	tests := []struct {
		name          string
		pathSetup     func() string
		skipOnWindows bool
		expected      bool
	}{
		{
			name:      "directory returns true",
			pathSetup: func() string { return testDir },
			expected:  true,
		},
		{
			name:      "file returns false",
			pathSetup: func() string { return testFile },
			expected:  false,
		},
		{
			name:          "permission error is swallowed",
			skipOnWindows: true,
			expected:      false,
			pathSetup: func() string {
				secureDir := t.TempDir()
				secretFile := filepath.Join(secureDir, "secret.txt")
				if err := os.WriteFile(secretFile, []byte("secret"), 0644); err != nil {
					t.Fatalf("failed to create secret file: %v", err)
				}
				if err := os.Chmod(secureDir, 0000); err != nil {
					t.Fatalf("failed to change permissions: %v", err)
				}
				t.Cleanup(func() {
					_ = os.Chmod(secureDir, 0755)
				})
				return secretFile
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skipOnWindows && runtime.GOOS == "windows" {
				t.Skip("Skipping permission tests on Windows")
			}

			path := tt.pathSetup()
			result := lxio.IsDirOK(path)
			if result != tt.expected {
				t.Errorf("IsDirOK(%q) expected %v, got %v", path, tt.expected, result)
			}
		})
	}
}

func TestMustBeDir(t *testing.T) {
	testDir, testFile, _, missingPath := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	tests := []struct {
		name          string
		pathSetup     func() string
		skipOnWindows bool
		expected      bool
		shouldPanic   bool
	}{
		{
			name:        "directory returns true",
			pathSetup:   func() string { return testDir },
			expected:    true,
			shouldPanic: false,
		},
		{
			name:        "file returns false",
			pathSetup:   func() string { return testFile },
			expected:    false,
			shouldPanic: false,
		},
		{
			name:        "missing path returns false",
			pathSetup:   func() string { return missingPath },
			expected:    false,
			shouldPanic: false,
		},
		{
			name:          "panics on permission error",
			skipOnWindows: true,
			shouldPanic:   true,
			pathSetup: func() string {
				secureDir := t.TempDir()
				secretFile := filepath.Join(secureDir, "secret.txt")
				if err := os.WriteFile(secretFile, []byte("secret"), 0644); err != nil {
					t.Fatalf("failed to create secret file: %v", err)
				}
				if err := os.Chmod(secureDir, 0000); err != nil {
					t.Fatalf("failed to change permissions: %v", err)
				}
				t.Cleanup(func() {
					_ = os.Chmod(secureDir, 0755)
				})
				return secretFile
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skipOnWindows && runtime.GOOS == "windows" {
				t.Skip("Skipping permission tests on Windows")
			}

			path := tt.pathSetup()

			if tt.shouldPanic {
				defer func() {
					if recover() == nil {
						t.Errorf("MustBeDir(%q) should panic", path)
					}
				}()
			}

			result := lxio.MustBeDir(path)
			if !tt.shouldPanic && result != tt.expected {
				t.Errorf("MustBeDir(%q) expected %v, got %v", path, tt.expected, result)
			}
		})
	}
}

// ======================================== IsFile Tests ========================================

func TestIsFile(t *testing.T) {
	testDir, testFile, testSymlink, missingPath := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	tests := []struct {
		name              string
		pathSetup         func() string
		skipOnWindows     bool
		expected          bool
		expectedErr       bool
		shouldHavePermErr bool
	}{
		{
			name:        "file returns true",
			pathSetup:   func() string { return testFile },
			expected:    true,
			expectedErr: false,
		},
		{
			name:        "directory returns false",
			pathSetup:   func() string { return testDir },
			expected:    false,
			expectedErr: false,
		},
		{
			name:        "missing path returns false",
			pathSetup:   func() string { return missingPath },
			expected:    false,
			expectedErr: false,
		},
		{
			name:        "symlink returns true",
			pathSetup:   func() string { return testSymlink },
			expected:    true,
			expectedErr: false,
		},
		{
			name:              "permission error",
			skipOnWindows:     true,
			expected:          false,
			expectedErr:       true,
			shouldHavePermErr: true,
			pathSetup: func() string {
				secureDir := t.TempDir()
				secretFile := filepath.Join(secureDir, "secret.txt")
				if err := os.WriteFile(secretFile, []byte("secret"), 0644); err != nil {
					t.Fatalf("failed to create secret file: %v", err)
				}
				if err := os.Chmod(secureDir, 0000); err != nil {
					t.Fatalf("failed to change permissions: %v", err)
				}
				t.Cleanup(func() {
					_ = os.Chmod(secureDir, 0755)
				})
				return secretFile
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "symlink returns true" && runtime.GOOS == "windows" {
				if _, err := os.Lstat(testSymlink); err != nil {
					t.Skip("Symlink not created (normal on Windows)")
				}
			}

			if tt.skipOnWindows && runtime.GOOS == "windows" {
				t.Skip("Skipping permission tests on Windows")
			}

			path := tt.pathSetup()
			result, err := lxio.IsFile(path)
			hasErr := err != nil

			if hasErr != tt.expectedErr {
				t.Errorf("IsFile(%q) error expectation failed: expected error=%v, got error=%v (%v)", path, tt.expectedErr, hasErr, err)
			}

			if tt.shouldHavePermErr && (err == nil || !errors.Is(err, os.ErrPermission)) {
				t.Errorf("IsFile(%q) expected permission error, got: %v", path, err)
			}

			if !hasErr && result != tt.expected {
				t.Errorf("IsFile(%q) expected %v, got %v", path, tt.expected, result)
			}
		})
	}
}

func TestIsFileOK(t *testing.T) {
	testDir, testFile, _, _ := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	tests := []struct {
		name          string
		pathSetup     func() string
		skipOnWindows bool
		expected      bool
	}{
		{
			name:      "file returns true",
			pathSetup: func() string { return testFile },
			expected:  true,
		},
		{
			name:      "directory returns false",
			pathSetup: func() string { return testDir },
			expected:  false,
		},
		{
			name:          "permission error is swallowed",
			skipOnWindows: true,
			expected:      false,
			pathSetup: func() string {
				secureDir := t.TempDir()
				secretFile := filepath.Join(secureDir, "secret.txt")
				if err := os.WriteFile(secretFile, []byte("secret"), 0644); err != nil {
					t.Fatalf("failed to create secret file: %v", err)
				}
				if err := os.Chmod(secureDir, 0000); err != nil {
					t.Fatalf("failed to change permissions: %v", err)
				}
				t.Cleanup(func() {
					_ = os.Chmod(secureDir, 0755)
				})
				return secretFile
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skipOnWindows && runtime.GOOS == "windows" {
				t.Skip("Skipping permission tests on Windows")
			}

			path := tt.pathSetup()
			result := lxio.IsFileOK(path)
			if result != tt.expected {
				t.Errorf("IsFileOK(%q) expected %v, got %v", path, tt.expected, result)
			}
		})
	}
}

func TestMustBeFile(t *testing.T) {
	testDir, testFile, _, missingPath := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	tests := []struct {
		name          string
		pathSetup     func() string
		skipOnWindows bool
		expected      bool
		shouldPanic   bool
	}{
		{
			name:        "file returns true",
			pathSetup:   func() string { return testFile },
			expected:    true,
			shouldPanic: false,
		},
		{
			name:        "directory returns false",
			pathSetup:   func() string { return testDir },
			expected:    false,
			shouldPanic: false,
		},
		{
			name:        "missing path returns false",
			pathSetup:   func() string { return missingPath },
			expected:    false,
			shouldPanic: false,
		},
		{
			name:          "panics on permission error",
			skipOnWindows: true,
			shouldPanic:   true,
			pathSetup: func() string {
				secureDir := t.TempDir()
				secretFile := filepath.Join(secureDir, "secret.txt")
				if err := os.WriteFile(secretFile, []byte("secret"), 0644); err != nil {
					t.Fatalf("failed to create secret file: %v", err)
				}
				if err := os.Chmod(secureDir, 0000); err != nil {
					t.Fatalf("failed to change permissions: %v", err)
				}
				t.Cleanup(func() {
					_ = os.Chmod(secureDir, 0755)
				})
				return secretFile
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skipOnWindows && runtime.GOOS == "windows" {
				t.Skip("Skipping permission tests on Windows")
			}

			path := tt.pathSetup()

			if tt.shouldPanic {
				defer func() {
					if recover() == nil {
						t.Errorf("MustBeFile(%q) should panic", path)
					}
				}()
			}

			result := lxio.MustBeFile(path)
			if !tt.shouldPanic && result != tt.expected {
				t.Errorf("MustBeFile(%q) expected %v, got %v", path, tt.expected, result)
			}
		})
	}
}

// ======================================== IsSymlink Tests ========================================

func TestIsSymlink(t *testing.T) {
	testDir, testFile, testSymlink, missingPath := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	tests := []struct {
		name              string
		pathSetup         func() string
		skipOnWindows     bool
		expected          bool
		expectedErr       bool
		shouldHavePermErr bool
	}{
		{
			name:        "symlink returns true",
			pathSetup:   func() string { return testSymlink },
			expected:    true,
			expectedErr: false,
		},
		{
			name:        "file returns false",
			pathSetup:   func() string { return testFile },
			expected:    false,
			expectedErr: false,
		},
		{
			name:        "directory returns false",
			pathSetup:   func() string { return testDir },
			expected:    false,
			expectedErr: false,
		},
		{
			name:        "missing path returns false",
			pathSetup:   func() string { return missingPath },
			expected:    false,
			expectedErr: false,
		},
		{
			name:              "permission error",
			skipOnWindows:     true,
			expected:          false,
			expectedErr:       true,
			shouldHavePermErr: true,
			pathSetup: func() string {
				secureDir := t.TempDir()
				secretFile := filepath.Join(secureDir, "secret.txt")
				if err := os.WriteFile(secretFile, []byte("secret"), 0644); err != nil {
					t.Fatalf("failed to create secret file: %v", err)
				}
				if err := os.Chmod(secureDir, 0000); err != nil {
					t.Fatalf("failed to change permissions: %v", err)
				}
				t.Cleanup(func() {
					_ = os.Chmod(secureDir, 0755)
				})
				return secretFile
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "symlink returns true" && runtime.GOOS == "windows" {
				if _, err := os.Lstat(testSymlink); err != nil {
					t.Skip("Symlink not created (normal on Windows)")
				}
			}

			if tt.skipOnWindows && runtime.GOOS == "windows" {
				t.Skip("Skipping permission tests on Windows")
			}

			path := tt.pathSetup()
			result, err := lxio.IsSymlink(path)
			hasErr := err != nil

			if hasErr != tt.expectedErr {
				t.Errorf("IsSymlink(%q) error expectation failed: expected error=%v, got error=%v (%v)", path, tt.expectedErr, hasErr, err)
			}

			if tt.shouldHavePermErr && (err == nil || !errors.Is(err, os.ErrPermission)) {
				t.Errorf("IsSymlink(%q) expected permission error, got: %v", path, err)
			}

			if !hasErr && result != tt.expected {
				t.Errorf("IsSymlink(%q) expected %v, got %v", path, tt.expected, result)
			}
		})
	}
}

func TestIsSymlinkOK(t *testing.T) {
	testDir, testFile, testSymlink, _ := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	tests := []struct {
		name          string
		pathSetup     func() string
		skipOnWindows bool
		expected      bool
	}{
		{
			name:      "symlink returns true",
			pathSetup: func() string { return testSymlink },
			expected:  true,
		},
		{
			name:      "file returns false",
			pathSetup: func() string { return testFile },
			expected:  false,
		},
		{
			name:          "permission error is swallowed",
			skipOnWindows: true,
			expected:      false,
			pathSetup: func() string {
				secureDir := t.TempDir()
				secretFile := filepath.Join(secureDir, "secret.txt")
				if err := os.WriteFile(secretFile, []byte("secret"), 0644); err != nil {
					t.Fatalf("failed to create secret file: %v", err)
				}
				if err := os.Chmod(secureDir, 0000); err != nil {
					t.Fatalf("failed to change permissions: %v", err)
				}
				t.Cleanup(func() {
					_ = os.Chmod(secureDir, 0755)
				})
				return secretFile
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "symlink returns true" && runtime.GOOS == "windows" {
				if _, err := os.Lstat(testSymlink); err != nil {
					t.Skip("Symlink not created (normal on Windows)")
				}
			}

			if tt.skipOnWindows && runtime.GOOS == "windows" {
				t.Skip("Skipping permission tests on Windows")
			}

			path := tt.pathSetup()
			result := lxio.IsSymlinkOK(path)
			if result != tt.expected {
				t.Errorf("IsSymlinkOK(%q) expected %v, got %v", path, tt.expected, result)
			}
		})
	}
}

func TestMustBeSymlink(t *testing.T) {
	testDir, testFile, testSymlink, missingPath := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	tests := []struct {
		name          string
		pathSetup     func() string
		skipOnWindows bool
		expected      bool
		shouldPanic   bool
	}{
		{
			name:        "symlink returns true",
			pathSetup:   func() string { return testSymlink },
			expected:    true,
			shouldPanic: false,
		},
		{
			name:        "file returns false",
			pathSetup:   func() string { return testFile },
			expected:    false,
			shouldPanic: false,
		},
		{
			name:        "directory returns false",
			pathSetup:   func() string { return testDir },
			expected:    false,
			shouldPanic: false,
		},
		{
			name:        "missing path returns false",
			pathSetup:   func() string { return missingPath },
			expected:    false,
			shouldPanic: false,
		},
		{
			name:          "panics on permission error",
			skipOnWindows: true,
			shouldPanic:   true,
			pathSetup: func() string {
				secureDir := t.TempDir()
				secretFile := filepath.Join(secureDir, "secret.txt")
				if err := os.WriteFile(secretFile, []byte("secret"), 0644); err != nil {
					t.Fatalf("failed to create secret file: %v", err)
				}
				if err := os.Chmod(secureDir, 0000); err != nil {
					t.Fatalf("failed to change permissions: %v", err)
				}
				t.Cleanup(func() {
					_ = os.Chmod(secureDir, 0755)
				})
				return secretFile
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "symlink returns true" && runtime.GOOS == "windows" {
				if _, err := os.Lstat(testSymlink); err != nil {
					t.Skip("Symlink not created (normal on Windows)")
				}
			}

			if tt.skipOnWindows && runtime.GOOS == "windows" {
				t.Skip("Skipping permission tests on Windows")
			}

			path := tt.pathSetup()

			if tt.shouldPanic {
				defer func() {
					if recover() == nil {
						t.Errorf("MustBeSymlink(%q) should panic", path)
					}
				}()
			}

			result := lxio.MustBeSymlink(path)
			if !tt.shouldPanic && result != tt.expected {
				t.Errorf("MustBeSymlink(%q) expected %v, got %v", path, tt.expected, result)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	testDir := t.TempDir()

	// Create test files and directories
	emptyFile := filepath.Join(testDir, "empty.txt")
	if err := os.WriteFile(emptyFile, []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	nonEmptyFile := filepath.Join(testDir, "nonempty.txt")
	if err := os.WriteFile(nonEmptyFile, []byte("content"), 0644); err != nil {
		t.Fatal(err)
	}

	emptyDir := filepath.Join(testDir, "empty_dir")
	if err := os.Mkdir(emptyDir, 0755); err != nil {
		t.Fatal(err)
	}

	nonEmptyDir := filepath.Join(testDir, "nonempty_dir")
	if err := os.Mkdir(nonEmptyDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(nonEmptyDir, "file.txt"), []byte("content"), 0644); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name              string
		pathSetup         func() string
		skipOnWindows     bool
		expected          bool
		expectedErr       bool
		shouldHavePermErr bool
	}{
		{
			name:        "empty file",
			pathSetup:   func() string { return emptyFile },
			expected:    true,
			expectedErr: false,
		},
		{
			name:        "nonempty file",
			pathSetup:   func() string { return nonEmptyFile },
			expected:    false,
			expectedErr: false,
		},
		{
			name:        "empty directory",
			pathSetup:   func() string { return emptyDir },
			expected:    true,
			expectedErr: false,
		},
		{
			name:        "nonempty directory",
			pathSetup:   func() string { return nonEmptyDir },
			expected:    false,
			expectedErr: false,
		},
		{
			name:        "nonexistent path",
			pathSetup:   func() string { return filepath.Join(testDir, "nonexistent") },
			expected:    false,
			expectedErr: true,
		},
		{
			name:              "permission error on file",
			skipOnWindows:     true,
			expected:          false,
			expectedErr:       true,
			shouldHavePermErr: true,
			pathSetup: func() string {
				secureDir := t.TempDir()
				secretFile := filepath.Join(secureDir, "secret.txt")
				if err := os.WriteFile(secretFile, []byte("secret"), 0644); err != nil {
					t.Fatalf("failed to create secret file: %v", err)
				}
				if err := os.Chmod(secureDir, 0000); err != nil {
					t.Fatalf("failed to change permissions: %v", err)
				}
				t.Cleanup(func() {
					_ = os.Chmod(secureDir, 0755)
				})
				return secretFile
			},
		},
		{
			name:              "permission error on directory",
			skipOnWindows:     true,
			expected:          false,
			expectedErr:       true,
			shouldHavePermErr: true,
			pathSetup: func() string {
				secureDir := t.TempDir()
				secretSubDir := filepath.Join(secureDir, "secret_dir")
				if err := os.Mkdir(secretSubDir, 0755); err != nil {
					t.Fatalf("failed to create secret directory: %v", err)
				}
				if err := os.Chmod(secureDir, 0000); err != nil {
					t.Fatalf("failed to change permissions: %v", err)
				}
				t.Cleanup(func() {
					_ = os.Chmod(secureDir, 0755)
				})
				return secretSubDir
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skipOnWindows && runtime.GOOS == "windows" {
				t.Skip("Skipping permission tests on Windows")
			}

			path := tt.pathSetup()
			result, err := lxio.IsEmpty(path)
			hasErr := err != nil

			if hasErr != tt.expectedErr {
				t.Errorf("IsEmpty(%q) error expectation failed: expected error=%v, got error=%v (%v)", path, tt.expectedErr, hasErr, err)
			}

			if tt.shouldHavePermErr && (err == nil || !errors.Is(err, os.ErrPermission)) {
				t.Errorf("IsEmpty(%q) expected permission error, got: %v", path, err)
			}

			if !hasErr && result != tt.expected {
				t.Errorf("IsEmpty(%q) expected %v, got %v", path, tt.expected, result)
			}
		})
	}
}

func TestIsEmptyOK(t *testing.T) {
	testDir := t.TempDir()

	// Create test files and directories
	emptyFile := filepath.Join(testDir, "empty.txt")
	if err := os.WriteFile(emptyFile, []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	nonEmptyFile := filepath.Join(testDir, "nonempty.txt")
	if err := os.WriteFile(nonEmptyFile, []byte("content"), 0644); err != nil {
		t.Fatal(err)
	}

	emptyDir := filepath.Join(testDir, "empty_dir")
	if err := os.Mkdir(emptyDir, 0755); err != nil {
		t.Fatal(err)
	}

	nonEmptyDir := filepath.Join(testDir, "nonempty_dir")
	if err := os.Mkdir(nonEmptyDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(nonEmptyDir, "file.txt"), []byte("content"), 0644); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name          string
		pathSetup     func() string
		skipOnWindows bool
		expected      bool
	}{
		{
			name:      "empty file",
			pathSetup: func() string { return emptyFile },
			expected:  true,
		},
		{
			name:      "nonempty file",
			pathSetup: func() string { return nonEmptyFile },
			expected:  false,
		},
		{
			name:      "empty directory",
			pathSetup: func() string { return emptyDir },
			expected:  true,
		},
		{
			name:      "nonempty directory",
			pathSetup: func() string { return nonEmptyDir },
			expected:  false,
		},
		{
			name:      "nonexistent path returns false",
			pathSetup: func() string { return filepath.Join(testDir, "nonexistent") },
			expected:  false,
		},
		{
			name:          "permission error is swallowed on file",
			skipOnWindows: true,
			expected:      false,
			pathSetup: func() string {
				secureDir := t.TempDir()
				secretFile := filepath.Join(secureDir, "secret.txt")
				if err := os.WriteFile(secretFile, []byte("secret"), 0644); err != nil {
					t.Fatalf("failed to create secret file: %v", err)
				}
				if err := os.Chmod(secureDir, 0000); err != nil {
					t.Fatalf("failed to change permissions: %v", err)
				}
				t.Cleanup(func() {
					_ = os.Chmod(secureDir, 0755)
				})
				return secretFile
			},
		},
		{
			name:          "permission error is swallowed on directory",
			skipOnWindows: true,
			expected:      false,
			pathSetup: func() string {
				secureDir := t.TempDir()
				secretSubDir := filepath.Join(secureDir, "secret_dir")
				if err := os.Mkdir(secretSubDir, 0755); err != nil {
					t.Fatalf("failed to create secret directory: %v", err)
				}
				if err := os.Chmod(secureDir, 0000); err != nil {
					t.Fatalf("failed to change permissions: %v", err)
				}
				t.Cleanup(func() {
					_ = os.Chmod(secureDir, 0755)
				})
				return secretSubDir
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skipOnWindows && runtime.GOOS == "windows" {
				t.Skip("Skipping permission tests on Windows")
			}

			path := tt.pathSetup()
			result := lxio.IsEmptyOK(path)
			if result != tt.expected {
				t.Errorf("IsEmptyOK(%q) expected %v, got %v", path, tt.expected, result)
			}
		})
	}
}

// ======================================== Size Tests ========================================

func TestSize(t *testing.T) {
	testDir := t.TempDir()

	// Create test files with specific sizes
	emptyFile := filepath.Join(testDir, "empty.txt")
	if err := os.WriteFile(emptyFile, []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	smallFile := filepath.Join(testDir, "small.txt")
	content := "hello"
	if err := os.WriteFile(smallFile, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	largeFile := filepath.Join(testDir, "large.txt")
	largeContent := string(make([]byte, 1024)) // 1024 bytes
	if err := os.WriteFile(largeFile, []byte(largeContent), 0644); err != nil {
		t.Fatal(err)
	}

	testSubDir := filepath.Join(testDir, "sub_dir")
	if err := os.Mkdir(testSubDir, 0755); err != nil {
		t.Fatal(err)
	}

	missingPath := filepath.Join(testDir, "nonexistent.txt")

	tests := []struct {
		name              string
		pathSetup         func() string
		skipOnWindows     bool
		expectedSize      int64
		isDirectory       bool
		expectedErr       bool
		shouldHavePermErr bool
	}{
		{
			name:         "empty file",
			pathSetup:    func() string { return emptyFile },
			expectedSize: 0,
			isDirectory:  false,
			expectedErr:  false,
		},
		{
			name:         "small file",
			pathSetup:    func() string { return smallFile },
			expectedSize: int64(len(content)),
			isDirectory:  false,
			expectedErr:  false,
		},
		{
			name:         "large file",
			pathSetup:    func() string { return largeFile },
			expectedSize: 1024,
			isDirectory:  false,
			expectedErr:  false,
		},
		{
			name:         "directory",
			pathSetup:    func() string { return testSubDir },
			expectedSize: 0,
			isDirectory:  true,
			expectedErr:  false,
		},
		{
			name:         "nonexistent path",
			pathSetup:    func() string { return missingPath },
			expectedSize: 0,
			isDirectory:  false,
			expectedErr:  true,
		},
		{
			name:              "permission error",
			skipOnWindows:     true,
			expectedSize:      0,
			expectedErr:       true,
			shouldHavePermErr: true,
			pathSetup: func() string {
				secureDir := t.TempDir()
				secretFile := filepath.Join(secureDir, "secret.txt")
				if err := os.WriteFile(secretFile, []byte("secret"), 0644); err != nil {
					t.Fatalf("failed to create secret file: %v", err)
				}
				if err := os.Chmod(secureDir, 0000); err != nil {
					t.Fatalf("failed to change permissions: %v", err)
				}
				t.Cleanup(func() {
					_ = os.Chmod(secureDir, 0755)
				})
				return secretFile
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skipOnWindows && runtime.GOOS == "windows" {
				t.Skip("Skipping permission tests on Windows")
			}

			path := tt.pathSetup()
			size, err := lxio.Size(path)
			hasErr := err != nil

			if hasErr != tt.expectedErr {
				t.Errorf("Size(%q) error expectation failed: expected error=%v, got error=%v (%v)", path, tt.expectedErr, hasErr, err)
			}

			if tt.shouldHavePermErr && (err == nil || !errors.Is(err, os.ErrPermission)) {
				t.Errorf("Size(%q) expected permission error, got: %v", path, err)
			}

			if !hasErr {
				if tt.isDirectory {
					// Directories have platform-specific size values, just check non-negative
					if size < 0 {
						t.Errorf("Size(%q) directory should have non-negative size, got %d", path, size)
					}
				} else if size != tt.expectedSize {
					t.Errorf("Size(%q) expected %d, got %d", path, tt.expectedSize, size)
				}
			}
		})
	}
}

func TestSizeOK(t *testing.T) {
	testDir := t.TempDir()

	// Create test files with specific sizes
	emptyFile := filepath.Join(testDir, "empty.txt")
	if err := os.WriteFile(emptyFile, []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	smallFile := filepath.Join(testDir, "small.txt")
	content := "hello"
	if err := os.WriteFile(smallFile, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	largeFile := filepath.Join(testDir, "large.txt")
	largeContent := string(make([]byte, 1024)) // 1024 bytes
	if err := os.WriteFile(largeFile, []byte(largeContent), 0644); err != nil {
		t.Fatal(err)
	}

	missingPath := filepath.Join(testDir, "nonexistent.txt")

	tests := []struct {
		name          string
		pathSetup     func() string
		skipOnWindows bool
		expectedSize  int64
		isDirectory   bool
	}{
		{
			name:         "empty file",
			pathSetup:    func() string { return emptyFile },
			expectedSize: 0,
			isDirectory:  false,
		},
		{
			name:         "small file",
			pathSetup:    func() string { return smallFile },
			expectedSize: int64(len(content)),
			isDirectory:  false,
		},
		{
			name:         "large file",
			pathSetup:    func() string { return largeFile },
			expectedSize: 1024,
			isDirectory:  false,
		},
		{
			name:         "nonexistent path",
			pathSetup:    func() string { return missingPath },
			expectedSize: 0,
			isDirectory:  false,
		},
		{
			name:          "permission error is swallowed",
			skipOnWindows: true,
			expectedSize:  0,
			pathSetup: func() string {
				secureDir := t.TempDir()
				secretFile := filepath.Join(secureDir, "secret.txt")
				if err := os.WriteFile(secretFile, []byte("secret"), 0644); err != nil {
					t.Fatalf("failed to create secret file: %v", err)
				}
				if err := os.Chmod(secureDir, 0000); err != nil {
					t.Fatalf("failed to change permissions: %v", err)
				}
				t.Cleanup(func() {
					_ = os.Chmod(secureDir, 0755)
				})
				return secretFile
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skipOnWindows && runtime.GOOS == "windows" {
				t.Skip("Skipping permission tests on Windows")
			}

			path := tt.pathSetup()
			size := lxio.SizeOK(path)
			if !tt.isDirectory && size != tt.expectedSize {
				t.Errorf("SizeOK(%q) expected %d, got %d", path, tt.expectedSize, size)
			}
			if tt.isDirectory && size < 0 {
				t.Errorf("SizeOK(%q) directory should have non-negative size, got %d", path, size)
			}
		})
	}
}

// ======================================== ModTime Tests ========================================

func TestModTime(t *testing.T) {
	testDir := t.TempDir()

	// Create a test file
	testFile := filepath.Join(testDir, "test.txt")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		t.Fatal(err)
	}

	// Create a test directory
	testSubDir := filepath.Join(testDir, "sub_dir")
	if err := os.Mkdir(testSubDir, 0755); err != nil {
		t.Fatal(err)
	}

	missingPath := filepath.Join(testDir, "nonexistent.txt")

	tests := []struct {
		name              string
		pathSetup         func() string
		skipOnWindows     bool
		expectedErr       bool
		shouldHavePermErr bool
		shouldBeNonZero   bool
	}{
		{
			name:            "file modtime",
			pathSetup:       func() string { return testFile },
			expectedErr:     false,
			shouldBeNonZero: true,
		},
		{
			name:            "directory modtime",
			pathSetup:       func() string { return testSubDir },
			expectedErr:     false,
			shouldBeNonZero: true,
		},
		{
			name:            "nonexistent path",
			pathSetup:       func() string { return missingPath },
			expectedErr:     true,
			shouldBeNonZero: false,
		},
		{
			name:              "permission error",
			skipOnWindows:     true,
			expectedErr:       true,
			shouldHavePermErr: true,
			pathSetup: func() string {
				secureDir := t.TempDir()
				secretFile := filepath.Join(secureDir, "secret.txt")
				if err := os.WriteFile(secretFile, []byte("secret"), 0644); err != nil {
					t.Fatalf("failed to create secret file: %v", err)
				}
				if err := os.Chmod(secureDir, 0000); err != nil {
					t.Fatalf("failed to change permissions: %v", err)
				}
				t.Cleanup(func() {
					_ = os.Chmod(secureDir, 0755)
				})
				return secretFile
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skipOnWindows && runtime.GOOS == "windows" {
				t.Skip("Skipping permission tests on Windows")
			}

			path := tt.pathSetup()
			modTime, err := lxio.ModTime(path)
			hasErr := err != nil

			if hasErr != tt.expectedErr {
				t.Errorf("ModTime(%q) error expectation failed: expected error=%v, got error=%v (%v)", path, tt.expectedErr, hasErr, err)
			}

			if tt.shouldHavePermErr && (err == nil || !errors.Is(err, os.ErrPermission)) {
				t.Errorf("ModTime(%q) expected permission error, got: %v", path, err)
			}

			if !hasErr {
				if tt.shouldBeNonZero && modTime == (time.Time{}) {
					t.Errorf("ModTime(%q) expected non-zero time, got zero time", path)
				}
				if !tt.shouldBeNonZero && modTime != (time.Time{}) {
					t.Errorf("ModTime(%q) expected zero time, got %v", path, modTime)
				}
			}
		})
	}
}

// ======================================== IsReadable Tests ========================================

func TestIsReadable(t *testing.T) {
	testDir := t.TempDir()

	// Create a readable file
	readableFile := filepath.Join(testDir, "readable.txt")
	if err := os.WriteFile(readableFile, []byte("content"), 0644); err != nil {
		t.Fatal(err)
	}

	// Create an unreadable file (if possible)
	unreadableFile := filepath.Join(testDir, "unreadable.txt")
	if err := os.WriteFile(unreadableFile, []byte("content"), 0000); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = os.Chmod(unreadableFile, 0644)
	})

	missingPath := filepath.Join(testDir, "nonexistent.txt")

	tests := []struct {
		name          string
		pathSetup     func() string
		skipOnWindows bool
		expected      bool
	}{
		{
			name:      "readable file",
			pathSetup: func() string { return readableFile },
			expected:  true,
		},
		{
			name:          "unreadable file",
			pathSetup:     func() string { return unreadableFile },
			skipOnWindows: true,
			expected:      false,
		},
		{
			name:      "nonexistent path",
			pathSetup: func() string { return missingPath },
			expected:  false,
		},
		{
			name:      "readable directory",
			pathSetup: func() string { return testDir },
			expected:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skipOnWindows && runtime.GOOS == "windows" {
				t.Skip("Skipping permission tests on Windows")
			}

			path := tt.pathSetup()
			result := lxio.IsReadable(path)
			if result != tt.expected {
				t.Errorf("IsReadable(%q) expected %v, got %v", path, tt.expected, result)
			}
		})
	}
}

// ======================================== IsWritable Tests ========================================

func TestIsWritable(t *testing.T) {
	testDir := t.TempDir()

	// Create a writable file
	writableFile := filepath.Join(testDir, "writable.txt")
	if err := os.WriteFile(writableFile, []byte("content"), 0644); err != nil {
		t.Fatal(err)
	}

	// Create a read-only file
	readOnlyFile := filepath.Join(testDir, "readonly.txt")
	if err := os.WriteFile(readOnlyFile, []byte("content"), 0444); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = os.Chmod(readOnlyFile, 0644)
	})

	missingPath := filepath.Join(testDir, "nonexistent.txt")

	tests := []struct {
		name          string
		pathSetup     func() string
		skipOnWindows bool
		expected      bool
	}{
		{
			name:      "writable file",
			pathSetup: func() string { return writableFile },
			expected:  true,
		},
		{
			name:          "read-only file",
			pathSetup:     func() string { return readOnlyFile },
			skipOnWindows: true,
			expected:      false,
		},
		{
			name:      "nonexistent path",
			pathSetup: func() string { return missingPath },
			expected:  false,
		},
		{
			name:      "writable directory",
			pathSetup: func() string { return testDir },
			expected:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skipOnWindows && runtime.GOOS == "windows" {
				t.Skip("Skipping permission tests on Windows")
			}

			path := tt.pathSetup()
			result := lxio.IsWritable(path)
			if result != tt.expected {
				t.Errorf("IsWritable(%q) expected %v, got %v", path, tt.expected, result)
			}
		})
	}
}

// ======================================== IsExecutable Tests ========================================

func TestIsExecutable(t *testing.T) {
	testDir := t.TempDir()

	// Create an executable file
	execFile := filepath.Join(testDir, "executable")
	if err := os.WriteFile(execFile, []byte("#!/bin/bash\necho test"), 0755); err != nil {
		t.Fatal(err)
	}

	// Create a non-executable file
	nonExecFile := filepath.Join(testDir, "nonexecutable.txt")
	if err := os.WriteFile(nonExecFile, []byte("content"), 0644); err != nil {
		t.Fatal(err)
	}

	missingPath := filepath.Join(testDir, "nonexistent.txt")

	tests := []struct {
		name          string
		pathSetup     func() string
		skipOnWindows bool
		expected      bool
	}{
		{
			name:      "executable file (owner bit set)",
			pathSetup: func() string { return execFile },
			expected:  true,
		},
		{
			name:          "non-executable file (owner bit not set)",
			pathSetup:     func() string { return nonExecFile },
			skipOnWindows: true,
			expected:      false,
		},
		{
			name:      "nonexistent path",
			pathSetup: func() string { return missingPath },
			expected:  false,
		},
		{
			name:          "directory (traversable)",
			pathSetup:     func() string { return testDir },
			skipOnWindows: true,
			expected:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skipOnWindows && runtime.GOOS == "windows" {
				t.Skip("Skipping permission tests on Windows")
			}

			path := tt.pathSetup()
			result := lxio.IsExecutable(path)
			if result != tt.expected {
				t.Errorf("IsExecutable(%q) expected %v, got %v", path, tt.expected, result)
			}
		})
	}
}
