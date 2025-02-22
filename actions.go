package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v55/github"
	"golang.org/x/oauth2"
)

func ListShaActions(sha string) map[string]bool {
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
	refs := map[string]bool{}
	for _, run := range runs.WorkflowRuns {
		localSHA := run.GetHeadSHA()
		if sha != localSHA {
			continue
		}

		ref := run.GetHeadBranch()
		tokens := strings.Split(ref, "-")
		token := tokens[0]
		refs[token] = true
	}

	return refs
}

func isTag(ref string) bool {
	return len(ref) > 0 && (ref[0] == 'v' || ref[0] == 'r' || ref[0] == 't')
}
