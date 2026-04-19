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

// ======================================== Exists Tests ========================================

func TestExists(t *testing.T) {
	testDir, testFile, testSymlink, missingPath := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	t.Run("file exists", func(t *testing.T) {
		ok, err := lxio.Exists(testFile)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !ok {
			t.Errorf("Exists() = %v, want true", ok)
		}
	})

	t.Run("directory exists", func(t *testing.T) {
		ok, err := lxio.Exists(testDir)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !ok {
			t.Errorf("Exists() = %v, want true", ok)
		}
	})

	t.Run("path does not exist", func(t *testing.T) {
		ok, err := lxio.Exists(missingPath)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if ok {
			t.Errorf("Exists() = %v, want false", ok)
		}
	})

	t.Run("symlink exists", func(t *testing.T) {
		// Check if symlink was created
		if _, err := os.Lstat(testSymlink); err != nil {
			t.Skip("Symlink not created (normal on Windows)")
		}

		// Exists uses os.Stat, which follows symlinks to the target
		ok, err := lxio.Exists(testSymlink)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !ok {
			t.Errorf("Exists() = %v, want true (symlink target exists)", ok)
		}
	})

	t.Run("permission error", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("Skipping permission tests on Windows")
		}

		secureDir := t.TempDir()
		secretFile := filepath.Join(secureDir, "secret.txt")
		err := os.WriteFile(secretFile, []byte("secret"), 0644)
		if err != nil {
			t.Fatalf("failed to create secret file: %v", err)
		}

		err = os.Chmod(secureDir, 0000)
		if err != nil {
			t.Fatalf("failed to change permissions: %v", err)
		}
		t.Cleanup(func() {
			_ = os.Chmod(secureDir, 0755)
		})

		// Exists returns error on permission denied
		ok, err := lxio.Exists(secretFile)
		if ok {
			t.Errorf("Exists() = %v, want false", ok)
		}
		if err == nil || errors.Is(err, os.ErrNotExist) {
			t.Errorf("Exists() expected a permission error, got: %v", err)
		}
	})
}

func TestExistsOK(t *testing.T) {
	testDir, testFile, _, missingPath := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	t.Run("file exists", func(t *testing.T) {
		ok := lxio.ExistsOK(testFile)
		if !ok {
			t.Errorf("ExistsOK() = %v, want true", ok)
		}
	})

	t.Run("path does not exist", func(t *testing.T) {
		ok := lxio.ExistsOK(missingPath)
		if ok {
			t.Errorf("ExistsOK() = %v, want false", ok)
		}
	})

	t.Run("permission error is swallowed", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("Skipping permission tests on Windows")
		}

		secureDir := t.TempDir()
		secretFile := filepath.Join(secureDir, "secret.txt")
		err := os.WriteFile(secretFile, []byte("secret"), 0644)
		if err != nil {
			t.Fatalf("failed to create secret file: %v", err)
		}

		err = os.Chmod(secureDir, 0000)
		if err != nil {
			t.Fatalf("failed to change permissions: %v", err)
		}
		t.Cleanup(func() {
			_ = os.Chmod(secureDir, 0755)
		})

		// ExistsOK swallows permission errors and returns false
		ok := lxio.ExistsOK(secretFile)
		if ok {
			t.Errorf("ExistsOK() = %v, want false", ok)
		}
	})
}

