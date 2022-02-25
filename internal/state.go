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
	Star  *LastProcessed            `json:"star"`
	Repos map[string]*LastProcessed `json:"repos"`
}

func (r *Backup) IsRepoProcessedRecently(repoName string) bool {
	return r.isProcessedRecently(r.meta.Repos[repoName])
}

func (r *Backup) SetRepoProcessedRecently(repoName string) {
	r.meta.Repos[repoName] = r.genProcessedRecently()
}

func (r *Backup) IsStarProcessedRecently() bool {
	return r.isProcessedRecently(r.meta.Star)
}

func (r *Backup) SetStarProcessedRecently() {
	r.meta.Star = r.genProcessedRecently()
}

func (r *Backup) metaPath() string {
	return fmt.Sprintf("%s/%s/github2dropbox/meta.json", r.backupDir, r.self.GetLogin())
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

func (r *Backup) genProcessedRecently() *LastProcessed {
	now := time.Now()
	return &LastProcessed{
		LastProcessedAt:      now.Unix(),
		LastProcessedAtTitle: now.Format("2006-01-02 15:04:05"),
	}
}
