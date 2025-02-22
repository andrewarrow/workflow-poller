# workflow-poller

A little go program to poll github actions every 6 seconds looking for
`completed` actions that match your current git HEAD SHA 

```
git rev-parse HEAD
```

Each poll returns a `map[string]bool` of the tags with your SHA now in completed
state.

It helps a lot to have your branch and sha in your prompt:

```
autoload -Uz vcs_info
precmd() {
    vcs_info
    if git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
        GITHEAD=$(git rev-parse --short HEAD 2>/dev/null)
    else
        GITHEAD=""
    fi
}

# Format the vcs_info_msg_0_ variable
zstyle ':vcs_info:git:*' formats '%F{cyan}(%b)%f'
zstyle ':vcs_info:*' enable git

# Set up the prompt with git HEAD
setopt prompt_subst
PROMPT='%F{blue}%~%f ${vcs_info_msg_0_}%F{yellow}${GITHEAD:+[${GITHEAD}]}%f $ '
```

In `main.go` you'll find sample logic to look for completed tags, and then add
the next needed tag.


