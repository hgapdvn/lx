package lxenv_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/nthanhhai2909/lx/lxenv"
)

// TestLoadEnv_FileDescriptorsClosed verifies that LoadEnv does not leak file
// descriptors when loading many files. With the old defer-in-loop behaviour all
// files stayed open until the function returned, which could exhaust the OS
// file-descriptor limit and produce "too many open files" errors.
func TestLoadEnv_FileDescriptorsClosed(t *testing.T) {
	const numFiles = 200

	dir := t.TempDir()
	paths := make([]string, numFiles)

	for i := 0; i < numFiles; i++ {
		p := filepath.Join(dir, fmt.Sprintf("env%d.env", i))
		content := fmt.Sprintf("TEST_FD_KEY_%d=value%d\n", i, i)
		if err := os.WriteFile(p, []byte(content), 0o600); err != nil {
			t.Fatalf("failed to write temp env file: %v", err)
		}
		paths[i] = p
	}

	if err := lxenv.LoadEnv(paths...); err != nil {
		t.Fatalf("LoadEnv returned unexpected error (possible fd leak): %v", err)
	}
}

// TestLoadEnv_NormalizeWhitespace verifies that normalize trims surrounding
// whitespace from unquoted values while preserving internal whitespace inside
// quoted values.
func TestLoadEnv_NormalizeWhitespace(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "ws.env")

	// KEY1: unquoted value followed by trailing spaces (no comment)
	// KEY2: quoted value with intentional leading/trailing spaces inside quotes
	// KEY3: unquoted value with an inline comment and extra spaces before #
	content := "KEY1=trimmed   \nKEY2=\"  spaced  \"\nKEY3=value   # inline comment\n"
	if err := os.WriteFile(p, []byte(content), 0o600); err != nil {
		t.Fatalf("failed to write temp env file: %v", err)
	}

	t.Cleanup(func() {
		os.Unsetenv("KEY1")
		os.Unsetenv("KEY2")
		os.Unsetenv("KEY3")
	})

	if err := lxenv.LoadEnv(p); err != nil {
		t.Fatalf("LoadEnv() unexpected error: %v", err)
	}

	tests := []struct{ key, want string }{
		{"KEY1", "trimmed"},         // trailing spaces on unquoted value must be stripped
		{"KEY2", "  spaced  "},      // internal spaces inside quotes must be preserved
		{"KEY3", "value"},           // inline comment and surrounding spaces must be stripped
	}
	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			if got := os.Getenv(tt.key); got != tt.want {
				t.Errorf("env[%q] = %q, want %q", tt.key, got, tt.want)
			}
		})
	}
}
