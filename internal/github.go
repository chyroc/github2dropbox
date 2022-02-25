package internal

import (
	"context"

	"github.com/google/go-github/v42/github"
)

func (r *Backup) AllRepo() ([]*github.Repository, error) {
	page := 1
	resp := []*github.Repository{}

	for page != 0 {
		repos, response, err := r.githubClient.Repositories.List(context.Background(), "", &github.RepositoryListOptions{
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

func (r *Backup) AllStar() ([]*github.StarredRepository, error) {
	page := 1
	var stars []*github.StarredRepository
	for page > 0 {
		starredRepository, resp, err := r.githubClient.Activity.ListStarred(context.Background(), "", &github.ActivityListStarredOptions{
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

func (r *Backup) SelfUser() (*github.User, error) {
	user, _, err := r.githubClient.Users.Get(context.Background(), "")
	return user, err
}
