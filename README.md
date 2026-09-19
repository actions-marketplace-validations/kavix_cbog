# 🛡️ cbog: Community Bot for Open-source Governance

[![Go Report Card](https://goreportcard.com/badge/github.com/kavix/cbog)](https://goreportcard.com/report/github.com/kavix/cbog)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Marketplace](https://img.shields.io/badge/Marketplace-cbog-blue.svg)](https://github.com/marketplace/actions/cbog-community-bot-for-open-source-governance)

> **Supercharge your open-source repo in 60 seconds.**  
> Automated triage, `/lgtm`, `/merge`, contributor self-assignment (`/claim`), PR size labeling, and one-click community guidelines scaffolding.

---

## ⚡ Quickstart (1 Step, Zero Config)

To add `cbog` to your repository, create **`.github/workflows/cbog.yml`**:

```yaml
name: cbog

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
        description: 'Mode'
        required: true
        default: 'bot'
        type: choice
        options: [bot, scaffold-oss-docs]

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
      - uses: kavix/cbog@v1
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
          mode: ${{ inputs.mode || 'bot' }}
```

**That's it! 🎉**  
`cbog` is now active on your issues and pull requests.

---

## 🚀 Auto-Generate Your Community Guidelines (1 Click)

Missing `CONTRIBUTING.md`, PR checklists, or modern Issue Forms? Let `cbog` write them for you:

1. In your repository, click the **Actions** tab.
2. Select **cbog** from the sidebar.
3. Click **Run workflow** $\rightarrow$ select **`scaffold-oss-docs`** $\rightarrow$ click **Run workflow**.

`cbog` will detect your project stack (Go, Python, TypeScript, Rust, etc.) and open an automated Pull Request containing:
- 📄 **`CONTRIBUTING.md`** (Tailored build instructions + slash command cheat sheet)
- 📋 **`.github/PULL_REQUEST_TEMPLATE.md`** (Structured PR checklist)
- 🐛 **`.github/ISSUE_TEMPLATE/`** (Modern YAML GitHub Issue Forms for Bug Reports & Feature Requests)
- ⚙️ **`.github/cbog.yml`** (Ready-to-tweak configuration file)

---

## 💬 Contributor & Maintainer Commands

Comment on any Issue or Pull Request to trigger commands:

| Command | What it does | Who can use it |
| :--- | :--- | :--- |
| **`/claim`** | Assigns the issue to you so you can start working on it | Everyone |
| **`/unclaim`** | Releases the issue back to the community if you're busy | Assignee |
| **`/assign [@user]`** | Assigns yourself or mentioned collaborators | Everyone |
| **`/unassign [@user]`** | Removes assignees | Everyone |
| **`/lgtm`** | Marks PR with `lgtm` ("Looks Good To Me") | Maintainers (non-author) |
| **`/lgtm cancel`** | Removes `lgtm` label | Maintainers |
| **`/approve`** | Submits formal GitHub PR approval + `approved` label | Maintainers (non-author) |
| **`/approve cancel`** | Removes approval label | Maintainers |
| **`/hold`** | Adds `do-not-merge/hold` to prevent accidental merges | Everyone |
| **`/hold cancel`** | Removes hold label | Everyone |
| **`/merge`** | Safely merges the PR (verifies CI checks & branch protection) | Maintainers |
| **`/merge squash`** | Merges with squash (or `merge`, `rebase`) | Maintainers |
| **`/close`** / **`/reopen`** | Closes or reopens an issue/PR | Author / Maintainers |
| **`/help`** | Displays this cheat sheet directly in the issue/PR | Everyone |

---

## 🧩 Included Features & Plugins

| Feature | Description |
| :--- | :--- |
| 🏷️ **Auto PR Sizing** | Automatically tags PRs with `size/XS` (<10 lines), `size/S`, `size/M`, `size/L`, or `size/XL` (>500 lines). |
| ✍️ **PR Title Linter** | Validates PR titles against Conventional Commits (`feat:`, `fix:`, `docs:`) and reports GitHub status checks. |
| 👋 **First-Timer Welcome** | Greets first-time contributors on their first issue or PR with helpful links. |
| 📂 **Auto-Labeler** | Automatically applies area labels based on modified paths (e.g. `docs/**` $\rightarrow$ `area/docs`). |
| ⚡ **Instant Go Engine** | Runs as a compiled static binary in **<100ms** with zero dependencies. |

---

## ⚙️ Customization (Optional)

If you want to customize messages, sizing thresholds, or disable specific plugins, add **`.github/cbog.yml`**:

```yaml
version: 1

plugins:
  # Slash commands: /lgtm, /approve, /hold, /merge
  commands:
    enabled: true
    allow_self_approval: false
    default_merge_method: squash # squash, merge, or rebase

  # Contributor self-assignment: /claim and /unclaim
  claim:
    enabled: true
    max_issues_per_user: 3

  # Welcome greeting for first-time contributors
  welcome:
    enabled: true
    issue_message: "👋 Welcome @{{.User}}! Thanks for opening an issue in {{.Repo}}!"
    pr_message: "🎉 Welcome @{{.User}}! Thanks for opening your first PR in {{.Repo}}!"

  # Automatic PR size labeling
  size:
    enabled: true
    xs: 10
    s: 30
    m: 100
    l: 500
    xl: 1000

  # Conventional commit title checker
  title_lint:
    enabled: true
    types: ["feat", "fix", "docs", "style", "refactor", "perf", "test", "chore", "revert"]

  # Automatic area labels
  auto_label:
    enabled: true
    rules:
      - label: "area/docs"
        paths: ["docs/**", "**/*.md"]
      - label: "area/ci"
        paths: [".github/**"]
```

---

## 📦 How to Publish This Action (For Maintainers)

1. Create a public repository named `cbog` under your GitHub account:
   ```bash
   cd /Users/kavindu2/Desktop/cbog
   git remote add origin https://github.com/kavix/cbog.git
   git push -u origin main
   ```
2. On GitHub, navigate to **Releases** $\rightarrow$ **Draft a new release**.
3. Choose tag **`v1.0.0`** (and point **`v1`** to it).
4. Check **"Publish this Action to the GitHub Marketplace"**.
5. Pick categories (*Community*, *Automation*) and click **Publish release**.

---

## 📄 License

[MIT](LICENSE) © 2026 kavix
