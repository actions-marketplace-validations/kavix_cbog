package commands

import (
	"fmt"

	"github.com/kavix/cbog/pkg/reactions"
	"github.com/kavix/cbog/pkg/types"
)

// HandleAssign handles /assign [@user...]
func HandleAssign(bCtx *types.BotContext, cmd types.ParsedCommand) error {
	assignees := cmd.Args
	if len(assignees) == 0 {
		assignees = []string{bCtx.Sender}
	}

	_, _, err := bCtx.Client.Issues.AddAssignees(bCtx.Ctx, bCtx.Owner, bCtx.Repo, bCtx.IssueNumber, assignees)
	if err != nil {
		_ = reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "-1")
		return fmt.Errorf("failed to add assignees: %w", err)
	}

	return reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "+1")
}

// HandleUnassign handles /unassign [@user...]
func HandleUnassign(bCtx *types.BotContext, cmd types.ParsedCommand) error {
	assignees := cmd.Args
	if len(assignees) == 0 {
		assignees = []string{bCtx.Sender}
	}

	_, _, err := bCtx.Client.Issues.RemoveAssignees(bCtx.Ctx, bCtx.Owner, bCtx.Repo, bCtx.IssueNumber, assignees)
	if err != nil {
		_ = reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "-1")
		return fmt.Errorf("failed to remove assignees: %w", err)
	}

	return reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "+1")
}
