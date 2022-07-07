package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func (r *Backup) Download(path string) error {
	return runCmd(r.DropboxCli,
		[]string{
			"download",
			"--token", r.DropboxToken,
			r.DropboxPath + path, // remote path
			path,                 // local path
		})
}

func (r *Backup) Upload(path string) error {
	return runCmd(r.DropboxCli,
		[]string{
			"upload",
			"--token", r.DropboxToken,
			path,
			r.DropboxPath + path,
		})
}

func (r *Backup) UploadMeta() error {
	bs, err := json.MarshalIndent(r.meta, "", "  ")
	if err != nil {
		return err
	}
	file := fmt.Sprintf("%s/%s/github2dropbox/meta.json", r.BackupDir,
		sanitizedFilePath(r.self.GetLogin()))
	if err = os.MkdirAll(filepath.Dir(file), 0o755); err != nil {
		return err
	}
	if err = ioutil.WriteFile(file, bs, 0o644); err != nil {
		return err
	}
	return r.Upload(file)
}
