package golog

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const (
	GOLOG_TEST_FILE   string = "golog_test.txt"
	GOLOG_TEST_STRING string = "Test String"
)

// Helper function for removing test log file (if it exists)
func RemoveFile() error {
	if _, err := os.Stat(GOLOG_TEST_FILE); err == nil {
		if removeErr := os.Remove(GOLOG_TEST_FILE); removeErr != nil {
			return removeErr
		}
	}
	return nil
}

func TestNew(t *testing.T) {
	// Start fresh
	if err := RemoveFile(); err != nil {
		t.Fatalf("Unable to remove existing test log file '%v'. Error: %v\n", err, GOLOG_TEST_FILE)
	}

	defer RemoveFile()

	logger, err := New(GOLOG_TEST_FILE, false)
	if err != nil {
		t.Error("Should not get any errors instantiating new logger. Error:", err)
	}

	logger.Debug.Println(GOLOG_TEST_STRING)
	logger.Info.Println(GOLOG_TEST_STRING)
	logger.Warning.Println(GOLOG_TEST_STRING)
	logger.Error.Println(GOLOG_TEST_STRING)
	logger.Critical.Println(GOLOG_TEST_STRING)

	// Manually read the file, verify we wrote what we intended to write
	contents, readErr := ioutil.ReadFile(GOLOG_TEST_FILE)
	if readErr != nil {
		t.Fatalf("Unable to read contents of '%v'\n", GOLOG_TEST_FILE)
	}

	splitContents := strings.Split(string(contents), "\n")
	requiredStrings := map[string]bool{
		"DEBUG":    false,
		"INFO":     false,
		"WARNING":  false,
		"ERROR":    false,
		"CRITICAL": false,
	}

	for key, _ := range requiredStrings {
		for _, line := range splitContents {
			if strings.HasPrefix(line, key) {
				requiredStrings[key] = true
			}
		}
	}

	for key, val := range requiredStrings {
		if !val {
			t.Errorf("Unable to find '%v' log line", key)
		}
	}
}
