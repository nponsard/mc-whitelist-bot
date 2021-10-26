#!/bin/sh
. ./.env

DESTDIR="./package/${NAME}_${VERSION}_windows_amd64"

mkdir -p $DESTDIR

make ${NAME}.exe
mv ${NAME}.exe $DESTDIR

zip -r ${DESTDIR}.zip $DESTDIR