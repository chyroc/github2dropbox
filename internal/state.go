package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Meta struct {
	LastProcessedAt      int64  `json:"last_processed_at"`
	LastProcessedAtTitle string `json:"last_processed_at_title"`
}

func (r *Backup) IsProcessedRecently(repoName string) bool {
	file := fmt.Sprintf("%s/%s/meta.json", r.backupDir, repoName)
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		_ = os.RemoveAll(file)
		return false
	}

	var meta Meta
	err = json.Unmarshal(bs, &meta)
	if err != nil {
		_ = os.RemoveAll(file)
		return false
	}

	if meta.LastProcessedAt == 0 {
		return false
	}
	if time.Now().Unix()-meta.LastProcessedAt < 60*60*24 {
		return true
	}
	return false
}

func (r *Backup) SetProcessedRecently(repoName string) error {
	file := fmt.Sprintf("%s/%s/meta.json", r.backupDir, repoName)
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	var meta Meta
	err = json.Unmarshal(bs, &meta)
	if err != nil {
		_ = os.RemoveAll(file)
	}

	meta.LastProcessedAt = time.Now().Unix()
	meta.LastProcessedAtTitle = time.Now().Format("2006-01-02 15:04:05")
	bs, err = json.MarshalIndent(meta, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(file, bs, 0o644)
	if err != nil {
		return err
	}
	return nil
}
