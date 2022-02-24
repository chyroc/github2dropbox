# github2dropbox

Backup GitHub Data to DropBox

## Usage

add file: `.github/workflows/github-backup.yml` to your project

```yaml
name: github-backup

on:
  push:
    branches: [ master ] # trigger on pushes to master
  pull_request: # trigger on pull requests

jobs:

  run:
    runs-on: ubuntu-latest
    timeout-minutes: 60 # timeout after 60 minutes
    steps:
      - name: Backup
        uses: chyroc/github2dropbox@a423fe2
        with:
          DROPBOX_TOKEN: ${{ secrets.DROPBOX_TOKEN }} # dropbox token
          GITHUB_TOKEN: ${{ secrets.G_TOKEN }} # github token
```