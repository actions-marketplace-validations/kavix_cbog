package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/v60/github"
	"github.com/kavindu/contributor-bot-action/pkg/commands"
	"github.com/kavindu/contributor-bot-action/pkg/lexer"
	"github.com/kavindu/contributor-bot-action/pkg/scaffolder"
	"github.com/kavindu/contributor-bot-action/pkg/types"
	"golang.org/x/oauth2"
)

func main() {
	token := getEnv("INPUT_GITHUB_TOKEN", os.Getenv("GITHUB_TOKEN"))
	if token == "" {
		log.Fatal("Missing GITHUB_TOKEN or INPUT_GITHUB_TOKEN")
	}

	mode := strings.ToLower(getEnv("INPUT_MODE", "bot"))
	mergeMethod := strings.ToLower(getEnv("INPUT_MERGE_METHOD", "squash"))
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

	// Mode 1: Scaffold OSS Documentation & Community Guidelines
	if mode == "scaffold-oss-docs" || mode == "scaffold" {
		log.Println("Running in OSS Scaffolder mode...")
		cwd, _ := os.Getwd()
		stack := scaffolder.DetectStack(cwd)
		log.Printf("Detected project tech stack: %s", stack)

		scaff := scaffolder.NewScaffolder(stack)
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

	// Mode 2: Bot / Slash Command Execution
	eventPath := os.Getenv("GITHUB_EVENT_PATH")
	if eventPath == "" {
		log.Fatal("GITHUB_EVENT_PATH environment variable is not set")
	}

	data, err := os.ReadFile(eventPath)
	if err != nil {
		log.Fatalf("Failed to read GITHUB_EVENT_PATH: %v", err)
	}

	eventName := os.Getenv("GITHUB_EVENT_NAME")
	if eventName != "issue_comment" {
		log.Printf("Ignoring event '%s'. Contributor bot only processes 'issue_comment'.", eventName)
		return
	}

	var event github.IssueCommentEvent
	if err := json.Unmarshal(data, &event); err != nil {
		log.Fatalf("Failed to parse issue_comment event: %v", err)
	}

	if event.GetAction() != "created" {
		log.Printf("Ignoring comment action '%s' (only 'created' is handled).", event.GetAction())
		return
	}

	comment := event.GetComment()
	sender := event.GetSender().GetLogin()
	issue := event.GetIssue()

	// Ignore bot's own comments to avoid infinite loops
	if strings.HasSuffix(strings.ToLower(sender), "[bot]") {
		log.Printf("Ignoring comment from bot user: %s", sender)
		return
	}

	commentBody := comment.GetBody()
	parsedCmds := lexer.ParseCommands(commentBody)
	if len(parsedCmds) == 0 {
		log.Println("No valid slash commands found in comment.")
		return
	}

	log.Printf("Found %d command(s) from @%s on issue #%d", len(parsedCmds), sender, issue.GetNumber())

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
		DefaultMergeMethod: mergeMethod,
	}

	if err := commands.Dispatch(bCtx, parsedCmds); err != nil {
		log.Fatalf("Command dispatch failed: %v", err)
	}

	log.Println("All commands executed successfully.")
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
