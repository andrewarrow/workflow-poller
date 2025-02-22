package main

import (
	"fmt"
	"time"
)

func main() {
	sha := GetGitCommitHash()
	env := "dev"
	for {
		m := ListShaActions(sha)
		if m["containerize"] && m["promote"] == false {
			fmt.Println("promoting")
			GitTag("promote", env)
		} else if m["containerize"] && m["promote"] {
			fmt.Println("deploying")
			GitTag(env, "")
		} else if m[env] {
			fmt.Println("done")
			break
		}
		fmt.Println("Waiting 6 seconds...")
		time.Sleep(time.Second * 6)
	}
}
