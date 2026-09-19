# 🤖 Contributor Bot & Universal OSS Scaffolder

[![Go Report Card](https://goreportcard.com/badge/github.com/kavindu/contributor-bot-action)](https://goreportcard.com/report/github.com/kavindu/contributor-bot-action)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A blazing-fast, Go-powered GitHub Action that automates CNCF/Kubernetes Prow–style slash commands (`/lgtm`, `/approve`, `/hold`, `/merge`, `/assign`) and auto-scaffolds modern open-source community guidelines for **any programming language**.

---

## ✨ Key Features

- **⚡ Blazing Fast (<100ms)**: Written in Go and compiled as a standalone static binary (`CGO_ENABLED=0`). Zero Node/npm dependencies, zero Docker overhead.
- **🛡️ Secure RBAC**: Prevents unauthorized merges and blocks authors from self-approving PRs.
- **👀 Real-Time Feedback**: Immediate emoji reactions (`👀` for processing, `🚀` for merge success, `+1`, `❌` for errors) and detailed diagnostic comments.
- **📦 Universal OSS Scaffolder**: Run `mode: scaffold-oss-docs` and the bot will inspect your repository (detecting Go, Python, Rust, Node, Java, etc.) and open a Pull Request with tailored `CONTRIBUTING.md`, PR checklists, and modern GitHub Issue Forms.
- **🧩 1-Click "Suggested Workflows" Setup**: Ready to be integrated into any repo or organization starter workflow catalog.

---

## 📋 Slash Command Reference

Contributors and maintainers can use these commands in issue or pull request comments:

| Command | Description | Eligible Roles |
| :--- | :--- | :--- |
| `/assign [@user...]` | Assign yourself or mentioned users to the issue/PR | Everyone |
| `/unassign [@user...]` | Remove assignees | Everyone |
| `/lgtm` | Apply `lgtm` label ("Looks Good To Me") | Maintainers (non-author) |
| `/lgtm cancel` | Remove `lgtm` label | Maintainers |
| `/approve` | Apply `approved` label and submit formal GitHub PR Review approval | Maintainers (non-author) |
| `/approve cancel` | Remove `approved` label | Maintainers |
| `/hold` | Apply `do-not-merge/hold` to prevent premature merges | Everyone |
| `/hold cancel` | Remove `do-not-merge/hold` label | Everyone |
| `/merge [squash\|merge\|rebase]` | Safely merge the PR once all checks pass | Maintainers |
| `/close` | Close an issue or PR | Author / Maintainers |
| `/reopen` | Reopen a closed issue or PR | Author / Maintainers |
| `/help` | Print the command cheat sheet | Everyone |

---

## 🚀 Quickstart: Adding to Your Repository

Add this file to `.github/workflows/contributor-bot.yml` in your repository:

```yaml
name: Contributor Bot & OSS Helper

on:
  issue_comment:
    types: [created]
  workflow_dispatch:
    inputs:
      mode:
        description: 'Run mode (bot or scaffold-oss-docs)'
        required: true
        default: 'bot'
        type: choice
        options:
          - bot
          - scaffold-oss-docs

jobs:
  contributor-bot:
    runs-on: ubuntu-latest
    permissions:
      issues: write
      pull-requests: write
      contents: write
      statuses: read
      checks: read
    steps:
      - name: Checkout Code
        uses: actions/checkout@v4

      - name: Run Contributor Bot
        uses: kavindu/contributor-bot-action@v1
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
          mode: ${{ inputs.mode || 'bot' }}
          merge-method: 'squash'
```

---

## 🛠️ Auto-Scaffolding Community Guidelines

To auto-generate your repository's `CONTRIBUTING.md`, PR template, and GitHub Issue Forms:

1. Go to the **Actions** tab in your repository.
2. Select **Contributor Bot & OSS Helper**.
3. Click **Run workflow** $\rightarrow$ select `scaffold-oss-docs`.
4. The bot will automatically create a branch and open a Pull Request ready for you to review and merge with `/merge`!

---

## 📦 Publishing to GitHub Marketplace

1. Push this repository to GitHub as a **public** repository.
2. Navigate to **Releases** $\rightarrow$ **Draft a new release**.
3. Create a tag (e.g., `v1.0.0` and `v1`).
4. Check the box **"Publish this Action to the GitHub Marketplace"**.
5. Select a category (e.g., *Community* and *Automation*), accept terms, and click **Publish release**.

---

## 🧪 Local Development & Testing

```bash
# Run unit tests
go test -v ./...

# Build binary
go build -o bin/bot ./cmd/bot
```

---

## 📄 License

MIT © 2026
