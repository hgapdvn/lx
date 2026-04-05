package lxio_test

import (
	"bytes"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	lxio "github.com/hgapdvn/lx/io"
)

func TestWrite(t *testing.T) {
	t.Run("write binary data", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "binary.bin")
		data := []byte{0xDE, 0xAD, 0xBE, 0xEF}

		err := lxio.Write(path, data)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if !bytes.Equal(got, data) {
			t.Errorf("expected %v, got %v", data, got)
		}
	})

	t.Run("write empty data", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "empty.bin")
		data := []byte{}

		err := lxio.Write(path, data)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if len(got) != 0 {
			t.Errorf("expected empty file, got %d bytes", len(got))
		}
	})

	t.Run("write single byte", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "single.bin")
		data := []byte{0xFF}

		err := lxio.Write(path, data)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if !bytes.Equal(got, data) {
			t.Errorf("expected %v, got %v", data, got)
		}
	})

	t.Run("write large data", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "large.bin")
		data := bytes.Repeat([]byte{0xAB}, 10000)

		err := lxio.Write(path, data)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if !bytes.Equal(got, data) {
			t.Errorf("data mismatch")
		}
	})

	t.Run("overwrite existing file", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "overwrite.bin")

		_ = lxio.Write(path, []byte("first"))
		err := lxio.Write(path, []byte("second"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(got) != "second" {
			t.Errorf("expected 'second', got %q", string(got))
		}
	})

	t.Run("file permissions", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "perm.bin")

		err := lxio.Write(path, []byte("test"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		info, err := os.Stat(path)
		if err != nil {
			t.Fatalf("failed to stat file: %v", err)
		}

		if info.IsDir() {
			t.Errorf("expected regular file, got directory")
		}

		if info.Mode()&0400 == 0 {
			t.Errorf("file not readable by owner")
		}

		if info.Mode()&0200 == 0 {
			t.Errorf("file not writable by owner")
		}
	})

	t.Run("error on invalid path", func(t *testing.T) {
		invalidPath := "/invalid/nonexistent/deeply/nested/path/file.bin"
		err := lxio.Write(invalidPath, []byte("data"))
		if err == nil {
			t.Error("expected error for invalid path, got nil")
		}
	})
}

func TestWriteString(t *testing.T) {
	t.Run("write basic string", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "basic.txt")
		data := "hello world"

		err := lxio.WriteString(path, data)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(got) != data {
			t.Errorf("expected %q, got %q", data, string(got))
		}
	})

	t.Run("write empty string", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "empty.txt")
		data := ""

		err := lxio.WriteString(path, data)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if len(got) != 0 {
			t.Errorf("expected empty file, got %q", string(got))
		}
	})

	t.Run("write string with newlines", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "multiline.txt")
		data := "line1\nline2\nline3"

		err := lxio.WriteString(path, data)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(got) != data {
			t.Errorf("expected %q, got %q", data, string(got))
		}
	})

	t.Run("write string with special characters", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "special.txt")
		data := "!@#$%^&*()"

		err := lxio.WriteString(path, data)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(got) != data {
			t.Errorf("expected %q, got %q", data, string(got))
		}
	})

	t.Run("write unicode string", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "unicode.txt")
		data := "Hello 世界 🌍"

		err := lxio.WriteString(path, data)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(got) != data {
			t.Errorf("expected %q, got %q", data, string(got))
		}
	})

	t.Run("write string with tabs", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "tabs.txt")
		data := "col1\tcol2\tcol3"

		err := lxio.WriteString(path, data)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(got) != data {
			t.Errorf("expected %q, got %q", data, string(got))
		}
	})

	t.Run("overwrite longer with shorter", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "overwrite.txt")

		_ = lxio.WriteString(path, "longerstring")
		err := lxio.WriteString(path, "short")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(got) != "short" {
			t.Errorf("expected 'short', got %q", string(got))
		}
	})

	t.Run("write with null bytes", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "nullbytes.txt")
		data := "before\x00after"

		err := lxio.WriteString(path, data)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(got) != data {
			t.Errorf("expected %q, got %q", data, string(got))
		}
	})

	t.Run("error on invalid path", func(t *testing.T) {
		invalidPath := "/invalid/nonexistent/deeply/nested/path/file.txt"
		err := lxio.WriteString(invalidPath, "data")
		if err == nil {
			t.Error("expected error for invalid path, got nil")
		}
	})
}

