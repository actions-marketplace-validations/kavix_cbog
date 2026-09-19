package commands

import (
	"fmt"

	"github.com/google/go-github/v60/github"
	"github.com/kavindu/cbog/pkg/permissions"
	"github.com/kavindu/cbog/pkg/reactions"
	"github.com/kavindu/cbog/pkg/types"
)

const LabelApproved = "approved"

// HandleApprove handles /approve and /approve cancel
func HandleApprove(bCtx *types.BotContext, cmd types.ParsedCommand) error {
	allowed, reason, err := permissions.CanApproveOrLGTM(bCtx)
	if err != nil {
		return fmt.Errorf("failed to check permissions: %w", err)
	}
	if !allowed {
		_ = reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "-1")
		return reactions.PostComment(bCtx, fmt.Sprintf("@%s: %s", bCtx.Sender, reason))
	}

	if cmd.IsCancel {
		_, err := bCtx.Client.Issues.RemoveLabelForIssue(bCtx.Ctx, bCtx.Owner, bCtx.Repo, bCtx.IssueNumber, LabelApproved)
		if err != nil {
			// Ignore if not present
		}
	} else {
		_, _, err := bCtx.Client.Issues.AddLabelsToIssue(bCtx.Ctx, bCtx.Owner, bCtx.Repo, bCtx.IssueNumber, []string{LabelApproved})
		if err != nil {
			_ = reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "-1")
			return fmt.Errorf("failed to add label %s: %w", LabelApproved, err)
		}

		// If this is a PR, also submit a formal review approval
		if bCtx.IsPR {
			event := "APPROVE"
			body := fmt.Sprintf("Approved via slash command by @%s", bCtx.Sender)
			reviewReq := &github.PullRequestReviewRequest{
				Event: &event,
				Body:  &body,
			}
			_, _, _ = bCtx.Client.PullRequests.CreateReview(bCtx.Ctx, bCtx.Owner, bCtx.Repo, bCtx.IssueNumber, reviewReq)
		}
	}

	return reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "+1")
}
