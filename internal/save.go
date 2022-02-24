package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/go-github/v42/github"
)

func (r *Backup) SaveRepoJson(repo *github.Repository) {
	file := fmt.Sprintf("%s/%s/repo.json", r.backupDir, repo.GetName())
	_ = os.MkdirAll(filepath.Dir(file), 0o755)
	bs, err := json.MarshalIndent(repo, "", "  ")
	if err != nil {
		return
	}
	_ = ioutil.WriteFile(file, bs, 0o644)
}

func (r *Backup) SaveZipFile(repo *github.Repository) {
	link, _, err := r.githubClient.Repositories.GetArchiveLink(context.Background(), *repo.Owner.Login, *repo.Name, github.Zipball, &github.RepositoryContentGetOptions{}, true)
	if err != nil {
		return
	}
	file := fmt.Sprintf("%s/%s/repo.zip", r.backupDir, repo.GetName())
	_ = os.MkdirAll(filepath.Dir(file), 0o755)
	err = downloadFile(file, link.String())
	return
}
