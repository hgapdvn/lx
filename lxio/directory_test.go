package lxio_test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"testing"

	"github.com/hgapdvn/lx/lxio"
	"github.com/hgapdvn/lx/lxslices"
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
		if runtime.GOOS == "windows" {
			t.Skip("skipping permission test on Windows")
		}
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
		defer func() { _ = os.Chmod(dir, 0755) }() // Restore for cleanup

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
		if runtime.GOOS == "windows" {
			t.Skip("skipping permission test on Windows")
		}
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
		defer func() { _ = os.Chmod(dir, 0755) }() // Restore for cleanup

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

func TestWalkDirs(t *testing.T) {
	t.Run("walk directories recursively", func(t *testing.T) {
		dir := t.TempDir()

		// Create structure:
		// dir/
		//   subdir/
		//     nested/

		_ = os.Mkdir(filepath.Join(dir, "subdir"), 0755)
		_ = os.Mkdir(filepath.Join(dir, "subdir", "nested"), 0755)

		var dirs []string
		err := lxio.WalkDirs(dir, func(path string) error {
			dirs = append(dirs, path)
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		sort.Strings(dirs)
		expected := []string{"subdir", "subdir/nested"}
		if !lxslices.Equal(dirs, expected) {
			t.Errorf("expected %v, got %v", expected, dirs)
		}
	})

	t.Run("walk directories from empty directory", func(t *testing.T) {
		dir := t.TempDir()

		var dirs []string
		err := lxio.WalkDirs(dir, func(path string) error {
			dirs = append(dirs, path)
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(dirs) != 0 {
			t.Errorf("expected empty slice, got %v", dirs)
		}
	})

	t.Run("walk directories ignores files", func(t *testing.T) {
		dir := t.TempDir()

		_ = os.WriteFile(filepath.Join(dir, "file1.txt"), []byte(""), 0644)
		_ = os.Mkdir(filepath.Join(dir, "subdir"), 0755)
		_ = os.WriteFile(filepath.Join(dir, "subdir", "file2.txt"), []byte(""), 0644)

		var dirs []string
		err := lxio.WalkDirs(dir, func(path string) error {
			dirs = append(dirs, path)
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := []string{"subdir"}
		if !lxslices.Equal(dirs, expected) {
			t.Errorf("expected %v, got %v", expected, dirs)
		}
	})

	t.Run("walk directories stops on error", func(t *testing.T) {
		dir := t.TempDir()

		_ = os.Mkdir(filepath.Join(dir, "dir1"), 0755)
		_ = os.Mkdir(filepath.Join(dir, "dir2"), 0755)

		callCount := 0
		testErr := os.ErrPermission
		err := lxio.WalkDirs(dir, func(path string) error {
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
		err := lxio.WalkDirs(nonExistentPath, func(path string) error {
			return nil
		})
		if err == nil {
			t.Error("expected error for non-existent directory, got nil")
		}
	})

	t.Run("walk directories with nested structure", func(t *testing.T) {
		dir := t.TempDir()

		// Create structure:
		// dir/
		//   a/
		//     b/
		//       c/
		//   d/
		//   e/
		//     f/

		_ = os.Mkdir(filepath.Join(dir, "a"), 0755)
		_ = os.Mkdir(filepath.Join(dir, "a", "b"), 0755)
		_ = os.Mkdir(filepath.Join(dir, "a", "b", "c"), 0755)
		_ = os.Mkdir(filepath.Join(dir, "d"), 0755)
		_ = os.Mkdir(filepath.Join(dir, "e"), 0755)
		_ = os.Mkdir(filepath.Join(dir, "e", "f"), 0755)

		var dirs []string
		err := lxio.WalkDirs(dir, func(path string) error {
			dirs = append(dirs, path)
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		sort.Strings(dirs)
		expected := []string{
			"a",
			"a/b",
			"a/b/c",
			"d",
			"e",
			"e/f",
		}
		if !lxslices.Equal(dirs, expected) {
			t.Errorf("expected %v, got %v", expected, dirs)
		}
	})

	t.Run("walk directories with mixed files and directories", func(t *testing.T) {
		dir := t.TempDir()

		// Create structure:
		// dir/
		//   file1.txt
		//   subdir1/
		//     file2.txt
		//     nested/
		//       file3.txt
		//   subdir2/
		//   file4.txt

		_ = os.WriteFile(filepath.Join(dir, "file1.txt"), []byte(""), 0644)
		_ = os.Mkdir(filepath.Join(dir, "subdir1"), 0755)
		_ = os.WriteFile(filepath.Join(dir, "subdir1", "file2.txt"), []byte(""), 0644)
		_ = os.Mkdir(filepath.Join(dir, "subdir1", "nested"), 0755)
		_ = os.WriteFile(filepath.Join(dir, "subdir1", "nested", "file3.txt"), []byte(""), 0644)
		_ = os.Mkdir(filepath.Join(dir, "subdir2"), 0755)
		_ = os.WriteFile(filepath.Join(dir, "file4.txt"), []byte(""), 0644)

		var dirs []string
		err := lxio.WalkDirs(dir, func(path string) error {
			dirs = append(dirs, path)
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		sort.Strings(dirs)
		expected := []string{
			"subdir1",
			"subdir1/nested",
			"subdir2",
		}
		if !lxslices.Equal(dirs, expected) {
			t.Errorf("expected %v, got %v", expected, dirs)
		}
	})

	t.Run("walk directories uses forward slashes", func(t *testing.T) {
		dir := t.TempDir()

		// Create nested directories
		_ = os.Mkdir(filepath.Join(dir, "a"), 0755)
		_ = os.Mkdir(filepath.Join(dir, "a", "b"), 0755)
		_ = os.Mkdir(filepath.Join(dir, "a", "b", "c"), 0755)

		var dirs []string
		err := lxio.WalkDirs(dir, func(path string) error {
			dirs = append(dirs, path)
			// Verify no backslashes (Windows separators)
			if strings.Contains(path, "\\") {
				return fmt.Errorf("path contains backslash: %s", path)
			}
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		sort.Strings(dirs)
		expected := []string{"a", "a/b", "a/b/c"}
		if !lxslices.Equal(dirs, expected) {
			t.Errorf("expected %v, got %v", expected, dirs)
		}
	})

	t.Run("walk directories only direct subdirectories at each level", func(t *testing.T) {
		dir := t.TempDir()

		// Create structure to verify we don't duplicate subdirs:
		// dir/
		//   sub/
		//     nested/

		_ = os.Mkdir(filepath.Join(dir, "sub"), 0755)
		_ = os.Mkdir(filepath.Join(dir, "sub", "nested"), 0755)

		callCount := 0
		dirCalls := make(map[string]int)

		err := lxio.WalkDirs(dir, func(path string) error {
			callCount++
			dirCalls[path]++
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if callCount != 2 {
			t.Errorf("expected 2 directory calls, got %d", callCount)
		}

		// Verify no duplicates
		for path, count := range dirCalls {
			if count > 1 {
				t.Errorf("directory %q was called %d times, expected 1", path, count)
			}
		}
	})
}

// ======================== ListFilesByExt Tests =========================

func TestListFilesByExt(t *testing.T) {
	t.Run("filter files by single extension", func(t *testing.T) {
		dir := t.TempDir()

		// Create test files
		_ = os.WriteFile(filepath.Join(dir, "readme.txt"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "config.txt"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "script.py"), []byte("content"), 0644)

		files, err := lxio.ListFilesByExt(dir, ".txt")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := []string{"config.txt", "readme.txt"}
		if !lxslices.Equal(files, expected) {
			t.Errorf("expected %v, got %v", expected, files)
		}
	})

	t.Run("filter files by multiple extensions", func(t *testing.T) {
		dir := t.TempDir()

		// Create test files
		_ = os.WriteFile(filepath.Join(dir, "doc.pdf"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "report.docx"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "notes.txt"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "image.png"), []byte("content"), 0644)

		files, err := lxio.ListFilesByExt(dir, ".pdf", ".docx", ".txt")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := []string{"doc.pdf", "notes.txt", "report.docx"}
		if !lxslices.Equal(files, expected) {
			t.Errorf("expected %v, got %v", expected, files)
		}
	})

	t.Run("case-insensitive extension matching", func(t *testing.T) {
		dir := t.TempDir()

		// Create test files with various cases
		_ = os.WriteFile(filepath.Join(dir, "document.PDF"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "report.Pdf"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "guide.pdf"), []byte("content"), 0644)

		files, err := lxio.ListFilesByExt(dir, ".pdf")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(files) != 3 {
			t.Errorf("expected 3 files, got %d", len(files))
		}
	})

	t.Run("no matching files returns empty slice", func(t *testing.T) {
		dir := t.TempDir()

		// Create test files
		_ = os.WriteFile(filepath.Join(dir, "script.py"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "code.go"), []byte("content"), 0644)

		files, err := lxio.ListFilesByExt(dir, ".pdf", ".docx")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(files) != 0 {
			t.Errorf("expected empty slice, got %v", files)
		}
	})

	t.Run("no extensions provided returns empty slice", func(t *testing.T) {
		dir := t.TempDir()

		// Create test files
		_ = os.WriteFile(filepath.Join(dir, "file.txt"), []byte("content"), 0644)

		files, err := lxio.ListFilesByExt(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(files) != 0 {
			t.Errorf("expected empty slice, got %v", files)
		}
	})
}

// ======================== Convenience Filter Function Tests =========================

func TestListPdfFiles(t *testing.T) {
	t.Run("list PDF files", func(t *testing.T) {
		dir := t.TempDir()

		// Create test files
		_ = os.WriteFile(filepath.Join(dir, "doc1.pdf"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "doc2.PDF"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "image.png"), []byte("content"), 0644)

		files, err := lxio.ListPdfFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(files) != 2 {
			t.Errorf("expected 2 PDF files, got %d", len(files))
		}
	})
}

// ======================== Document File Type Tests =========================

func TestListDocFiles(t *testing.T) {
	t.Run("list .doc files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "file.doc"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "file.docx"), []byte("content"), 0644)

		files, err := lxio.ListDocFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .doc file, got %d", len(files))
		}
	})
}

func TestListDocxFiles(t *testing.T) {
	t.Run("list .docx files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "file.doc"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "file.docx"), []byte("content"), 0644)

		files, err := lxio.ListDocxFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .docx file, got %d", len(files))
		}
	})
}

func TestListTxtFiles(t *testing.T) {
	t.Run("list .txt files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "readme.txt"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "notes.txt"), []byte("content"), 0644)

		files, err := lxio.ListTxtFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 2 {
			t.Errorf("expected 2 .txt files, got %d", len(files))
		}
	})
}

// ======================== Image File Type Tests =========================

func TestListImageFiles(t *testing.T) {
	t.Run("list image files", func(t *testing.T) {
		dir := t.TempDir()

		// Create test files
		_ = os.WriteFile(filepath.Join(dir, "photo.jpg"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "screenshot.png"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "graphic.svg"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "document.pdf"), []byte("content"), 0644)

		files, err := lxio.ListImageFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(files) != 3 {
			t.Errorf("expected 3 image files, got %d", len(files))
		}
	})
}

func TestListJpgFiles(t *testing.T) {
	t.Run("list .jpg files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "photo.jpg"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "photo.jpeg"), []byte("content"), 0644)

		files, err := lxio.ListJpgFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .jpg file, got %d", len(files))
		}
	})
}

func TestListPngFiles(t *testing.T) {
	t.Run("list .png files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "image.png"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "photo.jpg"), []byte("content"), 0644)

		files, err := lxio.ListPngFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .png file, got %d", len(files))
		}
	})
}

// ======================== Archive File Type Tests =========================

func TestListArchiveFiles(t *testing.T) {
	t.Run("list archive files", func(t *testing.T) {
		dir := t.TempDir()

		// Create test files
		_ = os.WriteFile(filepath.Join(dir, "archive.zip"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "backup.tar.gz"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "file.txt"), []byte("content"), 0644)

		files, err := lxio.ListArchiveFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(files) != 2 {
			t.Errorf("expected 2 archive files, got %d", len(files))
		}
	})
}

func TestListZipFiles(t *testing.T) {
	t.Run("list .zip files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "archive.zip"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "archive.rar"), []byte("content"), 0644)

		files, err := lxio.ListZipFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .zip file, got %d", len(files))
		}
	})
}

// ======================== Code File Type Tests =========================

func TestListCodeFiles(t *testing.T) {
	t.Run("list code files", func(t *testing.T) {
		dir := t.TempDir()

		// Create test files
		_ = os.WriteFile(filepath.Join(dir, "main.go"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "script.py"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "config.json"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "readme.txt"), []byte("content"), 0644)

		files, err := lxio.ListCodeFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(files) != 3 {
			t.Errorf("expected 3 code files, got %d", len(files))
		}
	})
}

func TestListGoFiles(t *testing.T) {
	t.Run("list .go files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "main.go"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "main.py"), []byte("content"), 0644)

		files, err := lxio.ListGoFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .go file, got %d", len(files))
		}
	})
}

func TestListPyFiles(t *testing.T) {
	t.Run("list .py files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "script.py"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "script.js"), []byte("content"), 0644)

		files, err := lxio.ListPyFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .py file, got %d", len(files))
		}
	})
}

func TestListJsonFiles(t *testing.T) {
	t.Run("list .json files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "config.json"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "package.json"), []byte("content"), 0644)

		files, err := lxio.ListJsonFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 2 {
			t.Errorf("expected 2 .json files, got %d", len(files))
		}
	})
}

// ======================== Audio File Type Tests =========================

func TestListAudioFiles(t *testing.T) {
	t.Run("list audio files", func(t *testing.T) {
		dir := t.TempDir()

		// Create test files
		_ = os.WriteFile(filepath.Join(dir, "song.mp3"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "track.flac"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "video.mp4"), []byte("content"), 0644)

		files, err := lxio.ListAudioFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(files) != 2 {
			t.Errorf("expected 2 audio files, got %d", len(files))
		}
	})
}

func TestListMp3Files(t *testing.T) {
	t.Run("list .mp3 files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "song.mp3"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "track.wav"), []byte("content"), 0644)

		files, err := lxio.ListMp3Files(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .mp3 file, got %d", len(files))
		}
	})
}

