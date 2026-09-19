package plugins

import (
	"regexp"
	"strings"
	"testing"
)

func TestTitleLintRegex(t *testing.T) {
	typesList := []string{"feat", "fix", "docs", "style", "refactor", "perf", "test", "chore", "revert"}
	pattern := `^(feat|fix|docs|style|refactor|perf|test|chore|revert)(\([a-zA-Z0-9_\-\./]+\))?(!)?:\s+.+$`
	re := regexp.MustCompile(pattern)

	validTitles := []string{
		"feat: add awesome new feature",
		"fix(core): resolve race condition in dispatcher",
		"docs: update slash commands cheat sheet",
		"refactor(plugins/size)!: breaking change to sizing algorithm",
		"chore: bump dependencies to v2",
	}

	for _, title := range validTitles {
		if !re.MatchString(title) {
			t.Errorf("expected valid title %q to match", title)
		}
	}

	invalidTitles := []string{
		"Update README.md",
		"WIP: working on thing",
		"fixing bug with parser",
		"feat:",
		"feat : missing colon spacing",
	}

	for _, title := range invalidTitles {
		if re.MatchString(title) {
			t.Errorf("expected invalid title %q to fail", title)
		}
	}
	_ = typesList
	_ = strings.Join
}
