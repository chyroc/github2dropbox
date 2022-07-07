package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// Test the case when the file path contains forward slashes.
func TestSanitizedFilePathWithForwardSlashesPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code should panic with forward slashes")
		}
	}()

	test_path := "directory/subdirectory/file.txt"

	sanitizedFilePath(test_path)
}

// Test the case when the file path contains OS specific separators.
// This is redundant and may be testing the same situation as above.
// But the code explicitly uses forward slashes until the final path
// is built.
func TestSanitizedFilePathPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code should panic with file separators")
		}
	}()

	test_path := filepath.Join("directory", "subdirectory", "file.txt")

	sanitizedFilePath(test_path)
}

// Test that a file path with no separators works.
func TestSanitizedFilePathWorks(t *testing.T) {
	test_path := filepath.Join("file.txt")

	result := sanitizedFilePath(test_path)

	if result != "file.txt" {
		t.Error("Expected file.txt got", result)
	}
}

// Test that replacing forward slashes with a replacement character
// works.
func TestReplaceSeparatorsFilePathWithForwardSlashesWorks(t *testing.T) {
	result := replaceSeparatorsFilePath("directory/file.txt", os.PathSeparator, "_")

	if result != "directory_file.txt" {
		t.Error("Expected directory_file.txt got", result)
	}
}

// Test that replacing OS-specific separator with a replacement
// character works.  This exercises the tests on the supported system
// when it's run on them.
func TestReplaceSeparatorsFilePathWithOSSeparatorsWorks(t *testing.T) {
	test_string := fmt.Sprintf("directory%cfile.txt", os.PathSeparator)
	result := replaceSeparatorsFilePath(test_string, os.PathSeparator, "_")

	if result != "directory_file.txt" {
		t.Error("Expected directory_file.txt got", result)
	}
}
