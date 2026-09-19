package plugins

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/google/go-github/v60/github"
	"github.com/kavindu/cbog/pkg/config"
)

// HandleAutoLabel inspects files changed in PR and applies area/component labels based on configured rules.
func HandleAutoLabel(ctx context.Context, client *github.Client, owner, repo string, pr *github.PullRequest, cfg config.AutoLabelConfig) error {
	if !cfg.Enabled || len(cfg.Rules) == 0 || pr == nil {
		return nil
	}

	prNum := pr.GetNumber()
	files, _, err := client.PullRequests.ListFiles(ctx, owner, repo, prNum, &github.ListOptions{PerPage: 100})
	if err != nil {
		return err
	}

	labelsToAdd := make(map[string]bool)
	for _, f := range files {
		filename := f.GetFilename()
		for _, rule := range cfg.Rules {
			for _, p := range rule.Paths {
				matched := matchPattern(p, filename)
				if matched {
					labelsToAdd[rule.Label] = true
				}
			}
		}
	}

	if len(labelsToAdd) == 0 {
		return nil
	}

	var labels []string
	for lbl := range labelsToAdd {
		labels = append(labels, lbl)
	}

	_, _, err = client.Issues.AddLabelsToIssue(ctx, owner, repo, prNum, labels)
	return err
}

func matchPattern(pattern, filename string) bool {
	if strings.HasPrefix(pattern, "**/*") {
		ext := strings.TrimPrefix(pattern, "**/*")
		return strings.HasSuffix(filename, ext)
	}
	if strings.HasSuffix(pattern, "/**") {
		dir := strings.TrimSuffix(pattern, "/**")
		return strings.HasPrefix(filename, dir+"/")
	}
	matched, _ := filepath.Match(pattern, filename)
	return matched
}
