package helper

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Helper function for printing a message to stdout + exit with given exit code
func CustomExit(exitMessage string, exitCode int) {
	log.Println(exitMessage)
	os.Exit(exitCode)
}

// Check whether a given filename is a dir
func IsDir(dir string) bool {
	f, err := os.Open(dir)
	if err != nil {
		return false
	}

	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return false
	}

	if mode := fi.Mode(); mode.IsDir() {
		return true
	}

	return false
}

// Checks whether a file is writeable
func IsWritable(filename string) bool {
	f, err := os.Create(filename)
	if err != nil {
		return false
	}

	defer f.Close()

	if _, err := f.Write(make([]byte, 0)); err != nil {
		return false
	}

	return true
}

// Subfolders returns a slice of subfolders (recursive), including the folder provided.
func Subfolders(path string) (paths []string) {
	filepath.Walk(path, func(newPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			name := info.Name()
			// skip folders that begin with a dot
			if shouldIgnoreFile(name) && name != "." && name != ".." {
				return filepath.SkipDir
			}
			paths = append(paths, newPath)
		}
		return nil
	})
	return paths
}

// shouldIgnoreFile determines if a file should be ignored.
// File names that begin with "." or "_" are ignored by the go tool.
func shouldIgnoreFile(name string) bool {
	return strings.HasPrefix(name, ".") || strings.HasPrefix(name, "_")
}