// ======================== Video File Type Tests =========================

func TestListVideoFiles(t *testing.T) {
	t.Run("list video files", func(t *testing.T) {
		dir := t.TempDir()

		// Create test files
		_ = os.WriteFile(filepath.Join(dir, "movie.mp4"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "film.mkv"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "clip.avi"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "song.mp3"), []byte("content"), 0644)

		files, err := lxio.ListVideoFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(files) != 3 {
			t.Errorf("expected 3 video files, got %d", len(files))
		}
	})
}

func TestListMp4Files(t *testing.T) {
	t.Run("list .mp4 files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "movie.mp4"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "movie.mkv"), []byte("content"), 0644)

		files, err := lxio.ListMp4Files(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .mp4 file, got %d", len(files))
		}
	})
}

// ======================== ListFilesFunc Tests =========================

func TestListFilesFunc(t *testing.T) {
	t.Run("filter files with predicate", func(t *testing.T) {
		dir := t.TempDir()

		// Create test files
		_ = os.WriteFile(filepath.Join(dir, "test_helper.go"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "test_main.go"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "main.go"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "config.json"), []byte("content"), 0644)

		// Filter files starting with "test_"
		files, err := lxio.ListFilesFunc(dir, func(name string) bool {
			return len(name) > 5 && name[:5] == "test_"
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := []string{"test_helper.go", "test_main.go"}
		if !lxslices.Equal(files, expected) {
			t.Errorf("expected %v, got %v", expected, files)
		}
	})

	t.Run("filter files by size using predicate", func(t *testing.T) {
		dir := t.TempDir()

		// Create test files with different names (length = size proxy)
		_ = os.WriteFile(filepath.Join(dir, "a.txt"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "bb.txt"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "ccccccc.txt"), []byte("content"), 0644)

		// Filter files with names longer than 6 characters
		files, err := lxio.ListFilesFunc(dir, func(name string) bool {
			return len(name) > 6
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(files) != 1 {
			t.Errorf("expected 1 file, got %d", len(files))
		}
	})
}

func TestListRtfFiles(t *testing.T) {
	t.Run("list .rtf files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "document.rtf"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "document.txt"), []byte("content"), 0644)

		files, err := lxio.ListRtfFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .rtf file, got %d", len(files))
		}
	})
}

func TestListOdtFiles(t *testing.T) {
	t.Run("list .odt files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "document.odt"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "document.docx"), []byte("content"), 0644)

		files, err := lxio.ListOdtFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .odt file, got %d", len(files))
		}
	})
}

