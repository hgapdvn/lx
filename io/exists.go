package lxio

// Exists returns true if the file exists.
// It swallows any errors and safely defaults to false.
func Exists(path string) bool {
	ok, _ := ExistsE(path)
	return ok
}
