package types

import (
	"context"

	"github.com/google/go-github/v60/github"
)

// CommandType defines the recognized slash command.
type CommandType string

const (
	CmdAssign   CommandType = "assign"
	CmdUnassign CommandType = "unassign"
	CmdLGTM     CommandType = "lgtm"
	CmdApprove  CommandType = "approve"
	CmdHold     CommandType = "hold"
	CmdMerge    CommandType = "merge"
	CmdClose    CommandType = "close"
	CmdReopen   CommandType = "reopen"
	CmdHelp     CommandType = "help"
)

// ParsedCommand represents an extracted command and its parameters.
type ParsedCommand struct {
	Type     CommandType
	Args     []string
	IsCancel bool
	Raw      string
}

// BotContext encapsulates all runtime metadata and clients.
type BotContext struct {
	Ctx         context.Context
	Client      *github.Client
	Owner       string
	Repo        string
	IssueNumber int
	CommentID   int64
	Sender      string
	IssueAuthor string
	IsPR        bool
	CommentBody string
	DefaultMergeMethod string
}
