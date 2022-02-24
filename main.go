package main

import (
	"fmt"

	"github.com/chyroc/github2dropbox/internal"
)

func main() {
	r := internal.NewBackup()

	fmt.Println("start download dropbox old data")
	_ = r.DownloadDropboxBackupDir()

	repos, err := r.AllRepo()
	if err != nil {
		panic(err)
	}

	for _, repo := range repos {
		fmt.Printf("[%s] start\n", repo.GetName())
		if r.IsProcessedRecently(repo.GetName()) {
			fmt.Printf("[%s] processed recently, skip\n", repo.GetName())
			continue
		}

		r.SaveRepoJson(repo)
		fmt.Printf("[%s] save repo json\n", repo.GetName())

		r.SaveZipFile(repo)
		fmt.Printf("[%s] save zip file\n", repo.GetName())

		_ = r.SetProcessedRecently(repo.GetName())

		fmt.Printf("[%s] upload to dropbox\n", repo.GetName())
		if err := r.UploadRepo(repo); err != nil {
			panic(err)
		}
	}
}
