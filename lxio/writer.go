package lxio

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
)

const defaultFileMode = 0644

// Newline contains the native line ending for the current operating system.
// It is "\r\n" on Windows, and "\n" on Linux/macOS.
var Newline = getNewline()

func getNewline() []byte {
	if runtime.GOOS == "windows" {
		return []byte("\r\n")
	}
	return []byte("\n")
}

// --------------------------------- Writing Data (Overwrites file) ---------------------------------

// Write writes data to the named file, creating it if necessary.
// If the file already exists, it is truncated (overwritten).
func Write(path string, data []byte) error {
	return os.WriteFile(path, data, defaultFileMode)
}

// WriteWithPerm writes data to the named file with custom file permissions.
// If the file already exists, it is truncated (overwritten).
//
// Example:
//
//     err := lxio.WriteWithPerm("/path/to/file", []byte("content"), 0755)
//     // Creates file with rwxr-xr-x permissions
//
func WriteWithPerm(path string, data []byte, perm os.FileMode) error {
	return os.WriteFile(path, data, perm)
}

// WriteString writes a string to the named file, creating it if necessary.
// If the file already exists, it is truncated (overwritten).
func WriteString(path string, data string) error {
	return os.WriteFile(path, []byte(data), defaultFileMode)
}

// WriteStringf writes a formatted string to the named file, creating it if necessary.
// If the file already exists, it is truncated (overwritten).
// The format and arguments work the same as fmt.Sprintf.
//
// Example:
//
//     err := lxio.WriteStringf("/path/to/file", "User: %s, Age: %d", "Alice", 30)
//     // File contents: "User: Alice, Age: 30"
//
func WriteStringf(path string, format string, args ...any) error {
	return WriteString(path, fmt.Sprintf(format, args...))
}

// WriteLinesBytes writes a slice of byte slices to a file, separating them with standard newlines (\n).
// If the file already exists, it is truncated.
func WriteLinesBytes(path string, lines [][]byte) error {
	data := bytes.Join(lines, []byte("\n"))
	return os.WriteFile(path, data, defaultFileMode)
}

// WriteLinesString writes a slice of strings to a file, separating them with standard newlines (\n).
// If the file already exists, it is truncated.
func WriteLinesString(path string, lines []string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, defaultFileMode)
	if err != nil {
		return err
	}
	defer f.Close()

	for i, line := range lines {
		if _, err := f.WriteString(line); err != nil {
			return err
		}
		// Add a newline after every line except the last one
		if i < len(lines)-1 {
			if _, err := f.WriteString("\n"); err != nil {
				return err
			}
		}
	}
	return nil
}

// --------------------------------- Appending Data (Keeps existing) ---------------------------------

// Append adds data to the end of the named file.
// If the file does not exist, it is created.
func Append(path string, data []byte) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, defaultFileMode)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	return err
}

// AppendString adds a string to the end of the named file.
// If the file does not exist, it is created.
func AppendString(path string, data string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, defaultFileMode)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(data)
	return err
}

// AppendLine writes a string followed by the OS-native newline (\r\n on Windows, \n on Unix).
// If the file does not exist, it is created. This is highly recommended for appending logs.
func AppendLine(path string, line string) error {
	data := append([]byte(line), Newline...)
	return Append(path, data)
}

// AppendLines appends multiple lines to the file, each followed by a newline.
// If the file does not exist, it is created.
// Each line is terminated with a newline character.
//
// Example:
//
//     lines := []string{"line1", "line2", "line3"}
//     err := lxio.AppendLines("/path/to/file", lines)
//     // File contents (appended):
//     // line1
//     // line2
//     // line3
//
func AppendLines(path string, lines []string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, defaultFileMode)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, line := range lines {
		data := append([]byte(line), Newline...)
		if _, err := f.Write(data); err != nil {
			return err
		}
	}
	return nil
}

// Truncate clears the contents of the named file without deleting the file itself.
// If the file does not exist, it is created.
func Truncate(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, defaultFileMode)
	if err != nil {
		return err
	}
	return f.Close()
}
