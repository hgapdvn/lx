package lxio_test

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/hgapdvn/lx/lxio"
)

// helper to write a temporary file and return its path
func createTempFile(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "testfile.txt")
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	return path
}

func TestRead(t *testing.T) {
	content := "hello world"
	path := createTempFile(t, content)

	// Test Success
	b, err := lxio.Read(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !bytes.Equal(b, []byte(content)) {
		t.Errorf("expected %q, got %q", content, string(b))
	}

	// Test Error (File not found)
	_, err = lxio.Read("non_existent_file.txt")
	if err == nil {
		t.Error("expected error for non-existent file, got nil")
	}
}

func TestReadString(t *testing.T) {
	content := "hello world string"
	path := createTempFile(t, content)

	// Test Success
	s, err := lxio.ReadString(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s != content {
		t.Errorf("expected %q, got %q", content, s)
	}

	// Test Error
	_, err = lxio.ReadString("non_existent_file.txt")
	if err == nil {
		t.Error("expected error for non-existent file, got nil")
	}
}

func TestReadLinesBytes(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected [][]byte
	}{
		{
			name:     "Unix newlines",
			content:  "line1\nline2\nline3\n",
			expected: [][]byte{[]byte("line1"), []byte("line2"), []byte("line3")},
		},
		{
			name:     "Windows newlines",
			content:  "line1\r\nline2\r\nline3\r\n",
			expected: [][]byte{[]byte("line1"), []byte("line2"), []byte("line3")},
		},
		{
			name:     "No trailing newline",
			content:  "line1\nline2",
			expected: [][]byte{[]byte("line1"), []byte("line2")},
		},
		{
			name:     "Empty file",
			content:  "",
			expected: [][]byte{}, // Split on empty string might yield [""] depending on implementation, but your logic removes trailing empty
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := createTempFile(t, tt.content)
			lines, err := lxio.ReadLinesBytes(path)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Handle empty slice comparison nicely
			if len(lines) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(lines, tt.expected) {
				t.Errorf("expected %q, got %q", tt.expected, lines)
			}
		})
	}

	t.Run("file not found", func(t *testing.T) {
		_, err := lxio.ReadLinesBytes("non_existent_file.txt")
		if err == nil {
			t.Error("expected error for non-existent file, got nil")
		}
	})
}

func TestReadLinesString(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected []string
	}{
		{
			name:     "Unix newlines",
			content:  "line1\nline2\nline3\n",
			expected: []string{"line1", "line2", "line3"},
		},
		{
			name:     "No trailing newline",
			content:  "line1\nline2",
			expected: []string{"line1", "line2"},
		},
		{
			name:     "Empty file",
			content:  "",
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := createTempFile(t, tt.content)
			lines, err := lxio.ReadLinesString(path)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if len(lines) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(lines, tt.expected) {
				t.Errorf("expected %q, got %q", tt.expected, lines)
			}
		})
	}

	t.Run("file not found", func(t *testing.T) {
		_, err := lxio.ReadLinesString("non_existent_file.txt")
		if err == nil {
			t.Error("expected error for non-existent file, got nil")
		}
	})

	t.Run("Windows line endings", func(t *testing.T) {
		windowsContent := "line1\r\nline2\r\nline3"
		windowsPath := createTempFile(t, windowsContent)
		lines, err := lxio.ReadLinesString(windowsPath)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expected := []string{"line1", "line2", "line3"}
		if !reflect.DeepEqual(lines, expected) {
			t.Errorf("expected %q, got %q", expected, lines)
		}
	})
}