func TestWriteLinesString(t *testing.T) {
	t.Run("write multiple lines", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "lines.txt")
		lines := []string{"one", "two", "three"}
		expected := "one\ntwo\nthree"

		err := lxio.WriteLinesString(path, lines)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(got) != expected {
			t.Errorf("expected %q, got %q", expected, string(got))
		}
	})

	t.Run("write single line", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "single.txt")
		lines := []string{"single"}
		expected := "single"

		err := lxio.WriteLinesString(path, lines)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(got) != expected {
			t.Errorf("expected %q, got %q", expected, string(got))
		}
	})

	t.Run("write empty slice", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "empty.txt")
		lines := []string{}

		err := lxio.WriteLinesString(path, lines)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if len(got) != 0 {
			t.Errorf("expected empty file, got %q", string(got))
		}
	})

	t.Run("write lines with empty strings", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "empty_lines.txt")
		lines := []string{"first", "", "third"}
		expected := "first\n\nthird"

		err := lxio.WriteLinesString(path, lines)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(got) != expected {
			t.Errorf("expected %q, got %q", expected, string(got))
		}
	})

	t.Run("write lines with unicode", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "unicode.txt")
		lines := []string{"hello", "世界", "🌍"}
		expected := "hello\n世界\n🌍"

		err := lxio.WriteLinesString(path, lines)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(got) != expected {
			t.Errorf("expected %q, got %q", expected, string(got))
		}
	})

	t.Run("write lines with spaces", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "spaces.txt")
		lines := []string{"  leading", "trailing  ", "  both  "}
		expected := "  leading\ntrailing  \n  both  "

		err := lxio.WriteLinesString(path, lines)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(got) != expected {
			t.Errorf("expected %q, got %q", expected, string(got))
		}
	})

	t.Run("overwrite with fewer lines", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "overwrite.txt")

		_ = lxio.WriteLinesString(path, []string{"a", "b", "c"})
		err := lxio.WriteLinesString(path, []string{"x"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(got) != "x" {
			t.Errorf("expected 'x', got %q", string(got))
		}
	})

	t.Run("error on invalid path", func(t *testing.T) {
		invalidPath := "/invalid/nonexistent/deeply/nested/path/file.txt"
		err := lxio.WriteLinesString(invalidPath, []string{"data"})
		if err == nil {
			t.Error("expected error for invalid path, got nil")
		}
	})
}

func TestWriteLinesBytes(t *testing.T) {
	t.Run("write multiple byte lines", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "bytes.txt")
		lines := [][]byte{[]byte("100"), []byte("200"), []byte("300")}
		expected := "100\n200\n300"

		err := lxio.WriteLinesBytes(path, lines)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(got) != expected {
			t.Errorf("expected %q, got %q", expected, string(got))
		}
	})

	t.Run("write single byte line", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "single.txt")
		lines := [][]byte{[]byte("single")}
		expected := "single"

		err := lxio.WriteLinesBytes(path, lines)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(got) != expected {
			t.Errorf("expected %q, got %q", expected, string(got))
		}
	})

	t.Run("write empty byte slice", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "empty.txt")
		lines := [][]byte{}

		err := lxio.WriteLinesBytes(path, lines)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if len(got) != 0 {
			t.Errorf("expected empty file, got %d bytes", len(got))
		}
	})

	t.Run("write bytes with empty lines", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "empty_bytes.txt")
		lines := [][]byte{[]byte("first"), []byte(""), []byte("third")}
		expected := "first\n\nthird"

		err := lxio.WriteLinesBytes(path, lines)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(got) != expected {
			t.Errorf("expected %q, got %q", expected, string(got))
		}
	})

	t.Run("write binary bytes", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "binary.txt")
		lines := [][]byte{{0xFF, 0xFE}, {0xAA, 0xBB}}
		expected := "\xff\xfe\n\xaa\xbb"

		err := lxio.WriteLinesBytes(path, lines)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(got) != expected {
			t.Errorf("expected %q, got %q", expected, string(got))
		}
	})

	t.Run("overwrite with more lines", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "overwrite.txt")

		_ = lxio.WriteLinesBytes(path, [][]byte{[]byte("a")})
		err := lxio.WriteLinesBytes(path, [][]byte{[]byte("x"), []byte("y"), []byte("z")})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		expected := "x\ny\nz"
		if string(got) != expected {
			t.Errorf("expected %q, got %q", expected, string(got))
		}
	})

	t.Run("error on invalid path", func(t *testing.T) {
		invalidPath := "/invalid/nonexistent/deeply/nested/path/file.txt"
		err := lxio.WriteLinesBytes(invalidPath, [][]byte{[]byte("data")})
		if err == nil {
			t.Error("expected error for invalid path, got nil")
		}
	})
}

