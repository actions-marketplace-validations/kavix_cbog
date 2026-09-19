package plugins

import (
	"context"
	"strings"

	"github.com/google/go-github/v60/github"
	"github.com/kavix/cbog/pkg/config"
)

var allSizeLabels = []string{"size/XS", "size/S", "size/M", "size/L", "size/XL"}

// HandlePRSize calculates lines changed in PR and applies corresponding size/ label.
func HandlePRSize(ctx context.Context, client *github.Client, owner, repo string, pr *github.PullRequest, cfg config.SizeConfig) error {
	if !cfg.Enabled || pr == nil {
		return nil
	}

	additions := pr.GetAdditions()
	deletions := pr.GetDeletions()
	total := additions + deletions

	var targetLabel string
	switch {
	case total < cfg.XS:
		targetLabel = "size/XS"
	case total < cfg.S:
		targetLabel = "size/S"
	case total < cfg.M:
		targetLabel = "size/M"
	case total < cfg.L:
		targetLabel = "size/L"
	default:
		targetLabel = "size/XL"
	}

	prNumber := pr.GetNumber()

	// Fetch current labels to remove previous size labels
	issue, _, err := client.Issues.Get(ctx, owner, repo, prNumber)
	if err == nil && issue != nil {
		for _, lbl := range issue.Labels {
			name := lbl.GetName()
			if strings.HasPrefix(name, "size/") && name != targetLabel {
				_, _ = client.Issues.RemoveLabelForIssue(ctx, owner, repo, prNumber, name)
			}
		}
	}

	_, _, err = client.Issues.AddLabelsToIssue(ctx, owner, repo, prNumber, []string{targetLabel})
	return err
}
