# cbog: Community Bot for Open-source Governance

[![Go Report Card](https://goreportcard.com/badge/github.com/kavix/cbog)](https://goreportcard.com/report/github.com/kavix/cbog)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Marketplace](https://img.shields.io/badge/Marketplace-cbog-blue.svg)](https://github.com/marketplace/actions/cbog)

> Supercharge any open-source repository in 60 seconds with automated triage, contributor self-assignment (`/claim`), `/lgtm`, `/merge`, and a 1-click governance scaffolder.

---

## Quickstart (1 Step, Zero Config)

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

`cbog` will immediately begin managing your issues and pull requests.

---

## 1-Click Governance Scaffolder

Generate a complete, enterprise-grade open-source governance suite in 30 seconds:

1. In your repository, navigate to the **Actions** tab.
2. Select **cbog** from the sidebar.
3. Click **Run workflow** -> select **`scaffold-oss-docs`** -> click **Run workflow**.

`cbog` will:
- **Auto-enable GitHub Discussions** on your repository.
- **Auto-enable automatic branch deletion** on pull request merge.
- Auto-detect your programming language (Go, Python, Rust, Node.js, Java) and open a Pull Request with all 10 standard community files:
  - `CONTRIBUTING.md` (Language-tailored build instructions and command reference)
  - `CODE_OF_CONDUCT.md` (Contributor Covenant v2.1)
  - `SECURITY.md` (Confidential vulnerability reporting policy)
  - `SUPPORT.md` (Community assistance guidelines)
  - `LICENSE` (Supports `mit`, `apache-2.0`, `bsd-3-clause`, `gpl-3.0`, `mpl-2.0`)
  - `.github/PULL_REQUEST_TEMPLATE.md` (Structured PR checklist)
  - `.github/ISSUE_TEMPLATE/bug_report.yml` (Structured Bug Report form)
  - `.github/ISSUE_TEMPLATE/feature_request.yml` (Structured Feature Request form)
  - `.github/ISSUE_TEMPLATE/config.yml` (Blank issue redirection to Discussions)
  - `.github/cbog.yml` (Ready-to-tweak plugin configuration)

---

## Contributor & Maintainer Commands

Comment on any Issue or Pull Request to trigger commands:

| Command | What it does | Who can use it |
| :--- | :--- | :--- |
| `/claim` | Self-assign an open issue to start working on it | Everyone |
| `/unclaim` | Release an assigned issue back to the community | Assigned Contributor |
| `/assign [@user]` | Assign yourself or mentioned collaborators | Everyone |
| `/unassign [@user]` | Remove assignees | Everyone |
| `/lgtm` | Mark PR with `lgtm` ("Looks Good To Me") | Maintainers (non-author) |
| `/lgtm cancel` | Remove `lgtm` label | Maintainers |
| `/approve` | Submit formal GitHub PR Review approval + `approved` label | Maintainers (non-author) |
| `/approve cancel` | Remove approval label | Maintainers |
| `/hold` | Add `do-not-merge/hold` to prevent accidental merges | Everyone |
| `/hold cancel` | Remove hold label | Everyone |
| `/merge` | Safely merge the PR once all status checks pass | Maintainers |
| `/merge squash` | Merge using squash (or `merge`, `rebase`) | Maintainers |
| `/close` / `/reopen` | Close or reopen an issue or PR | Author / Maintainers |
| `/help` | Print command reference cheat sheet | Everyone |

---

## Included Plugins & Automations

| Feature | Description |
| :--- | :--- |
| **Auto PR Sizing** | Automatically labels PRs with `size/XS` (<10 lines), `size/S`, `size/M`, `size/L`, or `size/XL` (>500 lines). |
| **PR Title Linter** | Validates PR titles against Conventional Commits (`feat:`, `fix:`, `docs:`) and sets GitHub status checks. |
| **First-Timer Welcome** | Greets first-time contributors on their first issue or PR with helpful links. |
| **Path-Based Auto-Labeler** | Automatically categorizes PRs by file paths (e.g. `docs/**` -> `documentation`). |
| **90-Day Stale Triage** | Automatically identifies and marks inactive issues/PRs with `lifecycle/stale`. |
| **Fast Go Engine** | Runs as a static compiled binary in under 100ms with zero runtime dependencies. |

---

## Configuration (`.github/cbog.yml`)

You can customize or disable any plugin by adding `.github/cbog.yml`:

```yaml
version: 1

# Custom bot display name and icon URL (.ico or image)
bot:
  name: "cbog"
  icon_url: "" # e.g. "https://example.com/favicon.ico"

plugins:
  commands:
    enabled: true
    allow_self_approval: false
    default_merge_method: squash # squash, merge, or rebase

  claim:
    enabled: true
    max_issues_per_user: 3

  welcome:
    enabled: true
    issue_message: "Welcome @{{.User}}. Thank you for opening an issue in {{.Repo}}!"
    pr_message: "Welcome @{{.User}}. Thank you for submitting your first PR in {{.Repo}}!"

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
      - label: "documentation"
        paths: ["docs/**", "**/*.md"]
      - label: "enhancement"
        paths: ["pkg/plugins/**", "pkg/commands/**"]
```

---

## Publishing to GitHub Marketplace

1. Ensure your repository is **public**.
2. Navigate to **Releases** -> **Draft a new release**.
3. Create release tag **`v1.0.0`** (and point **`v1`** to it).
4. Check **"Publish this Action to the GitHub Marketplace"**.
5. Select categories (**Community**, **Automation**) and click **Publish release**.

---

## Container Image (GitHub Packages)

`cbog` is also published as a container image on GitHub Packages:

```bash
docker pull ghcr.io/kavix/cbog:latest
```

You can run it directly:

```bash
docker run --rm \
  -e GITHUB_TOKEN="${GITHUB_TOKEN}" \
  -e GITHUB_REPOSITORY="owner/repo" \
  -e GITHUB_EVENT_NAME="issue_comment" \
  ghcr.io/kavix/cbog:latest
```

---

## Local Development & Testing

```bash
# Run test suite
go test -v ./...

# Compile binary
go build -o bin/cbog ./cmd/bot
```

---

## License

[MIT](LICENSE) (c) 2026 kavix
