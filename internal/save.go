package internal

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/go-github/v42/github"
)

func (r *Backup) repoJsonPath(repo *github.Repository) string {
	return fmt.Sprintf("%s/%s/repo/%s/%s.json", r.BackupDir, r.self.GetLogin(), repo.GetName(), repo.GetName())
}

func (r *Backup) repoIssueJsonPath(repo *github.Repository, issue *github.Issue) string {
	return fmt.Sprintf("%s/%s/repo/%s/issue/%d/%d.json", r.BackupDir, r.self.GetLogin(), repo.GetName(), issue.GetID(), issue.GetID())
}

func (r *Backup) repoIssueCommentJsonPath(repo *github.Repository, issue *github.Issue, comment *github.IssueComment) string {
	return fmt.Sprintf("%s/%s/repo/%s/issue/%d/comment/%d.json", r.BackupDir, r.self.GetLogin(), repo.GetName(), issue.GetID(), comment.GetID())
}

func (r *Backup) starJsonPath(star *github.StarredRepository) string {
	return fmt.Sprintf("%s/%s/star/%s.json", r.BackupDir, r.self.GetLogin(), strings.ReplaceAll(star.GetRepository().GetFullName(), "/", "_"))
}

func (r *Backup) followerJsonPath(user *github.User) string {
	return fmt.Sprintf("%s/%s/follower/%s.json", r.BackupDir, r.self.GetLogin(), user.GetLogin())
}

func (r *Backup) followingJsonPath(user *github.User) string {
	return fmt.Sprintf("%s/%s/following/%s.json", r.BackupDir, r.self.GetLogin(), user.GetLogin())
}

func (r *Backup) gistJsonPath(data *github.Gist) string {
	return fmt.Sprintf("%s/%s/gist/%s.json", r.BackupDir, r.self.GetLogin(), data.GetID())
}

func getPathBaseName(path string) string {
	base := filepath.Base(path)
	ext := filepath.Ext(base)
	return base[:len(base)-len(ext)]
}

func (r *Backup) repoZipPath(repo *github.Repository) string {
	return fmt.Sprintf("%s/%s/repo/%s/repo.zip", r.BackupDir, r.self.GetLogin(), repo.GetName())
}

func (r *Backup) SaveRepoZip(repo *github.Repository) {
	file := r.repoZipPath(repo)
	link, _, err := r.GithubClient.Repositories.GetArchiveLink(context.Background(), *repo.Owner.Login, *repo.Name, github.Zipball, &github.RepositoryContentGetOptions{}, true)
	if err != nil {
		return
	}
	_ = os.MkdirAll(filepath.Dir(file), 0o755)
	err = downloadFile(file, link.String())
	return
}
