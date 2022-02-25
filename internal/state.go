package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type LastProcessed struct {
	LastProcessedAt      int64  `json:"last_processed_at"`
	LastProcessedAtTitle string `json:"last_processed_at_title"`
}

type Meta struct {
	Stars      map[string]*LastProcessed `json:"stars"`
	Followers  map[string]*LastProcessed `json:"followers"`
	Followings map[string]*LastProcessed `json:"followings"`
	Repos      map[string]*LastProcessed `json:"repos"`
}

func (r *Backup) IsRepoProcessedRecently(repoName string) bool {
	return r.isProcessedRecently(r.meta.Repos[repoName])
}

func (r *Backup) SetRepoProcessedRecently(repoName string) {
	r.meta.Repos[repoName] = r.genProcessedRecently()
}

func (r *Backup) getProcessedRecentlyByTitle(title string) map[string]*LastProcessed {
	switch title {
	case "stars":
		return r.meta.Stars
	case "followers":
		return r.meta.Followers
	case "followings":
		return r.meta.Followings
	case "repos":
		return r.meta.Repos
	default:
		panic(fmt.Sprintf("unknown title: %s", title))
	}
}

const backupStars = "stars"
const backupFollowers = "followers"
const backupFollowings = "followings"
const backupRepos = "repos"

func (r *Backup) setProcessedRecentlyByTitle(title, name string) {
	switch title {
	case backupStars:
		if r.meta.Stars == nil {
			r.meta.Stars = make(map[string]*LastProcessed)
		}
		r.meta.Stars[name] = r.genProcessedRecently()
	case backupFollowers:
		if r.meta.Followers == nil {
			r.meta.Followers = make(map[string]*LastProcessed)
		}
		r.meta.Followers[name] = r.genProcessedRecently()
	case backupFollowings:
		if r.meta.Followings == nil {
			r.meta.Followings = make(map[string]*LastProcessed)
		}
		r.meta.Followings[name] = r.genProcessedRecently()
	case backupRepos:
		if r.meta.Repos == nil {
			r.meta.Repos = make(map[string]*LastProcessed)
		}
		r.meta.Repos[name] = r.genProcessedRecently()
	default:
		panic(fmt.Sprintf("unknown title: %s", title))
	}
}

func (r *Backup) metaPath() string {
	return fmt.Sprintf("%s/%s/github2dropbox/meta.json", r.BackupDir, r.self.GetLogin())
}

func (r *Backup) loadMeta() *Meta {
	file := r.metaPath()
	meta := Meta{Repos: map[string]*LastProcessed{}}

	bs, err := ioutil.ReadFile(file)
	if err != nil {
		_ = os.RemoveAll(file)
	} else if err = json.Unmarshal(bs, &meta); err != nil {
		_ = os.RemoveAll(file)
	}

	return &meta
}

func (r *Backup) saveMeta(meta *Meta) error {
	file := r.metaPath()

	bs, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(file, bs, 0o644)
	if err != nil {
		return err
	}
	return nil
}

func (r *Backup) isProcessedRecently(lp *LastProcessed) bool {
	if lp == nil || lp.LastProcessedAt == 0 {
		return false
	}

	return time.Now().Unix()-lp.LastProcessedAt < 60*60*24
}

func (r *Backup) isProcessedRecentlyBYTitle(title, name string) bool {
	lps := r.getProcessedRecentlyByTitle(title)
	if lps == nil {
		return false
	}

	lp := lps[name]
	if lp == nil || lp.LastProcessedAt == 0 {
		return false
	}

	return time.Now().Unix()-lp.LastProcessedAt < 60*60*24
}

func (r *Backup) genProcessedRecently() *LastProcessed {
	now := time.Now()
	return &LastProcessed{
		LastProcessedAt:      now.Unix(),
		LastProcessedAtTitle: now.Format("2006-01-02 15:04:05"),
	}
}
