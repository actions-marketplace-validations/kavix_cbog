package commands

import (
	"fmt"

	"github.com/google/go-github/v60/github"
	"github.com/kavix/cbog/pkg/permissions"
	"github.com/kavix/cbog/pkg/reactions"
	"github.com/kavix/cbog/pkg/types"
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

		if bCtx.IsPR {
			pr, _, err := bCtx.Client.PullRequests.Get(bCtx.Ctx, bCtx.Owner, bCtx.Repo, bCtx.IssueNumber)
			if err == nil && pr != nil && pr.GetHead().GetSHA() != "" {
				contextName := "cbog/approved"
				state := "pending"
				desc := "Approval removed"
				_, _, _ = bCtx.Client.Repositories.CreateStatus(bCtx.Ctx, bCtx.Owner, bCtx.Repo, pr.GetHead().GetSHA(), &github.RepoStatus{
					State:       &state,
					Context:     &contextName,
					Description: &desc,
				})
			}
		}
	} else {
		_, _, err := bCtx.Client.Issues.AddLabelsToIssue(bCtx.Ctx, bCtx.Owner, bCtx.Repo, bCtx.IssueNumber, []string{LabelApproved})
		if err != nil {
			_ = reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "-1")
			return fmt.Errorf("failed to add label %s: %w", LabelApproved, err)
		}

		// If this is a PR, submit formal review approval AND set status check
		if bCtx.IsPR {
			event := "APPROVE"
			body := fmt.Sprintf("Approved on behalf of @%s via /approve", bCtx.Sender)
			reviewReq := &github.PullRequestReviewRequest{
				Event: &event,
				Body:  &body,
			}
			_, _, _ = bCtx.Client.PullRequests.CreateReview(bCtx.Ctx, bCtx.Owner, bCtx.Repo, bCtx.IssueNumber, reviewReq)

			// Set green commit status check: cbog/approved (Approved by @user)
			pr, _, err := bCtx.Client.PullRequests.Get(bCtx.Ctx, bCtx.Owner, bCtx.Repo, bCtx.IssueNumber)
			if err == nil && pr != nil && pr.GetHead().GetSHA() != "" {
				contextName := "cbog/approved"
				state := "success"
				desc := fmt.Sprintf("Approved by @%s", bCtx.Sender)
				_, _, _ = bCtx.Client.Repositories.CreateStatus(bCtx.Ctx, bCtx.Owner, bCtx.Repo, pr.GetHead().GetSHA(), &github.RepoStatus{
					State:       &state,
					Context:     &contextName,
					Description: &desc,
				})
			}
		}
	}

	return reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "+1")
}
