#!/bin/sh
. ./.env

DESTDIR="./package/${NAME}_${VERSION}_linux-any_amd64"

mkdir -p $DESTDIR

make ${NAME}
mv ${NAME} $DESTDIR

tar -czf ${DESTDIR}.tar.gz $DESTDIR