func TestAppend(t *testing.T) {
	t.Run("append to existing file", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "append.bin")

		_ = lxio.Write(path, []byte("A"))
		err := lxio.Append(path, []byte("B"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if !bytes.Equal(got, []byte("AB")) {
			t.Errorf("expected 'AB', got %v", got)
		}
	})

	t.Run("append to non-existent file", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "new.bin")

		err := lxio.Append(path, []byte("first"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if !bytes.Equal(got, []byte("first")) {
			t.Errorf("expected 'first', got %v", got)
		}
	})

	t.Run("append empty bytes", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "empty.bin")

		_ = lxio.Write(path, []byte("data"))
		err := lxio.Append(path, []byte(""))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if !bytes.Equal(got, []byte("data")) {
			t.Errorf("expected 'data', got %v", got)
		}
	})

	t.Run("append multiple times", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "multiple.bin")

		_ = lxio.Write(path, []byte("start"))
		_ = lxio.Append(path, []byte("_"))
		_ = lxio.Append(path, []byte("middle"))
		err := lxio.Append(path, []byte("_end"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if !bytes.Equal(got, []byte("start_middle_end")) {
			t.Errorf("expected 'start_middle_end', got %v", got)
		}
	})

	t.Run("append large data", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "large.bin")

		_ = lxio.Write(path, []byte("x"))
		err := lxio.Append(path, bytes.Repeat([]byte("y"), 5000))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if len(got) != 5001 {
			t.Errorf("expected 5001 bytes, got %d", len(got))
		}
	})

	t.Run("error on invalid path", func(t *testing.T) {
		invalidPath := "/invalid/nonexistent/deeply/nested/path/file.bin"
		err := lxio.Append(invalidPath, []byte("data"))
		if err == nil {
			t.Error("expected error for invalid path, got nil")
		}
	})
}

func TestAppendString(t *testing.T) {
	t.Run("append string to existing file", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "append.txt")

		_ = lxio.WriteString(path, "hello")
		err := lxio.AppendString(path, " world")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(got) != "hello world" {
			t.Errorf("expected 'hello world', got %q", string(got))
		}
	})

	t.Run("append string to non-existent file", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "new.txt")

		err := lxio.AppendString(path, "new")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(got) != "new" {
			t.Errorf("expected 'new', got %q", string(got))
		}
	})

	t.Run("append empty string", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "empty.txt")

		_ = lxio.WriteString(path, "data")
		err := lxio.AppendString(path, "")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(got) != "data" {
			t.Errorf("expected 'data', got %q", string(got))
		}
	})

	t.Run("append unicode strings", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "unicode.txt")

		_ = lxio.WriteString(path, "start")
		err := lxio.AppendString(path, "中文🎉")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(got) != "start中文🎉" {
			t.Errorf("expected 'start中文🎉', got %q", string(got))
		}
	})

	t.Run("error on invalid path", func(t *testing.T) {
		invalidPath := "/invalid/nonexistent/deeply/nested/path/file.txt"
		err := lxio.AppendString(invalidPath, "data")
		if err == nil {
			t.Error("expected error for invalid path, got nil")
		}
	})
}

func TestAppendLine(t *testing.T) {
	t.Run("append single line", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "logs.txt")

		err := lxio.AppendLine(path, "log1")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}

		newline := "\n"
		if runtime.GOOS == "windows" {
			newline = "\r\n"
		}

		expected := "log1" + newline
		if string(got) != expected {
			t.Errorf("expected %q, got %q", expected, string(got))
		}
	})

	t.Run("append multiple lines", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "logs.txt")

		_ = lxio.AppendLine(path, "log1")
		_ = lxio.AppendLine(path, "log2")
		err := lxio.AppendLine(path, "log3")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}

		newline := "\n"
		if runtime.GOOS == "windows" {
			newline = "\r\n"
		}

		expected := "log1" + newline + "log2" + newline + "log3" + newline
		if string(got) != expected {
			t.Errorf("expected %q, got %q", expected, string(got))
		}
	})

	t.Run("append line to non-existent file", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "new_logs.txt")

		err := lxio.AppendLine(path, "first")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}

		newline := "\n"
		if runtime.GOOS == "windows" {
			newline = "\r\n"
		}

		expected := "first" + newline
		if string(got) != expected {
			t.Errorf("expected %q, got %q", expected, string(got))
		}
	})

	t.Run("append empty lines", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "empty_logs.txt")

		_ = lxio.AppendLine(path, "")
		err := lxio.AppendLine(path, "")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}

		newline := "\n"
		if runtime.GOOS == "windows" {
			newline = "\r\n"
		}

		expected := newline + newline
		if string(got) != expected {
			t.Errorf("expected %q, got %q", expected, string(got))
		}
	})

	t.Run("append unicode lines", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "unicode_logs.txt")

		_ = lxio.AppendLine(path, "line1")
		_ = lxio.AppendLine(path, "世界")
		err := lxio.AppendLine(path, "🌍")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}

		newline := "\n"
		if runtime.GOOS == "windows" {
			newline = "\r\n"
		}

		expected := "line1" + newline + "世界" + newline + "🌍" + newline
		if string(got) != expected {
			t.Errorf("expected %q, got %q", expected, string(got))
		}
	})

	t.Run("error on invalid path", func(t *testing.T) {
		invalidPath := "/invalid/nonexistent/deeply/nested/path/file.txt"
		err := lxio.AppendLine(invalidPath, "data")
		if err == nil {
			t.Error("expected error for invalid path, got nil")
		}
	})
}
