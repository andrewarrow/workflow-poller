#!/bin/bash

curl https://aroma.trypura.io/things/$1/shadow | jq .$2
