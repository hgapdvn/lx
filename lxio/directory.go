package lxio

import (
	"os"
	"path/filepath"
	"sort"
)

// ListFiles returns a sorted slice of file names (not directories) in the given directory.
// The slice contains only the base names of files, not full paths.
// It returns an error if the directory doesn't exist or cannot be read.
//
// Example:
//
//	files, err := lxio.ListFiles("/path/to/dir")
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, file := range files {
//		fmt.Println(file)
//	}
func ListFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}

	sort.Strings(files)
	return files, nil
}

// ListDirs returns a sorted slice of directory names (not files) in the given directory.
// The slice contains only the base names of directories, not full paths.
// It returns an error if the directory doesn't exist or cannot be read.
//
// Example:
//
//	dirs, err := lxio.ListDirs("/path/to/dir")
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, dir := range dirs {
//		fmt.Println(dir)
//	}
func ListDirs(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var dirs []string
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		}
	}

	sort.Strings(dirs)
	return dirs, nil
}

// ListAll returns a sorted slice of all entry names (both files and directories) in the given directory.
// The slice contains only the base names of entries, not full paths.
// It returns an error if the directory doesn't exist or cannot be read.
//
// Example:
//
//	entries, err := lxio.ListAll("/path/to/dir")
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, entry := range entries {
//		fmt.Println(entry)
//	}
func ListAll(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var names []string
	for _, entry := range entries {
		names = append(names, entry.Name())
	}

	sort.Strings(names)
	return names, nil
}

// WalkFiles recursively walks through the root directory and all subdirectories,
// calling fn for each file found. The path passed to fn is relative to root.
// If fn returns a non-nil error, the walk stops and that error is returned.
// It returns an error if the root directory doesn't exist or cannot be read.
//
// Example:
//
//	err := lxio.WalkFiles("/path/to/dir", func(path string) error {
//		fmt.Println(path) // e.g., "subdir/file.txt"
//		return nil
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
func WalkFiles(root string, fn func(path string) error) error {
	return walkFilesInternal(root, "", fn)
}

// walkFilesInternal is a helper function for WalkFiles.
// relRoot tracks the relative path from the original root.
func walkFilesInternal(currentPath string, relRoot string, fn func(path string) error) error {
	entries, err := os.ReadDir(currentPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		var relPath string
		if relRoot == "" {
			relPath = entry.Name()
		} else {
			relPath = filepath.Join(relRoot, entry.Name())
		}

		if entry.IsDir() {
			// Recurse into subdirectory
			fullPath := filepath.Join(currentPath, entry.Name())
			if err := walkFilesInternal(fullPath, relPath, fn); err != nil {
				return err
			}
		} else {
			// Call fn for this file
			if err := fn(relPath); err != nil {
				return err
			}
		}
	}

	return nil
}
