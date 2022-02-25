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

	r.EnableRepo = os.Getenv("INPUT_ENABLE_REPO") == "true"
	r.EnableStar = os.Getenv("INPUT_ENABLE_STAR") == "true"
	r.EnableFollower = os.Getenv("INPUT_ENABLE_FOLLOWER") == "true"
	r.EnableFollowing = os.Getenv("INPUT_ENABLE_FOLLOWING") == "true"
	r.EnableGist = os.Getenv("INPUT_ENABLE_GIST") == "true"
	r.EnableIssue = os.Getenv("INPUT_ENABLE_ISSUE") == "true"
	r.EnableIssueComment = os.Getenv("ENABLE_ISSUE_COMMENT") == "true"

	if r.DropboxToken == "" {
		r.DropboxToken = os.Getenv("INPUT_DROPBOX_TOKEN")
	}
	if s, _ := exec.LookPath(r.DropboxCli); s != "" {
		r.DropboxCli = s
	} else {
		r.DropboxCli = "/bin/dropbox-cli"
	}

	return r
}

func main() {
	opt := NewOption()
	r := internal.NewBackup(opt)

	r.Run()
}
