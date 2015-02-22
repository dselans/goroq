package helper

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Helper function for printing a message to stdout + exit with given exit code
func CustomExit(exitMessage string, exitCode int) {
	fmt.Println(exitMessage)
	os.Exit(exitCode)
}

// Check if file exists
func FileExists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}

	return true
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
	var file *os.File
	var err error

	// If the file already exists, attempt to open for APPEND; otherwise
	// do a Create(), followed by a Remove()
	if FileExists(filename) {
		file, err = os.OpenFile(filename, os.O_APPEND, 0666)
	} else {
		file, err = os.Create(filename)
	}

	if err != nil {
		return false
	}

	file.Close()

	return true
}

// This probably needs to be a bit more sophisticated
func ExecCmd(command string, args ...string) ([]byte, error) {
	out, err := exec.Command(command, args...).CombinedOutput()
	if err != nil {
		return out, err
	}

	return out, nil
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