func TestMustExist(t *testing.T) {
	testDir, testFile, _, _ := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	t.Run("file exists", func(t *testing.T) {
		ok := lxio.MustExist(testFile)
		if !ok {
			t.Errorf("MustExist() = %v, want true", ok)
		}
	})

	t.Run("path does not exist returns false", func(t *testing.T) {
		missingPath := filepath.Join(testDir, "nonexistent.txt")
		ok := lxio.MustExist(missingPath)
		if ok {
			t.Errorf("MustExist() = %v, want false", ok)
		}
	})

	t.Run("panics on permission error", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("Skipping permission tests on Windows")
		}

		secureDir := t.TempDir()
		secretFile := filepath.Join(secureDir, "secret.txt")
		err := os.WriteFile(secretFile, []byte("secret"), 0644)
		if err != nil {
			t.Fatalf("failed to create secret file: %v", err)
		}

		err = os.Chmod(secureDir, 0000)
		if err != nil {
			t.Fatalf("failed to change permissions: %v", err)
		}
		t.Cleanup(func() {
			_ = os.Chmod(secureDir, 0755)
		})

		defer func() {
			if recover() == nil {
				t.Errorf("MustExist() should panic on permission error")
			}
		}()
		lxio.MustExist(secretFile)
	})
}

func TestNotExists(t *testing.T) {
	testDir, testFile, _, missingPath := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	t.Run("file exists", func(t *testing.T) {
		ok, err := lxio.NotExists(testFile)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if ok {
			t.Errorf("NotExists() = %v, want false", ok)
		}
	})

	t.Run("path does not exist", func(t *testing.T) {
		ok, err := lxio.NotExists(missingPath)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !ok {
			t.Errorf("NotExists() = %v, want true", ok)
		}
	})

	t.Run("permission error", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("Skipping permission tests on Windows")
		}

		secureDir := t.TempDir()
		secretFile := filepath.Join(secureDir, "secret.txt")
		err := os.WriteFile(secretFile, []byte("secret"), 0644)
		if err != nil {
			t.Fatalf("failed to create secret file: %v", err)
		}

		err = os.Chmod(secureDir, 0000)
		if err != nil {
			t.Fatalf("failed to change permissions: %v", err)
		}
		t.Cleanup(func() {
			_ = os.Chmod(secureDir, 0755)
		})

		// NotExists returns error on permission denied
		// Returns (false, error)—conservative when existence is ambiguous
		ok, err := lxio.NotExists(secretFile)
		if ok {
			t.Errorf("NotExists() = %v, want false", ok)
		}
		if err == nil || errors.Is(err, os.ErrNotExist) {
			t.Errorf("NotExists() expected a permission error, got: %v", err)
		}
	})
}

func TestNotExistsOK(t *testing.T) {
	testDir, testFile, _, missingPath := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	t.Run("file exists", func(t *testing.T) {
		ok := lxio.NotExistsOK(testFile)
		if ok {
			t.Errorf("NotExistsOK() = %v, want false", ok)
		}
	})

	t.Run("path does not exist", func(t *testing.T) {
		ok := lxio.NotExistsOK(missingPath)
		if !ok {
			t.Errorf("NotExistsOK() = %v, want true", ok)
		}
	})

	t.Run("permission error is swallowed", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("Skipping permission tests on Windows")
		}

		secureDir := t.TempDir()
		secretFile := filepath.Join(secureDir, "secret.txt")
		err := os.WriteFile(secretFile, []byte("secret"), 0644)
		if err != nil {
			t.Fatalf("failed to create secret file: %v", err)
		}

		err = os.Chmod(secureDir, 0000)
		if err != nil {
			t.Fatalf("failed to change permissions: %v", err)
		}
		t.Cleanup(func() {
			_ = os.Chmod(secureDir, 0755)
		})

		// NotExistsOK swallows permission errors and returns false
		// Conservative: when we can't determine existence, assume the file exists or is inaccessible
		ok := lxio.NotExistsOK(secretFile)
		if ok {
			t.Errorf("NotExistsOK() = %v, want false", ok)
		}
	})
}

