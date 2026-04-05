package lxio_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/hgapdvn/lx/lxio"
)

func TestIsDir_DirectoryExists(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.Remove(tempDir)

	result := lxio.IsDir(tempDir)

	if !result {
		t.Errorf("Expected true for existing directory")
	}
}

func TestIsDir_FileExists(t *testing.T) {
	// Create a temporary file
	tempFile, err := os.CreateTemp("", "testfile-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	result := lxio.IsDir(tempFile.Name())

	if result {
		t.Errorf("Expected false for file (not a directory)")
	}
}

func TestIsDir_PathDoesNotExist(t *testing.T) {
	nonExistentPath := filepath.Join(os.TempDir(), "nonexistent-dir-xyz")

	result := lxio.IsDir(nonExistentPath)

	if result {
		t.Errorf("Expected false for non-existent path")
	}
}

func TestIsDir_EmptyPath(t *testing.T) {
	result := lxio.IsDir("")

	if result {
		t.Errorf("Expected false for empty path")
	}
}

func TestIsDir_RelativePath_Exists(t *testing.T) {
	// Create a directory in current directory
	tempDir, err := os.MkdirTemp(".", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	dirName := filepath.Base(tempDir)
	defer os.Remove(tempDir)

	result := lxio.IsDir(dirName)

	if !result {
		t.Errorf("Expected true for relative path that is a directory")
	}
}

func TestIsDir_RelativePath_DoesNotExist(t *testing.T) {
	result := lxio.IsDir("nonexistent-relative-dir")

	if result {
		t.Errorf("Expected false for non-existent relative path")
	}
}

func TestIsDir_NestedPath_Exists(t *testing.T) {
	// Create nested directories
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	nestedPath := filepath.Join(tempDir, "level1", "level2", "level3")
	err = os.MkdirAll(nestedPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create nested path: %v", err)
	}

	result := lxio.IsDir(nestedPath)

	if !result {
		t.Errorf("Expected true for existing nested directory")
	}
}

func TestIsDir_NestedPath_DoesNotExist(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	nonExistentPath := filepath.Join(tempDir, "nonexistent", "nested", "path")

	result := lxio.IsDir(nonExistentPath)

	if result {
		t.Errorf("Expected false for non-existent nested path")
	}
}

func TestIsDir_Symlink_ToDirectory(t *testing.T) {
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

	result := lxio.IsDir(symlinkPath)

	if !result {
		t.Errorf("Expected true for symlink to directory")
	}
}

func TestIsDir_Symlink_ToFile(t *testing.T) {
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

	result := lxio.IsDir(symlinkPath)

	if result {
		t.Errorf("Expected false for symlink to file")
	}
}

func TestIsDir_Symlink_BrokenLink(t *testing.T) {
	// Create a symlink to non-existent path
	nonExistentTarget := filepath.Join(os.TempDir(), "nonexistent-target-xyz")
	symlinkPath := filepath.Join(os.TempDir(), "broken-link")

	err := os.Symlink(nonExistentTarget, symlinkPath)
	if err != nil {
		t.Skip("Skipping symlink test - cannot create symlinks on this system")
	}
	defer os.Remove(symlinkPath)

	result := lxio.IsDir(symlinkPath)

	if result {
		t.Errorf("Expected false for broken symlink")
	}
}

func TestIsDir_CurrentDirectory(t *testing.T) {
	// "." is the current directory
	result := lxio.IsDir(".")

	if !result {
		t.Errorf("Expected true for current directory '.'")
	}
}

func TestIsDir_ParentDirectory(t *testing.T) {
	// ".." is the parent directory
	result := lxio.IsDir("..")

	if !result {
		t.Errorf("Expected true for parent directory '..'")
	}
}

func TestIsDir_DotHiddenDirectory(t *testing.T) {
	// Create a hidden directory (starting with dot)
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	hiddenDirPath := filepath.Join(tempDir, ".hidden")
	err = os.Mkdir(hiddenDirPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create hidden directory: %v", err)
	}

	result := lxio.IsDir(hiddenDirPath)

	if !result {
		t.Errorf("Expected true for hidden directory")
	}
}

func TestIsDir_EmptyDirectory(t *testing.T) {
	// Create an empty directory
	tempDir, err := os.MkdirTemp("", "empty-dir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.Remove(tempDir)

	result := lxio.IsDir(tempDir)

	if !result {
		t.Errorf("Expected true for empty directory")
	}
}

func TestIsDir_DirectoryWithFiles(t *testing.T) {
	// Create a directory with files
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create files in the directory
	for i := 0; i < 5; i++ {
		filePath := filepath.Join(tempDir, "file"+string(rune('0'+i))+".txt")
		err := os.WriteFile(filePath, []byte("test"), 0644)
		if err != nil {
			t.Fatalf("Failed to create file: %v", err)
		}
	}

	result := lxio.IsDir(tempDir)

	if !result {
		t.Errorf("Expected true for directory with files")
	}
}

func TestIsDir_DirectoryWithSubdirectories(t *testing.T) {
	// Create a directory with subdirectories
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create subdirectories
	for i := 0; i < 3; i++ {
		subPath := filepath.Join(tempDir, "subdir"+string(rune('0'+i)))
		err := os.Mkdir(subPath, 0755)
		if err != nil {
			t.Fatalf("Failed to create subdirectory: %v", err)
		}
	}

	result := lxio.IsDir(tempDir)

	if !result {
		t.Errorf("Expected true for directory with subdirectories")
	}
}

func TestIsDir_SwallowsErrors(t *testing.T) {
	// Create a directory with restricted permissions
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Change permissions to remove read access
	err = os.Chmod(tempDir, 0000)
	if err != nil {
		t.Skip("Skipping permission test - cannot change directory permissions on this system")
	}
	defer os.Chmod(tempDir, 0755) // Restore for cleanup

	result := lxio.IsDir(tempDir)

	// The function should return false and not panic
	_ = result // Just verify it doesn't panic
}
