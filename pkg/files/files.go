package files

import (
	"log"
	"os"
	"path"
	"strings"
)

// Parse a path, replaces ~ with the home directory
func ParsePath(p string) string {

	// use the right separator (windows)

	parts := strings.Split(p, "/")
	joinedPath := path.Join(parts...)

	// get home folder

	homedir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal("Could not find the home dir : ", err)
	}

	return path.Clean(strings.ReplaceAll(joinedPath, "~", homedir))
}

// Ensure that the folder of the path exists (name before the last foward slash)
func EnsureFolder(p string) {
	base := path.Dir(p)
	os.MkdirAll(base, 0755)
}

// Checks if the file exists
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
