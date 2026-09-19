package permissions

import (
	"context"

	"github.com/google/go-github/v60/github"
	"github.com/kavix/cbog/pkg/types"
)

// HasWriteAccess checks if a given user has write or admin permissions in the repository.
func HasWriteAccess(ctx context.Context, client *github.Client, owner, repo, username string) (bool, error) {
	perm, _, err := client.Repositories.GetPermissionLevel(ctx, owner, repo, username)
	if err != nil {
		return false, err
	}

	level := perm.GetPermission()
	return level == "admin" || level == "write" || level == "maintain", nil
}

// CanApproveOrLGTM checks if the user is authorized to /lgtm or /approve.
// By default, PR authors cannot approve or LGTM their own PRs.
func CanApproveOrLGTM(bCtx *types.BotContext) (bool, string, error) {
	if bCtx.Sender == bCtx.IssueAuthor {
		return false, "PR authors cannot approve or /lgtm their own pull requests.", nil
	}

	hasWrite, err := HasWriteAccess(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.Sender)
	if err != nil {
		return false, "", err
	}
	if !hasWrite {
		return false, "Only repository maintainers with write access can use this command.", nil
	}

	return true, "", nil
}

// CanMerge checks if the user is authorized to /merge.
func CanMerge(bCtx *types.BotContext) (bool, string, error) {
	hasWrite, err := HasWriteAccess(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.Sender)
	if err != nil {
		return false, "", err
	}
	if !hasWrite {
		return false, "Only repository maintainers with write access can merge pull requests.", nil
	}
	return true, "", nil
}
