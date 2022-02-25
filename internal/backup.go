package internal

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/go-github/v42/github"
	"golang.org/x/oauth2"
)

type Backup struct {
	*Option
	GithubClient *github.Client

	self *github.User
	meta *Meta
}

type Option struct {
	BackupDir    string
	DropboxToken string
	DropboxPath  string
	DropboxCli   string
	GithubToken  string

	EnableStar         bool
	EnableFollower     bool
	EnableFollowing    bool
	EnableRepo         bool
	EnableGist         bool
	EnableIssue        bool
	EnableIssueComment bool
}

func NewBackup(option *Option) *Backup {
	if option.DropboxToken == "" {
		panic(errors.New("dropbox token is empty"))
	}
	return &Backup{
		Option: option,
		GithubClient: github.NewClient(
			oauth2.NewClient(
				context.Background(),
				oauth2.StaticTokenSource(
					&oauth2.Token{AccessToken: option.GithubToken}))),
	}
}

func (r *Backup) Init() error {
	// GitHub user
	{
		user, err := r.SelfUser()
		if err != nil {
			fmt.Printf("get github user, fail: %s\n", err)
			return err
		}
		r.self = user
		if r.self.GetLogin() == "" {
			fmt.Printf("get github user, fail: %s\n", "no name")
			return errors.New("no name")
		}
	}

	// fetch meta json
	{
		if err := r.DownloadMeta(); err != nil {
			fmt.Printf("download meta, fail: %s\n", err)
			return err
		}
		r.meta = r.loadMeta()
	}

	return nil
}
