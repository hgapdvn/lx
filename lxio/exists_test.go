package lxio_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/hgapdvn/lx/lxio"
)

func TestExists_FileExists(t *testing.T) {
	// Create a temporary file
	tempFile, err := os.CreateTemp("", "testfile-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	result := lxio.Exists(tempFile.Name())

	if !result {
		t.Errorf("Expected true for existing file")
	}
}

func TestExists_DirectoryExists(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.Remove(tempDir)

	result := lxio.Exists(tempDir)

	if !result {
		t.Errorf("Expected true for existing directory")
	}
}

func TestExists_FileDoesNotExist(t *testing.T) {
	nonExistentPath := filepath.Join(os.TempDir(), "nonexistent-file-that-should-not-exist.txt")

	result := lxio.Exists(nonExistentPath)

	if result {
		t.Errorf("Expected false for non-existent file")
	}
}

func TestExists_EmptyPath(t *testing.T) {
	result := lxio.Exists("")

	if result {
		t.Errorf("Expected false for empty path")
	}
}

func TestExists_RelativePath_Exists(t *testing.T) {
	// Create a file in the current directory
	tempFile, err := os.CreateTemp(".", "testfile-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	filename := filepath.Base(tempFile.Name())
	defer os.Remove(filename)
	defer tempFile.Close()

	result := lxio.Exists(filename)

	if !result {
		t.Errorf("Expected true for relative path that exists")
	}
}

func TestExists_RelativePath_DoesNotExist(t *testing.T) {
	result := lxio.Exists("nonexistent-relative-path.txt")

	if result {
		t.Errorf("Expected false for non-existent relative path")
	}
}

func TestExists_NestedPath_Exists(t *testing.T) {
	// Create a nested directory structure
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	nestedPath := filepath.Join(tempDir, "nested", "path")
	err = os.MkdirAll(nestedPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create nested path: %v", err)
	}

	result := lxio.Exists(nestedPath)

	if !result {
		t.Errorf("Expected true for existing nested path")
	}
}

func TestExists_NestedPath_DoesNotExist(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	nonExistentNestedPath := filepath.Join(tempDir, "nonexistent", "nested", "path")

	result := lxio.Exists(nonExistentNestedPath)

	if result {
		t.Errorf("Expected false for non-existent nested path")
	}
}

func TestExists_Symlink_ToExistingFile(t *testing.T) {
	// Create a temporary file
	tempFile, err := os.CreateTemp("", "testfile-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	// Create a symlink
	symlinkPath := filepath.Join(os.TempDir(), "testlink")
	err = os.Symlink(tempFile.Name(), symlinkPath)
	if err != nil {
		t.Skip("Skipping symlink test - cannot create symlinks on this system")
	}
	defer os.Remove(symlinkPath)

	result := lxio.Exists(symlinkPath)

	if !result {
		t.Errorf("Expected true for symlink to existing file")
	}
}

func TestExists_Symlink_ToNonExistentFile(t *testing.T) {
	// Create a symlink to non-existent file
	nonExistentTarget := filepath.Join(os.TempDir(), "nonexistent-target-for-symlink.txt")
	symlinkPath := filepath.Join(os.TempDir(), "broken-link")

	err := os.Symlink(nonExistentTarget, symlinkPath)
	if err != nil {
		t.Skip("Skipping symlink test - cannot create symlinks on this system")
	}
	defer os.Remove(symlinkPath)

	result := lxio.Exists(symlinkPath)

	if result {
		t.Errorf("Expected false for broken symlink")
	}
}

func TestExists_SpecialPath_DevNull(t *testing.T) {
	result := lxio.Exists("/dev/null")

	if !result {
		t.Errorf("Expected true for /dev/null")
	}
}

func TestExists_DotPath(t *testing.T) {
	// "." represents the current directory
	result := lxio.Exists(".")

	if !result {
		t.Errorf("Expected true for '.' path (current directory)")
	}
}

func TestExists_DotDotPath(t *testing.T) {
	// ".." represents the parent directory
	result := lxio.Exists("..")

	if !result {
		t.Errorf("Expected true for '..' path (parent directory)")
	}
}

func TestExists_SwallowsErrors(t *testing.T) {
	// Verify that Exists returns false when ExistsE would error
	// by checking a path that might cause permission denied on some systems
	nonExistentPath := filepath.Join(os.TempDir(), "nonexistent-test-file.txt")

	result := lxio.Exists(nonExistentPath)

	// Should return false without panicking or erroring
	if result {
		t.Errorf("Expected false for non-existent path")
	}
}