func TestListJpegFiles(t *testing.T) {
	t.Run("list .jpeg files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "photo.jpeg"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "photo.jpg"), []byte("content"), 0644)

		files, err := lxio.ListJpegFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .jpeg file, got %d", len(files))
		}
	})
}

func TestListGifFiles(t *testing.T) {
	t.Run("list .gif files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "animation.gif"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "image.png"), []byte("content"), 0644)

		files, err := lxio.ListGifFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .gif file, got %d", len(files))
		}
	})
}

func TestListBmpFiles(t *testing.T) {
	t.Run("list .bmp files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "image.bmp"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "image.png"), []byte("content"), 0644)

		files, err := lxio.ListBmpFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .bmp file, got %d", len(files))
		}
	})
}

func TestListTiffFiles(t *testing.T) {
	t.Run("list .tiff files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "scan.tiff"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "photo.jpg"), []byte("content"), 0644)

		files, err := lxio.ListTiffFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .tiff file, got %d", len(files))
		}
	})
}

func TestListSvgFiles(t *testing.T) {
	t.Run("list .svg files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "logo.svg"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "logo.png"), []byte("content"), 0644)

		files, err := lxio.ListSvgFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .svg file, got %d", len(files))
		}
	})
}

func TestListWebpFiles(t *testing.T) {
	t.Run("list .webp files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "image.webp"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "image.png"), []byte("content"), 0644)

		files, err := lxio.ListWebpFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .webp file, got %d", len(files))
		}
	})
}

