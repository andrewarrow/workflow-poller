#!/bin/bash
time=$(ts 2>&1 | tr -d '\0')  # Capture both stdout/stderr and remove null bytes
tag_name="dev-$time"

git tag $tag_name

echo "Created tag: $tag_name"

git push origin $tag_name
