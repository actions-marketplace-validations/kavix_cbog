package config

import (
	"context"
	"log"

	"github.com/google/go-github/v60/github"
	"gopkg.in/yaml.v3"
)

// BotIdentityConfig allows users to customize the bot display name and icon URL (.ico or image).
type BotIdentityConfig struct {
	Name    string `yaml:"name"`
	IconURL string `yaml:"icon_url"`
}

// Config represents the complete .github/cbog.yml configuration file.
type Config struct {
	Version int               `yaml:"version"`
	Bot     BotIdentityConfig `yaml:"bot"`
	Plugins PluginsConfig     `yaml:"plugins"`
}

type PluginsConfig struct {
	Commands  CommandsConfig  `yaml:"commands"`
	Claim     ClaimConfig     `yaml:"claim"`
	Welcome   WelcomeConfig   `yaml:"welcome"`
	Size      SizeConfig      `yaml:"size"`
	TitleLint TitleLintConfig `yaml:"title_lint"`
	AutoLabel AutoLabelConfig `yaml:"auto_label"`
}

type CommandsConfig struct {
	Enabled            bool   `yaml:"enabled"`
	AllowSelfApproval  bool   `yaml:"allow_self_approval"`
	DefaultMergeMethod string `yaml:"default_merge_method"`
}

type ClaimConfig struct {
	Enabled           bool     `yaml:"enabled"`
	RequireLabels     []string `yaml:"require_labels"`     // e.g. ["good first issue", "help wanted"]
	MaxIssuesPerUser  int      `yaml:"max_issues_per_user"` // 0 = unlimited
}

type WelcomeConfig struct {
	Enabled      bool   `yaml:"enabled"`
	IssueMessage string `yaml:"issue_message"`
	PRMessage    string `yaml:"pr_message"`
}

type SizeConfig struct {
	Enabled bool `yaml:"enabled"`
	XS      int  `yaml:"xs"`
	S       int  `yaml:"s"`
	M       int  `yaml:"m"`
	L       int  `yaml:"l"`
	XL      int  `yaml:"xl"`
}

type TitleLintConfig struct {
	Enabled bool     `yaml:"enabled"`
	Types   []string `yaml:"types"`
}

type AutoLabelConfig struct {
	Enabled bool             `yaml:"enabled"`
	Rules   []AutoLabelRule  `yaml:"rules"`
}

type AutoLabelRule struct {
	Label string   `yaml:"label"`
	Paths []string `yaml:"paths"`
}

// DefaultConfig returns ready-to-use default settings.
func DefaultConfig() *Config {
	return &Config{
		Version: 1,
		Bot: BotIdentityConfig{
			Name:    "cbog",
			IconURL: "",
		},
		Plugins: PluginsConfig{
			Commands: CommandsConfig{
				Enabled:            true,
				AllowSelfApproval:  false,
				DefaultMergeMethod: "squash",
			},
			Claim: ClaimConfig{
				Enabled:          true,
				RequireLabels:    []string{},
				MaxIssuesPerUser: 3,
			},
			Welcome: WelcomeConfig{
				Enabled:      true,
				IssueMessage: "Welcome @{{.User}}. Thank you for opening an issue in {{.Repo}}. A maintainer will review it shortly.",
				PRMessage:    "Welcome @{{.User}}. Thank you for submitting your first pull request to {{.Repo}}. Please ensure your changes pass all tests and follow our contributing guidelines.",
			},
			Size: SizeConfig{
				Enabled: true,
				XS:      10,
				S:       30,
				M:       100,
				L:       500,
				XL:      1000,
			},
			TitleLint: TitleLintConfig{
				Enabled: true,
				Types:   []string{"feat", "fix", "docs", "style", "refactor", "perf", "test", "chore", "revert"},
			},
			AutoLabel: AutoLabelConfig{
				Enabled: true,
				Rules: []AutoLabelRule{
					{Label: "area/docs", Paths: []string{"docs/**", "**/*.md"}},
					{Label: "area/ci", Paths: []string{".github/**"}},
				},
			},
		},
	}
}

// LoadFromRepo fetches .github/cbog.yml from the target repository, or returns defaults if not found.
func LoadFromRepo(ctx context.Context, client *github.Client, owner, repo, ref string) *Config {
	cfg := DefaultConfig()

	opts := &github.RepositoryContentGetOptions{}
	if ref != "" {
		opts.Ref = ref
	}

	content, _, _, err := client.Repositories.GetContents(ctx, owner, repo, ".github/cbog.yml", opts)
	if err != nil || content == nil {
		log.Println(".github/cbog.yml not found in repository. Using default configuration.")
		return cfg
	}

	raw, err := content.GetContent()
	if err != nil {
		log.Printf("Failed to read .github/cbog.yml content: %v", err)
		return cfg
	}

	if err := yaml.Unmarshal([]byte(raw), cfg); err != nil {
		log.Printf("Failed to parse .github/cbog.yml: %v. Using defaults.", err)
		return cfg
	}

	log.Println("Successfully loaded custom .github/cbog.yml from repository.")
	return cfg
}
