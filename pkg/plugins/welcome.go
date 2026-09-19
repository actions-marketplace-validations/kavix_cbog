package plugins

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v60/github"
	"github.com/kavindu/cbog/pkg/config"
)

// HandleWelcome checks if an issue or PR author is a first-time contributor and posts a greeting.
func HandleWelcome(ctx context.Context, client *github.Client, owner, repo string, issue *github.Issue, isPR bool, cfg config.WelcomeConfig) error {
	if !cfg.Enabled || issue == nil {
		return nil
	}

	assoc := issue.GetAuthorAssociation()
	// Only greet first-timers or users with no previous associations
	if assoc != "FIRST_TIME_CONTRIBUTOR" && assoc != "FIRST_TIMER" {
		return nil
	}

	user := issue.GetUser().GetLogin()
	var messageTmpl string
	if isPR {
		messageTmpl = cfg.PRMessage
	} else {
		messageTmpl = cfg.IssueMessage
	}

	msg := strings.ReplaceAll(messageTmpl, "{{.User}}", user)
	msg = strings.ReplaceAll(msg, "{{.Repo}}", repo)

	_, _, err := client.Issues.CreateComment(ctx, owner, repo, issue.GetNumber(), &github.IssueComment{
		Body: github.String(fmt.Sprintf("%s\n\n*Powered by [cbog](https://github.com/kavindu/cbog)*", msg)),
	})
	return err
}
