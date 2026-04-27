package lxio_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/hgapdvn/lx/lxio"
)

func TestCopyFile(t *testing.T) {
	tests := []struct {
		name     string
		srcSetup func(t *testing.T, dir string) string
		dstPath  func(dir string) string
		wantErr  bool
		checkFn  func(t *testing.T, src, dst string) bool
	}{
		{
			name: "copy simple file",
			srcSetup: func(t *testing.T, dir string) string {
				path := filepath.Join(dir, "source.txt")
				_ = os.WriteFile(path, []byte("hello world"), 0644)
				return path
			},
			dstPath: func(dir string) string { return filepath.Join(dir, "dest.txt") },
			wantErr: false,
			checkFn: func(t *testing.T, src, dst string) bool {
				srcContent, _ := os.ReadFile(src)
				dstContent, _ := os.ReadFile(dst)
				return string(srcContent) == string(dstContent)
			},
		},
		{
			name: "copy overwrites existing file",
			srcSetup: func(t *testing.T, dir string) string {
				path := filepath.Join(dir, "source.txt")
				_ = os.WriteFile(path, []byte("new content"), 0644)
				return path
			},
			dstPath: func(dir string) string {
				path := filepath.Join(dir, "dest.txt")
				_ = os.WriteFile(path, []byte("old content"), 0644)
				return path
			},
			wantErr: false,
			checkFn: func(t *testing.T, src, dst string) bool {
				content, _ := os.ReadFile(dst)
				return string(content) == "new content"
			},
		},
		{
			name: "copy nonexistent source returns error",
			srcSetup: func(t *testing.T, dir string) string {
				return filepath.Join(dir, "nonexistent.txt")
			},
			dstPath: func(dir string) string { return filepath.Join(dir, "dest.txt") },
			wantErr: true,
		},
		{
			name: "copy to invalid destination directory returns error",
			srcSetup: func(t *testing.T, dir string) string {
				path := filepath.Join(dir, "source.txt")
				_ = os.WriteFile(path, []byte("content"), 0644)
				return path
			},
			dstPath: func(dir string) string {
				return filepath.Join(dir, "nonexistent_dir", "file.txt")
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			src := tt.srcSetup(t, dir)
			dst := tt.dstPath(dir)

			err := lxio.CopyFile(src, dst)
			if (err != nil) != tt.wantErr {
				t.Errorf("CopyFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tt.checkFn != nil && !tt.checkFn(t, src, dst) {
				t.Errorf("CopyFile() result check failed")
			}
		})
	}
}

func TestMoveFile(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(t *testing.T, dir string) (string, string)
		wantErr bool
		checkFn func(t *testing.T, src, dst string) bool
	}{
		{
			name: "move simple file",
			setup: func(t *testing.T, dir string) (string, string) {
				src := filepath.Join(dir, "source.txt")
				dst := filepath.Join(dir, "dest.txt")
				_ = os.WriteFile(src, []byte("content"), 0644)
				return src, dst
			},
			wantErr: false,
			checkFn: func(t *testing.T, src, dst string) bool {
				// Source should not exist after move
				_, errSrc := os.Stat(src)
				// Destination should exist
				_, errDst := os.Stat(dst)
				return os.IsNotExist(errSrc) && errDst == nil
			},
		},
		{
			name: "move to different directory",
			setup: func(t *testing.T, dir string) (string, string) {
				subdir := filepath.Join(dir, "subdir")
				_ = os.Mkdir(subdir, 0755)
				src := filepath.Join(dir, "file.txt")
				dst := filepath.Join(subdir, "file.txt")
				_ = os.WriteFile(src, []byte("content"), 0644)
				return src, dst
			},
			wantErr: false,
			checkFn: func(t *testing.T, src, dst string) bool {
				_, errSrc := os.Stat(src)
				_, errDst := os.Stat(dst)
				return os.IsNotExist(errSrc) && errDst == nil
			},
		},
		{
			name: "move nonexistent source returns error",
			setup: func(t *testing.T, dir string) (string, string) {
				return filepath.Join(dir, "nonexistent.txt"), filepath.Join(dir, "dest.txt")
			},
			wantErr: true,
		},
		{
			name: "move to invalid destination directory returns error",
			setup: func(t *testing.T, dir string) (string, string) {
				src := filepath.Join(dir, "source.txt")
				_ = os.WriteFile(src, []byte("content"), 0644)
				dst := filepath.Join(dir, "nonexistent_dir", "file.txt")
				return src, dst
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			src, dst := tt.setup(t, dir)

			err := lxio.MoveFile(src, dst)
			if (err != nil) != tt.wantErr {
				t.Errorf("MoveFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tt.checkFn != nil && !tt.checkFn(t, src, dst) {
				t.Errorf("MoveFile() result check failed")
			}
		})
	}
}

func TestRemoveFile(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(t *testing.T, dir string) string
		wantErr bool
		checkFn func(t *testing.T, path string) bool
	}{
		{
			name: "remove existing file",
			setup: func(t *testing.T, dir string) string {
				path := filepath.Join(dir, "file.txt")
				_ = os.WriteFile(path, []byte("content"), 0644)
				return path
			},
			wantErr: false,
			checkFn: func(t *testing.T, path string) bool {
				_, err := os.Stat(path)
				return os.IsNotExist(err)
			},
		},
		{
			name: "remove nonexistent file returns error",
			setup: func(t *testing.T, dir string) string {
				return filepath.Join(dir, "nonexistent.txt")
			},
			wantErr: true,
		},
		{
			name: "remove file with content",
			setup: func(t *testing.T, dir string) string {
				path := filepath.Join(dir, "large.bin")
				_ = os.WriteFile(path, make([]byte, 10000), 0644)
				return path
			},
			wantErr: false,
			checkFn: func(t *testing.T, path string) bool {
				_, err := os.Stat(path)
				return os.IsNotExist(err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			path := tt.setup(t, dir)

			err := lxio.RemoveFile(path)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tt.checkFn != nil && !tt.checkFn(t, path) {
				t.Errorf("RemoveFile() result check failed")
			}
		})
	}
}

func TestCreateDir(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(t *testing.T, dir string) (string, os.FileMode)
		wantErr bool
		checkFn func(t *testing.T, path string) bool
	}{
		{
			name: "create simple directory",
			setup: func(t *testing.T, dir string) (string, os.FileMode) {
				return filepath.Join(dir, "newdir"), 0755
			},
			wantErr: false,
			checkFn: func(t *testing.T, path string) bool {
				info, err := os.Stat(path)
				return err == nil && info.IsDir()
			},
		},
		{
			name: "create directory with nested parents auto-created",
			setup: func(t *testing.T, dir string) (string, os.FileMode) {
				return filepath.Join(dir, "a", "b", "c"), 0755
			},
			wantErr: false,
			checkFn: func(t *testing.T, path string) bool {
				info, err := os.Stat(path)
				return err == nil && info.IsDir()
			},
		},
		{
			name: "create directory that already exists succeeds",
			setup: func(t *testing.T, dir string) (string, os.FileMode) {
				path := filepath.Join(dir, "existing")
				_ = os.Mkdir(path, 0755)
				return path, 0755
			},
			wantErr: false,
			checkFn: func(t *testing.T, path string) bool {
				info, err := os.Stat(path)
				return err == nil && info.IsDir()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			path, perm := tt.setup(t, dir)

			err := lxio.CreateDir(path, perm)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tt.checkFn != nil && !tt.checkFn(t, path) {
				t.Errorf("CreateDir() result check failed")
			}
		})
	}
}

