package scaffolder

import (
	"strings"
	"testing"
)

func TestScaffolderGenerateFiles(t *testing.T) {
	stacks := []TechStack{StackGo, StackNode, StackPython, StackRust, StackUniversal}

	for _, stack := range stacks {
		scaff := NewScaffolder(stack, "kavix", "mit", "my-test-bot", "https://example.com/icon.ico")
		files, err := scaff.GenerateFiles()
		if err != nil {
			t.Fatalf("failed to generate files for stack %s: %v", stack, err)
		}

		// Verify essential files exist
		expectedFiles := []string{
			"CONTRIBUTING.md",
			"CODE_OF_CONDUCT.md",
			"SECURITY.md",
			"SUPPORT.md",
			"LICENSE",
			".github/PULL_REQUEST_TEMPLATE.md",
			".github/ISSUE_TEMPLATE/bug_report.yml",
			".github/ISSUE_TEMPLATE/feature_request.yml",
			".github/ISSUE_TEMPLATE/config.yml",
			".github/cbog.yml",
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
		if !strings.Contains(contrib, "/lgtm") || !strings.Contains(contrib, "/merge") || !strings.Contains(contrib, "/claim") {
			t.Errorf("CONTRIBUTING.md is missing command cheat sheet for stack %s", stack)
		}

		// Verify custom bot identity is formatted into .github/cbog.yml
		cbogYaml := files[".github/cbog.yml"]
		if !strings.Contains(cbogYaml, "my-test-bot") || !strings.Contains(cbogYaml, "https://example.com/icon.ico") {
			t.Errorf(".github/cbog.yml missing custom bot identity: %s", cbogYaml)
		}
	}
}

func TestLicenseTypes(t *testing.T) {
	licenses := map[string]string{
		"mit":          "MIT License",
		"apache-2.0":   "Apache License",
		"bsd-3-clause": "BSD 3-Clause License",
		"gpl-3.0":      "GNU GENERAL PUBLIC LICENSE",
		"mpl-2.0":      "Mozilla Public License",
	}

	for licKey, expectedSubstring := range licenses {
		scaff := NewScaffolder(StackGo, "kavix", licKey, "cbog", "")
		files, err := scaff.GenerateFiles()
		if err != nil {
			t.Fatalf("failed to generate for license %s: %v", licKey, err)
		}

		licContent, ok := files["LICENSE"]
		if !ok || !strings.Contains(licContent, expectedSubstring) {
			t.Errorf("license %s expected to contain %q, got:\n%s", licKey, expectedSubstring, licContent)
		}
	}
}
