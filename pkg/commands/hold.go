package commands

import (
	"fmt"

	"github.com/kavindu/contributor-bot-action/pkg/reactions"
	"github.com/kavindu/contributor-bot-action/pkg/types"
)

const LabelHold = "do-not-merge/hold"

// HandleHold handles /hold and /hold cancel
func HandleHold(bCtx *types.BotContext, cmd types.ParsedCommand) error {
	if cmd.IsCancel {
		_, err := bCtx.Client.Issues.RemoveLabelForIssue(bCtx.Ctx, bCtx.Owner, bCtx.Repo, bCtx.IssueNumber, LabelHold)
		if err != nil {
			// Ignore if not present
		}
	} else {
		_, _, err := bCtx.Client.Issues.AddLabelsToIssue(bCtx.Ctx, bCtx.Owner, bCtx.Repo, bCtx.IssueNumber, []string{LabelHold})
		if err != nil {
			_ = reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "-1")
			return fmt.Errorf("failed to add label %s: %w", LabelHold, err)
		}
	}

	return reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "+1")
}