func TestCreateFile(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(t *testing.T, dir string) string
		wantErr bool
		checkFn func(t *testing.T, path string) bool
	}{
		{
			name: "create simple file",
			setup: func(t *testing.T, dir string) string {
				return filepath.Join(dir, "file.txt")
			},
			wantErr: false,
			checkFn: func(t *testing.T, path string) bool {
				info, err := os.Stat(path)
				return err == nil && !info.IsDir()
			},
		},
		{
			name: "create file with one level of nested directories",
			setup: func(t *testing.T, dir string) string {
				return filepath.Join(dir, "subdir", "file.txt")
			},
			wantErr: false,
			checkFn: func(t *testing.T, path string) bool {
				_, err := os.Stat(path)
				return err == nil
			},
		},
		{
			name: "create file with deeply nested directories",
			setup: func(t *testing.T, dir string) string {
				return filepath.Join(dir, "a", "b", "c", "d", "e", "file.txt")
			},
			wantErr: false,
			checkFn: func(t *testing.T, path string) bool {
				info, err := os.Stat(path)
				return err == nil && !info.IsDir() && info.Size() == 0
			},
		},
		{
			name: "create file truncates existing file",
			setup: func(t *testing.T, dir string) string {
				path := filepath.Join(dir, "file.txt")
				_ = os.WriteFile(path, []byte("old content"), 0644)
				return path
			},
			wantErr: false,
			checkFn: func(t *testing.T, path string) bool {
				info, err := os.Stat(path)
				return err == nil && info.Size() == 0
			},
		},
		{
			name: "create file with custom permissions",
			setup: func(t *testing.T, dir string) string {
				return filepath.Join(dir, "perm_file.txt")
			},
			wantErr: false,
			checkFn: func(t *testing.T, path string) bool {
				info, err := os.Stat(path)
				if err != nil {
					return false
				}
				if runtime.GOOS == "windows" {
					// Windows doesn't support Unix permission bits
					return !info.IsDir()
				}
				// Check permissions (mode mask)
				return (info.Mode() & 0777) == 0644
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			path := tt.setup(t, dir)

			err := lxio.CreateFile(path, 0644)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tt.checkFn != nil && !tt.checkFn(t, path) {
				t.Errorf("CreateFile() result check failed")
			}
		})
	}
}

