package internal

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/go-github/v42/github"
)

func (r *Backup) repoJsonPath(repo *github.Repository) string {
	return fmt.Sprintf("%s/%s/repo/%s/%s.json", r.BackupDir, sanitizedFilePath(r.self.GetLogin()), sanitizedFilePath(repo.GetName()), sanitizedFilePath(repo.GetName()))
}

func (r *Backup) repoGitZipPath(repo *github.Repository) string {
	return fmt.Sprintf("%s/%s/repo/%s/%s.git.zip", r.BackupDir, sanitizedFilePath(r.self.GetLogin()), sanitizedFilePath(repo.GetName()), sanitizedFilePath(repo.GetName()))
}

func (r *Backup) repoIssueJsonPath(repo *github.Repository, issue *github.Issue) string {
	return fmt.Sprintf("%s/%s/repo/%s/issue/%d/%d.json", r.BackupDir, sanitizedFilePath(r.self.GetLogin()), sanitizedFilePath(repo.GetName()), issue.GetID(), issue.GetID())
}

func (r *Backup) repoIssueCommentJsonPath(repo *github.Repository, issue *github.Issue, comment *github.IssueComment) string {
	return fmt.Sprintf("%s/%s/repo/%s/issue/%d/comment/%d.json", r.BackupDir, sanitizedFilePath(r.self.GetLogin()), sanitizedFilePath(repo.GetName()), issue.GetID(), comment.GetID())
}

func (r *Backup) starJsonPath(star *github.StarredRepository) string {
	return fmt.Sprintf("%s/%s/star/%s.json", r.BackupDir, sanitizedFilePath(r.self.GetLogin()), replaceSeparatorsFilePath(star.GetRepository().GetFullName(), os.PathSeparator, "_"))
}

func (r *Backup) followerJsonPath(user *github.User) string {
	return fmt.Sprintf("%s/%s/follower/%s.json", r.BackupDir, sanitizedFilePath(r.self.GetLogin()), sanitizedFilePath(user.GetLogin()))
}

func (r *Backup) followingJsonPath(user *github.User) string {
	return fmt.Sprintf("%s/%s/following/%s.json", r.BackupDir, sanitizedFilePath(r.self.GetLogin()), sanitizedFilePath(user.GetLogin()))
}

func (r *Backup) gistJsonPath(data *github.Gist) string {
	return fmt.Sprintf("%s/%s/gist/%s.json", r.BackupDir, sanitizedFilePath(r.self.GetLogin()), sanitizedFilePath(data.GetID()))
}

func getPathBaseName(path string) string {
	base := filepath.Base(path)
	ext := filepath.Ext(base)
	return base[:len(base)-len(ext)]
}

func (r *Backup) repoZipPath(repo *github.Repository) string {
	return fmt.Sprintf("%s/%s/repo/%s/repo.zip", r.BackupDir, sanitizedFilePath(r.self.GetLogin()), sanitizedFilePath(repo.GetName()))
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
