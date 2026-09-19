package commands

import (
	"fmt"

	"github.com/kavix/cbog/pkg/reactions"
	"github.com/kavix/cbog/pkg/types"
)

// HandleClaim handles /claim on issues
func HandleClaim(bCtx *types.BotContext, cmd types.ParsedCommand) error {
	if bCtx.IsPR {
		return reactions.PostComment(bCtx, fmt.Sprintf("@%s: Pull requests cannot be claimed. Use `/claim` on open issues!", bCtx.Sender))
	}

	// Check if already assigned
	issue, _, err := bCtx.Client.Issues.Get(bCtx.Ctx, bCtx.Owner, bCtx.Repo, bCtx.IssueNumber)
	if err != nil {
		return err
	}

	for _, assignee := range issue.Assignees {
		if assignee.GetLogin() == bCtx.Sender {
			_ = reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "+1")
			return reactions.PostComment(bCtx, fmt.Sprintf("@%s: You are already assigned to this issue!", bCtx.Sender))
		}
		if len(issue.Assignees) > 0 {
			_ = reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "confused")
			return reactions.PostComment(bCtx, fmt.Sprintf("@%s: This issue is already claimed by @%s. If they are no longer working on it, ask a maintainer to unassign them.", bCtx.Sender, issue.Assignees[0].GetLogin()))
		}
	}

	// Assign the user
	_, _, err = bCtx.Client.Issues.AddAssignees(bCtx.Ctx, bCtx.Owner, bCtx.Repo, bCtx.IssueNumber, []string{bCtx.Sender})
	if err != nil {
		_ = reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "-1")
		return fmt.Errorf("failed to claim issue: %w", err)
	}

	_ = reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "rocket")
	return reactions.PostComment(bCtx, fmt.Sprintf("@%s has claimed this issue. Please feel free to open a draft pull request when you are ready to share your work.", bCtx.Sender))
}

// HandleUnclaim handles /unclaim
func HandleUnclaim(bCtx *types.BotContext, cmd types.ParsedCommand) error {
	_, _, err := bCtx.Client.Issues.RemoveAssignees(bCtx.Ctx, bCtx.Owner, bCtx.Repo, bCtx.IssueNumber, []string{bCtx.Sender})
	if err != nil {
		_ = reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "-1")
		return err
	}

	_ = reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "+1")
	return reactions.PostComment(bCtx, fmt.Sprintf("@%s has unassigned themselves from this issue. It is now open for other contributors to claim with `/claim`.", bCtx.Sender))
}
