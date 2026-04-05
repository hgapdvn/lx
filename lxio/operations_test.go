package lxio_test

import (
	"os"
	"path/filepath"
	"testing"

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
