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
