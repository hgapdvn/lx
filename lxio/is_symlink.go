package lxio

import (
	"os"
)

// IsSymlink returns true if the path exists and is a symbolic link.
// It safely returns false if the path doesn't exist, is a normal file/dir,
// or if there is a permission error.
func IsSymlink(path string) bool {
	// CRITICAL: Must use Lstat, not Stat.
	// Stat follows the link, Lstat looks at the link itself.
	info, err := os.Lstat(path)
	if err != nil {
		return false
	}

	// We use a bitwise AND to check if the Symlink bit is set in the file mode
	return info.Mode()&os.ModeSymlink != 0
}
