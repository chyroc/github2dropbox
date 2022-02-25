package main

import (
	"fmt"
	"os"

	"github.com/chyroc/github2dropbox/internal"
)

func main() {
	r := internal.NewBackup()

	err := r.Init()
	if err != nil {
		fmt.Printf("[backup] init fail: %s\n", err)
		os.Exit(1)
	}

	// fmt.Println("start download dropbox old data")
	// _ = r.DownloadDropboxBackupDir()

	r.Run()
}