func TestMustNotExist(t *testing.T) {
	testDir, testFile, _, missingPath := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	t.Run("path does not exist returns true", func(t *testing.T) {
		ok := lxio.MustNotExist(missingPath)
		if !ok {
			t.Errorf("MustNotExist() = %v, want true", ok)
		}
	})

	t.Run("file exists returns false", func(t *testing.T) {
		ok := lxio.MustNotExist(testFile)
		if ok {
			t.Errorf("MustNotExist() = %v, want false", ok)
		}
	})

	t.Run("directory exists returns false", func(t *testing.T) {
		ok := lxio.MustNotExist(testDir)
		if ok {
			t.Errorf("MustNotExist() = %v, want false", ok)
		}
	})

	t.Run("panics on permission error", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("Skipping permission tests on Windows")
		}

		secureDir := t.TempDir()
		secretFile := filepath.Join(secureDir, "secret.txt")
		err := os.WriteFile(secretFile, []byte("secret"), 0644)
		if err != nil {
			t.Fatalf("failed to create secret file: %v", err)
		}

		err = os.Chmod(secureDir, 0000)
		if err != nil {
			t.Fatalf("failed to change permissions: %v", err)
		}
		t.Cleanup(func() {
			_ = os.Chmod(secureDir, 0755)
		})

		defer func() {
			if recover() == nil {
				t.Errorf("MustNotExist() should panic on permission error")
			}
		}()
		lxio.MustNotExist(secretFile)
	})
}

// ======================================== IsDir Tests ========================================

func TestIsDir(t *testing.T) {
	testDir, testFile, _, missingPath := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	t.Run("directory returns true", func(t *testing.T) {
		ok, err := lxio.IsDir(testDir)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !ok {
			t.Errorf("IsDir() = %v, want true", ok)
		}
	})

	t.Run("file returns false", func(t *testing.T) {
		ok, err := lxio.IsDir(testFile)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if ok {
			t.Errorf("IsDir() = %v, want false", ok)
		}
	})

	t.Run("missing path returns false", func(t *testing.T) {
		ok, err := lxio.IsDir(missingPath)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if ok {
			t.Errorf("IsDir() = %v, want false", ok)
		}
	})

	t.Run("permission error", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("Skipping permission tests on Windows")
		}

		secureDir := t.TempDir()
		secretFile := filepath.Join(secureDir, "secret.txt")
		err := os.WriteFile(secretFile, []byte("secret"), 0644)
		if err != nil {
			t.Fatalf("failed to create secret file: %v", err)
		}

		err = os.Chmod(secureDir, 0000)
		if err != nil {
			t.Fatalf("failed to change permissions: %v", err)
		}
		t.Cleanup(func() {
			_ = os.Chmod(secureDir, 0755)
		})

		// IsDir returns error on permission denied
		ok, err := lxio.IsDir(secretFile)
		if ok {
			t.Errorf("IsDir() = %v, want false", ok)
		}
		if err == nil || errors.Is(err, os.ErrNotExist) {
			t.Errorf("IsDir() expected a permission error, got: %v", err)
		}
	})
}

func TestIsDirOK(t *testing.T) {
	testDir, testFile, _, _ := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	t.Run("directory returns true", func(t *testing.T) {
		ok := lxio.IsDirOK(testDir)
		if !ok {
			t.Errorf("IsDirOK() = %v, want true", ok)
		}
	})

	t.Run("file returns false", func(t *testing.T) {
		ok := lxio.IsDirOK(testFile)
		if ok {
			t.Errorf("IsDirOK() = %v, want false", ok)
		}
	})

	t.Run("permission error is swallowed", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("Skipping permission tests on Windows")
		}

		secureDir := t.TempDir()
		secretFile := filepath.Join(secureDir, "secret.txt")
		err := os.WriteFile(secretFile, []byte("secret"), 0644)
		if err != nil {
			t.Fatalf("failed to create secret file: %v", err)
		}

		err = os.Chmod(secureDir, 0000)
		if err != nil {
			t.Fatalf("failed to change permissions: %v", err)
		}
		t.Cleanup(func() {
			_ = os.Chmod(secureDir, 0755)
		})

		// IsDirOK swallows permission errors and returns false
		ok := lxio.IsDirOK(secretFile)
		if ok {
			t.Errorf("IsDirOK() = %v, want false", ok)
		}
	})
}

