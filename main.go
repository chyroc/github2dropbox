package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/chyroc/github2dropbox/internal"
)

func NewOption() *internal.Option {
	r := new(internal.Option)

	r.DropboxCli = "dropbox-cli"
	r.GithubToken = os.Getenv("INPUT_GITHUB_TOKEN")
	r.DropboxPath = strings.TrimRight(os.Getenv("INPUT_DROPBOX_PATH"), "/") + "/"
	r.DropboxToken = os.Getenv("DROPBOX_TOKEN_BACKUP")
	r.BackupDir = "GitHub"

	if r.GithubToken == "" {
		r.GithubToken = os.Getenv("GITHUB_TOKEN")
	}
	if r.DropboxToken == "" {
		r.DropboxToken = os.Getenv("DROPBOX_TOKEN")
	}
	if s, _ := exec.LookPath(r.DropboxCli); s != "" {
		r.DropboxCli = s
	} else {
		r.DropboxCli = "/bin/dropbox-cli"
	}

	return r
}

func main() {
	r := internal.NewBackup(NewOption())

	r.Run()
}
