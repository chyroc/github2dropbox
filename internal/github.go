package internal

import (
	"context"

	"github.com/google/go-github/v42/github"
)

func (r *Backup) AllRepo() ([]*github.Repository, error) {
	page := 1
	resp := []*github.Repository{}

	for page != 0 {
		repos, response, err := r.GithubClient.Repositories.List(context.Background(), "", &github.RepositoryListOptions{
			Visibility:  "all",
			Affiliation: "owner",
			ListOptions: github.ListOptions{
				Page:    page,
				PerPage: 100,
			},
		})
		if err != nil {
			return nil, err
		}
		resp = append(resp, repos...)
		page = response.NextPage
	}
	return resp, nil
}

func (r *Backup) AllIssueByRepo(repo *github.Repository) ([]*github.Issue, error) {
	page := 1
	resp := []*github.Issue{}

	for page != 0 {
		repos, response, err := r.GithubClient.Issues.ListByRepo(context.Background(), r.self.GetLogin(), repo.GetName(), &github.IssueListByRepoOptions{
			State: "all",
			ListOptions: github.ListOptions{
				Page:    page,
				PerPage: 100,
			},
		})
		if err != nil {
			return nil, err
		}
		resp = append(resp, repos...)
		page = response.NextPage
	}
	return resp, nil
}

func (r *Backup) AllStar() ([]*github.StarredRepository, error) {
	page := 1
	var stars []*github.StarredRepository
	for page > 0 {
		starredRepository, resp, err := r.GithubClient.Activity.ListStarred(context.Background(), "", &github.ActivityListStarredOptions{
			ListOptions: github.ListOptions{
				Page:    page,
				PerPage: 100,
			},
		})
		if err != nil {
			return nil, err
		}
		stars = append(stars, starredRepository...)
		page = resp.NextPage
	}
	return stars, nil
}

func (r *Backup) AllFollower() ([]*github.User, error) {
	page := 1
	var stars []*github.User
	for page > 0 {
		starredRepository, resp, err := r.GithubClient.Users.ListFollowers(context.Background(), "", &github.ListOptions{
			Page:    page,
			PerPage: 100,
		})
		if err != nil {
			return nil, err
		}
		stars = append(stars, starredRepository...)
		page = resp.NextPage
	}
	return stars, nil
}

func (r *Backup) AllFollowing() ([]*github.User, error) {
	page := 1
	var stars []*github.User
	for page > 0 {
		starredRepository, resp, err := r.GithubClient.Users.ListFollowing(context.Background(), "", &github.ListOptions{
			Page:    page,
			PerPage: 100,
		})
		if err != nil {
			return nil, err
		}
		stars = append(stars, starredRepository...)
		page = resp.NextPage
	}
	return stars, nil
}

func (r *Backup) AllGist() ([]*github.Gist, error) {
	page := 1
	var dataList []*github.Gist
	for page > 0 {
		starredRepository, resp, err := r.GithubClient.Gists.List(context.Background(), r.self.GetLogin(), &github.GistListOptions{
			ListOptions: github.ListOptions{
				Page:    page,
				PerPage: 100,
			},
		})
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, starredRepository...)
		page = resp.NextPage
	}
	return dataList, nil
}

func (r *Backup) AllIssueComment(repo string, id int) ([]*github.IssueComment, error) {
	page := 1
	var dataList []*github.IssueComment
	for page > 0 {
		starredRepository, resp, err := r.GithubClient.Issues.ListComments(context.Background(), r.self.GetLogin(), repo, id, &github.IssueListCommentsOptions{
			ListOptions: github.ListOptions{
				Page:    page,
				PerPage: 100,
			},
		})
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, starredRepository...)
		page = resp.NextPage
	}
	return dataList, nil
}

func (r *Backup) SelfUser() (*github.User, error) {
	user, _, err := r.GithubClient.Users.Get(context.Background(), "")
	return user, err
}