func TestMustBeDir(t *testing.T) {
	testDir, testFile, _, missingPath := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	t.Run("directory returns true", func(t *testing.T) {
		ok := lxio.MustBeDir(testDir)
		if !ok {
			t.Errorf("MustBeDir() = %v, want true", ok)
		}
	})

	t.Run("file returns false", func(t *testing.T) {
		ok := lxio.MustBeDir(testFile)
		if ok {
			t.Errorf("MustBeDir() = %v, want false", ok)
		}
	})

	t.Run("missing path returns false", func(t *testing.T) {
		ok := lxio.MustBeDir(missingPath)
		if ok {
			t.Errorf("MustBeDir() = %v, want false", ok)
		}
	})

	t.Run("panics on permission error", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("Skipping permission tests on Windows")
		}

		secureDir := t.TempDir()
		secretFile := filepath.Join(secureDir, "secret.txt")
		err := os.WriteFile(secretFile, []byte("secret"), 0644)
		if err != nil {
			t.Fatalf("failed to create secret file: %v", err)
		}

		err = os.Chmod(secureDir, 0000)
		if err != nil {
			t.Fatalf("failed to change permissions: %v", err)
		}
		t.Cleanup(func() {
			_ = os.Chmod(secureDir, 0755)
		})

		defer func() {
			if recover() == nil {
				t.Errorf("MustBeDir() should panic on permission error")
			}
		}()
		lxio.MustBeDir(secretFile)
	})
}

// ======================================== IsFile Tests ========================================

func TestIsFile(t *testing.T) {
	testDir, testFile, testSymlink, missingPath := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	t.Run("file returns true", func(t *testing.T) {
		ok, err := lxio.IsFile(testFile)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !ok {
			t.Errorf("IsFile() = %v, want true", ok)
		}
	})

	t.Run("directory returns false", func(t *testing.T) {
		ok, err := lxio.IsFile(testDir)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if ok {
			t.Errorf("IsFile() = %v, want false", ok)
		}
	})

	t.Run("missing path returns false", func(t *testing.T) {
		ok, err := lxio.IsFile(missingPath)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if ok {
			t.Errorf("IsFile() = %v, want false", ok)
		}
	})

	t.Run("symlink returns true", func(t *testing.T) {
		// Check if symlink was created
		if _, err := os.Lstat(testSymlink); err != nil {
			t.Skip("Symlink not created (normal on Windows)")
		}

		// IsFile uses os.Stat, which follows symlinks to the target file
		ok, err := lxio.IsFile(testSymlink)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !ok {
			t.Errorf("IsFile() = %v, want true (symlink target is a file)", ok)
		}
	})

	t.Run("permission error", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("Skipping permission tests on Windows")
		}

		secureDir := t.TempDir()
		secretFile := filepath.Join(secureDir, "secret.txt")
		err := os.WriteFile(secretFile, []byte("secret"), 0644)
		if err != nil {
			t.Fatalf("failed to create secret file: %v", err)
		}

		err = os.Chmod(secureDir, 0000)
		if err != nil {
			t.Fatalf("failed to change permissions: %v", err)
		}
		t.Cleanup(func() {
			_ = os.Chmod(secureDir, 0755)
		})

		// IsFile returns error on permission denied
		ok, err := lxio.IsFile(secretFile)
		if ok {
			t.Errorf("IsFile() = %v, want false", ok)
		}
		if err == nil || errors.Is(err, os.ErrNotExist) {
			t.Errorf("IsFile() expected a permission error, got: %v", err)
		}
	})
}