func TestListIcoFiles(t *testing.T) {
	t.Run("list .ico files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "favicon.ico"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "favicon.png"), []byte("content"), 0644)

		files, err := lxio.ListIcoFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .ico file, got %d", len(files))
		}
	})
}

func TestListRarFiles(t *testing.T) {
	t.Run("list .rar files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "archive.rar"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "archive.zip"), []byte("content"), 0644)

		files, err := lxio.ListRarFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .rar file, got %d", len(files))
		}
	})
}

func TestListTarFiles(t *testing.T) {
	t.Run("list .tar files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "archive.tar"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "archive.zip"), []byte("content"), 0644)

		files, err := lxio.ListTarFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .tar file, got %d", len(files))
		}
	})
}

func TestListJsFiles(t *testing.T) {
	t.Run("list .js files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "script.js"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "script.ts"), []byte("content"), 0644)

		files, err := lxio.ListJsFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .js file, got %d", len(files))
		}
	})
}

func TestListTsFiles(t *testing.T) {
	t.Run("list .ts files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "app.ts"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "app.js"), []byte("content"), 0644)

		files, err := lxio.ListTsFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .ts file, got %d", len(files))
		}
	})
}

func TestListJavaFiles(t *testing.T) {
	t.Run("list .java files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "Main.java"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "Main.cpp"), []byte("content"), 0644)

		files, err := lxio.ListJavaFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .java file, got %d", len(files))
		}
	})
}

