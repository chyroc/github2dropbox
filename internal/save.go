package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/go-github/v42/github"
)

func (r *Backup) repoPath(repo *github.Repository) string {
	return fmt.Sprintf("%s/%s/repo/%s", r.backupDir, r.self.GetLogin(), repo.GetName())
}

func (r *Backup) repoJsonPath(repo *github.Repository) string {
	return fmt.Sprintf("%s/%s/repo/%s/repo.json", r.backupDir, r.self.GetLogin(), repo.GetName())
}

func (r *Backup) starJsonPath(star *github.StarredRepository) string {
	return fmt.Sprintf("%s/%s/star/%s.json", r.backupDir, r.self.GetLogin(), strings.ReplaceAll(star.GetRepository().GetFullName(), "/", "_"))
}

func (r *Backup) starJsonDirPath() string {
	return fmt.Sprintf("%s/%s/star", r.backupDir, r.self.GetLogin())
}

func (r *Backup) repoZipPath(repo *github.Repository) string {
	return fmt.Sprintf("%s/%s/repo/%s/repo.zip", r.backupDir, r.self.GetLogin(), repo.GetName())
}

func (r *Backup) SaveRepoJson(repo *github.Repository) {
	file := r.repoJsonPath(repo)
	_ = os.MkdirAll(filepath.Dir(file), 0o755)
	bs, err := json.MarshalIndent(repo, "", "  ")
	if err != nil {
		return
	}
	_ = ioutil.WriteFile(file, bs, 0o644)
}

func (r *Backup) SaveStarsJson(stars []*github.StarredRepository) error {
	for idx, v := range stars {
		file := r.starJsonPath(v)
		if idx == 0 {
			if err := os.MkdirAll(filepath.Dir(file), 0o755); err != nil {
				return err
			}
		}
		bs, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			return err
		}
		_ = ioutil.WriteFile(file, bs, 0o644)
	}
	return nil
}

func (r *Backup) SaveRepoZip(repo *github.Repository) {
	file := r.repoZipPath(repo)
	link, _, err := r.githubClient.Repositories.GetArchiveLink(context.Background(), *repo.Owner.Login, *repo.Name, github.Zipball, &github.RepositoryContentGetOptions{}, true)
	if err != nil {
		return
	}
	_ = os.MkdirAll(filepath.Dir(file), 0o755)
	err = downloadFile(file, link.String())
	return
}
