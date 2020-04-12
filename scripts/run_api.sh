#!/usr/bin/env bash

docker run -p 8080:8080 --env-file prod.env -t minesweeper-api