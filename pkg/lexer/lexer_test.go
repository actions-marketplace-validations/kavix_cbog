package lexer

import (
	"testing"

	"github.com/kavindu/contributor-bot-action/pkg/types"
)

func TestParseCommands(t *testing.T) {
	input := `
Hello, thank you for the PR!
> /merge
The above quote should be ignored.

` + "```" + `
/lgtm
` + "```" + `
Code block above should also be ignored.

/lgtm
/approve
/hold cancel
/assign @alice @bob
/merge squash
/help
`

	cmds := ParseCommands(input)
	if len(cmds) != 6 {
		t.Fatalf("expected 6 commands, got %d", len(cmds))
	}

	if cmds[0].Type != types.CmdLGTM || cmds[0].IsCancel {
		t.Errorf("unexpected cmd 0: %+v", cmds[0])
	}
	if cmds[1].Type != types.CmdApprove || cmds[1].IsCancel {
		t.Errorf("unexpected cmd 1: %+v", cmds[1])
	}
	if cmds[2].Type != types.CmdHold || !cmds[2].IsCancel {
		t.Errorf("unexpected cmd 2: %+v", cmds[2])
	}
	if cmds[3].Type != types.CmdAssign || len(cmds[3].Args) != 2 || cmds[3].Args[0] != "alice" || cmds[3].Args[1] != "bob" {
		t.Errorf("unexpected cmd 3: %+v", cmds[3])
	}
	if cmds[4].Type != types.CmdMerge || len(cmds[4].Args) != 1 || cmds[4].Args[0] != "squash" {
		t.Errorf("unexpected cmd 4: %+v", cmds[4])
	}
	if cmds[5].Type != types.CmdHelp {
		t.Errorf("unexpected cmd 5: %+v", cmds[5])
	}
}