func TestListCppFiles(t *testing.T) {
	t.Run("list .cpp files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "main.cpp"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "main.c"), []byte("content"), 0644)

		files, err := lxio.ListCppFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .cpp file, got %d", len(files))
		}
	})
}

func TestListCFiles(t *testing.T) {
	t.Run("list .c files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "program.c"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "program.cpp"), []byte("content"), 0644)

		files, err := lxio.ListCFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .c file, got %d", len(files))
		}
	})
}

func TestListHFiles(t *testing.T) {
	t.Run("list .h files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "header.h"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "header.cpp"), []byte("content"), 0644)

		files, err := lxio.ListHFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .h file, got %d", len(files))
		}
	})
}

func TestListRbFiles(t *testing.T) {
	t.Run("list .rb files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "script.rb"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "script.py"), []byte("content"), 0644)

		files, err := lxio.ListRbFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .rb file, got %d", len(files))
		}
	})
}

func TestListPhpFiles(t *testing.T) {
	t.Run("list .php files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "index.php"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "index.html"), []byte("content"), 0644)

		files, err := lxio.ListPhpFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .php file, got %d", len(files))
		}
	})
}

func TestListXmlFiles(t *testing.T) {
	t.Run("list .xml files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "data.xml"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "data.json"), []byte("content"), 0644)

		files, err := lxio.ListXmlFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .xml file, got %d", len(files))
		}
	})
}

