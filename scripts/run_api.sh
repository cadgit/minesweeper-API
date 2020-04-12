#!/usr/bin/env bash

docker run -p 8080:8080 --env-file prod.env -d -t minesweeper-api