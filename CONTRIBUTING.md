# Contributing to cbog

Thank you for contributing to **cbog**! We welcome bug reports, feature proposals, and pull requests from all open-source contributors.

---

## Code of Conduct & Security

- **Code of Conduct**: Please review and follow our [Code of Conduct](CODE_OF_CONDUCT.md).
- **Security Policy**: For confidential vulnerability disclosure, see [SECURITY.md](SECURITY.md).

---

## Development Workflow

1. **Fork & Clone** the repository.
2. **Create a topic branch**:
   ```bash
   git checkout -b feat/your-feature-name
   ```
3. **Run tests & verify build**:
   ```bash
   go test -v ./...
   go build -v ./cmd/bot
   ```
4. **Commit your changes**:
   Please follow [Conventional Commits](https://www.conventionalcommits.org/) (e.g. `feat(plugins): add new slash command`).
5. **Open a Pull Request**:
   Fill in the checklist and link relevant issues (`Fixes #1`).

---

## Contributor Bot Commands

You can interact with issues and pull requests using slash commands:

| Command | Description | Eligible Roles |
| :--- | :--- | :--- |
| `/claim` | Self-assign an open issue | Everyone |
| `/unclaim` | Release an assigned issue back to the community | Assigned Contributor |
| `/assign [@user...]` | Assign yourself or mentioned collaborators | Everyone |
| `/unassign [@user...]` | Remove assignees | Everyone |
| `/lgtm` | Apply `lgtm` label ("Looks Good To Me") | Maintainers (non-author) |
| `/lgtm cancel` | Remove `lgtm` label | Maintainers |
| `/approve` | Submit formal review approval & `approved` label | Maintainers (non-author) |
| `/approve cancel` | Remove `approved` label | Maintainers |
| `/hold` | Add `do-not-merge/hold` to prevent accidental merges | Everyone |
| `/hold cancel` | Remove `do-not-merge/hold` label | Everyone |
| `/merge [squash\|merge\|rebase]` | Safely merge pull request once checks pass | Maintainers |
| `/close` / `/reopen` | Close or reopen an issue or PR | Author / Maintainers |
| `/help` | Print command reference table | Everyone |
