package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/go-github/v42/github"
)

func (r *Backup) Run() {
	if err := r.Init(); err != nil {
		os.Exit(1)
		return
	}

	if r.EnableStar {
		_ = r.SaveStar()
	} else {
		fmt.Println("Star is disabled")
	}
	if r.EnableFollower {
		_ = r.SaveFollower()
	} else {
		fmt.Println("Follower is disabled")
	}
	if r.EnableFollowing {
		_ = r.SaveFollowing()
	} else {
		fmt.Println("Following is disabled")
	}
	if r.EnableRepo {
		_ = r.SaveRepos(r.EnableIssue)
	} else {
		fmt.Println("Repo is disabled")
	}
	if r.EnableGist {

	} else {
		fmt.Println("Gist is disabled")
	}
}

func (r *Backup) DownloadMeta() error {
	return r.Download(r.metaPath())
}

func (r *Backup) SaveRepos(issuesEnabled bool) error {
	return saveDataList(r, backupRepos, r.AllRepo, r.repoJsonPath, 1, func(data *github.Repository) {
		r.SaveRepoZip(data)
	})
}

func (r *Backup) SaveStar() error {
	return saveDataList(r, backupStars, r.AllStar, r.starJsonPath, 0)
}

func (r *Backup) SaveFollower() error {
	return saveDataList(r, backupFollowers, r.AllFollower, r.followerJsonPath, 0)
}

func (r *Backup) SaveFollowing() error {
	return saveDataList(r, backupFollowings, r.AllFollowing, r.followingJsonPath, 0)
}

func saveDataList[T any](
	r *Backup,
	title string,
	listFunc func() ([]T, error),
	genPath func(T) string,
	uploadDepth int,
	additionalFunc ...func(T),
) error {
	dataList, err := listFunc()
	if err != nil {
		fmt.Printf("[%s] get data, fail: %s\n", title, err)
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
			jsonPath := genPath(v)
			name := getPathBaseName(jsonPath)

			if r.isProcessedRecentlyBYTitle(title, name) {
				fmt.Printf("[%s:%s] processed recently, skip\n", title, name)
				continue
			}

			saveData(title, name, jsonPath, v, additionalFunc...)
			r.setProcessedRecentlyByTitle(title, name)

			if err = r.Upload(genUploadPath(jsonPath, uploadDepth)); err != nil {
				fmt.Printf("[%s] upload to dropbox fail[ignore err]: %s\n", title, err)
			} else {
				fmt.Printf("[%s] upload to dropbox success\n", title)
			}

			_ = r.UploadMeta()
		}
	}

	return nil
}

func saveData[T any](title, name, jsonPath string, data T, additionalFunc ...func(T)) {
	// json
	{
		bs, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			fmt.Printf("[%s:%s] save json fail: %s\n", title, name, err)
		} else {
			_ = ioutil.WriteFile(jsonPath, bs, 0o644)
		}
	}

	// additional
	for _, v := range additionalFunc {
		v(data)
	}
}

func genUploadPath(jsonPath string, depth int) string {
	for depth > 0 {
		depth--
		jsonPath = filepath.Dir(jsonPath)
	}
	return jsonPath
}
