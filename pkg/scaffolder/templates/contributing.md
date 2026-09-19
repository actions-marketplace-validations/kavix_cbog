# Contributing Guidelines

Welcome to the project! We love contributions from the open-source community.

---

## 1. Development Workflow

1. **Fork & Clone** the repository.
2. **Create a branch**:
   ```bash
   git checkout -b feat/your-feature-name
   ```
3. **Write clean code & tests**:
   {{ .TestInstructions }}
4. **Commit your changes**:
   Please follow [Conventional Commits](https://www.conventionalcommits.org/) (e.g. `feat: add awesome feature`, `fix: resolve crash on startup`).
5. **Open a Pull Request**:
   Fill in the PR template checklist and link relevant issues.

---

## 2. Bot Automation & Slash Commands

This repository uses an automated contributor bot. Anyone can interact with issues and PRs using slash commands:

| Command | Description | Eligible Roles |
| :--- | :--- | :--- |
| `/assign [@user...]` | Assign yourself or others to an issue/PR | Everyone |
| `/unassign [@user...]` | Remove assignees | Everyone |
| `/lgtm` | Apply `lgtm` label indicating "Looks Good To Me" | Maintainers |
| `/lgtm cancel` | Remove `lgtm` label | Maintainers |
| `/approve` | Add `approved` label and submit formal review approval | Maintainers |
| `/approve cancel` | Remove `approved` label | Maintainers |
| `/hold` | Add `do-not-merge/hold` to prevent premature merges | Everyone |
| `/hold cancel` | Remove `do-not-merge/hold` label | Everyone |
| `/merge [squash\|merge\|rebase]` | Merge the pull request (verifies CI checks) | Maintainers |
| `/close` / `/reopen` | Close or reopen an issue or PR | Author / Maintainers |
| `/help` | Display command cheat sheet | Everyone |

---

## 3. Code of Conduct

Please be respectful, kind, and collaborative. We are committed to providing a friendly and welcoming environment for everyone.