func TestIsFileOK(t *testing.T) {
	testDir, testFile, _, _ := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	t.Run("file returns true", func(t *testing.T) {
		ok := lxio.IsFileOK(testFile)
		if !ok {
			t.Errorf("IsFileOK() = %v, want true", ok)
		}
	})

	t.Run("directory returns false", func(t *testing.T) {
		ok := lxio.IsFileOK(testDir)
		if ok {
			t.Errorf("IsFileOK() = %v, want false", ok)
		}
	})

	t.Run("permission error is swallowed", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("Skipping permission tests on Windows")
		}

		secureDir := t.TempDir()
		secretFile := filepath.Join(secureDir, "secret.txt")
		err := os.WriteFile(secretFile, []byte("secret"), 0644)
		if err != nil {
			t.Fatalf("failed to create secret file: %v", err)
		}

		err = os.Chmod(secureDir, 0000)
		if err != nil {
			t.Fatalf("failed to change permissions: %v", err)
		}
		t.Cleanup(func() {
			_ = os.Chmod(secureDir, 0755)
		})

		// IsFileOK swallows permission errors and returns false
		ok := lxio.IsFileOK(secretFile)
		if ok {
			t.Errorf("IsFileOK() = %v, want false", ok)
		}
	})
}

func TestMustBeFile(t *testing.T) {
	testDir, testFile, _, missingPath := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	t.Run("file returns true", func(t *testing.T) {
		ok := lxio.MustBeFile(testFile)
		if !ok {
			t.Errorf("MustBeFile() = %v, want true", ok)
		}
	})

	t.Run("directory returns false", func(t *testing.T) {
		ok := lxio.MustBeFile(testDir)
		if ok {
			t.Errorf("MustBeFile() = %v, want false", ok)
		}
	})

	t.Run("missing path returns false", func(t *testing.T) {
		ok := lxio.MustBeFile(missingPath)
		if ok {
			t.Errorf("MustBeFile() = %v, want false", ok)
		}
	})

	t.Run("panics on permission error", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("Skipping permission tests on Windows")
		}

		secureDir := t.TempDir()
		secretFile := filepath.Join(secureDir, "secret.txt")
		err := os.WriteFile(secretFile, []byte("secret"), 0644)
		if err != nil {
			t.Fatalf("failed to create secret file: %v", err)
		}

		err = os.Chmod(secureDir, 0000)
		if err != nil {
			t.Fatalf("failed to change permissions: %v", err)
		}
		t.Cleanup(func() {
			_ = os.Chmod(secureDir, 0755)
		})

		defer func() {
			if recover() == nil {
				t.Errorf("MustBeFile() should panic on permission error")
			}
		}()
		lxio.MustBeFile(secretFile)
	})
}

// ======================================== IsSymlink Tests ========================================

func TestIsSymlink(t *testing.T) {
	testDir, testFile, testSymlink, missingPath := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	t.Run("symlink returns true", func(t *testing.T) {
		// Check if symlink was created
		if _, err := os.Lstat(testSymlink); err != nil {
			t.Skip("Symlink not created (normal on Windows)")
		}

		ok, err := lxio.IsSymlink(testSymlink)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !ok {
			t.Errorf("IsSymlink() = %v, want true", ok)
		}
	})

	t.Run("file returns false", func(t *testing.T) {
		ok, err := lxio.IsSymlink(testFile)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if ok {
			t.Errorf("IsSymlink() = %v, want false", ok)
		}
	})

	t.Run("directory returns false", func(t *testing.T) {
		ok, err := lxio.IsSymlink(testDir)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if ok {
			t.Errorf("IsSymlink() = %v, want false", ok)
		}
	})

	t.Run("missing path returns false", func(t *testing.T) {
		ok, err := lxio.IsSymlink(missingPath)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if ok {
			t.Errorf("IsSymlink() = %v, want false", ok)
		}
	})

	t.Run("permission error", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("Skipping permission tests on Windows")
		}

		secureDir := t.TempDir()
		secretFile := filepath.Join(secureDir, "secret.txt")
		err := os.WriteFile(secretFile, []byte("secret"), 0644)
		if err != nil {
			t.Fatalf("failed to create secret file: %v", err)
		}

		err = os.Chmod(secureDir, 0000)
		if err != nil {
			t.Fatalf("failed to change permissions: %v", err)
		}
		t.Cleanup(func() {
			_ = os.Chmod(secureDir, 0755)
		})

		// IsSymlink returns error on permission denied
		ok, err := lxio.IsSymlink(secretFile)
		if ok {
			t.Errorf("IsSymlink() = %v, want false", ok)
		}
		if err == nil || errors.Is(err, os.ErrNotExist) {
			t.Errorf("IsSymlink() expected a permission error, got: %v", err)
		}
	})
}

