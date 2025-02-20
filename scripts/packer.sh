#!/bin/bash
time=$(ts 2>&1 | tr -d '\0') 
tag_name="packer-$time"

git tag $tag_name

echo "Created tag: $tag_name"

git push origin $tag_name