// ========================= RemoveAll Tests =========================

func TestRemoveAll(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(t *testing.T, dir string) string
		wantErr bool
		checkFn func(t *testing.T, path string) bool
	}{
		{
			name: "remove empty directory",
			setup: func(t *testing.T, dir string) string {
				path := filepath.Join(dir, "empty")
				_ = os.Mkdir(path, 0755)
				return path
			},
			wantErr: false,
			checkFn: func(t *testing.T, path string) bool {
				_, err := os.Stat(path)
				return os.IsNotExist(err)
			},
		},
		{
			name: "remove directory with files",
			setup: func(t *testing.T, dir string) string {
				path := filepath.Join(dir, "withfiles")
				_ = os.Mkdir(path, 0755)
				_ = os.WriteFile(filepath.Join(path, "file1.txt"), []byte("content"), 0644)
				_ = os.WriteFile(filepath.Join(path, "file2.txt"), []byte("content"), 0644)
				return path
			},
			wantErr: false,
			checkFn: func(t *testing.T, path string) bool {
				_, err := os.Stat(path)
				return os.IsNotExist(err)
			},
		},
		{
			name: "remove nested directories recursively",
			setup: func(t *testing.T, dir string) string {
				path := filepath.Join(dir, "a", "b", "c")
				_ = os.MkdirAll(path, 0755)
				_ = os.WriteFile(filepath.Join(path, "file.txt"), []byte("content"), 0644)
				return filepath.Join(dir, "a")
			},
			wantErr: false,
			checkFn: func(t *testing.T, path string) bool {
				_, err := os.Stat(path)
				return os.IsNotExist(err)
			},
		},
		{
			name: "remove single file",
			setup: func(t *testing.T, dir string) string {
				path := filepath.Join(dir, "file.txt")
				_ = os.WriteFile(path, []byte("content"), 0644)
				return path
			},
			wantErr: false,
			checkFn: func(t *testing.T, path string) bool {
				_, err := os.Stat(path)
				return os.IsNotExist(err)
			},
		},
		{
			name: "remove nonexistent path (no error)",
			setup: func(t *testing.T, dir string) string {
				return filepath.Join(dir, "nonexistent")
			},
			wantErr: false,
			checkFn: func(t *testing.T, path string) bool {
				_, err := os.Stat(path)
				return os.IsNotExist(err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			path := tt.setup(t, dir)

			err := lxio.RemoveAll(path)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tt.checkFn != nil && !tt.checkFn(t, path) {
				t.Errorf("RemoveAll() result check failed")
			}
		})
	}
}

// ========================= RemoveIfExists Tests =========================

func TestRemoveIfExists(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(t *testing.T, dir string) string
		wantErr bool
		checkFn func(t *testing.T, path string) bool
	}{
		{
			name: "remove existing file",
			setup: func(t *testing.T, dir string) string {
				path := filepath.Join(dir, "file.txt")
				_ = os.WriteFile(path, []byte("content"), 0644)
				return path
			},
			wantErr: false,
			checkFn: func(t *testing.T, path string) bool {
				_, err := os.Stat(path)
				return os.IsNotExist(err)
			},
		},
		{
			name: "remove nonexistent file (no error)",
			setup: func(t *testing.T, dir string) string {
				return filepath.Join(dir, "nonexistent.txt")
			},
			wantErr: false,
			checkFn: func(t *testing.T, path string) bool {
				_, err := os.Stat(path)
				return os.IsNotExist(err)
			},
		},
		{
			name: "remove existing directory recursively",
			setup: func(t *testing.T, dir string) string {
				path := filepath.Join(dir, "mydir")
				_ = os.Mkdir(path, 0755)
				_ = os.WriteFile(filepath.Join(path, "file.txt"), []byte("content"), 0644)
				return path
			},
			wantErr: false,
			checkFn: func(t *testing.T, path string) bool {
				_, err := os.Stat(path)
				return os.IsNotExist(err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			path := tt.setup(t, dir)

			err := lxio.RemoveIfExists(path)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveIfExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tt.checkFn != nil && !tt.checkFn(t, path) {
				t.Errorf("RemoveIfExists() result check failed")
			}
		})
	}
}

