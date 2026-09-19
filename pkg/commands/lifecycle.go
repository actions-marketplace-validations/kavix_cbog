package commands

import (
	"fmt"

	"github.com/google/go-github/v60/github"
	"github.com/kavix/cbog/pkg/permissions"
	"github.com/kavix/cbog/pkg/reactions"
	"github.com/kavix/cbog/pkg/types"
)

// HandleClose closes an issue or PR
func HandleClose(bCtx *types.BotContext, cmd types.ParsedCommand) error {
	hasWrite, _ := permissions.HasWriteAccess(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.Sender)
	isAuthor := bCtx.Sender == bCtx.IssueAuthor

	if !hasWrite && !isAuthor {
		_ = reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "-1")
		return reactions.PostComment(bCtx, fmt.Sprintf("@%s: Only the author or maintainers can close this issue/PR.", bCtx.Sender))
	}

	state := "closed"
	_, _, err := bCtx.Client.Issues.Edit(bCtx.Ctx, bCtx.Owner, bCtx.Repo, bCtx.IssueNumber, &github.IssueRequest{
		State: &state,
	})
	if err != nil {
		_ = reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "-1")
		return err
	}

	return reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "+1")
}

// HandleReopen reopens an issue or PR
func HandleReopen(bCtx *types.BotContext, cmd types.ParsedCommand) error {
	hasWrite, _ := permissions.HasWriteAccess(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.Sender)
	isAuthor := bCtx.Sender == bCtx.IssueAuthor

	if !hasWrite && !isAuthor {
		_ = reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "-1")
		return reactions.PostComment(bCtx, fmt.Sprintf("@%s: Only the author or maintainers can reopen this issue/PR.", bCtx.Sender))
	}

	state := "open"
	_, _, err := bCtx.Client.Issues.Edit(bCtx.Ctx, bCtx.Owner, bCtx.Repo, bCtx.IssueNumber, &github.IssueRequest{
		State: &state,
	})
	if err != nil {
		_ = reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "-1")
		return err
	}

	return reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "+1")
}
