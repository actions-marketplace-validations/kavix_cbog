package lexer

import (
	"regexp"
	"strings"

	"github.com/kavix/cbog/pkg/types"
)

var (
	mentionRegex = regexp.MustCompile(`@([a-zA-Z0-9_\-]+)`)
)

// ParseCommands scans comment body and extracts valid slash commands.
func ParseCommands(body string) []types.ParsedCommand {
	var commands []types.ParsedCommand
	inCodeBlock := false

	lines := strings.Split(body, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Toggle multi-line code blocks
		if strings.HasPrefix(trimmed, "```") {
			inCodeBlock = !inCodeBlock
			continue
		}
		if inCodeBlock {
			continue
		}

		// Skip quoted lines
		if strings.HasPrefix(trimmed, ">") {
			continue
		}

		// Look for commands starting with /
		if !strings.HasPrefix(trimmed, "/") {
			continue
		}

		parts := strings.Fields(trimmed)
		if len(parts) == 0 {
			continue
		}

		cmdWord := strings.ToLower(strings.TrimPrefix(parts[0], "/"))
		args := parts[1:]

		switch types.CommandType(cmdWord) {
		case types.CmdLGTM:
			isCancel := len(args) > 0 && strings.ToLower(args[0]) == "cancel"
			commands = append(commands, types.ParsedCommand{
				Type:     types.CmdLGTM,
				Args:     args,
				IsCancel: isCancel,
				Raw:      trimmed,
			})

		case types.CmdApprove:
			isCancel := len(args) > 0 && strings.ToLower(args[0]) == "cancel"
			commands = append(commands, types.ParsedCommand{
				Type:     types.CmdApprove,
				Args:     args,
				IsCancel: isCancel,
				Raw:      trimmed,
			})

		case types.CmdHold:
			isCancel := len(args) > 0 && strings.ToLower(args[0]) == "cancel"
			commands = append(commands, types.ParsedCommand{
				Type:     types.CmdHold,
				Args:     args,
				IsCancel: isCancel,
				Raw:      trimmed,
			})

		case types.CmdAssign:
			assignees := ExtractMentions(args)
			commands = append(commands, types.ParsedCommand{
				Type:     types.CmdAssign,
				Args:     assignees,
				IsCancel: false,
				Raw:      trimmed,
			})

		case types.CmdUnassign:
			assignees := ExtractMentions(args)
			commands = append(commands, types.ParsedCommand{
				Type:     types.CmdUnassign,
				Args:     assignees,
				IsCancel: false,
				Raw:      trimmed,
			})

		case types.CmdClaim:
			commands = append(commands, types.ParsedCommand{
				Type:     types.CmdClaim,
				Args:     args,
				IsCancel: false,
				Raw:      trimmed,
			})

		case types.CmdUnclaim:
			commands = append(commands, types.ParsedCommand{
				Type:     types.CmdUnclaim,
				Args:     args,
				IsCancel: false,
				Raw:      trimmed,
			})

		case types.CmdMerge:
			commands = append(commands, types.ParsedCommand{
				Type:     types.CmdMerge,
				Args:     args,
				IsCancel: false,
				Raw:      trimmed,
			})

		case types.CmdClose:
			commands = append(commands, types.ParsedCommand{
				Type:     types.CmdClose,
				Args:     args,
				IsCancel: false,
				Raw:      trimmed,
			})

		case types.CmdReopen:
			commands = append(commands, types.ParsedCommand{
				Type:     types.CmdReopen,
				Args:     args,
				IsCancel: false,
				Raw:      trimmed,
			})

		case types.CmdHelp:
			commands = append(commands, types.ParsedCommand{
				Type:     types.CmdHelp,
				Args:     args,
				IsCancel: false,
				Raw:      trimmed,
			})
		}
	}

	return commands
}

// ExtractMentions extracts github usernames without the @ prefix.
func ExtractMentions(args []string) []string {
	var mentions []string
	joined := strings.Join(args, " ")
	matches := mentionRegex.FindAllStringSubmatch(joined, -1)
	for _, m := range matches {
		if len(m) > 1 {
			mentions = append(mentions, m[1])
		}
	}
	return mentions
}
