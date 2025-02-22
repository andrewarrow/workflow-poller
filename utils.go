package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func GetGitCommitHash() string {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}
func GitTag(prefix, env string) {
	var tagName string

	if env == "" {
		tagName = fmt.Sprintf("%s-%d", prefix, time.Now().UnixNano())
	} else {
		tagName = fmt.Sprintf("%s-%s-%d", prefix, env, time.Now().UnixNano())
	}
	cmd := exec.Command("git", "tag", tagName)
	cmd.Output()
	cmd = exec.Command("git", "push", "origin", tagName)
	cmd.Output()
}
