# github2dropbox

Backup GitHub Data to DropBox: [View on Marketplace](https://github.com/marketplace/actions/github2dropbox)

Support:

- Star
- Follower
- Following
- Repo
- Issue
- Issue Comment
- Gist
- .git

## Usage

### 1. Create a new repository to run backup action

### 2. Get the access token from [DropBox](https://www.dropbox.com/developers/apps)

- 2.1 Create a new app
- 2.2 Add permission
  - files.metadata.write
  - files.content.write
  - and then submit the change.
- 2.3 Set Access token expiration to No-Expiration
- 2.4 Click Generate Access Token
  - and copy the access token.

### 3. Generate the access token from [Personal Access Token](https://github.com/settings/tokens)

- 3.1 Expiration: No-Expiration
- 3.2 Choose ALL Permissions

### 4. Config the backup action

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

## Backup File Structure

```text
GitHub/
  <username>/
    repo/
      <repo>/
        repo.json
        repo.zip
        issue/
          <id>/
            <id.json>
            comment/
              <id.json>
    star/
      <repo.json>
    follower/
      <user.json>
    following/
      <user.json>
    github2dropbox/
      meta.json
```

## Change Log


- 2022-02-24 v0.1.0
  - Initial release