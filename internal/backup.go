package internal

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"strings"

	"github.com/google/go-github/v42/github"
	"golang.org/x/oauth2"
)

type Backup struct {
	githubClient *github.Client
	backupDir    string
	dropboxToken string
	dropboxPath  string
	dropboxCli   string
	self         *github.User
	meta         *Meta
}

func NewBackup() *Backup {
	cli := new(Backup)

	var (
		dropboxCli   = "dropbox-cli"
		dropboxPath  = strings.TrimRight(os.Getenv("INPUT_DROPBOX_PATH"), "/") + "/"
		githubToken  = os.Getenv("INPUT_GITHUB_TOKEN")
		dropboxToken = os.Getenv("DROPBOX_TOKEN")
		backupDir    = "GitHub"
	)
	if githubToken == "" {
		githubToken = os.Getenv("GITHUB_TOKEN")
	}
	if s, _ := exec.LookPath(dropboxCli); s != "" {
		dropboxCli = s
	} else {
		dropboxCli = "/bin/dropbox-cli"
	}

	httpClient := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	))

	cli.githubClient = github.NewClient(httpClient)
	cli.backupDir = backupDir
	cli.dropboxToken = dropboxToken
	cli.dropboxPath = dropboxPath
	cli.dropboxCli = dropboxCli

	return cli
}

func (r *Backup) Init() error {
	if r.dropboxToken == "" {
		return errors.New("dropbox token is empty")
	}

	user, err := r.SelfUser()
	if err != nil {
		return err
	}
	r.self = user
	if r.self.GetLogin() == "" {
		return errors.New("No login found")
	}

	r.meta = r.loadMeta()

	return nil
}