// ========================= CopyDir Tests =========================

func TestCopyDir(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(t *testing.T, dir string) (string, string)
		wantErr bool
		checkFn func(t *testing.T, src, dst string) bool
	}{
		{
			name: "copy empty directory",
			setup: func(t *testing.T, dir string) (string, string) {
				src := filepath.Join(dir, "src_empty")
				dst := filepath.Join(dir, "dst_empty")
				_ = os.Mkdir(src, 0755)
				return src, dst
			},
			wantErr: false,
			checkFn: func(t *testing.T, src, dst string) bool {
				dstInfo, err := os.Stat(dst)
				return err == nil && dstInfo.IsDir()
			},
		},
		{
			name: "copy directory with files",
			setup: func(t *testing.T, dir string) (string, string) {
				src := filepath.Join(dir, "src_files")
				dst := filepath.Join(dir, "dst_files")
				_ = os.Mkdir(src, 0755)
				_ = os.WriteFile(filepath.Join(src, "file1.txt"), []byte("content1"), 0644)
				_ = os.WriteFile(filepath.Join(src, "file2.txt"), []byte("content2"), 0644)
				return src, dst
			},
			wantErr: false,
			checkFn: func(t *testing.T, src, dst string) bool {
				// Check destination exists and has files
				dstInfo, err := os.Stat(dst)
				if err != nil || !dstInfo.IsDir() {
					return false
				}
				file1, _ := os.ReadFile(filepath.Join(dst, "file1.txt"))
				file2, _ := os.ReadFile(filepath.Join(dst, "file2.txt"))
				return string(file1) == "content1" && string(file2) == "content2"
			},
		},
		{
			name: "copy nested directories",
			setup: func(t *testing.T, dir string) (string, string) {
				src := filepath.Join(dir, "src_nested")
				dst := filepath.Join(dir, "dst_nested")
				_ = os.Mkdir(src, 0755)
				_ = os.Mkdir(filepath.Join(src, "subdir1"), 0755)
				_ = os.Mkdir(filepath.Join(src, "subdir1", "subdir2"), 0755)
				_ = os.WriteFile(filepath.Join(src, "subdir1", "subdir2", "file.txt"), []byte("nested"), 0644)
				return src, dst
			},
			wantErr: false,
			checkFn: func(t *testing.T, src, dst string) bool {
				content, _ := os.ReadFile(filepath.Join(dst, "subdir1", "subdir2", "file.txt"))
				return string(content) == "nested"
			},
		},
		{
			name: "copy fails if source not directory",
			setup: func(t *testing.T, dir string) (string, string) {
				src := filepath.Join(dir, "file.txt")
				dst := filepath.Join(dir, "dst")
				_ = os.WriteFile(src, []byte("content"), 0644)
				return src, dst
			},
			wantErr: true,
		},
		{
			name: "copy fails if destination exists",
			setup: func(t *testing.T, dir string) (string, string) {
				src := filepath.Join(dir, "src_existing")
				dst := filepath.Join(dir, "dst_existing")
				_ = os.Mkdir(src, 0755)
				_ = os.Mkdir(dst, 0755)
				return src, dst
			},
			wantErr: true,
		},
		{
			name: "copy fails if source not exists",
			setup: func(t *testing.T, dir string) (string, string) {
				src := filepath.Join(dir, "nonexistent")
				dst := filepath.Join(dir, "dst")
				return src, dst
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			src, dst := tt.setup(t, dir)

			err := lxio.CopyDir(src, dst)
			if (err != nil) != tt.wantErr {
				t.Errorf("CopyDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tt.checkFn != nil && !tt.checkFn(t, src, dst) {
				t.Errorf("CopyDir() result check failed")
			}
		})
	}
}

// ========================= Touch Tests =========================

