#!/bin/sh
. ./.env

rm -rf publish
mkdir -p publish
cp package/*${VERSION}*.* publish