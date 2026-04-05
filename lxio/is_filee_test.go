package lxio_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/hgapdvn/lx/lxio"
)

func TestIsFileE_RegularFileExists(t *testing.T) {
	// Create a temporary file
	tempFile, err := os.CreateTemp("", "testfile-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	result, err := lxio.IsFileE(tempFile.Name())

	if err != nil {
		t.Errorf("Expected no error for existing file, got: %v", err)
	}
	if !result {
		t.Errorf("Expected true for existing regular file")
	}
}

func TestIsFileE_DirectoryExists(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.Remove(tempDir)

	result, err := lxio.IsFileE(tempDir)

	if err != nil {
		t.Errorf("Expected no error for existing directory, got: %v", err)
	}
	if result {
		t.Errorf("Expected false for directory (not a regular file)")
	}
}

func TestIsFileE_FileDoesNotExist(t *testing.T) {
	nonExistentPath := filepath.Join(os.TempDir(), "nonexistent-file-xyz.txt")

	result, err := lxio.IsFileE(nonExistentPath)

	if err != nil {
		t.Errorf("Expected no error for non-existent path, got: %v", err)
	}
	if result {
		t.Errorf("Expected false for non-existent file")
	}
}

func TestIsFileE_EmptyPath(t *testing.T) {
	result, err := lxio.IsFileE("")

	if err != nil {
		t.Errorf("Expected no error for empty path, got: %v", err)
	}
	if result {
		t.Errorf("Expected false for empty path")
	}
}

func TestIsFileE_RelativePath_Exists(t *testing.T) {
	// Create a file in current directory
	tempFile, err := os.CreateTemp(".", "testfile-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	filename := filepath.Base(tempFile.Name())
	defer os.Remove(filename)
	tempFile.Close()

	result, err := lxio.IsFileE(filename)

	if err != nil {
		t.Errorf("Expected no error for relative path that exists, got: %v", err)
	}
	if !result {
		t.Errorf("Expected true for relative path that is a file")
	}
}

func TestIsFileE_RelativePath_DoesNotExist(t *testing.T) {
	result, err := lxio.IsFileE("nonexistent-relative-file.txt")

	if err != nil {
		t.Errorf("Expected no error for non-existent relative path, got: %v", err)
	}
	if result {
		t.Errorf("Expected false for non-existent relative path")
	}
}

func TestIsFileE_NestedPath_File(t *testing.T) {
	// Create nested directories with a file
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

	filePath := filepath.Join(nestedPath, "file.txt")
	err = os.WriteFile(filePath, []byte("test"), 0644)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	result, err := lxio.IsFileE(filePath)

	if err != nil {
		t.Errorf("Expected no error for nested file, got: %v", err)
	}
	if !result {
		t.Errorf("Expected true for nested file")
	}
}

func TestIsFileE_NestedPath_DoesNotExist(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	nonExistentPath := filepath.Join(tempDir, "nonexistent", "nested", "file.txt")

	result, err := lxio.IsFileE(nonExistentPath)

	if err != nil {
		t.Errorf("Expected no error for non-existent nested path, got: %v", err)
	}
	if result {
		t.Errorf("Expected false for non-existent nested path")
	}
}

func TestIsFileE_Symlink_ToFile(t *testing.T) {
	// Create a temporary file
	tempFile, err := os.CreateTemp("", "testfile-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	// Create a symlink to the file
	symlinkPath := filepath.Join(os.TempDir(), "testlink-file")
	err = os.Symlink(tempFile.Name(), symlinkPath)
	if err != nil {
		t.Skip("Skipping symlink test - cannot create symlinks on this system")
	}
	defer os.Remove(symlinkPath)

	result, err := lxio.IsFileE(symlinkPath)

	if err != nil {
		t.Errorf("Expected no error for symlink to file, got: %v", err)
	}
	if !result {
		t.Errorf("Expected true for symlink to file")
	}
}

func TestIsFileE_Symlink_ToDirectory(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.Remove(tempDir)

	// Create a symlink to the directory
	symlinkPath := filepath.Join(os.TempDir(), "testlink-dir")
	err = os.Symlink(tempDir, symlinkPath)
	if err != nil {
		t.Skip("Skipping symlink test - cannot create symlinks on this system")
	}
	defer os.Remove(symlinkPath)

	result, err := lxio.IsFileE(symlinkPath)

	if err != nil {
		t.Errorf("Expected no error for symlink to directory, got: %v", err)
	}
	if result {
		t.Errorf("Expected false for symlink to directory")
	}
}

func TestIsFileE_Symlink_BrokenLink(t *testing.T) {
	// Create a symlink to non-existent file
	nonExistentTarget := filepath.Join(os.TempDir(), "nonexistent-target-xyz.txt")
	symlinkPath := filepath.Join(os.TempDir(), "broken-link")

	err := os.Symlink(nonExistentTarget, symlinkPath)
	if err != nil {
		t.Skip("Skipping symlink test - cannot create symlinks on this system")
	}
	defer os.Remove(symlinkPath)

	result, err := lxio.IsFileE(symlinkPath)

	if err != nil {
		t.Errorf("Expected no error for broken symlink, got: %v", err)
	}
	if result {
		t.Errorf("Expected false for broken symlink")
	}
}

func TestIsFileE_SpecialFile_DevNull(t *testing.T) {
	result, err := lxio.IsFileE("/dev/null")

	if err != nil {
		t.Errorf("Expected no error for /dev/null, got: %v", err)
	}
	if result {
		t.Errorf("Expected false for /dev/null (special file, not regular file)")
	}
}

func TestIsFileE_EmptyFile(t *testing.T) {
	// Create an empty file
	tempFile, err := os.CreateTemp("", "empty-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	result, err := lxio.IsFileE(tempFile.Name())

	if err != nil {
		t.Errorf("Expected no error for empty file, got: %v", err)
	}
	if !result {
		t.Errorf("Expected true for empty file")
	}
}

func TestIsFileE_LargeFile(t *testing.T) {
	// Create a large file
	tempFile, err := os.CreateTemp("", "large-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write 1MB of data
	largeData := make([]byte, 1024*1024)
	_, err = tempFile.Write(largeData)
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tempFile.Close()

	result, err := lxio.IsFileE(tempFile.Name())

	if err != nil {
		t.Errorf("Expected no error for large file, got: %v", err)
	}
	if !result {
		t.Errorf("Expected true for large file")
	}
}

func TestIsFileE_DotFile(t *testing.T) {
	// Create a hidden file (dot file)
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	dotFilePath := filepath.Join(tempDir, ".hidden")
	err = os.WriteFile(dotFilePath, []byte("test"), 0644)
	if err != nil {
		t.Fatalf("Failed to create dot file: %v", err)
	}

	result, err := lxio.IsFileE(dotFilePath)

	if err != nil {
		t.Errorf("Expected no error for dot file, got: %v", err)
	}
	if !result {
		t.Errorf("Expected true for dot file")
	}
}

func TestIsFileE_FileWithoutExtension(t *testing.T) {
	// Create a file without extension
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	filePath := filepath.Join(tempDir, "noextension")
	err = os.WriteFile(filePath, []byte("test"), 0644)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	result, err := lxio.IsFileE(filePath)

	if err != nil {
		t.Errorf("Expected no error for file without extension, got: %v", err)
	}
	if !result {
		t.Errorf("Expected true for file without extension")
	}
}

func TestIsFileE_DotPath(t *testing.T) {
	// "." is the current directory, not a file
	result, err := lxio.IsFileE(".")

	if err != nil {
		t.Errorf("Expected no error for current directory, got: %v", err)
	}
	if result {
		t.Errorf("Expected false for current directory '.'")
	}
}

func TestIsFileE_DotDotPath(t *testing.T) {
	// ".." is the parent directory, not a file
	result, err := lxio.IsFileE("..")

	if err != nil {
		t.Errorf("Expected no error for parent directory, got: %v", err)
	}
	if result {
		t.Errorf("Expected false for parent directory '..'")
	}
}

func TestIsFileE_ReturnType_TrueNoError(t *testing.T) {
	// Test that when path is a file, we get (true, nil)
	tempFile, err := os.CreateTemp("", "testfile-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	result, err := lxio.IsFileE(tempFile.Name())

	if result != true {
		t.Errorf("Expected result to be true")
	}
	if err != nil {
		t.Errorf("Expected err to be nil")
	}
}

func TestIsFileE_ReturnType_FalseNoError_Directory(t *testing.T) {
	// Test that when path is a directory, we get (false, nil)
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.Remove(tempDir)

	result, err := lxio.IsFileE(tempDir)

	if result != false {
		t.Errorf("Expected result to be false for directory")
	}
	if err != nil {
		t.Errorf("Expected err to be nil")
	}
}

func TestIsFileE_ReturnType_FalseNoError_NonExistent(t *testing.T) {
	// Test that when path doesn't exist, we get (false, nil)
	nonExistentPath := filepath.Join(os.TempDir(), "nonexistent-xyz.txt")

	result, err := lxio.IsFileE(nonExistentPath)

	if result != false {
		t.Errorf("Expected result to be false for non-existent path")
	}
	if err != nil {
		t.Errorf("Expected err to be nil")
	}
}

func TestIsFileE_PermissionDenied(t *testing.T) {
	// Create a temporary file
	tempFile, err := os.CreateTemp("", "testfile-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	// Change permissions to remove read access
	err = os.Chmod(tempFile.Name(), 0000)
	if err != nil {
		t.Skip("Skipping permission test - cannot change file permissions on this system")
	}
	defer os.Chmod(tempFile.Name(), 0644) // Restore for cleanup

	result, err := lxio.IsFileE(tempFile.Name())

	// On some systems (like running as root), permission checks are bypassed
	if err != nil {
		// Expected path: permission denied error
		if result {
			t.Errorf("Expected result to be false when error is returned for permission denied")
		}
	} else {
		// Some systems allow reading file info even with 0000 permissions (e.g., running as root)
		if !result {
			t.Errorf("Expected result to be true when no error is returned")
		}
	}
}

func TestIsFileE_PermissionDeniedOnDirectory(t *testing.T) {
	// Create a temporary directory with a file inside
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	filePath := filepath.Join(tempDir, "testfile.txt")
	err = os.WriteFile(filePath, []byte("test"), 0644)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	// Remove read permissions from directory
	err = os.Chmod(tempDir, 0000)
	if err != nil {
		t.Skip("Skipping permission test - cannot change directory permissions on this system")
	}
	defer os.Chmod(tempDir, 0755) // Restore for cleanup

	result, err := lxio.IsFileE(filePath)

	// We expect an error for permission denied on the directory
	if err == nil {
		t.Errorf("Expected error for permission denied on directory, got: result=%v, err=nil", result)
	}
}

func TestIsFileE_IsRegularFile_NotSymlink(t *testing.T) {
	// Create a regular file
	tempFile, err := os.CreateTemp("", "testfile-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	result, err := lxio.IsFileE(tempFile.Name())

	if err != nil {
		t.Errorf("Expected no error for regular file, got: %v", err)
	}
	if !result {
		t.Errorf("Expected true for regular file")
	}

	// Verify IsRegular returns true for regular files
	info, err := os.Stat(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to stat file: %v", err)
	}
	if !info.Mode().IsRegular() {
		t.Errorf("Expected IsRegular to return true for regular file")
	}
}
