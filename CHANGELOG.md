# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-09-19

### Added
- **Slash Commands Engine**: Support for `/lgtm`, `/approve`, `/hold`, `/merge [squash|merge|rebase]`, `/assign`, `/unassign`, `/close`, `/reopen`, and `/help`.
- **Contributor Self-Assignment**: Contributors can self-assign issues with `/claim` and release them with `/unclaim`.
- **PR Size Labeler**: Automatic labeling with `size/XS`, `size/S`, `size/M`, `size/L`, and `size/XL` based on line diffs.
- **Conventional Commits Linter**: Validation of PR titles against Conventional Commits with status check reporting (`cbog/title-lint`).
- **First-Time Contributor Welcome**: Automated friendly onboarding messages for new contributors.
- **Path-Based Auto-Labeling**: Area labels automatically applied based on modified file paths.
- **Universal OSS Scaffolder**: 1-click generation of `CONTRIBUTING.md`, `CODE_OF_CONDUCT.md`, `SECURITY.md`, `SUPPORT.md`, `LICENSE`, PR checklists, and modern Issue Forms.
- **Pluggable Architecture**: Customizable via `.github/cbog.yml`.
- **90-Day Stale Triage**: Scheduled workflow to manage aging issues and pull requests.
- **Zero-Dependency Static Go Binary**: Blazing fast execution in <100ms.
