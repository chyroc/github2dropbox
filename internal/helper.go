package internal

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"
)

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
