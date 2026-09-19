package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/v60/github"
	"github.com/kavix/cbog/pkg/commands"
	"github.com/kavix/cbog/pkg/config"
	"github.com/kavix/cbog/pkg/lexer"
	"github.com/kavix/cbog/pkg/plugins"
	"github.com/kavix/cbog/pkg/scaffolder"
	"github.com/kavix/cbog/pkg/types"
	"golang.org/x/oauth2"
)

func main() {
	token := getEnv("INPUT_GITHUB_TOKEN", os.Getenv("GITHUB_TOKEN"))
	if token == "" {
		log.Fatal("Missing GITHUB_TOKEN or INPUT_GITHUB_TOKEN")
	}

	mode := strings.ToLower(getEnv("INPUT_MODE", "bot"))
	repoFull := os.Getenv("GITHUB_REPOSITORY")
	if repoFull == "" {
		log.Fatal("GITHUB_REPOSITORY environment variable is not set")
	}

	parts := strings.Split(repoFull, "/")
	if len(parts) != 2 {
		log.Fatalf("Invalid GITHUB_REPOSITORY: %s", repoFull)
	}
	owner, repo := parts[0], parts[1]

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Load repository configuration (.github/cbog.yml)
	cfg := config.LoadFromRepo(ctx, client, owner, repo, "")

	// Mode 1: Scaffold OSS Community Guidelines & Templates
	if mode == "scaffold-oss-docs" || mode == "scaffold" {
		log.Println("Running in OSS Scaffolder mode...")
		cwd, _ := os.Getwd()
		stack := scaffolder.DetectStack(cwd)
		log.Printf("Detected project tech stack: %s", stack)

		licenseType := getEnv("INPUT_LICENSE", "mit")
		scaff := scaffolder.NewScaffolder(stack, owner, licenseType)
		files, err := scaff.GenerateFiles()
		if err != nil {
			log.Fatalf("Failed to generate scaffold files: %v", err)
		}

		prURL, err := scaffolder.CreateScaffoldPR(ctx, client, owner, repo, files)
		if err != nil {
			log.Fatalf("Failed to create scaffolding PR: %v", err)
		}

		log.Printf("Successfully created scaffolding PR: %s", prURL)
		fmt.Printf("::set-output name=pr-url::%s\n", prURL)
		return
	}

	// Mode 2: Event-driven Bot & Plugin Subsystems
	eventPath := os.Getenv("GITHUB_EVENT_PATH")
	if eventPath == "" {
		log.Fatal("GITHUB_EVENT_PATH environment variable is not set")
	}

	data, err := os.ReadFile(eventPath)
	if err != nil {
		log.Fatalf("Failed to read GITHUB_EVENT_PATH: %v", err)
	}

	eventName := os.Getenv("GITHUB_EVENT_NAME")
	log.Printf("Processing event: %s", eventName)

	switch eventName {
	case "issue_comment":
		handleIssueComment(ctx, client, owner, repo, data, cfg)

	case "issues":
		handleIssuesEvent(ctx, client, owner, repo, data, cfg)

	case "pull_request":
		handlePullRequestEvent(ctx, client, owner, repo, data, cfg)

	default:
		log.Printf("Event %s is not actively watched by cbog. Skipping.", eventName)
	}
}

func handleIssueComment(ctx context.Context, client *github.Client, owner, repo string, data []byte, cfg *config.Config) {
	var event github.IssueCommentEvent
	if err := json.Unmarshal(data, &event); err != nil {
		log.Fatalf("Failed to parse issue_comment event: %v", err)
	}

	if event.GetAction() != "created" {
		return
	}

	sender := event.GetSender().GetLogin()
	if strings.HasSuffix(strings.ToLower(sender), "[bot]") {
		return
	}

	comment := event.GetComment()
	issue := event.GetIssue()
	commentBody := comment.GetBody()

	parsedCmds := lexer.ParseCommands(commentBody)
	if len(parsedCmds) == 0 {
		return
	}

	log.Printf("Found %d slash command(s) from @%s on #%d", len(parsedCmds), sender, issue.GetNumber())

	bCtx := &types.BotContext{
		Ctx:                ctx,
		Client:             client,
		Owner:              owner,
		Repo:               repo,
		IssueNumber:        issue.GetNumber(),
		CommentID:          comment.GetID(),
		Sender:             sender,
		IssueAuthor:        issue.GetUser().GetLogin(),
		IsPR:               issue.IsPullRequest(),
		CommentBody:        commentBody,
		DefaultMergeMethod: cfg.Plugins.Commands.DefaultMergeMethod,
		Config:             cfg,
	}

	if err := commands.Dispatch(bCtx, parsedCmds); err != nil {
		log.Fatalf("Command dispatch failed: %v", err)
	}
}

func handleIssuesEvent(ctx context.Context, client *github.Client, owner, repo string, data []byte, cfg *config.Config) {
	var event github.IssuesEvent
	if err := json.Unmarshal(data, &event); err != nil {
		log.Fatalf("Failed to parse issues event: %v", err)
	}

	action := event.GetAction()
	issue := event.GetIssue()

	if action == "opened" {
		_ = plugins.HandleWelcome(ctx, client, owner, repo, issue, false, cfg.Plugins.Welcome)
	}
}

func handlePullRequestEvent(ctx context.Context, client *github.Client, owner, repo string, data []byte, cfg *config.Config) {
	var event github.PullRequestEvent
	if err := json.Unmarshal(data, &event); err != nil {
		log.Fatalf("Failed to parse pull_request event: %v", err)
	}

	action := event.GetAction()
	pr := event.GetPullRequest()

	if action == "opened" {
		// Convert PR user to Issue representation for welcome greeting
		issue := &github.Issue{
			Number:            pr.Number,
			User:              pr.User,
			AuthorAssociation: pr.AuthorAssociation,
		}
		_ = plugins.HandleWelcome(ctx, client, owner, repo, issue, true, cfg.Plugins.Welcome)
	}

	if action == "opened" || action == "synchronize" || action == "reopened" || action == "edited" {
		_ = plugins.HandlePRSize(ctx, client, owner, repo, pr, cfg.Plugins.Size)
		_ = plugins.HandleTitleLint(ctx, client, owner, repo, pr, cfg.Plugins.TitleLint)
		_ = plugins.HandleAutoLabel(ctx, client, owner, repo, pr, cfg.Plugins.AutoLabel)
	}
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
