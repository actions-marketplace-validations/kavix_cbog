package commands

import (
	"log"

	"github.com/kavix/cbog/pkg/types"
)

// Dispatch executes each extracted command.
func Dispatch(bCtx *types.BotContext, cmds []types.ParsedCommand) error {
	for _, cmd := range cmds {
		log.Printf("Executing command: %s (args: %v, cancel: %t)", cmd.Type, cmd.Args, cmd.IsCancel)
		var err error
		switch cmd.Type {
		case types.CmdAssign:
			err = HandleAssign(bCtx, cmd)
		case types.CmdUnassign:
			err = HandleUnassign(bCtx, cmd)
		case types.CmdClaim:
			err = HandleClaim(bCtx, cmd)
		case types.CmdUnclaim:
			err = HandleUnclaim(bCtx, cmd)
		case types.CmdLGTM:
			err = HandleLGTM(bCtx, cmd)
		case types.CmdApprove:
			err = HandleApprove(bCtx, cmd)
		case types.CmdHold:
			err = HandleHold(bCtx, cmd)
		case types.CmdMerge:
			err = HandleMerge(bCtx, cmd)
		case types.CmdClose:
			err = HandleClose(bCtx, cmd)
		case types.CmdReopen:
			err = HandleReopen(bCtx, cmd)
		case types.CmdHelp:
			err = HandleHelp(bCtx)
		default:
			log.Printf("Unknown command: %s", cmd.Type)
		}

		if err != nil {
			log.Printf("Error executing command %s: %v", cmd.Type, err)
		}
	}
	return nil
}
