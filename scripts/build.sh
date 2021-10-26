#!/bin/sh
. ./.env
go build -ldflags "-X main.version=${VERSION}"