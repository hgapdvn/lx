package lxio

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// ------------------------------- Helper functions for reading small amount of data. -------------------------------

// Read reads the entire file into memory.
// CAUTION: Do not use this for large files (e.g., > 100MB), as it will load the entire
// file into RAM, potentially causing an Out-Of-Memory (OOM) error.
func Read(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// ReadString reads the entire file into memory and returns it as a string.
func ReadString(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ReadLinesBytes reads the entire file and returns a slice of byte slices, one for each line.
// Use this for small files where you want to avoid string allocation overhead.
func ReadLinesBytes(path string) ([][]byte, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Fast path: use bytes.Split instead of strings.Split
	lines := bytes.Split(b, []byte("\n"))

	// Clean up Windows \r
	for i, line := range lines {
		lines[i] = bytes.TrimSuffix(line, []byte("\r"))
	}

	// Remove the last empty element if the file ends with a newline
	// (len(line) == 0 checks if the byte slice is empty)
	if len(lines) > 0 && len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}

	return lines, nil
}

// ReadLinesString reads the entire file and returns a slice of strings, one for each line.
// Use this for small files. For large files, use ReadLines with a callback.
func ReadLinesString(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// --------------------------------- Helper functions for reading large amount of data. --------------------------------

// ForEachChunk reads data from r in chunks and calls fn for each chunk.
// The provided slice is reused between calls, so fn must not retain it.
func ForEachChunk(r io.Reader, chunkSize int, fn func([]byte) error) error {
	if chunkSize <= 0 {
		chunkSize = 32 * 1024
	}

	buf := make([]byte, chunkSize)

	for {
		n, err := r.Read(buf)

		if n > 0 {
			if err := fn(buf[:n]); err != nil {
				return err
			}
		}

		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

// ForEachLine reads a file line by line and calls fn for each line.
// If fn returns an error, iteration stops and that error is returned.
//
// Lines up to 1MB are supported. Larger lines will cause scanning to fail.
func ForEachLine(path string, fn func(line string) error) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Increase the maximum token size.
	//
	// bufio.Scanner has a default maximum line size of 64KB. If a line exceeds
	// this limit, scanning will fail with "bufio.Scanner: token too long".
	//
	// We provide:
	//   - an initial buffer capacity of 64KB
	//   - a maximum allowed line size of 1MB
	//
	// The scanner will automatically grow the buffer as needed up to 1MB.
	// If lines may exceed this size, consider increasing the limit or using
	// bufio.Reader instead.
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	for scanner.Scan() {
		line := scanner.Text()
		if err := fn(line); err != nil {
			return err
		}
	}

	return scanner.Err()
}

// --------------------------------- First N / Last N / Count Lines  ---------------------------------

// ReadFirstN reads the first n lines from a file.
// Returns an error if the file cannot be opened or read.
// If the file has fewer than n lines, returns all available lines.
// Returns empty slice if n is 0 or the file is empty.
//
// Example:
//
//     lines, err := lxio.ReadFirstN("/path/to/file", 10)
//     // lines: first 10 lines of file (or fewer if file is smaller)
//
func ReadFirstN(path string, n int) ([]string, error) {
	if n <= 0 {
		return []string{}, nil
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	var lines []string
	count := 0

	for scanner.Scan() && count < n {
		lines = append(lines, scanner.Text())
		count++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// ReadLastN reads the last n lines from a file using a memory-efficient rolling buffer.
// Returns an error if the file cannot be opened or read.
// If the file has fewer than n lines, returns all available lines.
// Returns empty slice if n is 0 or the file is empty.
//
// Performance: Uses a circular buffer that keeps only the last n lines in memory,
// making it memory-efficient even for very large files.
//
// Example:
//
//     lines, err := lxio.ReadLastN("/path/to/file", 10)
//     // lines: last 10 lines of file (or fewer if file is smaller)
//
func ReadLastN(path string, n int) ([]string, error) {
	if n <= 0 {
		return []string{}, nil
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	// Use a rolling/circular buffer that stores at most n lines
	buffer := make([]string, 0, n)
	index := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(buffer) < n {
			// Buffer not full yet, just append
			buffer = append(buffer, line)
		} else {
			// Buffer is full, overwrite oldest entry (circular)
			buffer[index%n] = line
		}
		index++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Reconstruct the lines in correct order if buffer was full
	if len(buffer) < n {
		// Buffer never filled, return as-is
		return buffer, nil
	}

	// Buffer was full, rotate to get correct order
	result := make([]string, n)
	for i := 0; i < n; i++ {
		result[i] = buffer[(index+i)%n]
	}

	return result, nil
}

// CountLines counts the number of lines in a file without loading the entire content into memory.
// Returns an error if the file cannot be opened or read.
// An empty file has 0 lines. A file with content but no trailing newline counts as 1 line.
//
// Example:
//
//     count, err := lxio.CountLines("/path/to/file")
//     if err != nil {
//         // handle error
//     }
//     fmt.Printf("File has %d lines\n", count)
//
func CountLines(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	buf := make([]byte, 32*1024)
	count := 0
	var lastByte byte
	var readAny bool

	for {
		n, err := file.Read(buf)
		if n > 0 {
			readAny = true
			for i := 0; i < n; i++ {
				if buf[i] == '\n' {
					count++
				}
			}
			lastByte = buf[n-1]
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}
	}

	// If file is not empty and doesn't end with newline → add last line
	if readAny && lastByte != '\n' {
		count++
	}

	return count, nil
}
