package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func (r *Backup) DownloadDropboxBackupDir() error {
	return runCmd(r.dropboxCli,
		[]string{
			"download",
			"--token", r.dropboxToken,
			r.dropboxPath + r.backupDir, // remote path
			r.backupDir,                 // local path
		})
}

func (r *Backup) Upload(path string) error {
	return runCmd(r.dropboxCli,
		[]string{
			"upload",
			"--token", r.dropboxToken,
			path,
			r.dropboxPath + path,
		})
}

func (r *Backup) UploadMeta() error {
	bs, err := json.MarshalIndent(r.meta, "", "  ")
	if err != nil {
		return err
	}
	file := fmt.Sprintf("%s/%s/github2dropbox/meta.json", r.backupDir, r.self.GetLogin())
	if err = os.MkdirAll(filepath.Dir(file), 0o755); err != nil {
		return err
	}
	if err = ioutil.WriteFile(file, bs, 0o644); err != nil {
		return err
	}
	return r.Upload(file)
}