func TestListYamlFiles(t *testing.T) {
	t.Run("list .yaml files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "config.yaml"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "config.yml"), []byte("content"), 0644)

		files, err := lxio.ListYamlFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .yaml file, got %d", len(files))
		}
	})
}

func TestListYmlFiles(t *testing.T) {
	t.Run("list .yml files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "config.yml"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "config.yaml"), []byte("content"), 0644)

		files, err := lxio.ListYmlFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .yml file, got %d", len(files))
		}
	})
}

func TestListCsvFiles(t *testing.T) {
	t.Run("list .csv files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "data.csv"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "data.json"), []byte("content"), 0644)

		files, err := lxio.ListCsvFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .csv file, got %d", len(files))
		}
	})
}

func TestListWavFiles(t *testing.T) {
	t.Run("list .wav files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "sound.wav"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "sound.mp3"), []byte("content"), 0644)

		files, err := lxio.ListWavFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .wav file, got %d", len(files))
		}
	})
}

func TestListFlacFiles(t *testing.T) {
	t.Run("list .flac files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "audio.flac"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "audio.mp3"), []byte("content"), 0644)

		files, err := lxio.ListFlacFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .flac file, got %d", len(files))
		}
	})
}

func TestListAacFiles(t *testing.T) {
	t.Run("list .aac files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "music.aac"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "music.mp3"), []byte("content"), 0644)

		files, err := lxio.ListAacFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .aac file, got %d", len(files))
		}
	})
}

func TestListOggFiles(t *testing.T) {
	t.Run("list .ogg files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "track.ogg"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "track.mp3"), []byte("content"), 0644)

		files, err := lxio.ListOggFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .ogg file, got %d", len(files))
		}
	})
}

func TestListM4aFiles(t *testing.T) {
	t.Run("list .m4a files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "song.m4a"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "song.mp3"), []byte("content"), 0644)

		files, err := lxio.ListM4aFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .m4a file, got %d", len(files))
		}
	})
}

func TestListAviFiles(t *testing.T) {
	t.Run("list .avi files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "video.avi"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "video.mp4"), []byte("content"), 0644)

		files, err := lxio.ListAviFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .avi file, got %d", len(files))
		}
	})
}

func TestListMkvFiles(t *testing.T) {
	t.Run("list .mkv files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "movie.mkv"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "movie.mp4"), []byte("content"), 0644)

		files, err := lxio.ListMkvFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .mkv file, got %d", len(files))
		}
	})
}

func TestListMovFiles(t *testing.T) {
	t.Run("list .mov files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "clip.mov"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "clip.mp4"), []byte("content"), 0644)

		files, err := lxio.ListMovFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .mov file, got %d", len(files))
		}
	})
}

func TestListWmvFiles(t *testing.T) {
	t.Run("list .wmv files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "video.wmv"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "video.mp4"), []byte("content"), 0644)

		files, err := lxio.ListWmvFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .wmv file, got %d", len(files))
		}
	})
}

func TestListFlvFiles(t *testing.T) {
	t.Run("list .flv files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "stream.flv"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "stream.mp4"), []byte("content"), 0644)

		files, err := lxio.ListFlvFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .flv file, got %d", len(files))
		}
	})
}

func TestListWebmFiles(t *testing.T) {
	t.Run("list .webm files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "video.webm"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "video.mp4"), []byte("content"), 0644)

		files, err := lxio.ListWebmFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .webm file, got %d", len(files))
		}
	})
}

func TestListPagesFiles(t *testing.T) {
	t.Run("list .pages files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "document.pages"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "document.docx"), []byte("content"), 0644)

		files, err := lxio.ListPagesFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .pages file, got %d", len(files))
		}
	})
}

func TestListGzFiles(t *testing.T) {
	t.Run("list .gz files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "archive.gz"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "archive.zip"), []byte("content"), 0644)

		files, err := lxio.ListGzFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .gz file, got %d", len(files))
		}
	})
}

