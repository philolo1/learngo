#!/bin/bash
GOOS=linux GOARCH=amd64 go build -o hello main.go
zip lambda-handler.zip hello
aws lambda update-function-code --function-name goFun --zip-file fileb://lambda-handler.zip --no-cli-pager
