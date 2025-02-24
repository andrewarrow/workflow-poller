package main

import (
	"fmt"
	"time"
)

func main() {
	sha := GetGitCommitHash()
	env := "dev"
	readyToDeploy := false
	alreadyDeployed := false
	for {
		m := ListShaActions(sha)
		if m["containerize"] && readyToDeploy == false {
			fmt.Println("promoting")
			GitTag("promote", env)
			readyToDeploy = true
		} else if m["promote"] && readyToDeploy == true && alreadyDeployed == false {
			fmt.Println("deploying")
			GitTag(env, "")
			alreadyDeployed = true
		} else if m[env] {
			fmt.Println("done")
			break
		}
		fmt.Println("Waiting 6 seconds...")
		time.Sleep(time.Second * 6)
	}
}