func TestForEachChunk(t *testing.T) {
	content := "123456789" // 9 bytes total

	t.Run("read in multiple chunks", func(t *testing.T) {
		r := strings.NewReader(content)
		var result []byte

		err := lxio.ForEachChunk(r, 4, func(chunk []byte) error {
			result = append(result, chunk...)
			return nil
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if string(result) != content {
			t.Errorf("expected %q, got %q", content, string(result))
		}
	})

	t.Run("stop early on error", func(t *testing.T) {
		r := strings.NewReader(content)
		expectedErr := errors.New("stop early")
		calls := 0

		err := lxio.ForEachChunk(r, 2, func(chunk []byte) error {
			calls++
			if calls == 2 { // Stop on second chunk
				return expectedErr
			}
			return nil
		})

		if !errors.Is(err, expectedErr) {
			t.Errorf("expected error %v, got %v", expectedErr, err)
		}
		if calls != 2 {
			t.Errorf("expected 2 calls, got %d", calls)
		}
	})

	t.Run("negative chunk size defaults to 32KB", func(t *testing.T) {
		r := strings.NewReader(content)
		err := lxio.ForEachChunk(r, -1, func(chunk []byte) error {
			// It should read everything in one go since 32KB > 9 bytes
			if len(chunk) != len(content) {
				t.Errorf("expected chunk size %d, got %d", len(content), len(chunk))
			}
			return nil
		})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("zero chunk size defaults to 32KB", func(t *testing.T) {
		r := strings.NewReader(content)
		called := false
		err := lxio.ForEachChunk(r, 0, func(chunk []byte) error {
			called = true
			// It should read everything in one go since 32KB > 9 bytes
			if len(chunk) != len(content) {
				t.Errorf("expected chunk size %d, got %d", len(content), len(chunk))
			}
			return nil
		})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !called {
			t.Error("expected callback to be called")
		}
	})

	t.Run("empty reader", func(t *testing.T) {
		r := strings.NewReader("")
		called := false
		err := lxio.ForEachChunk(r, 4, func(chunk []byte) error {
			called = true
			return nil
		})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if called {
			t.Error("expected callback not to be called for empty reader")
		}
	})
}

func TestForEachLine(t *testing.T) {
	content := "line1\nline2\nline3"
	path := createTempFile(t, content)

	t.Run("read all lines", func(t *testing.T) {
		var result []string
		err := lxio.ForEachLine(path, func(line string) error {
			result = append(result, line)
			return nil
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expected := []string{"line1", "line2", "line3"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %q, got %q", expected, result)
		}
	})

	t.Run("stop on error", func(t *testing.T) {
		expectedErr := errors.New("stop early")
		var result []string

		err := lxio.ForEachLine(path, func(line string) error {
			result = append(result, line)
			if line == "line2" {
				return expectedErr
			}
			return nil
		})

		if !errors.Is(err, expectedErr) {
			t.Errorf("expected error %v, got %v", expectedErr, err)
		}

		// It should have only captured line1 and line2
		expected := []string{"line1", "line2"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %q, got %q", expected, result)
		}
	})

	t.Run("file not found", func(t *testing.T) {
		err := lxio.ForEachLine("missing_file.txt", func(line string) error {
			return nil
		})
		if err == nil {
			t.Error("expected error for non-existent file, got nil")
		}
	})

	t.Run("Windows line endings", func(t *testing.T) {
		windowsContent := "line1\r\nline2\r\nline3"
		windowsPath := createTempFile(t, windowsContent)
		var result []string
		err := lxio.ForEachLine(windowsPath, func(line string) error {
			result = append(result, line)
			return nil
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expected := []string{"line1", "line2", "line3"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %q, got %q", expected, result)
		}
	})
}

// =============================== ReadFirstN Tests ===============================

func TestReadFirstN(t *testing.T) {
	tests := []struct {
		name      string
		content   string
		n         int
		expected  []string
		shouldErr bool
	}{
		{
			name:     "read first 3 lines",
			content:  "line1\nline2\nline3\nline4\nline5",
			n:        3,
			expected: []string{"line1", "line2", "line3"},
		},
		{
			name:     "read first 1 line",
			content:  "only line",
			n:        1,
			expected: []string{"only line"},
		},
		{
			name:     "n greater than file lines",
			content:  "line1\nline2",
			n:        10,
			expected: []string{"line1", "line2"},
		},
		{
			name:     "n is 0",
			content:  "line1\nline2",
			n:        0,
			expected: []string{},
		},
		{
			name:     "n is negative",
			content:  "line1\nline2",
			n:        -5,
			expected: []string{},
		},
		{
			name:     "empty file",
			content:  "",
			n:        5,
			expected: []string{},
		},
		{
			name:     "file with trailing newline",
			content:  "line1\nline2\nline3\n",
			n:        2,
			expected: []string{"line1", "line2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := createTempFile(t, tt.content)
			result, err := lxio.ReadFirstN(path, tt.n)

			if tt.shouldErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tt.shouldErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !tt.shouldErr {
				// Handle nil vs empty slice comparison
				if len(result) != len(tt.expected) {
					t.Errorf("expected %d lines, got %d", len(tt.expected), len(result))
				}
				for i, line := range result {
					if i >= len(tt.expected) || line != tt.expected[i] {
						t.Errorf("expected %q, got %q at index %d", tt.expected, result, i)
						break
					}
				}
			}
		})
	}
}

func TestReadFirstN_NonExistentFile(t *testing.T) {
	result, err := lxio.ReadFirstN("/nonexistent/path/file.txt", 5)
	if err == nil {
		t.Errorf("expected error for nonexistent file, got nil")
	}
	if result != nil {
		t.Errorf("expected nil result for nonexistent file, got %v", result)
	}
}

// =============================== ReadLastN Tests ================================

func TestReadLastN(t *testing.T) {
	tests := []struct {
		name      string
		content   string
		n         int
		expected  []string
		shouldErr bool
	}{
		{
			name:     "read last 3 lines",
			content:  "line1\nline2\nline3\nline4\nline5",
			n:        3,
			expected: []string{"line3", "line4", "line5"},
		},
		{
			name:     "read last 1 line",
			content:  "only line",
			n:        1,
			expected: []string{"only line"},
		},
		{
			name:     "n greater than file lines",
			content:  "line1\nline2",
			n:        10,
			expected: []string{"line1", "line2"},
		},
		{
			name:     "n is 0",
			content:  "line1\nline2",
			n:        0,
			expected: []string{},
		},
		{
			name:     "n is negative",
			content:  "line1\nline2",
			n:        -5,
			expected: []string{},
		},
		{
			name:     "empty file",
			content:  "",
			n:        5,
			expected: []string{},
		},
		{
			name:     "file with trailing newline",
			content:  "line1\nline2\nline3\n",
			n:        2,
			expected: []string{"line2", "line3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := createTempFile(t, tt.content)
			result, err := lxio.ReadLastN(path, tt.n)

			if tt.shouldErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tt.shouldErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !tt.shouldErr {
				// Handle nil vs empty slice comparison
				if len(result) != len(tt.expected) {
					t.Errorf("expected %d lines, got %d", len(tt.expected), len(result))
				}
				for i, line := range result {
					if i >= len(tt.expected) || line != tt.expected[i] {
						t.Errorf("expected %q, got %q at index %d", tt.expected, result, i)
						break
					}
				}
			}
		})
	}
}

func TestReadLastN_NonExistentFile(t *testing.T) {
	result, err := lxio.ReadLastN("/nonexistent/path/file.txt", 5)
	if err == nil {
		t.Errorf("expected error for nonexistent file, got nil")
	}
	if result != nil {
		t.Errorf("expected nil result for nonexistent file, got %v", result)
	}
}

// =============================== CountLines Tests ===============================

func TestCountLines(t *testing.T) {
	tests := []struct {
		name      string
		content   string
		expected  int
		shouldErr bool
		isLarge   bool // Flag for large file tests
	}{
		{
			name:     "count 5 lines with trailing newlines",
			content:  "line1\nline2\nline3\nline4\nline5\n",
			expected: 5,
		},
		{
			name:     "single line with newline",
			content:  "only line\n",
			expected: 1,
		},
		{
			name:     "two lines with newline",
			content:  "line1\nline2\n",
			expected: 2,
		},
		{
			name:     "empty file",
			content:  "",
			expected: 0,
		},
		{
			name:     "file with trailing newline",
			content:  "line1\nline2\nline3\n",
			expected: 3,
		},
		{
			name:     "Windows line endings",
			content:  "line1\r\nline2\r\nline3\r\n",
			expected: 3,
		},
		{
			name:     "mixed line endings",
			content:  "line1\nline2\r\nline3\n",
			expected: 3,
		},
		{
			name:     "large file with 10000 lines",
			isLarge:  true,
			expected: 10000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var path string
			if tt.isLarge {
				// Generate large file dynamically
				testDir := t.TempDir()
				path = filepath.Join(testDir, "large.txt")
				content := strings.Repeat("line\n", tt.expected)
				if err := os.WriteFile(path, []byte(content), 0644); err != nil {
					t.Fatalf("failed to create test file: %v", err)
				}
			} else {
				path = createTempFile(t, tt.content)
			}

			result, err := lxio.CountLines(path)

			if tt.shouldErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tt.shouldErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !tt.shouldErr && result != tt.expected {
				t.Errorf("expected %d lines, got %d", tt.expected, result)
			}
		})
	}
}

func TestCountLines_NonExistentFile(t *testing.T) {
	result, err := lxio.CountLines("/nonexistent/path/file.txt")
	if err == nil {
		t.Errorf("expected error for nonexistent file, got nil")
	}
	if result != 0 {
		t.Errorf("expected 0 lines for nonexistent file, got %d", result)
	}
}

// Benchmark ReadLastN with a large file to demonstrate memory efficiency
func BenchmarkReadLastN_Large(b *testing.B) {
	testDir := b.TempDir()
	path := filepath.Join(testDir, "large_bench.txt")

	// Create a file with 100,000 lines
	sb := &strings.Builder{}
	for i := 0; i < 100000; i++ {
		sb.WriteString("line " + strconv.Itoa(i) + "\n")
	}
	if err := os.WriteFile(path, []byte(sb.String()), 0644); err != nil {
		b.Fatalf("failed to create bench file: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := lxio.ReadLastN(path, 100)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

// Benchmark ReadLastN with varying tail sizes
func BenchmarkReadLastN_Various(b *testing.B) {
	testDir := b.TempDir()
	path := filepath.Join(testDir, "bench.txt")

	// Create file with 50,000 lines
	sb := &strings.Builder{}
	for i := 0; i < 50000; i++ {
		sb.WriteString("line " + strconv.Itoa(i) + "\n")
	}
	if err := os.WriteFile(path, []byte(sb.String()), 0644); err != nil {
		b.Fatalf("failed to create bench file: %v", err)
	}

	sizes := []int{10, 100, 1000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("LastN_%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_, err := lxio.ReadLastN(path, size)
				if err != nil {
					b.Fatalf("unexpected error: %v", err)
				}
			}
		})
	}
}