func TestIsSymlinkOK(t *testing.T) {
	testDir, testFile, testSymlink, _ := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	t.Run("symlink returns true", func(t *testing.T) {
		// Check if symlink was created
		if _, err := os.Lstat(testSymlink); err != nil {
			t.Skip("Symlink not created (normal on Windows)")
		}

		ok := lxio.IsSymlinkOK(testSymlink)
		if !ok {
			t.Errorf("IsSymlinkOK() = %v, want true", ok)
		}
	})

	t.Run("file returns false", func(t *testing.T) {
		ok := lxio.IsSymlinkOK(testFile)
		if ok {
			t.Errorf("IsSymlinkOK() = %v, want false", ok)
		}
	})

	t.Run("permission error is swallowed", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("Skipping permission tests on Windows")
		}

		secureDir := t.TempDir()
		secretFile := filepath.Join(secureDir, "secret.txt")
		err := os.WriteFile(secretFile, []byte("secret"), 0644)
		if err != nil {
			t.Fatalf("failed to create secret file: %v", err)
		}

		err = os.Chmod(secureDir, 0000)
		if err != nil {
			t.Fatalf("failed to change permissions: %v", err)
		}
		t.Cleanup(func() {
			_ = os.Chmod(secureDir, 0755)
		})

		// IsSymlinkOK swallows permission errors and returns false
		ok := lxio.IsSymlinkOK(secretFile)
		if ok {
			t.Errorf("IsSymlinkOK() = %v, want false", ok)
		}
	})
}

func TestMustBeSymlink(t *testing.T) {
	testDir, testFile, testSymlink, missingPath := setupTestEnvironment(t)
	defer os.RemoveAll(testDir)

	t.Run("symlink returns true", func(t *testing.T) {
		// Check if symlink was created
		if _, err := os.Lstat(testSymlink); err != nil {
			t.Skip("Symlink not created (normal on Windows)")
		}

		ok := lxio.MustBeSymlink(testSymlink)
		if !ok {
			t.Errorf("MustBeSymlink() = %v, want true", ok)
		}
	})

	t.Run("file returns false", func(t *testing.T) {
		ok := lxio.MustBeSymlink(testFile)
		if ok {
			t.Errorf("MustBeSymlink() = %v, want false", ok)
		}
	})

	t.Run("directory returns false", func(t *testing.T) {
		ok := lxio.MustBeSymlink(testDir)
		if ok {
			t.Errorf("MustBeSymlink() = %v, want false", ok)
		}
	})

	t.Run("missing path returns false", func(t *testing.T) {
		ok := lxio.MustBeSymlink(missingPath)
		if ok {
			t.Errorf("MustBeSymlink() = %v, want false", ok)
		}
	})

	t.Run("panics on permission error", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("Skipping permission tests on Windows")
		}

		secureDir := t.TempDir()
		secretFile := filepath.Join(secureDir, "secret.txt")
		err := os.WriteFile(secretFile, []byte("secret"), 0644)
		if err != nil {
			t.Fatalf("failed to create secret file: %v", err)
		}

		err = os.Chmod(secureDir, 0000)
		if err != nil {
			t.Fatalf("failed to change permissions: %v", err)
		}
		t.Cleanup(func() {
			_ = os.Chmod(secureDir, 0755)
		})

		defer func() {
			if recover() == nil {
				t.Errorf("MustBeSymlink() should panic on permission error")
			}
		}()
		lxio.MustBeSymlink(secretFile)
	})
}
