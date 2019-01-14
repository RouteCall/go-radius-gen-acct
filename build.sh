#!/usr/bin/env bash

pGOOS=(linux freebsd)
GOARCH=amd64 
GOBUILDVERSION="$(go run go-radius-gen-acct.go -v 2> /dev/null | grep -Eo '[0-9]+\.[0-9]+\.[0-9]+')"

for GOOS in "${pGOOS[@]}"; do
  GOOS="$GOOS" GOARCH="$GOARCH" go build -o go-radius-gen-acct-"$GOBUILDVERSION"-"$GOOS"-"$GOARCH"
done
