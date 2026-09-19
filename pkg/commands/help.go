package commands

import (
	"github.com/kavix/cbog/pkg/reactions"
	"github.com/kavix/cbog/pkg/types"
)

const HelpText = `### Contributor Bot Commands

| Command | Description | Eligible Roles |
| :--- | :--- | :--- |
| ` + "`/claim`" + ` | Self-assign an open issue | Everyone |
| ` + "`/unclaim`" + ` | Release an assigned issue back to the community | Assigned Contributor |
| ` + "`/assign [@user...]`" + ` | Assign self or mentioned users | Everyone |
| ` + "`/unassign [@user...]`" + ` | Unassign self or mentioned users | Everyone |
| ` + "`/lgtm [cancel]`" + ` | Add or remove ` + "`lgtm`" + ` label | Maintainers (non-author) |
| ` + "`/approve [cancel]`" + ` | Add or remove ` + "`approved`" + ` label & submit formal review | Maintainers (non-author) |
| ` + "`/hold [cancel]`" + ` | Set or remove ` + "`do-not-merge/hold`" + ` label | Everyone |
| ` + "`/merge [squash\\|merge\\|rebase]`" + ` | Merge the pull request | Maintainers |
| ` + "`/close`" + ` | Close issue or pull request | Author / Maintainers |
| ` + "`/reopen`" + ` | Reopen issue or pull request | Author / Maintainers |
| ` + "`/help`" + ` | Show this help menu | Everyone |
`

// HandleHelp posts the help cheat-sheet
func HandleHelp(bCtx *types.BotContext) error {
	_ = reactions.AddReaction(bCtx.Ctx, bCtx.Client, bCtx.Owner, bCtx.Repo, bCtx.CommentID, "+1")
	return reactions.PostComment(bCtx, HelpText)
}
