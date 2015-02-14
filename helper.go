package main

import (
	"fmt"
	"os"
)

// Helper function for exiting the program with a message
func HelperCustomExit(exitMessage string, exitCode int) {
	fmt.Println(exitMessage)
	os.Exit(exitCode)
}

// Helper method to determine
func HelperIsDir(dir string) bool {
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

func HelperIsWritable(filename string) bool {
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
