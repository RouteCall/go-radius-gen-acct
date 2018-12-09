#!/usr/bin/env bash

GOBUILDVERSION="$(go run go-radius-gen-acct.go -v 2> /dev/null | grep -Eo '[0-9]+\.[0-9]+\.[0-9]+')"; GOOS=linux; GOARCH=amd64; GOOS="$GOOS" GOARCH="$GOARCH" go build -o go-radius-gen-acct-"$GOBUILDVERSION"-"$GOOS"-"$GOARCH"
