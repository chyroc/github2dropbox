package internal

import (
	"context"
	"os"
	"strings"

	"github.com/google/go-github/v42/github"
	"golang.org/x/oauth2"
)

type Backup struct {
	githubClient *github.Client
	backupDir    string
	dropboxToken string
	dropboxPath  string
}

func NewBackup() *Backup {
	cli := new(Backup)

	var (
		dropboxPath  = strings.TrimRight(os.Getenv("INPUT_DROPBOX_PATH"), "/") + "/"
		githubToken  = os.Getenv("INPUT_GITHUB_TOKEN")
		dropboxToken = os.Getenv("DROPBOX_TOKEN")
		backupDir    = "GitHub"
	)

	httpClient := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	))

	cli.githubClient = github.NewClient(httpClient)
	cli.backupDir = backupDir
	cli.dropboxToken = dropboxToken
	cli.dropboxPath = dropboxPath

	return cli
}
