package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v55/github"
	"golang.org/x/oauth2"
)

func main2() {
	// Get GitHub token from environment variable
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Println("Please set GITHUB_TOKEN environment variable")
		os.Exit(1)
	}

	// Create an authenticated GitHub client
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Repository details
	owner := "puradev" // Replace with repository owner
	repo := "aroma"    // Replace with repository name

	// List workflow runs
	opts := &github.ListWorkflowRunsOptions{
		ListOptions: github.ListOptions{
			PerPage: 30,
		},
	}

	runs, _, err := client.Actions.ListRepositoryWorkflowRuns(ctx, owner, repo, opts)
	if err != nil {
		fmt.Printf("Error fetching workflow runs: %v\n", err)
		os.Exit(1)
	}

	// Print workflow runs information
	fmt.Printf("Found %d workflow runs:\n\n", runs.GetTotalCount())
	for _, run := range runs.WorkflowRuns {
		if run.GetStatus() == "completed" {
			continue
		}
		fmt.Printf("Status: %s\n", run.GetStatus())
		fmt.Printf("Created At: %s\n", run.GetCreatedAt().Format("2006-01-02 15:04:05"))
		sha := run.GetHeadSHA()
		fmt.Printf("SHA: %s\n", sha)

		commit, _, err := client.Repositories.GetCommit(ctx, owner, repo, sha, &github.ListOptions{})
		if err == nil {
			message := strings.Split(commit.GetCommit().GetMessage(), "\n")[0] // Get first line of commit message
			fmt.Printf("Commit Message: %s\n", message)
		}

		ref := run.GetHeadBranch()
		fmt.Println("ref", ref)
		if isTag(ref) {
			fmt.Printf("Tag: %s\n", ref)
		} else {
			// Try to get associated tags for this commit
			tags, _, err := client.Repositories.ListTags(ctx, owner, repo, &github.ListOptions{})
			if err == nil {
				for _, tag := range tags {
					if tag.GetCommit().GetSHA() == sha {
						fmt.Printf("Tag: %s\n", tag.GetName())
						break
					}
				}
			}
		}

		fmt.Printf("------------------\n")
	}

}
