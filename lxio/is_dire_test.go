package lxio_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/hgapdvn/lx/lxio"
)

func TestIsDirE_DirectoryExists(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.Remove(tempDir)

	result, err := lxio.IsDirE(tempDir)

	if err != nil {
		t.Errorf("Expected no error for existing directory, got: %v", err)
	}
	if !result {
		t.Errorf("Expected true for existing directory")
	}
}

func TestIsDirE_FileExists(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testfile-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	result, err := lxio.IsDirE(tempFile.Name())

	if err != nil {
		t.Errorf("Expected no error for existing file, got: %v", err)
	}
	if result {
		t.Errorf("Expected false for file (not a directory)")
	}
}

func TestIsDirE_PathDoesNotExist(t *testing.T) {
	nonExistentPath := filepath.Join(os.TempDir(), "nonexistent-dir-xyz")

	result, err := lxio.IsDirE(nonExistentPath)

	if err != nil {
		t.Errorf("Expected no error for non-existent path, got: %v", err)
	}
	if result {
		t.Errorf("Expected false for non-existent path")
	}
}

func TestIsDirE_EmptyPath(t *testing.T) {
	result, err := lxio.IsDirE("")

	if err != nil {
		t.Errorf("Expected no error for empty path, got: %v", err)
	}
	if result {
		t.Errorf("Expected false for empty path")
	}
}

func TestIsDirE_RelativePath_Exists(t *testing.T) {
	tempDir, err := os.MkdirTemp(".", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	dirName := filepath.Base(tempDir)
	defer os.Remove(tempDir)

	result, err := lxio.IsDirE(dirName)

	if err != nil {
		t.Errorf("Expected no error for relative path that exists, got: %v", err)
	}
	if !result {
		t.Errorf("Expected true for relative path that is a directory")
	}
}

func TestIsDirE_NestedPath_Exists(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	nestedPath := filepath.Join(tempDir, "level1", "level2")
	err = os.MkdirAll(nestedPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create nested path: %v", err)
	}

	result, err := lxio.IsDirE(nestedPath)

	if err != nil {
		t.Errorf("Expected no error for nested directory, got: %v", err)
	}
	if !result {
		t.Errorf("Expected true for existing nested directory")
	}
}

func TestIsDirE_Symlink_ToDirectory(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.Remove(tempDir)

	symlinkPath := filepath.Join(os.TempDir(), "testlink-dir")
	err = os.Symlink(tempDir, symlinkPath)
	if err != nil {
		t.Skip("Skipping symlink test - cannot create symlinks on this system")
	}
	defer os.Remove(symlinkPath)

	result, err := lxio.IsDirE(symlinkPath)

	if err != nil {
		t.Errorf("Expected no error for symlink to directory, got: %v", err)
	}
	if !result {
		t.Errorf("Expected true for symlink to directory")
	}
}

func TestIsDirE_Symlink_ToFile(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testfile-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	symlinkPath := filepath.Join(os.TempDir(), "testlink-file")
	err = os.Symlink(tempFile.Name(), symlinkPath)
	if err != nil {
		t.Skip("Skipping symlink test - cannot create symlinks on this system")
	}
	defer os.Remove(symlinkPath)

	result, err := lxio.IsDirE(symlinkPath)

	if err != nil {
		t.Errorf("Expected no error for symlink to file, got: %v", err)
	}
	if result {
		t.Errorf("Expected false for symlink to file")
	}
}

func TestIsDirE_CurrentDirectory(t *testing.T) {
	result, err := lxio.IsDirE(".")

	if err != nil {
		t.Errorf("Expected no error for current directory, got: %v", err)
	}
	if !result {
		t.Errorf("Expected true for current directory '.'")
	}
}

func TestIsDirE_ReturnType_TrueNoError(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.Remove(tempDir)

	result, err := lxio.IsDirE(tempDir)

	if result != true {
		t.Errorf("Expected result to be true")
	}
	if err != nil {
		t.Errorf("Expected err to be nil")
	}
}

func TestIsDirE_ReturnType_FalseNoError_File(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testfile-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	result, err := lxio.IsDirE(tempFile.Name())

	if result != false {
		t.Errorf("Expected result to be false for file")
	}
	if err != nil {
		t.Errorf("Expected err to be nil")
	}
}

func TestIsDirE_PermissionDeniedOnParent(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	subDir := filepath.Join(tempDir, "subdir")
	err = os.Mkdir(subDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create subdirectory: %v", err)
	}

	err = os.Chmod(tempDir, 0000)
	if err != nil {
		t.Skip("Skipping permission test - cannot change directory permissions")
	}
	defer os.Chmod(tempDir, 0755)

	result, err := lxio.IsDirE(subDir)

	if err == nil {
		t.Errorf("Expected error for permission denied on parent, got: result=%v, err=nil", result)
	}
}
