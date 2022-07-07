package internal

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Check if a string is safe for substituting in a path.  If the
// string has substrings like ../ in it, it's not safe.  Check for
// OS-specific separator and the forward slash.  The code uses forward
// slashes before final substitution with filepath.FromSlash.  If
// we're getting unsafe strings, things have gone wrong, just panic.
// If the subpath had no separators, return the string.
func sanitizedFilePath(subpath string) string {
	if strings.Contains(subpath, fmt.Sprintf("%c", os.PathSeparator)) {
		panic(fmt.Sprintf("Unsafe result returned from github API: %s", subpath))
	}

	// Also check for a forward slash
	if strings.Contains(subpath, "/") {
		panic(fmt.Sprintf("Unsafe result returned from github API: %s", subpath))
	}

	return subpath
}

// Replace all separator characters in a string
// This replaces both OS-specific file path separators and foward slashes.
func replaceSeparatorsFilePath(path string, osPathSeparator rune, replacement string) string {
	path = strings.ReplaceAll(path, fmt.Sprintf("%c", osPathSeparator), replacement)
	return strings.ReplaceAll(path, "/", replacement)
}

func runCmd(cmd string, args []string) error {
	pwd, _ := os.Getwd()

	command := exec.Command(cmd, args...)
	command.Env = os.Environ()
	command.Dir = pwd
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	return err
}

func downloadFile(filename string, url string) error {
	output, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer output.Close()

	response, err := httpClientBigTimeout.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		return err
	}
	return nil
}

var httpClientBigTimeout = &http.Client{
	Timeout: time.Hour,
}
