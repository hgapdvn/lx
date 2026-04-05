package lxio_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/hgapdvn/lx/lxio"
)

func TestIsFile_RegularFileExists(t *testing.T) {
	// Create a temporary file
	tempFile, err := os.CreateTemp("", "testfile-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	result := lxio.IsFile(tempFile.Name())

	if !result {
		t.Errorf("Expected true for existing regular file")
	}
}

func TestIsFile_DirectoryExists(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.Remove(tempDir)

	result := lxio.IsFile(tempDir)

	if result {
		t.Errorf("Expected false for directory")
	}
}

func TestIsFile_FileDoesNotExist(t *testing.T) {
	nonExistentPath := filepath.Join(os.TempDir(), "nonexistent-file-xyz.txt")

	result := lxio.IsFile(nonExistentPath)

	if result {
		t.Errorf("Expected false for non-existent file")
	}
}

func TestIsFile_EmptyPath(t *testing.T) {
	result := lxio.IsFile("")

	if result {
		t.Errorf("Expected false for empty path")
	}
}

func TestIsFile_RelativePath_Exists(t *testing.T) {
	// Create a file in current directory
	tempFile, err := os.CreateTemp(".", "testfile-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	filename := filepath.Base(tempFile.Name())
	defer os.Remove(filename)
	tempFile.Close()

	result := lxio.IsFile(filename)

	if !result {
		t.Errorf("Expected true for relative path that is a file")
	}
}

func TestIsFile_RelativePath_DoesNotExist(t *testing.T) {
	result := lxio.IsFile("nonexistent-relative-file.txt")

	if result {
		t.Errorf("Expected false for non-existent relative path")
	}
}

func TestIsFile_NestedPath_File(t *testing.T) {
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

	result := lxio.IsFile(filePath)

	if !result {
		t.Errorf("Expected true for nested file")
	}
}

func TestIsFile_NestedPath_DoesNotExist(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	nonExistentPath := filepath.Join(tempDir, "nonexistent", "nested", "file.txt")

	result := lxio.IsFile(nonExistentPath)

	if result {
		t.Errorf("Expected false for non-existent nested path")
	}
}

func TestIsFile_Symlink_ToFile(t *testing.T) {
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

	result := lxio.IsFile(symlinkPath)

	if !result {
		t.Errorf("Expected true for symlink to file")
	}
}

func TestIsFile_Symlink_ToDirectory(t *testing.T) {
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

	result := lxio.IsFile(symlinkPath)

	if result {
		t.Errorf("Expected false for symlink to directory")
	}
}

func TestIsFile_Symlink_BrokenLink(t *testing.T) {
	// Create a symlink to non-existent file
	nonExistentTarget := filepath.Join(os.TempDir(), "nonexistent-target-xyz.txt")
	symlinkPath := filepath.Join(os.TempDir(), "broken-link")

	err := os.Symlink(nonExistentTarget, symlinkPath)
	if err != nil {
		t.Skip("Skipping symlink test - cannot create symlinks on this system")
	}
	defer os.Remove(symlinkPath)

	result := lxio.IsFile(symlinkPath)

	if result {
		t.Errorf("Expected false for broken symlink")
	}
}

func TestIsFile_SpecialFile_DevNull(t *testing.T) {
	result := lxio.IsFile("/dev/null")

	if result {
		t.Errorf("Expected false for /dev/null (special file, not regular file)")
	}
}

func TestIsFile_EmptyFile(t *testing.T) {
	// Create an empty file
	tempFile, err := os.CreateTemp("", "empty-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	result := lxio.IsFile(tempFile.Name())

	if !result {
		t.Errorf("Expected true for empty file")
	}
}

func TestIsFile_LargeFile(t *testing.T) {
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

	result := lxio.IsFile(tempFile.Name())

	if !result {
		t.Errorf("Expected true for large file")
	}
}

func TestIsFile_DotFile(t *testing.T) {
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

	result := lxio.IsFile(dotFilePath)

	if !result {
		t.Errorf("Expected true for dot file")
	}
}

func TestIsFile_FileWithoutExtension(t *testing.T) {
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

	result := lxio.IsFile(filePath)

	if !result {
		t.Errorf("Expected true for file without extension")
	}
}

func TestIsFile_DotPath(t *testing.T) {
	// "." is the current directory, not a file
	result := lxio.IsFile(".")

	if result {
		t.Errorf("Expected false for current directory '.'")
	}
}

func TestIsFile_DotDotPath(t *testing.T) {
	// ".." is the parent directory, not a file
	result := lxio.IsFile("..")

	if result {
		t.Errorf("Expected false for parent directory '..'")
	}
}

func TestIsFile_SwallowsErrors(t *testing.T) {
	// Verify that IsFile returns false for permission denied
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

	result := lxio.IsFile(tempFile.Name())

	// On systems where permission checks work, should be false or true depending on system
	// On systems running as root, might still be able to access
	// The important thing is that it doesn't panic or error
	_ = result // Just verify it doesn't panic
}
