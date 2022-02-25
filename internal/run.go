package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func (r *Backup) Run() {
	_ = r.SaveFollower()
	_ = r.SaveFollowing()
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
	return saveDataList(r, "star", r.meta.Star, func(lp *LastProcessed) {
		r.meta.Star = lp
	}, r.AllStar, r.starJsonPath)
}

func (r *Backup) SaveFollower() error {
	return saveDataList(r, "follower", r.meta.Follower, func(lp *LastProcessed) {
		r.meta.Follower = lp
	}, r.AllFollower, r.followerJsonPath)
}

func (r *Backup) SaveFollowing() error {
	return saveDataList(r, "following", r.meta.Following, func(lp *LastProcessed) {
		r.meta.Follower = lp
	}, r.AllFollowing, r.followingJsonPath)
}

func saveDataList[T any](r *Backup,
	title string,
	lastProcessed *LastProcessed,
	setLastProcessed func(lp *LastProcessed),
	listFunc func() ([]T, error),
	genPath func(T) string,
) error {
	if r.isProcessedRecently(lastProcessed) {
		fmt.Printf("[%s] processed recently, skip\n", title)
		return nil
	}

	dataList, err := listFunc()
	if err != nil {
		return err
	}
	fmt.Printf("[%s] get data, count: %d\n", title, len(dataList))

	dir := filepath.Dir(genPath(dataList[0]))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	// save json
	{
		for _, v := range dataList {
			bs, err := json.MarshalIndent(v, "", "  ")
			if err != nil {
				fmt.Printf("[%s] save json fail: %s\n", title, err)
				return err
			}
			_ = ioutil.WriteFile(genPath(v), bs, 0o644)
		}
	}

	setLastProcessed(r.genProcessedRecently())

	if err = r.Upload(dir); err != nil {
		fmt.Printf("[%s] upload to dropbox fail[ignore err]: %s\n", title, err)
	} else {
		fmt.Printf("[%s] upload to dropbox success\n", title)
	}

	return r.UploadMeta()
}
