# 🛡️ cbog: Community Bot for Open-source Governance

[![Go Report Card](https://goreportcard.com/badge/github.com/kavindu/cbog)](https://goreportcard.com/report/github.com/kavindu/cbog)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**`cbog`** is a high-performance, modular GitHub Action engineered in **Go** to solve the most pressing challenges of open-source software maintainers and contributors:
- **Triage Bottlenecks**: Maintainers spend hours assigning issues, checking PR sizes, and verifying commit conventions.
- **Contributor Friction**: New contributors often don't know where to start or wait days for an issue to be assigned.
- **Premature Merges & Broken CI**: Merging PRs before reviews or required checks finish breaks the main branch.
- **Missing OSS Governance**: Many repositories lack clean `CONTRIBUTING.md`, PR checklists, or structured Issue Forms.

---

## 🧩 Modular Plugin Ecosystem

Every feature in `cbog` is decoupled into an isolated, configurable plugin toggled via `.github/cbog.yml`:

| Plugin | What It Solves | Trigger |
| :--- | :--- | :--- |
| **`commands`** | CNCF/Prow slash commands (`/lgtm`, `/approve`, `/hold`, `/merge`, `/assign`) | `issue_comment` |
| **`claim`** | Contributors can self-assign issues with `/claim` and release with `/unclaim` | `issue_comment` |
| **`welcome`** | Greets first-time contributors with friendly onboarding links | `issues (opened)`, `pull_request (opened)` |
| **`size`** | Automatically calculates PR line changes and tags `size/XS` ... `size/XL` | `pull_request (opened, sync)` |
| **`title_lint`** | Validates PR titles against Conventional Commits and sets commit status | `pull_request (opened, sync, edited)` |
| **`auto_label`** | Applies area labels based on modified file paths (e.g. `docs/**` $\rightarrow$ `area/docs`) | `pull_request (opened, sync)` |
| **`scaffolder`** | Automatically detects tech stack and opens a PR with `CONTRIBUTING.md`, PR checklists, Issue Forms, and `.github/cbog.yml` | `workflow_dispatch` |

---

## 📋 Slash Command Cheat Sheet

| Command | Description | Eligible Roles |
| :--- | :--- | :--- |
| `/claim` | Self-assign an open issue | Everyone |
| `/unclaim` | Release an assigned issue back to the community | Assigned Contributor |
| `/assign [@user...]` | Assign yourself or mentioned collaborators | Everyone |
| `/unassign [@user...]` | Remove assignees | Everyone |
| `/lgtm` | Apply `lgtm` label ("Looks Good To Me") | Maintainers (non-author) |
| `/lgtm cancel` | Remove `lgtm` label | Maintainers |
| `/approve` | Apply `approved` label and submit formal GitHub PR Review | Maintainers (non-author) |
| `/approve cancel` | Remove `approved` label | Maintainers |
| `/hold` | Apply `do-not-merge/hold` to freeze merging | Everyone |
| `/hold cancel` | Remove `do-not-merge/hold` label | Everyone |
| `/merge [squash\|merge\|rebase]` | Safely merge the PR once all status checks pass | Maintainers |
| `/close` / `/reopen` | Close or reopen an issue or PR | Author / Maintainers |
| `/help` | Print the command reference table | Everyone |

---

## ⚙️ Repository Configuration (`.github/cbog.yml`)

Anyone can drop `.github/cbog.yml` into their repository to customize behavior:

```yaml
version: 1

plugins:
  commands:
    enabled: true
    allow_self_approval: false
    default_merge_method: squash

  claim:
    enabled: true
    max_issues_per_user: 3

  welcome:
    enabled: true
    issue_message: "👋 Welcome @{{.User}}! Thanks for opening an issue in {{.Repo}}!"
    pr_message: "🎉 Welcome @{{.User}}! Thanks for opening your first PR in {{.Repo}}!"

  size:
    enabled: true
    xs: 10
    s: 30
    m: 100
    l: 500
    xl: 1000

  title_lint:
    enabled: true
    types: ["feat", "fix", "docs", "style", "refactor", "perf", "test", "chore", "revert"]

  auto_label:
    enabled: true
    rules:
      - label: "area/docs"
        paths: ["docs/**", "**/*.md"]
      - label: "area/ci"
        paths: [".github/**"]
```

---

## 🚀 30-Second Setup in Any Repository

Create `.github/workflows/cbog.yml` in your repository:

```yaml
name: cbog Open-Source Helper

on:
  issue_comment:
    types: [created]
  issues:
    types: [opened]
  pull_request:
    types: [opened, synchronize, reopened, edited]
  workflow_dispatch:
    inputs:
      mode:
        description: 'Run mode'
        required: true
        default: 'bot'
        type: choice
        options:
          - bot
          - scaffold-oss-docs

jobs:
  cbog:
    runs-on: ubuntu-latest
    permissions:
      issues: write
      pull-requests: write
      contents: write
      statuses: write
      checks: read
    steps:
      - uses: actions/checkout@v4
      - name: Run cbog
        uses: kavindu/cbog@v1
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
          mode: ${{ inputs.mode || 'bot' }}
```

---

## 🛠️ Auto-Scaffolding Community Guidelines

To scaffold complete community guidelines (`CONTRIBUTING.md`, PR templates, Issue Forms, `.github/cbog.yml`):

1. Go to the **Actions** tab in your repository.
2. Select **cbog Open-Source Helper**.
3. Click **Run workflow** $\rightarrow$ choose `scaffold-oss-docs`.
4. `cbog` will detect your project stack and open an automated Pull Request!

---

## 🧪 Local Testing

```bash
# Run unit tests
go test -v ./...

# Build binary
go build -o bin/cbog ./cmd/bot
```

---

## 📄 License

MIT © 2026
