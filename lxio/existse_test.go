package lxio_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/hgapdvn/lx/lxio"
)

func TestExistsE_FileExists(t *testing.T) {
	// Create a temporary file
	tempFile, err := os.CreateTemp("", "testfile-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	exists, err := lxio.ExistsE(tempFile.Name())

	if err != nil {
		t.Errorf("Expected no error for existing file, got: %v", err)
	}
	if !exists {
		t.Errorf("Expected true for existing file")
	}
}

func TestExistsE_DirectoryExists(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.Remove(tempDir)

	exists, err := lxio.ExistsE(tempDir)

	if err != nil {
		t.Errorf("Expected no error for existing directory, got: %v", err)
	}
	if !exists {
		t.Errorf("Expected true for existing directory")
	}
}

func TestExistsE_FileDoesNotExist(t *testing.T) {
	nonExistentPath := filepath.Join(os.TempDir(), "nonexistent-file-that-should-not-exist.txt")

	exists, err := lxio.ExistsE(nonExistentPath)

	if err != nil {
		t.Errorf("Expected no error for non-existent path, got: %v", err)
	}
	if exists {
		t.Errorf("Expected false for non-existent file")
	}
}

func TestExistsE_EmptyPath(t *testing.T) {
	exists, err := lxio.ExistsE("")

	if err != nil {
		t.Errorf("Expected no error for empty path, got: %v", err)
	}
	if exists {
		t.Errorf("Expected false for empty path")
	}
}

func TestExistsE_RelativePath_Exists(t *testing.T) {
	// Create a file in the current directory
	tempFile, err := os.CreateTemp(".", "testfile-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	filename := filepath.Base(tempFile.Name())
	defer os.Remove(filename)
	defer tempFile.Close()

	exists, err := lxio.ExistsE(filename)

	if err != nil {
		t.Errorf("Expected no error for relative path that exists, got: %v", err)
	}
	if !exists {
		t.Errorf("Expected true for relative path that exists")
	}
}

func TestExistsE_RelativePath_DoesNotExist(t *testing.T) {
	exists, err := lxio.ExistsE("nonexistent-relative-path.txt")

	if err != nil {
		t.Errorf("Expected no error for non-existent relative path, got: %v", err)
	}
	if exists {
		t.Errorf("Expected false for non-existent relative path")
	}
}

func TestExistsE_NestedPath_Exists(t *testing.T) {
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

	exists, err := lxio.ExistsE(nestedPath)

	if err != nil {
		t.Errorf("Expected no error for existing nested path, got: %v", err)
	}
	if !exists {
		t.Errorf("Expected true for existing nested path")
	}
}

func TestExistsE_NestedPath_DoesNotExist(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	nonExistentNestedPath := filepath.Join(tempDir, "nonexistent", "nested", "path")

	exists, err := lxio.ExistsE(nonExistentNestedPath)

	if err != nil {
		t.Errorf("Expected no error for non-existent nested path, got: %v", err)
	}
	if exists {
		t.Errorf("Expected false for non-existent nested path")
	}
}

func TestExistsE_Symlink_ToExistingFile(t *testing.T) {
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

	exists, err := lxio.ExistsE(symlinkPath)

	if err != nil {
		t.Errorf("Expected no error for symlink to existing file, got: %v", err)
	}
	if !exists {
		t.Errorf("Expected true for symlink to existing file")
	}
}

func TestExistsE_Symlink_ToNonExistentFile(t *testing.T) {
	// Create a symlink to non-existent file
	nonExistentTarget := filepath.Join(os.TempDir(), "nonexistent-target-for-symlink.txt")
	symlinkPath := filepath.Join(os.TempDir(), "broken-link")

	err := os.Symlink(nonExistentTarget, symlinkPath)
	if err != nil {
		t.Skip("Skipping symlink test - cannot create symlinks on this system")
	}
	defer os.Remove(symlinkPath)

	exists, err := lxio.ExistsE(symlinkPath)

	if err != nil {
		t.Errorf("Expected no error for broken symlink, got: %v", err)
	}
	if exists {
		t.Errorf("Expected false for broken symlink")
	}
}

func TestExistsE_SpecialPath_DevNull(t *testing.T) {
	exists, err := lxio.ExistsE("/dev/null")

	if err != nil {
		t.Errorf("Expected no error for /dev/null, got: %v", err)
	}
	if !exists {
		t.Errorf("Expected true for /dev/null")
	}
}

func TestExistsE_ReturnType_TrueNoError(t *testing.T) {
	// Test that when file exists, we get (true, nil)
	tempFile, err := os.CreateTemp("", "testfile-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	exists, err := lxio.ExistsE(tempFile.Name())

	if exists != true {
		t.Errorf("Expected exists to be true")
	}
	if err != nil {
		t.Errorf("Expected err to be nil")
	}
}

func TestExistsE_ReturnType_FalseNoError(t *testing.T) {
	// Test that when file doesn't exist, we get (false, nil)
	nonExistentPath := filepath.Join(os.TempDir(), "nonexistent-file-xyz.txt")

	exists, err := lxio.ExistsE(nonExistentPath)

	if exists != false {
		t.Errorf("Expected exists to be false")
	}
	if err != nil {
		t.Errorf("Expected err to be nil")
	}
}

func TestExistsE_ErrorType_IsErrNotExist(t *testing.T) {
	// When file doesn't exist, verify it returns os.ErrNotExist
	nonExistentPath := filepath.Join(os.TempDir(), "definitely-does-not-exist-12345.txt")

	_, err := lxio.ExistsE(nonExistentPath)

	// Should not return ErrNotExist since we handle it internally
	if err != nil {
		t.Errorf("Expected nil error for non-existent path, got: %v", err)
	}
}

func TestExistsE_DotPath(t *testing.T) {
	// "." represents the current directory
	exists, err := lxio.ExistsE(".")

	if err != nil {
		t.Errorf("Expected no error for '.' path, got: %v", err)
	}
	if !exists {
		t.Errorf("Expected true for '.' path (current directory)")
	}
}

func TestExistsE_DotDotPath(t *testing.T) {
	// ".." represents the parent directory
	exists, err := lxio.ExistsE("..")

	if err != nil {
		t.Errorf("Expected no error for '..' path, got: %v", err)
	}
	if !exists {
		t.Errorf("Expected true for '..' path (parent directory)")
	}
}

func TestExistsE_DistinguishesErrNotExist(t *testing.T) {
	// Verify that os.ErrNotExist is handled specially
	nonExistentPath := filepath.Join(os.TempDir(), "test-not-exist-path.txt")

	exists, err := lxio.ExistsE(nonExistentPath)

	// Should not error even though file doesn't exist
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if exists {
		t.Errorf("Expected false when file doesn't exist")
	}
}

func TestExistsE_PermissionDenied(t *testing.T) {
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

	exists, err := lxio.ExistsE(tempFile.Name())

	// On some systems (like running as root), permission checks are bypassed
	// In that case, we should see exists=true, err=nil
	// On other systems, we expect an error for permission denied (not os.ErrNotExist)
	if err != nil {
		// Expected path: permission denied error
		if exists {
			t.Errorf("Expected exists to be false when error is returned for permission denied")
		}
	} else {
		// Some systems allow reading file info even with 0000 permissions (e.g., running as root)
		// In this case, we just verify the function returns without error
		if !exists {
			t.Errorf("Expected exists to be true when no error is returned")
		}
	}
}

func TestExistsE_PermissionDeniedOnDirectory(t *testing.T) {
	// Create a temporary directory with a subdirectory
	tempDir, err := os.MkdirTemp("", "testdir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a subdirectory
	subDir := filepath.Join(tempDir, "subdir")
	err = os.Mkdir(subDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create subdirectory: %v", err)
	}

	// Create a file inside the subdirectory
	testFile := filepath.Join(subDir, "testfile.txt")
	err = os.WriteFile(testFile, []byte("test"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Remove read permissions from parent directory
	err = os.Chmod(subDir, 0000)
	if err != nil {
		t.Skip("Skipping permission test - cannot change directory permissions on this system")
	}
	defer os.Chmod(subDir, 0755) // Restore for cleanup

	exists, err := lxio.ExistsE(testFile)

	// We expect an error for permission denied
	if err == nil {
		t.Errorf("Expected error for permission denied on directory, got: exists=%v, err=nil", exists)
	}
}

func TestExistsE_SymlinkWithBrokenLink_ReturnsErrorPath(t *testing.T) {
	// Note: On most systems, broken symlinks still exist but point to non-existent targets
	// This tests the error handling path more broadly
	symlinkPath := filepath.Join(os.TempDir(), "test-symlink-error")

	// Create temp file to point to
	tempFile, err := os.CreateTemp("", "temp-target-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	tempFile.Close()

	// Create symlink to the temp file
	err = os.Symlink(tempFile.Name(), symlinkPath)
	if err != nil {
		t.Skip("Skipping symlink test - cannot create symlinks on this system")
	}
	defer os.Remove(symlinkPath)

	// Verify symlink exists
	exists, err := lxio.ExistsE(symlinkPath)
	if err != nil {
		t.Errorf("Expected no error for existing symlink, got: %v", err)
	}
	if !exists {
		t.Errorf("Expected true for existing symlink")
	}

	// Now delete the target file
	os.Remove(tempFile.Name())

	// Broken symlink still exists (stat follows links by default, so this will return os.ErrNotExist)
	exists, err = lxio.ExistsE(symlinkPath)

	if err != nil {
		t.Errorf("Expected no error for broken symlink, got: %v", err)
	}
	if exists {
		t.Errorf("Expected false for broken symlink with deleted target")
	}
}
