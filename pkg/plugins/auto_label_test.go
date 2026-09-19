package plugins

import (
	"testing"
)

func TestMatchPattern(t *testing.T) {
	tests := []struct {
		pattern  string
		filename string
		expected bool
	}{
		{"**/*.md", "README.md", true},
		{"**/*.md", "docs/architecture.md", true},
		{"**/*.md", "src/main.go", false},
		{".github/**", ".github/workflows/ci.yml", true},
		{".github/**", ".github/cbog.yml", true},
		{".github/**", "pkg/commands/assign.go", false},
		{"docs/**", "docs/index.html", true},
	}

	for _, tt := range tests {
		result := matchPattern(tt.pattern, tt.filename)
		if result != tt.expected {
			t.Errorf("matchPattern(%q, %q) = %v; want %v", tt.pattern, tt.filename, result, tt.expected)
		}
	}
}
