package internal

import (
	"github.com/google/go-github/v42/github"
)

func (r *Backup) DownloadDropboxBackupDir() error {
	return runCmd("/bin/dropbox-cli",
		[]string{
			"download",
			"--token", r.dropboxToken,
			r.dropboxPath + r.backupDir, // remote path
			r.backupDir,                 // local path
		})
}

func (r *Backup) UploadRepo(repo *github.Repository) error {
	return runCmd("/bin/dropbox-cli",
		[]string{
			"upload",
			"--token", r.dropboxToken,
			r.backupDir + "/" + repo.GetName(),
			r.dropboxPath + r.backupDir + "/" + repo.GetName(),
		})
}
