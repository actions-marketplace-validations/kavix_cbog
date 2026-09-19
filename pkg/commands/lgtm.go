package commands

import (
	"fmt"

	"github.com/kavindu/cbog/pkg/permissions"
	"github.com/kavindu/cbog/pkg/reactions"
	"github.com/kavindu/cbog/pkg/types"
)

const LabelLGTM = "lgtm"

// HandleLGTM handles /lgtm and /lgtm cancel
func HandleLGTM(bCtx *types.BotContext, cmd types.ParsedCommand) error {
	allowed, reason, err := permissions.CanApproveOrLGTM(bCtx)
	if err != nil {
		return fmt.Errorf("failed to check permissions: %w", err)
	}
	if !allowed {
		_ = reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "-1")
		return reactions.PostComment(bCtx, fmt.Sprintf("@%s: %s", bCtx.Sender, reason))
	}

	if cmd.IsCancel {
		_, err := bCtx.Client.Issues.RemoveLabelForIssue(bCtx.Ctx, bCtx.Owner, bCtx.Repo, bCtx.IssueNumber, LabelLGTM)
		if err != nil {
			// Ignored if label doesn't exist
		}
	} else {
		_, _, err := bCtx.Client.Issues.AddLabelsToIssue(bCtx.Ctx, bCtx.Owner, bCtx.Repo, bCtx.IssueNumber, []string{LabelLGTM})
		if err != nil {
			_ = reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "-1")
			return fmt.Errorf("failed to add label %s: %w", LabelLGTM, err)
		}
	}

	return reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "+1")
}
