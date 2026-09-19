package commands

import (
	"fmt"
	"strings"

	"github.com/google/go-github/v60/github"
	"github.com/kavix/cbog/pkg/permissions"
	"github.com/kavix/cbog/pkg/reactions"
	"github.com/kavix/cbog/pkg/types"
)

// HandleMerge handles /merge [squash|merge|rebase]
func HandleMerge(bCtx *types.BotContext, cmd types.ParsedCommand) error {
	if !bCtx.IsPR {
		return reactions.PostComment(bCtx, fmt.Sprintf("@%s: The `/merge` command can only be executed on pull requests.", bCtx.Sender))
	}

	allowed, reason, err := permissions.CanMerge(bCtx)
	if err != nil {
		return fmt.Errorf("failed to check merge permissions: %w", err)
	}
	if !allowed {
		_ = reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "-1")
		return reactions.PostComment(bCtx, fmt.Sprintf("@%s: %s", bCtx.Sender, reason))
	}

	// Check if hold label is present
	issue, _, err := bCtx.Client.Issues.Get(bCtx.Ctx, bCtx.Owner, bCtx.Repo, bCtx.IssueNumber)
	if err == nil && issue != nil {
		for _, lbl := range issue.Labels {
			if lbl.GetName() == LabelHold {
				_ = reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "-1")
				return reactions.PostComment(bCtx, fmt.Sprintf("@%s: Cannot merge PR because the `%s` label is active. Remove it with `/hold cancel` first.", bCtx.Sender, LabelHold))
			}
		}
	}

	// Add 'eyes' reaction indicating merge processing
	_ = reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "eyes")

	// Determine merge method
	method := bCtx.DefaultMergeMethod
	if len(cmd.Args) > 0 {
		argMethod := strings.ToLower(cmd.Args[0])
		if argMethod == "squash" || argMethod == "merge" || argMethod == "rebase" {
			method = argMethod
		}
	}
	if method == "" {
		method = "squash"
	}

	commitMsg := fmt.Sprintf("Auto-merge triggered by @%s via /merge", bCtx.Sender)
	res, _, err := bCtx.Client.PullRequests.Merge(bCtx.Ctx, bCtx.Owner, bCtx.Repo, bCtx.IssueNumber, commitMsg, &github.PullRequestOptions{
		MergeMethod: method,
	})

	if err != nil {
		_ = reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "-1")
		return reactions.PostComment(bCtx, fmt.Sprintf("⚠️ **Merge failed for @%s**: `%v`.\nPlease ensure required status checks have passed, approvals are satisfied, and there are no merge conflicts.", bCtx.Sender, err))
	}

	if res.GetMerged() {
		return reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "rocket")
	}

	return reactions.PostComment(bCtx, fmt.Sprintf("⚠️ Merge could not be completed: %s", res.GetMessage()))
}
