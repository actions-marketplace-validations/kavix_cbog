# Contributing Guidelines

Thank you for your interest in contributing to this project. This document provides guidelines for contributing code, documentation, and reporting issues.

---

## 1. Code of Conduct & Security

- **Code of Conduct**: By participating in this project, you agree to abide by our [Code of Conduct](CODE_OF_CONDUCT.md).
- **Security Disclosures**: Please report security vulnerabilities privately according to our [Security Policy](SECURITY.md).

---

## 2. Development Workflow

1. **Fork & Clone** the repository.
2. **Create a topic branch**:
   ```bash
   git checkout -b feat/your-feature-name
   ```
3. **Write code & run tests**:
   {{ .TestInstructions }}
4. **Commit your changes**:
   Please follow the [Conventional Commits](https://www.conventionalcommits.org/) specification (e.g., `feat: add new capability`, `fix: handle edge case in parser`).
5. **Open a Pull Request**:
   Fill in the pull request template checklist and link relevant issues (`Fixes #123`).

---

## 3. Contributor Bot Commands

This repository uses automated governance tooling. You can interact with issues and pull requests using slash commands in comments:

| Command | Description | Eligible Roles |
| :--- | :--- | :--- |
| `/claim` | Self-assign an open issue to start working on it | Everyone |
| `/unclaim` | Release an assigned issue back to the community | Assigned Contributor |
| `/assign [@user...]` | Assign yourself or mentioned collaborators | Everyone |
| `/unassign [@user...]` | Remove assignees | Everyone |
| `/lgtm` | Apply `lgtm` label ("Looks Good To Me") | Maintainers (non-author) |
| `/lgtm cancel` | Remove `lgtm` label | Maintainers |
| `/approve` | Add `approved` label and submit formal review approval | Maintainers (non-author) |
| `/approve cancel` | Remove `approved` label | Maintainers |
| `/hold` | Add `do-not-merge/hold` to prevent merging | Everyone |
| `/hold cancel` | Remove `do-not-merge/hold` label | Everyone |
| `/merge [squash\|merge\|rebase]` | Safely merge pull request once checks pass | Maintainers |
| `/close` / `/reopen` | Close or reopen an issue or pull request | Author / Maintainers |
| `/help` | Print command reference | Everyone |