func TestListBz2Files(t *testing.T) {
	t.Run("list .bz2 files", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.WriteFile(filepath.Join(dir, "archive.bz2"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "archive.zip"), []byte("content"), 0644)

		files, err := lxio.ListBz2Files(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(files) != 1 {
			t.Errorf("expected 1 .bz2 file, got %d", len(files))
		}
	})
}

func TestListTarGzFiles(t *testing.T) {
	t.Run("list .tar.gz files", func(t *testing.T) {
		dir := t.TempDir()
		// Create .tar.gz file with proper naming
		_ = os.WriteFile(filepath.Join(dir, "backup.tar.gz"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "backup.gz"), []byte("content"), 0644)

		files, err := lxio.ListTarGzFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		// Note: ListTarGzFiles only returns files with .tar.gz extension
		// The test verifies it can filter for .tar.gz properly
		if len(files) != 1 {
			t.Logf("files: %v", files)
			t.Errorf("expected 1 .tar.gz file, got %d", len(files))
		}
	})
}

// ======================== Category Combination Tests =========================

func TestListDocumentFiles_Combined(t *testing.T) {
	t.Run("list document files combined", func(t *testing.T) {
		dir := t.TempDir()

		// Create test files
		_ = os.WriteFile(filepath.Join(dir, "doc.pdf"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "report.docx"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "notes.txt"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "guide.rtf"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "sheet.xlsx"), []byte("content"), 0644)

		files, err := lxio.ListDocumentFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(files) != 4 {
			t.Errorf("expected 4 document files (pdf, docx, txt, rtf), got %d", len(files))
		}
	})
}

func TestListImageFiles_Combined(t *testing.T) {
	t.Run("list all image files combined", func(t *testing.T) {
		dir := t.TempDir()

		// Create test files
		_ = os.WriteFile(filepath.Join(dir, "photo.jpg"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "pic.png"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "logo.svg"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "icon.ico"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "document.pdf"), []byte("content"), 0644)

		files, err := lxio.ListImageFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(files) != 4 {
			t.Errorf("expected 4 image files, got %d", len(files))
		}
	})
}

func TestListArchiveFiles_Combined(t *testing.T) {
	t.Run("list all archive files combined", func(t *testing.T) {
		dir := t.TempDir()

		// Create test files
		_ = os.WriteFile(filepath.Join(dir, "file.zip"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "file.rar"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "file.tar"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "file.gz"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "backup.tar.gz"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "file.txt"), []byte("content"), 0644)

		files, err := lxio.ListArchiveFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(files) != 5 {
			t.Errorf("expected 5 archive files, got %d", len(files))
		}
	})
}

func TestListCodeFiles_Combined(t *testing.T) {
	t.Run("list all code files combined", func(t *testing.T) {
		dir := t.TempDir()

		// Create test files
		_ = os.WriteFile(filepath.Join(dir, "main.go"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "script.py"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "app.js"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "config.json"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "data.xml"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "readme.md"), []byte("content"), 0644)

		files, err := lxio.ListCodeFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(files) != 5 {
			t.Errorf("expected 5 code files, got %d", len(files))
		}
	})
}

func TestListAudioFiles_Combined(t *testing.T) {
	t.Run("list all audio files combined", func(t *testing.T) {
		dir := t.TempDir()

		// Create test files
		_ = os.WriteFile(filepath.Join(dir, "song.mp3"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "track.wav"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "audio.flac"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "clip.m4a"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "video.mp4"), []byte("content"), 0644)

		files, err := lxio.ListAudioFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(files) != 4 {
			t.Errorf("expected 4 audio files, got %d", len(files))
		}
	})
}

func TestListVideoFiles_Combined(t *testing.T) {
	t.Run("list all video files combined", func(t *testing.T) {
		dir := t.TempDir()

		// Create test files
		_ = os.WriteFile(filepath.Join(dir, "movie.mp4"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "film.mkv"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "clip.avi"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "video.mov"), []byte("content"), 0644)
		_ = os.WriteFile(filepath.Join(dir, "song.mp3"), []byte("content"), 0644)

		files, err := lxio.ListVideoFiles(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(files) != 4 {
			t.Errorf("expected 4 video files, got %d", len(files))
		}
	})
}