func TestTouch(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(t *testing.T, dir string) string
		wantErr bool
		checkFn func(t *testing.T, path string) bool
	}{
		{
			name: "touch creates new file",
			setup: func(t *testing.T, dir string) string {
				return filepath.Join(dir, "new_file.txt")
			},
			wantErr: false,
			checkFn: func(t *testing.T, path string) bool {
				info, err := os.Stat(path)
				return err == nil && !info.IsDir() && info.Size() == 0
			},
		},
		{
			name: "touch updates modification time on existing file",
			setup: func(t *testing.T, dir string) string {
				path := filepath.Join(dir, "existing.txt")
				_ = os.WriteFile(path, []byte("content"), 0644)
				// Set old modification time
				_ = os.Chtimes(path, oldTime, oldTime)
				return path
			},
			wantErr: false,
			checkFn: func(t *testing.T, path string) bool {
				info, err := os.Stat(path)
				if err != nil {
					return false
				}
				// Check file was modified recently (within last second)
				return info.ModTime().After(oldTime)
			},
		},
		{
			name: "touch creates file with correct permissions",
			setup: func(t *testing.T, dir string) string {
				return filepath.Join(dir, "perm_file.txt")
			},
			wantErr: false,
			checkFn: func(t *testing.T, path string) bool {
				info, err := os.Stat(path)
				if err != nil || runtime.GOOS == "windows" {
					return err == nil
				}
				// Check file has readable/writable permissions
				return (info.Mode() & 0644) != 0
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			path := tt.setup(t, dir)

			err := lxio.Touch(path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Touch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tt.checkFn != nil && !tt.checkFn(t, path) {
				t.Errorf("Touch() result check failed")
			}
		})
	}
}

// ========================= Rename Tests =========================

func TestRename(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(t *testing.T, dir string) (string, string)
		wantErr bool
		checkFn func(t *testing.T, oldpath, newpath string) bool
	}{
		{
			name: "rename file",
			setup: func(t *testing.T, dir string) (string, string) {
				oldpath := filepath.Join(dir, "old.txt")
				newpath := filepath.Join(dir, "new.txt")
				_ = os.WriteFile(oldpath, []byte("content"), 0644)
				return oldpath, newpath
			},
			wantErr: false,
			checkFn: func(t *testing.T, oldpath, newpath string) bool {
				// Old path should not exist
				_, errOld := os.Stat(oldpath)
				// New path should exist
				_, errNew := os.Stat(newpath)
				return os.IsNotExist(errOld) && errNew == nil
			},
		},
		{
			name: "rename directory",
			setup: func(t *testing.T, dir string) (string, string) {
				oldpath := filepath.Join(dir, "olddir")
				newpath := filepath.Join(dir, "newdir")
				_ = os.Mkdir(oldpath, 0755)
				return oldpath, newpath
			},
			wantErr: false,
			checkFn: func(t *testing.T, oldpath, newpath string) bool {
				_, errOld := os.Stat(oldpath)
				_, errNew := os.Stat(newpath)
				return os.IsNotExist(errOld) && errNew == nil
			},
		},
		{
			name: "rename to different directory",
			setup: func(t *testing.T, dir string) (string, string) {
				subdir := filepath.Join(dir, "subdir")
				_ = os.Mkdir(subdir, 0755)
				oldpath := filepath.Join(dir, "file.txt")
				newpath := filepath.Join(subdir, "file.txt")
				_ = os.WriteFile(oldpath, []byte("content"), 0644)
				return oldpath, newpath
			},
			wantErr: false,
			checkFn: func(t *testing.T, oldpath, newpath string) bool {
				_, errOld := os.Stat(oldpath)
				_, errNew := os.Stat(newpath)
				return os.IsNotExist(errOld) && errNew == nil
			},
		},
		{
			name: "rename nonexistent file returns error",
			setup: func(t *testing.T, dir string) (string, string) {
				return filepath.Join(dir, "nonexistent.txt"), filepath.Join(dir, "new.txt")
			},
			wantErr: true,
		},
		{
			name: "rename to invalid directory returns error",
			setup: func(t *testing.T, dir string) (string, string) {
				oldpath := filepath.Join(dir, "file.txt")
				newpath := filepath.Join(dir, "nonexistent_dir", "file.txt")
				_ = os.WriteFile(oldpath, []byte("content"), 0644)
				return oldpath, newpath
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			oldpath, newpath := tt.setup(t, dir)

			err := lxio.Rename(oldpath, newpath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rename() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tt.checkFn != nil && !tt.checkFn(t, oldpath, newpath) {
				t.Errorf("Rename() result check failed")
			}
		})
	}
}

var oldTime = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
