package lxio_test

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"testing"

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

func TestLxioHelpers(t *testing.T) {
	dirPath, filePath, symlinkPath, missingPath := setupTestEnvironment(t)

	// We use an anonymous struct to define our test cases clearly
	tests := []struct {
		name        string
		path        string
		wantExist   bool
		wantFile    bool
		wantDir     bool
		wantSymlink bool
	}{
		{
			name:        "Regular File",
			path:        filePath,
			wantExist:   true,
			wantFile:    true,
			wantDir:     false,
			wantSymlink: false,
		},
		{
			name:        "Directory",
			path:        dirPath,
			wantExist:   true,
			wantFile:    false,
			wantDir:     true,
			wantSymlink: false,
		},
		{
			name:        "Missing Path",
			path:        missingPath,
			wantExist:   false,
			wantFile:    false,
			wantDir:     false,
			wantSymlink: false,
		},
	}

	// Only add the symlink test if we successfully created it
	if _, err := os.Lstat(symlinkPath); err == nil {
		tests = append(tests, struct {
			name        string
			path        string
			wantExist   bool
			wantFile    bool
			wantDir     bool
			wantSymlink bool
		}{
			name: "Symlink",
			path: symlinkPath,
			// Exists and IsFile use os.Stat, which follows the symlink to the target file.
			// Therefore, they will return true because the target is a valid file.
			wantExist:   true,
			wantFile:    true,
			wantDir:     false,
			wantSymlink: true, // Uses Lstat, so it knows it is a link
		})
	}

	// Run all the test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// --- Test Quiet Methods ---
			if got := lxio.Exists(tt.path); got != tt.wantExist {
				t.Errorf("Exists() = %v, want %v", got, tt.wantExist)
			}
			if got := lxio.IsFile(tt.path); got != tt.wantFile {
				t.Errorf("IsFile() = %v, want %v", got, tt.wantFile)
			}
			if got := lxio.IsDir(tt.path); got != tt.wantDir {
				t.Errorf("IsDir() = %v, want %v", got, tt.wantDir)
			}
			if got := lxio.IsSymlink(tt.path); got != tt.wantSymlink {
				t.Errorf("IsSymlink() = %v, want %v", got, tt.wantSymlink)
			}

			// --- Test Loud (Error) Methods ---
			gotE, err := lxio.ExistsE(tt.path)
			if err != nil {
				t.Errorf("ExistsE() unexpected error: %v", err)
			}
			if gotE != tt.wantExist {
				t.Errorf("ExistsE() = %v, want %v", gotE, tt.wantExist)
			}

			gotFileE, err := lxio.IsFileE(tt.path)
			if err != nil {
				t.Errorf("IsFileE() unexpected error: %v", err)
			}
			if gotFileE != tt.wantFile {
				t.Errorf("IsFileE() = %v, want %v", gotFileE, tt.wantFile)
			}

			gotDirE, err := lxio.IsDirE(tt.path)
			if err != nil {
				t.Errorf("IsDirE() unexpected error: %v", err)
			}
			if gotDirE != tt.wantDir {
				t.Errorf("IsDirE() = %v, want %v", gotDirE, tt.wantDir)
			}

			gotSymlinkE, err := lxio.IsSymlinkE(tt.path)
			if err != nil {
				t.Errorf("IsSymlinkE() unexpected error: %v", err)
			}
			if gotSymlinkE != tt.wantSymlink {
				t.Errorf("IsSymlinkE() = %v, want %v", gotSymlinkE, tt.wantSymlink)
			}
		})
	}
}

func TestLxioHelpers_PermissionError(t *testing.T) {
	// Skip this test on Windows.
	// Windows handles file permissions differently and chmod 000 often doesn't
	// behave the same way as it does on Unix-like systems.
	if runtime.GOOS == "windows" {
		t.Skip("Skipping permission tests on Windows")
	}

	// 1. Create a secure directory
	secureDir := t.TempDir()

	// 2. Create a file inside that we will try to access
	secretFile := filepath.Join(secureDir, "secret.txt")
	err := os.WriteFile(secretFile, []byte("secret"), 0644)
	if err != nil {
		t.Fatalf("failed to create secret file: %v", err)
	}

	// 3. Remove all permissions from the directory (chmod 000)
	// This makes it impossible to stat anything inside it.
	err = os.Chmod(secureDir, 0000)
	if err != nil {
		t.Fatalf("failed to change permissions: %v", err)
	}

	// Ensure we restore permissions at the end of the test so t.TempDir()
	// can successfully clean up the directory.
	t.Cleanup(func() {
		os.Chmod(secureDir, 0755)
	})

	// --- Test the Quiet functions ---
	// They should all swallow the permission error and return false.
	if got := lxio.Exists(secretFile); got != false {
		t.Errorf("Exists() = %v, want false (due to permission error)", got)
	}
	if got := lxio.IsFile(secretFile); got != false {
		t.Errorf("IsFile() = %v, want false", got)
	}
	if got := lxio.IsDir(secretFile); got != false {
		t.Errorf("IsDir() = %v, want false", got)
	}

	// --- Test the Loud (E) functions ---
	// They should all return false AND surface the permission error.

	gotE, errE := lxio.ExistsE(secretFile)
	if gotE != false {
		t.Errorf("ExistsE() = %v, want false", gotE)
	}
	if errE == nil || errors.Is(errE, os.ErrNotExist) {
		t.Errorf("ExistsE() expected a permission error, got: %v", errE)
	}

	gotFileE, errFileE := lxio.IsFileE(secretFile)
	if gotFileE != false {
		t.Errorf("IsFileE() = %v, want false", gotFileE)
	}
	if errFileE == nil || errors.Is(errFileE, os.ErrNotExist) {
		t.Errorf("IsFileE() expected a permission error, got: %v", errFileE)
	}

	gotDirE, errDirE := lxio.IsDirE(secretFile)
	if gotDirE != false {
		t.Errorf("IsDirE() = %v, want false", gotDirE)
	}
	if errDirE == nil || errors.Is(errDirE, os.ErrNotExist) {
		t.Errorf("IsDirE() expected a permission error, got: %v", errDirE)
	}

	if got := lxio.IsSymlink(secretFile); got != false {
		t.Errorf("IsSymlink() = %v, want false", got)
	}

	gotSymlinkE, errSymlinkE := lxio.IsSymlinkE(secretFile)
	if gotSymlinkE != false {
		t.Errorf("IsSymlinkE() = %v, want false", gotSymlinkE)
	}
	if errSymlinkE == nil || errors.Is(errSymlinkE, os.ErrNotExist) {
		t.Errorf("IsSymlinkE() expected a permission error, got: %v", errSymlinkE)
	}
}
