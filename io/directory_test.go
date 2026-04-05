package lxio_test

import (
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/hgapdvn/lx/io"
	"github.com/hgapdvn/lx/slices"
)

func TestListFiles(t *testing.T) {
	t.Run("list files from directory with mixed content", func(t *testing.T) {
		dir := t.TempDir()

		// Create test files
		_ = os.WriteFile(filepath.Join(dir, "file1.txt"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "file2.txt"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "file3.go"), []byte("content"), 0644)

		// Create subdirectory (should be ignored)
		_ = os.Mkdir(filepath.Join(dir, "subdir"), 0755)

		files, err := lxio.ListFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := []string{"file1.txt", "file2.txt", "file3.go"}
		if !lxslices.Equal(files, expected) {
			t.Errorf("expected %v, got %v", expected, files)
		}
	})

	t.Run("list files from empty directory", func(t *testing.T) {
		dir := t.TempDir()

		files, err := lxio.ListFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(files) != 0 {
			t.Errorf("expected empty slice, got %v", files)
		}
	})

	t.Run("list files ignores subdirectories", func(t *testing.T) {
		dir := t.TempDir()

		_ = os.WriteFile(filepath.Join(dir, "file.txt"), []byte("content"), 0644)
		_ = os.Mkdir(filepath.Join(dir, "dir1"), 0755)
		_ = os.Mkdir(filepath.Join(dir, "dir2"), 0755)

		files, err := lxio.ListFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := []string{"file.txt"}
		if !lxslices.Equal(files, expected) {
			t.Errorf("expected %v, got %v", expected, files)
		}
	})

	t.Run("list files returns sorted results", func(t *testing.T) {
		dir := t.TempDir()

		// Create files in non-alphabetical order
		_ = os.WriteFile(filepath.Join(dir, "zebra.txt"), []byte(""), 0644)
		_ = os.WriteFile(filepath.Join(dir, "apple.txt"), []byte(""), 0644)
		_ = os.WriteFile(filepath.Join(dir, "mango.txt"), []byte(""), 0644)

		files, err := lxio.ListFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := []string{"apple.txt", "mango.txt", "zebra.txt"}
		if !lxslices.Equal(files, expected) {
			t.Errorf("expected sorted %v, got %v", expected, files)
		}
	})

	t.Run("list files with various file extensions", func(t *testing.T) {
		dir := t.TempDir()

		_ = os.WriteFile(filepath.Join(dir, "file.go"), []byte(""), 0644)
		_ = os.WriteFile(filepath.Join(dir, "file.txt"), []byte(""), 0644)
		_ = os.WriteFile(filepath.Join(dir, "file.json"), []byte(""), 0644)
		_ = os.WriteFile(filepath.Join(dir, "file.yaml"), []byte(""), 0644)
		_ = os.WriteFile(filepath.Join(dir, "README"), []byte(""), 0644)

		files, err := lxio.ListFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(files) != 5 {
			t.Errorf("expected 5 files, got %d", len(files))
		}
	})

	t.Run("error on non-existent directory", func(t *testing.T) {
		nonExistentPath := "/nonexistent/path/that/does/not/exist"
		files, err := lxio.ListFiles(nonExistentPath)
		if err == nil {
			t.Error("expected error for non-existent directory, got nil")
		}
		if files != nil {
			t.Errorf("expected nil slice on error, got %v", files)
		}
	})

	t.Run("error on permission denied", func(t *testing.T) {
		if os.Geteuid() == 0 {
			t.Skip("skipping permission test when running as root")
		}

		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "file.txt"), []byte(""), 0644)

		// Remove read permission
		err := os.Chmod(dir, 0000)
		if err != nil {
			t.Skipf("unable to remove permissions: %v", err)
		}
		defer os.Chmod(dir, 0755) // Restore for cleanup

		_, err = lxio.ListFiles(dir)
		if err == nil {
			t.Error("expected error on permission denied, got nil")
		}
	})
}

