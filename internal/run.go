package internal

import (
	"fmt"
)

func (r *Backup) Run() {
	_ = r.SaveStar()
	_ = r.SaveRepos()
}

func (r *Backup) SaveRepos() error {
	repos, err := r.AllRepo()
	if err != nil {
		fmt.Printf("[repo] get repo fail[ignore err]: %s\n", err)
		return err
	}
	fmt.Printf("[repo] get repo, count: %d\n", len(repos))

	for _, repo := range repos {
		fmt.Printf("[repo:%s] start\n", repo.GetName())

		if r.IsRepoProcessedRecently(repo.GetName()) {
			fmt.Printf("[repo:%s] processed recently, skip\n", repo.GetName())
			continue
		}

		r.SaveRepoJson(repo)
		r.SaveRepoZip(repo)
		r.SetRepoProcessedRecently(repo.GetName())

		if err := r.Upload(r.repoPath(repo)); err != nil {
			fmt.Printf("[repo:%s] upload to dropbox fail[ignore err]: %s\n", repo.GetName(), err)
		} else {
			fmt.Printf("[repo:%s] upload to dropbox success\n", repo.GetName())
		}
		// _ = r.UploadMeta()
		fmt.Println(r.UploadMeta())
	}

	return nil
}

func (r *Backup) SaveStar() error {
	if r.IsStarProcessedRecently() {
		fmt.Printf("[star] processed recently, skip\n")
		return nil
	}

	stars, err := r.AllStar()
	if err != nil {
		return err
	}
	fmt.Printf("[star] get star, count: %d\n", len(stars))

	if err = r.SaveStarsJson(stars); err != nil {
		fmt.Printf("[star] save json fail: %s\n", err)
		return err
	}
	r.SetStarProcessedRecently()

	if err = r.Upload(r.starJsonDirPath()); err != nil {
		fmt.Printf("[star] upload to dropbox fail[ignore err]: %s\n", err)
	} else {
		fmt.Printf("[star] upload to dropbox success\n")
	}

	return r.UploadMeta()
}
