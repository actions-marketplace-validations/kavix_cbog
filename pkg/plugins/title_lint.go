package plugins

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/google/go-github/v60/github"
	"github.com/kavix/cbog/pkg/config"
)

// HandleTitleLint validates PR title format against Conventional Commits.
func HandleTitleLint(ctx context.Context, client *github.Client, owner, repo string, pr *github.PullRequest, cfg config.TitleLintConfig) error {
	if !cfg.Enabled || pr == nil {
		return nil
	}

	title := pr.GetTitle()
	sha := pr.GetHead().GetSHA()
	if sha == "" {
		return nil
	}

	typesList := cfg.Types
	if len(typesList) == 0 {
		typesList = []string{"feat", "fix", "docs", "style", "refactor", "perf", "test", "chore", "revert"}
	}

	pattern := fmt.Sprintf(`^(%s)(\([a-zA-Z0-9_\-\./]+\))?(!)?:\s+.+$`, strings.Join(typesList, "|"))
	re, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	contextName := "cbog/title-lint"
	isValid := re.MatchString(title)

	var state, desc string
	if isValid {
		state = "success"
		desc = "PR title follows Conventional Commits specification"
	} else {
		state = "failure"
		desc = fmt.Sprintf("Title must follow '%s: <description>' format", strings.Join(typesList[:3], "|"))
	}

	status := &github.RepoStatus{
		State:       &state,
		Context:     &contextName,
		Description: &desc,
	}

	_, _, err = client.Repositories.CreateStatus(ctx, owner, repo, sha, status)
	return err
}
