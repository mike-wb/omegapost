#!/usr/bin/env bash

echo "** Building omegapost web"
go build -v ./cmd/web && echo " * Complete"

