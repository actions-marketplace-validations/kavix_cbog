package scaffolder

import (
	"strings"
	"testing"
)

func TestScaffolderGenerateFiles(t *testing.T) {
	stacks := []TechStack{StackGo, StackNode, StackPython, StackRust, StackUniversal}

	for _, stack := range stacks {
		scaff := NewScaffolder(stack)
		files, err := scaff.GenerateFiles()
		if err != nil {
			t.Fatalf("failed to generate files for stack %s: %v", stack, err)
		}

		// Verify essential files exist
		expectedFiles := []string{
			"CONTRIBUTING.md",
			".github/PULL_REQUEST_TEMPLATE.md",
			".github/ISSUE_TEMPLATE/bug_report.yml",
			".github/ISSUE_TEMPLATE/feature_request.yml",
			".github/ISSUE_TEMPLATE/config.yml",
		}

		for _, ef := range expectedFiles {
			content, exists := files[ef]
			if !exists {
				t.Errorf("missing expected file %s for stack %s", ef, stack)
			}
			if len(content) == 0 {
				t.Errorf("file %s is empty for stack %s", ef, stack)
			}
		}

		// Verify bot slash command cheat sheet is inside CONTRIBUTING.md
		contrib := files["CONTRIBUTING.md"]
		if !strings.Contains(contrib, "/lgtm") || !strings.Contains(contrib, "/merge") || !strings.Contains(contrib, "/assign") {
			t.Errorf("CONTRIBUTING.md is missing command cheat sheet for stack %s", stack)
		}
	}
}
