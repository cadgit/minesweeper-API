#!/bin/bash

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ../builds/minesweeper_api ../main/main.go

docker build --no-cache -t minesweeper-api ../.