func TestListDirs(t *testing.T) {
	t.Run("list directories from directory with mixed content", func(t *testing.T) {
		dir := t.TempDir()

		_ = os.Mkdir(filepath.Join(dir, "dir1"), 0755)
		_ = os.Mkdir(filepath.Join(dir, "dir2"), 0755)
		_ = os.Mkdir(filepath.Join(dir, "dir3"), 0755)

		// Create files (should be ignored)
		_ = os.WriteFile(filepath.Join(dir, "file.txt"), []byte(""), 0644)

		dirs, err := lxio.ListDirs(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := []string{"dir1", "dir2", "dir3"}
		if !lxslices.Equal(dirs, expected) {
			t.Errorf("expected %v, got %v", expected, dirs)
		}
	})

	t.Run("list directories from empty directory", func(t *testing.T) {
		dir := t.TempDir()

		dirs, err := lxio.ListDirs(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(dirs) != 0 {
			t.Errorf("expected empty slice, got %v", dirs)
		}
	})

	t.Run("list directories ignores files", func(t *testing.T) {
		dir := t.TempDir()

		_ = os.Mkdir(filepath.Join(dir, "dir1"), 0755)
		_ = os.WriteFile(filepath.Join(dir, "file1.txt"), []byte(""), 0644)
		_ = os.WriteFile(filepath.Join(dir, "file2.txt"), []byte(""), 0644)

		dirs, err := lxio.ListDirs(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := []string{"dir1"}
		if !lxslices.Equal(dirs, expected) {
			t.Errorf("expected %v, got %v", expected, dirs)
		}
	})

	t.Run("list directories returns sorted results", func(t *testing.T) {
		dir := t.TempDir()

		_ = os.Mkdir(filepath.Join(dir, "zebra"), 0755)
		_ = os.Mkdir(filepath.Join(dir, "apple"), 0755)
		_ = os.Mkdir(filepath.Join(dir, "mango"), 0755)

		dirs, err := lxio.ListDirs(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := []string{"apple", "mango", "zebra"}
		if !lxslices.Equal(dirs, expected) {
			t.Errorf("expected sorted %v, got %v", expected, dirs)
		}
	})

	t.Run("error on non-existent directory", func(t *testing.T) {
		nonExistentPath := "/nonexistent/path/that/does/not/exist"
		dirs, err := lxio.ListDirs(nonExistentPath)
		if err == nil {
			t.Error("expected error for non-existent directory, got nil")
		}
		if dirs != nil {
			t.Errorf("expected nil slice on error, got %v", dirs)
		}
	})

	t.Run("error on permission denied", func(t *testing.T) {
		if os.Geteuid() == 0 {
			t.Skip("skipping permission test when running as root")
		}

		dir := t.TempDir()
		_ = os.Mkdir(filepath.Join(dir, "subdir"), 0755)

		// Remove read permission
		err := os.Chmod(dir, 0000)
		if err != nil {
			t.Skipf("unable to remove permissions: %v", err)
		}
		defer os.Chmod(dir, 0755) // Restore for cleanup

		_, err = lxio.ListDirs(dir)
		if err == nil {
			t.Error("expected error on permission denied, got nil")
		}
	})
}

func TestListAll(t *testing.T) {
	t.Run("list all entries from directory", func(t *testing.T) {
		dir := t.TempDir()

		_ = os.WriteFile(filepath.Join(dir, "file1.txt"), []byte(""), 0644)
		_ = os.WriteFile(filepath.Join(dir, "file2.txt"), []byte(""), 0644)
		_ = os.Mkdir(filepath.Join(dir, "dir1"), 0755)
		_ = os.Mkdir(filepath.Join(dir, "dir2"), 0755)

		entries, err := lxio.ListAll(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := []string{"dir1", "dir2", "file1.txt", "file2.txt"}
		if !lxslices.Equal(entries, expected) {
			t.Errorf("expected %v, got %v", expected, entries)
		}
	})

	t.Run("list all entries from empty directory", func(t *testing.T) {
		dir := t.TempDir()

		entries, err := lxio.ListAll(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(entries) != 0 {
			t.Errorf("expected empty slice, got %v", entries)
		}
	})

	t.Run("list all entries returns sorted results", func(t *testing.T) {
		dir := t.TempDir()

		_ = os.WriteFile(filepath.Join(dir, "zebra.txt"), []byte(""), 0644)
		_ = os.Mkdir(filepath.Join(dir, "apple"), 0755)
		_ = os.WriteFile(filepath.Join(dir, "mango.txt"), []byte(""), 0644)
		_ = os.Mkdir(filepath.Join(dir, "banana"), 0755)

		entries, err := lxio.ListAll(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := []string{"apple", "banana", "mango.txt", "zebra.txt"}
		if !lxslices.Equal(entries, expected) {
			t.Errorf("expected sorted %v, got %v", expected, entries)
		}
	})

	t.Run("error on non-existent directory", func(t *testing.T) {
		nonExistentPath := "/nonexistent/path/that/does/not/exist"
		entries, err := lxio.ListAll(nonExistentPath)
		if err == nil {
			t.Error("expected error for non-existent directory, got nil")
		}
		if entries != nil {
			t.Errorf("expected nil slice on error, got %v", entries)
		}
	})
}

func TestWalkFiles(t *testing.T) {
	t.Run("walk files recursively", func(t *testing.T) {
		dir := t.TempDir()

		// Create structure:
		// dir/
		//   file1.txt
		//   subdir/
		//     file2.txt
		//     nested/
		//       file3.txt

		_ = os.WriteFile(filepath.Join(dir, "file1.txt"), []byte(""), 0644)
		_ = os.Mkdir(filepath.Join(dir, "subdir"), 0755)
		_ = os.WriteFile(filepath.Join(dir, "subdir", "file2.txt"), []byte(""), 0644)
		_ = os.Mkdir(filepath.Join(dir, "subdir", "nested"), 0755)
		_ = os.WriteFile(filepath.Join(dir, "subdir", "nested", "file3.txt"), []byte(""), 0644)

		var files []string
		err := lxio.WalkFiles(dir, func(path string) error {
			files = append(files, path)
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		sort.Strings(files)
		expected := []string{"file1.txt", "subdir/file2.txt", "subdir/nested/file3.txt"}
		if !lxslices.Equal(files, expected) {
			t.Errorf("expected %v, got %v", expected, files)
		}
	})

	t.Run("walk files from empty directory", func(t *testing.T) {
		dir := t.TempDir()

		var files []string
		err := lxio.WalkFiles(dir, func(path string) error {
			files = append(files, path)
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(files) != 0 {
			t.Errorf("expected empty slice, got %v", files)
		}
	})

	t.Run("walk files ignores directories", func(t *testing.T) {
		dir := t.TempDir()

		_ = os.WriteFile(filepath.Join(dir, "file1.txt"), []byte(""), 0644)
		_ = os.Mkdir(filepath.Join(dir, "emptydir"), 0755)

		var files []string
		err := lxio.WalkFiles(dir, func(path string) error {
			files = append(files, path)
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := []string{"file1.txt"}
		if !lxslices.Equal(files, expected) {
			t.Errorf("expected %v, got %v", expected, files)
		}
	})

	t.Run("walk files stops on error", func(t *testing.T) {
		dir := t.TempDir()

		_ = os.WriteFile(filepath.Join(dir, "file1.txt"), []byte(""), 0644)
		_ = os.WriteFile(filepath.Join(dir, "file2.txt"), []byte(""), 0644)

		callCount := 0
		testErr := os.ErrPermission
		err := lxio.WalkFiles(dir, func(path string) error {
			callCount++
			if callCount == 1 {
				return testErr
			}
			return nil
		})

		if err != testErr {
			t.Errorf("expected error %v, got %v", testErr, err)
		}
		if callCount != 1 {
			t.Errorf("expected 1 call before stop, got %d", callCount)
		}
	})

	t.Run("error on non-existent directory", func(t *testing.T) {
		nonExistentPath := "/nonexistent/path/that/does/not/exist"
		err := lxio.WalkFiles(nonExistentPath, func(path string) error {
			return nil
		})
		if err == nil {
			t.Error("expected error for non-existent directory, got nil")
		}
	})

	t.Run("walk files with nested directories", func(t *testing.T) {
		dir := t.TempDir()

		// Create structure:
		// dir/
		//   a.txt
		//   sub1/
		//     b.txt
		//     sub2/
		//       c.txt
		//       d.txt
		//     e.txt
		//   f.txt

		_ = os.WriteFile(filepath.Join(dir, "a.txt"), []byte(""), 0644)
		_ = os.Mkdir(filepath.Join(dir, "sub1"), 0755)
		_ = os.WriteFile(filepath.Join(dir, "sub1", "b.txt"), []byte(""), 0644)
		_ = os.Mkdir(filepath.Join(dir, "sub1", "sub2"), 0755)
		_ = os.WriteFile(filepath.Join(dir, "sub1", "sub2", "c.txt"), []byte(""), 0644)
		_ = os.WriteFile(filepath.Join(dir, "sub1", "sub2", "d.txt"), []byte(""), 0644)
		_ = os.WriteFile(filepath.Join(dir, "sub1", "e.txt"), []byte(""), 0644)
		_ = os.WriteFile(filepath.Join(dir, "f.txt"), []byte(""), 0644)

		var files []string
		err := lxio.WalkFiles(dir, func(path string) error {
			files = append(files, path)
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		sort.Strings(files)
		expected := []string{
			"a.txt",
			"f.txt",
			"sub1/b.txt",
			"sub1/e.txt",
			"sub1/sub2/c.txt",
			"sub1/sub2/d.txt",
		}
		if !lxslices.Equal(files, expected) {
			t.Errorf("expected %v, got %v", expected, files)
		}
	})
}
