package reactions

import (
	"context"

	"github.com/google/go-github/v60/github"
	"github.com/kavindu/cbog/pkg/types"
)

// AddReaction adds an emoji reaction to the comment that triggered the action.
func AddReaction(ctx context.Context, client *github.Client, owner, repo string, commentID int64, content string) error {
	if commentID == 0 {
		return nil
	}
	_, _, err := client.Reactions.CreateIssueCommentReaction(ctx, owner, repo, commentID, content)
	return err
}

// PostComment posts an informative markdown comment to the issue or PR.
func PostComment(bCtx *types.BotContext, body string) error {
	_, _, err := bCtx.Client.Issues.CreateComment(bCtx.Ctx, bCtx.Owner, bCtx.Repo, bCtx.IssueNumber, &github.IssueComment{
		Body: github.String(body),
	})
	return err
}
